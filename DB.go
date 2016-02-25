package ChronosDB 
import (
    "fmt"
    "time"
//    "log"
    "encoding/json"
    "strconv"
)

var prefix = "CHRONOSDBv1\t"

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
                //log.Fatalf("Connect failed: %s\n", err.Error()) 
                continue    
            }

            element := make( map[string]string)
            keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
 
            element["timestamp"] = strconv.FormatInt(timestamp, 10)
            element["value"] = strconv.FormatFloat(value, 'f', 6, 64)
            element["tags"] = ""
 
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

                
                element := make( map[string]string)
                element["timestamp"] = strconv.FormatInt(timestamp, 10)
                element["value"] = strconv.FormatFloat(value, 'f', 6, 64)
                element["tags"] = ""
 
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
            fmt.Println(inputData)
        }


    }
}


func QueryTimeSeriesData(name string, start int64, stop int64)  {
    fmt.Println(time.Now())
    queryTimeSeries(prefix , name , start , stop )
    fmt.Println(time.Now())
}

//func AddDataPoint(timestamp unit32, data []string

