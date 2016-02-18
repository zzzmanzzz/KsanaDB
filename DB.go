package ChronosDB 
import (
    sjson "github.com/bitly/go-simplejson"
    "fmt"
    "log"
    "encoding/json"
//    "time"
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
            keyname, offset := generateTimeSeriesData(name , timestamp)
            SetTimeSeries(keyname, strconv.FormatFloat(value, 'f', 6, 64), offset, nil)
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

                keyname, offset := generateTimeSeriesData(name , timestamp)
                inputData[keyname] = append(inputData[keyname], value)
                inputData[keyname] = append(inputData[keyname], offset)
            }
            //TODO: add a function to bulk insert                
            for k := range inputData {
                tag := []string{}
                _, err := BulkSetTimeSeries(k, inputData[k], tag)
                if err != nil {
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
                    continue    
                }
            }
            fmt.Println(inputData)
        }


    }
}
/*
func QueryTimeSeriesData(query string) () {
    js, err := sjson.NewJson([]byte(query))
        if err != nil {                                                  
            log.Fatalf("Connect failed: %s\n", err.Error())              
                return                     
        }
   return js 
}
*/
func generateTimeSeriesData(name string, timestamp int64) (string, int64 ) {
     zeroOclock , offset := getDateStartSec(timestamp)
     keyname := prefix + name + "\t" + strconv.FormatInt(zeroOclock, 10)
     return keyname, offset
}
/*
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
*/
//func AddDataPoint(timestamp unit32, data []string

