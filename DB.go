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
   GetLink("127.0.0.1", 6379) 
}

func SetData(data string) {
    InputArray, err := ParseJson(data)
  
    if (err != nil) {
        return    
    }

    for _, data := range InputArray {
        hashdata := data.(map[string]interface{})
        name := ""
        dataPoints := hashdata["datapoints"]

        if hashdata["name"] == nil {
            continue    
        } else {
            name = hashdata["name"].(string) 
        }

        if dataPoints == nil {
            value, _ := (hashdata["value"].(json.Number)).Float64()

            if hashdata["timestamp"] == nil {
                //log.Fatalf("Connect failed: %s\n", err.Error()) 
                continue    
            }
            timestamp, err := (hashdata["timestamp"].(json.Number)).Int64()
 
            if err != nil {
                continue
            }


            element := make( map[string]interface{})
            keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
            tagSeq := getTagSeq(hashdata["tags"].(map[string]interface{}), prefix, name) 

 
            element["timestamp"] = strconv.FormatInt(timestamp, 10)
            element["value"] = strconv.FormatFloat(value, 'f', 6, 64)
            element["tags"] = tagSeq
 
            jstring, _ := json.Marshal(element)
            input := string(jstring[:])
            SetTimeSeries(keyname, fmt.Sprint(input), offset)
        } else {
            inputData := make(map[string][]interface{})
            for _, rowdata := range dataPoints.([]interface{}) {
                data := rowdata.([]interface{})
                timestamp, errT := (data[0].(json.Number)).Int64()
                value, errV := (data[1].(json.Number)).Float64()

                if errT != nil || errV != nil {
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
                    continue    
                }

               //FIXME: should move outside for loop 
                tagSeq := getTagSeq(hashdata["tags"].(map[string]interface{}), prefix, name) 
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

func QueryTimeSeriesData(name string, start int64, stop int64, tags []string, aggreationFunction string, timeRange int, unit string) ([]map[string]interface{} , error) {
    fmt.Println(time.Now())
    rawData := queryTimeSeries(prefix , name , start , stop )
    data, err := queryWorker(rawData, start, tags, aggreationFunction, unit, timeRange)
    fmt.Println(time.Now())
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
      d, err := getSeqByKV(prefix, name, filterList) 
      return d, err
}


