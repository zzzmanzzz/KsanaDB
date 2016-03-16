
package KsanaDB
import (
    sjson "github.com/bitly/go-simplejson"
    "encoding/json"
    "strconv"
)

func ParseQueryJson(data string) (*Query, error) {
    var q *Query
    b := []byte(data)  
    err := json.Unmarshal(b, &q)
    if err != nil {                                                  
            return nil, err             
    }
    return q, nil
}

func ParseDataJson(data string) ([]DataPoint, error) {
    var DataPoints []DataPoint   
    b := []byte(data)  
    err := json.Unmarshal(b, &DataPoints)
    if err != nil {                                                  
            return nil, err             
    }
    return DataPoints, nil
}

func ParseJson(data string) ([]interface{}, error) {
    js, err := sjson.NewJson([]byte(data))
        if err != nil {                                                  
                return nil, err             
        }
    InputArray,_ := js.Array()
        return InputArray, nil
}

func ParseJsonHash(data string) (int64, float64, []string, error) {
     var d map[string]*json.RawMessage
     err := json.Unmarshal([]byte(data), &d)

     if err != nil {                                                  
         return 0, 0, []string{}, err             
     }

     var ts string
     err = json.Unmarshal(*d["timestamp"], &ts) 
     timestamp, err := strconv.ParseInt(ts, 10, 64)
     if err != nil {                                                  
         return 0, 0, []string{}, err             
     }
  
     var v string
     err = json.Unmarshal(*d["value"], &v) 
     value, err := strconv.ParseFloat(v, 64)
     if err != nil {                                                  
         return 0, 0, []string{}, err             
     }

     var tags []string
     err = json.Unmarshal(*d["tags"], &tags)   
     if err != nil {                                                  
         return 0, 0, []string{}, err             
     }

    return timestamp, value, tags, nil
}
