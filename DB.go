package ChronosDB 
import (
    sjson "github.com/bitly/go-simplejson"
    "fmt"
    "log"
    "encoding/json"
)

func Connect() {
   GetLink("127.0.0.1", 6379) 
}

func SetData(data string) {
    js, err := sjson.NewJson([]byte(data))
        if err != nil {                                                                                                    
            log.Fatalf("Connect failed: %s\n", err.Error())                                                                
                return                                                                                                     
        }

    InputArray,_ := js.Array()    
    for _, data := range InputArray {
        hashdata  := data.(map[string]interface{})
        name := hashdata["name"]
        dataPoints := hashdata["datapoints"]

        if name == nil {
            continue    
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

            //TODO: add a function to insert                
            fmt.Println(name)
            fmt.Println(value)

            getDateStartSec(timestamp)
        } else {
            
            //TODO: add a function to bulk insert                

            fmt.Println(name)
            fmt.Println(dataPoints)
        }


    }
}

func getDateStartSec(timestamp int64) {
     fmt.Println(timestamp)
}

//func AddDataPoint(timestamp unit32, data []string

