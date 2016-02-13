package ChronosDB 
import (
    sjson "github.com/bitly/go-simplejson"
    "fmt"
    "log"
    "encoding/json"
    "time"
    "strconv"
)

var prefix = "CHRONOSDB\t"

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
        } else {
            name = name.(string)    
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
            zeroOclock , offset := getDateStartSec(timestamp)
            keyname := prefix + name.(string) + "\t" + strconv.FormatInt(zeroOclock, 10)
            
            SetTimeSeries(keyname, strconv.FormatFloat(value, 'f', 6, 64), offset, nil)
            fmt.Println(keyname)
            fmt.Println(offset)
        } else {
            
            //TODO: add a function to bulk insert                

            fmt.Println(name)
            fmt.Println(dataPoints)
        }


    }
}

func getDateStartSec(timestamp int64) (int64, int64 ) {
     const shortForm = "2006-01-02"
     tm := time.Unix(timestamp/1000, 0)
     DateStart :=  tm.Format(shortForm)

     st, _ := time.Parse(shortForm, DateStart)
     dateZeroOclock := st.UTC().Unix() * 1000
     //fmt.Println(st.UTC().Unix())
     //fmt.Println(timestamp)
     //fmt.Println(timestamp - 1000 * st.UTC().Unix() )
     return dateZeroOclock, timestamp - dateZeroOclock
}

//func AddDataPoint(timestamp unit32, data []string

