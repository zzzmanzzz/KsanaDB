package ChronosDB 
import (
    "testing"
   // "time"
)

func Test_relativeToAbsoluteTime(t *testing.T) {
    ret, err := relativeToAbsoluteTime("5", "M")
    if err != nil {
        t.Error("relativeToAbsoluteTime err")    
    } else {
        ret = ret     
    }
}
