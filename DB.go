package ChronosDB 
import (
    sjson "github.com/bitly/go-simplejson"
    "fmt"
    "log"
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
        hashdata, _  := data.(map[string]interface{})
        name := hashdata["name"]
        dataPoints := hashdata["datapoints"]

        if name == nil {
            continue    
        }

        if dataPoints == nil {
            value := hashdata["value"]
            timestamp := hashdata["timestamp"]
            TODO: add a function to bulk insert                
            fmt.Println(name)
            fmt.Println(value)
            fmt.Println(timestamp)
        } else {
            
            TODO: add a function to insert                
            fmt.Println(name)
            fmt.Println(dataPoints)
        }


    }
}

//func AddDataPoint(timestamp unit32, data []string

