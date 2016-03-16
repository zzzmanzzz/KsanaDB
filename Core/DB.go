package KsanaDB
import (
    "fmt"
    "log"
    "encoding/json"
    "strconv"
    "time"
    "errors"
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

            //fmt.Println(inputData)
        }
    }
    return nil
}

func QueryData(q *Query) (map[string][]map[string]interface{} , error) {
    var start int64 
    var end int64 
    var err error

    if q.Metric.Name == nil {
        return nil, errors.New("Need set Metric name")    
    }

    tNow := time.Now()

    if q.StartAbsolute == nil {
        if q.StartRelative == nil {
            return nil, errors.New("Need set absolute start time")    
        } else if q.StartRelative.Unit == nil || q.StartRelative.Value == nil {
            return nil, errors.New("Need set relative start time")    
        } else {
           rstart, err := q.StartRelative.Value.Int64()
           if err != nil {
               return nil, err 
           }
           start, err = relativeToAbsoluteTime(tNow, int(rstart), *q.StartRelative.Unit)
           if err != nil {
               return nil, err 
           }
        }
    } else {
        start, err = q.StartAbsolute.Int64()
        if err != nil {
            return nil, err 
        }
    }

    if q.EndAbsolute == nil {
        if q.EndRelative == nil {
            return nil, errors.New("Need set absolute end time")    
        } else if q.EndRelative.Unit == nil || q.EndRelative.Value == nil {
            return nil, errors.New("Need set relative end time")    
        } else {
           rend, err := q.EndRelative.Value.Int64() 
           if err != nil {
               return nil, err 
           }
           end, err = relativeToAbsoluteTime(tNow, int(rend), *q.EndRelative.Unit)
           if err != nil {
               return nil, err 
           }
        }
    } else {
        end, err = q.EndAbsolute.Int64()
        if err != nil {
            return nil, err 
        }
    }
    tagFilter := []string {}
    groupByTag := q.Metric.GroupBy
    aggreationFunction := q.Metric.Aggregator.Name
    unit := *q.Metric.Aggregator.Sampling.Unit
    timeRange, err := q.Metric.Aggregator.Sampling.Value.Int64()
    if err != nil {
        return nil, err 
    }

    ret, err :=  QueryTimeSeriesData(*q.Metric.Name, start, end, tagFilter, groupByTag, aggreationFunction, int(timeRange), unit)
    return ret, err
}

func QueryTimeSeriesData(name string, start int64, stop int64, tagFilter []string, groupByTag []string, aggreationFunction string, timeRange int, unit string) (map[string][]map[string]interface{} , error) {
    
    groupBy := map[string][]string{}

    if len(groupByTag) > 0 {
        for _,t := range groupByTag {
            tmp := GetMetricsTag(name, "TagSeq", t)
            groupBy[t] = tmp[t]
        }
    }

    tagFilterSeq, err := GetFilterSeq(name, tagFilter)
    if err != nil {
        return nil, err    
    }

    rawData := queryTimeSeries(prefix , name , start , stop )
    if len(rawData) == 0 {
        return map[string][]map[string]interface{}{}, nil    
    }
    data, err := queryWorker(rawData, start, tagFilterSeq, groupBy, aggreationFunction, unit, timeRange)
    return data, err
}

func GetMetricsTag(name string, target string, keyName string)(map[string][]string)  {
    var data string
    switch target {
        case "All", "TagKey", "TagValue", "TagSeq" :
            data = getTags(prefix, name, target, keyName)
      //  default:
           
    }
    var ret map[string][]string
    json.Unmarshal([]byte(data), &ret)
    return ret
}

func GetFilterSeq(name string, filterList []string) ([]string, error){
      if len(filterList) == 0 {
          ret := []string{} 
          return ret, nil   
      } 
      d, err := getSeqByKV(prefix, name, filterList) 
      return d, err
}


