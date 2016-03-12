
package KsanaDB
import (
    sjson "github.com/bitly/go-simplejson"
    "log"
    "encoding/json"
    "strconv"
)

func ParseDataJson(data string) ([]DataPoint, error) {
    var DataPoints []DataPoint   
    b := []byte(data)  
    err := json.Unmarshal(b, &DataPoints)
    if err != nil {                                                  
            log.Fatalf("Connect failed: %s\n", err.Error())              
            return nil, err             
    }
    return DataPoints, nil
}

func ParseJson(data string) ([]interface{}, error) {
    js, err := sjson.NewJson([]byte(data))
        if err != nil {                                                  
            log.Fatalf("Connect failed: %s\n", err.Error())              
                return nil, err             
        }
    InputArray,_ := js.Array()
        return InputArray, nil
}

func ParseJsonHash(data string) (int64, float64, []string, error) {
     var d map[string]*json.RawMessage
     err := json.Unmarshal([]byte(data), &d)

     if err != nil {                                                  
         log.Fatalf("Connect failed: %s\n", err.Error())              
         return 0, 0, []string{}, err             
     }

     var ts string
     err = json.Unmarshal(*d["timestamp"], &ts) 
     timestamp, err := strconv.ParseInt(ts, 10, 64)
     if err != nil {                                                  
         log.Fatalf("Connect failed: %s\n", err.Error())              
         return 0, 0, []string{}, err             
     }
  
     var v string
     err = json.Unmarshal(*d["value"], &v) 
     value, err := strconv.ParseFloat(v, 64)
     if err != nil {                                                  
         log.Fatalf("Connect failed: %s\n", err.Error())              
         return 0, 0, []string{}, err             
     }

     var tags []string
     err = json.Unmarshal(*d["tags"], &tags)   
     if err != nil {                                                  
         log.Fatalf("Connect failed: %s\n", err.Error())              
         return 0, 0, []string{}, err             
     }

    return timestamp, value, tags, nil
}
