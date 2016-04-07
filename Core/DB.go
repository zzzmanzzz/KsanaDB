package KsanaDB
import (
    "fmt"
    "log"
    "encoding/json"
    "strconv"
    "time"
    "errors"
    "strings"
)

var prefix = "KSANADBv1\t"

func SetData(data string) *error {
    InputArray, err := ParseDataJson(data)
  
    if (err != nil) {
        log.Println(err.Error())
        return &err
    }

    for _, hashdata := range InputArray {
        name := ""
        dataPoints := hashdata.Datapoints

        if hashdata.Name == "" {
            continue    
        } else {
            name = hashdata.Name 
        }

        if isMetricNameValidate(name) != true {
            continue    
        }

        if dataPoints == nil {
            value, _ := hashdata.Value.Float64()

            if hashdata.Timestamp == nil {
                log.Println(err.Error())
                continue    
            }
            timestamp, err := hashdata.Timestamp.Int64()
 
            if err != nil {
                log.Println(err.Error())
                continue
            }


            element := make( map[string]interface{})
            keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
            tagSeq := getTagSeq(hashdata.Tags, prefix, name) 

 
            element["timestamp"] = strconv.FormatInt(timestamp, 10)
            element["value"] = strconv.FormatFloat(value, 'f', 6, 64)
            element["tags"] = tagSeq
 
            jstring, _ := json.Marshal(element)
            input := string(jstring[:])
            SetTimeSeries(keyname, fmt.Sprint(input), offset)
        } else {
            inputData := make(map[string][]interface{})
            tagSeq := getTagSeq(hashdata.Tags, prefix, name) 

            for _, data := range dataPoints {
                timestamp, errT := data[0].Int64()
                value, errV := data[1].Float64()

                if errT != nil || errV != nil {
                    log.Println(err.Error())
                    continue    
                }

                element := make( map[string]interface{})
                element["timestamp"] = strconv.FormatInt(timestamp, 10)
                element["value"] = strconv.FormatFloat(value, 'f', 6, 64)
                element["tags"] = tagSeq
 
                jstring, _ := json.Marshal(element)
                input := string(jstring[:])
                keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
                inputData[keyname] = append(inputData[keyname], offset)
                inputData[keyname] = append(inputData[keyname], input)
            }

            for k := range inputData {
                _, err := BulkSetTimeSeries(k, inputData[k])
                if err != nil {
                    log.Println(err.Error())
                    continue    
                }
            }
        }
    }
    return nil
}

func QueryData(q *Query) (string , error) {
    var start int64 
    var end int64 
    var err error

    if q.Metric.Name == nil {
        return "", errors.New("Need set Metric name")    
    }

    tNow := time.Now()

    if q.StartAbsolute == nil {
        if q.StartRelative == nil {
            return "", errors.New("Need set absolute start time")    
        } else {
           start, err = getQueryTime(tNow, q.StartRelative.Unit, q.StartRelative.Value) 
           if err != nil {
               return "", err 
           }
        }
    } else {
        start, err = q.StartAbsolute.Int64()
        if err != nil {
            return "", err 
        }
    }

    if q.EndAbsolute == nil {
        if q.EndRelative == nil {
            return "", errors.New("Need set absolute end time")    
        } else {
           end, err = getQueryTime(tNow, q.EndRelative.Unit, q.EndRelative.Value) 
           if err != nil {
               return "", err 
           }
        }
    } else {
        end, err = q.EndAbsolute.Int64()
        if err != nil {
            return "", err 
        }
    }

    if start > end {
        return "", errors.New(fmt.Sprintf("end time(%d) little than start time(%d)", end, start))
    }


    tagFilter := []string{} 
    for k, v := range(q.Metric.Tags) {
        tagFilter = append(tagFilter, fmt.Sprintf("%s\t%s", k,v))    
    }
    groupByTag := q.Metric.GroupBy
    aggreationFunction := q.Metric.Aggregator.Name

    var unit string
    if q.Metric.Aggregator.Sampling.Unit != nil {
        unit = *q.Metric.Aggregator.Sampling.Unit
    }

    var timeRange int64
    if q.Metric.Aggregator.Sampling.Value != nil {
        timeRange, err = q.Metric.Aggregator.Sampling.Value.Int64()
            if err != nil {
                return "", err 
            }
    }

    ret, err :=  QueryTimeSeriesData(*q.Metric.Name, start, end, tagFilter, groupByTag, aggreationFunction, int(timeRange), unit)
    return ret, err
}

func QueryTimeSeriesData(name string, start int64, stop int64, tagFilter []string, groupByTag []string, aggreationFunction string, timeRange int, unit string) (string , error) {
    
    groupBy := map[string][]string{}
    var reverseHash map[string]string

    if len(groupByTag) > 0 {
        for _,t := range groupByTag {
            tmp := GetMetricsTagSeq(name, t)
            if len(tmp.Seq[t]) > 0 {
                groupBy[t] = tmp.Seq[t]
            } else {
                return "", errors.New("Input Group by tag(s) not Exist")    
            }
            reverseHash = tmp.Val
        }
    }

    tagFilterSeq, err := GetFilterSeq(name, tagFilter)
    if err != nil {
        return "", err    
    }

    for _, sq := range (tagFilterSeq) { 
        if sq == "" {
            return "", errors.New("Input Filter Tag(s) not Exist")
        }
    }

    rawData, err := queryTimeSeries(prefix , name , start , stop )

    if err != nil {
        return "{}", err
    }

    if len(rawData) == 0 {
        return "{}", nil    
    }


    data, err := queryWorker(rawData, start, tagFilterSeq, groupBy, aggreationFunction, unit, timeRange)

    ret, err := generateOutputData(data, reverseHash, name, start, stop, tagFilter, groupByTag, aggreationFunction, timeRange, unit)

    fmt.Print("Find record(s): ") 
    fmt.Println(len(rawData)) 
    return ret, err
}

func GetMetricsTag(name string, target string, keyName string) (string, error)  {
    var result map[string][]string
    ret :=  map[string]map[string][]string{}
    var data string
    switch target {
        case "All", "TagKey", "TagValue" :
            data = getTags(prefix, name, target, keyName)
        default:
            return "", errors.New("target should be All, TagKey or TagValue") 
    }
    json.Unmarshal([]byte(data), &result)
    ret["result"] = result
    jret, err := json.Marshal(ret)                                                                                  
    return string(jret), err 
}

func GetMetricsTagSeq(name string, keyName string) (AllTagSeqType)  {
    data := getTags(prefix, name, "TagSeq", keyName)
    var ret AllTagSeqType 
    json.Unmarshal([]byte(data), &ret)
    return ret
}

func GetMetric() (string, error)  {
    data, err := getMetric(prefix)
    if err != nil {
        return "", err    
    }
    var result []string 
    ret := map[string] []string{}
    json.Unmarshal([]byte(data), &result)
    ret["result"] = result
    jret, err := json.Marshal(ret)                                                                                  
    return string(jret), err 
}

func DeleteMetric(name string) (string, error)  {
    if isMetricNameValidate(name) != true {
        return "", errors.New("Metric name can't content /t")    
    }
    data, err := getMetricKeys(prefix, name)
    if err != nil {
        return "", err    
    }
    deleteKeys(data)
    return name, err 
}

func GetFilterSeq(name string, filterList []string) ([]string, error){
      if len(filterList) == 0 {
          ret := []string{} 
          return ret, nil   
      } 
      d, err := getSeqByKV(prefix, name, filterList) 
      return d, err
}

func generateOutputData(result map[string][]map[string]interface{}, reverseHash map[string]string, name string, start int64, stop int64, tagFilter []string, groupByTag []string, aggregationFunction string, timeRange int, unit string) (string , error) {
    resultData := ResultType{}
    resultData.Group = []GroupType{}

    resultData.Name = name
    resultData.GroupBy = groupByTag

    resultData.Start = start
    resultData.End = stop

    resultData.AggregateFunction = aggregationFunction
    resultData.TimeRange = timeRange
    resultData.TimeUnit = unit

    if len(tagFilter) > 0 {
       resultData.Filter = map[string]string{}
       for _,d := range(tagFilter) {
           ka := strings.Split(d, "\t")
           resultData.Filter[ka[0]] = ka[1]
       }
    }

    for k,v := range(result) {
       gp := GroupType{}
       gp.Tags = map[string]string{}
       gp.Values = [][]interface{}{}

       if k != "single" {
           ka := strings.Split(k, "\t")
           for _, s := range(ka) {
               tagPair := reverseHash[s]
               tagKV := strings.Split(tagPair, "\t")
               gp.Tags[tagKV[0]] = tagKV[1]
           }
       }

       for _,d := range(v) {
           ele := []interface{}{d["timestamp"], d["value"]}
           gp.Values = append(gp.Values, ele)
       }
       resultData.Group = append(resultData.Group, gp)
    }
    jret, err := json.Marshal(resultData)
    return string(jret), err
}
