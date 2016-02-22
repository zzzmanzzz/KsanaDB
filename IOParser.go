
package ChronosDB 
import (
    sjson "github.com/bitly/go-simplejson"
    "log"
//    "encoding/json"
//    "strconv"
)
func ParseJson(data string) ([]interface{}, error) {
    js, err := sjson.NewJson([]byte(data))
        if err != nil {                                                  
            log.Fatalf("Connect failed: %s\n", err.Error())              
            return nil, err             
        }
    InputArray,_ := js.Array()
    return InputArray, nil
}
