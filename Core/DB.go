package KsanaDB
import (
    "fmt"
    "time"
//    "log"
    "encoding/json"
    "strconv"
)

var prefix = "KSANADBv1\t"

func Connect() {
   InitRedis("tcp", "127.0.0.1:6379") 
}

func SetData(data string) {
    InputArray, err := ParseDataJson(data)
  
    if (err != nil) {
        return    
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
                //log.Fatalf("Connect failed: %s\n", err.Error()) 
                continue    
            }
            timestamp, err := hashdata.Timestamp.Int64()
 
            if err != nil {
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
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
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
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
                    continue    
                }
            }
            //fmt.Println(inputData)
        }
    }
}

func QueryTimeSeriesData(name string, start int64, stop int64, tagFilter []string, groupByTag []string, aggreationFunction string, timeRange int, unit string) ([]map[string]interface{} , error) {
    fmt.Println(time.Now())

    
    groupBy := map[string][]string{}

    if len(groupByTag) > 0 {
        for _,t := range groupByTag {
            tmp := GetMetricsTag(name, "TagSeq", t)
            groupBy[t] = tmp[t]
        }
    }

    fmt.Println(groupBy)


    tagFilterSeq, err := GetFilterSeq(name, tagFilter)
    if err != nil {
        return nil, err    
    }

    fmt.Println(tagFilterSeq)
    rawData := queryTimeSeries(prefix , name , start , stop )
    data, err := queryWorker(rawData, start, tagFilterSeq, groupBy, aggreationFunction, unit, timeRange)
    fmt.Println(time.Now())
    fmt.Println("Raw Data length")
    fmt.Println(len(rawData))
    fmt.Println("Data length")
    fmt.Println(len(data))
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


