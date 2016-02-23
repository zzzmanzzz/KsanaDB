package ChronosDB 
import (
    "testing"
    "time"
    "strconv"
)

func Test_relativeToAbsoluteTime(t *testing.T) {
    input := "5"
    now := time.Now()
    ret, err := relativeToAbsoluteTime(now, input, "M")
    if err != nil {
        t.Error("relativeToAbsoluteTime err")    
    } else {
        diff, err := strconv.ParseInt(input, 10, 64)
        if err != nil {
            t.Error("relativeToAbsoluteTime parse int err")   
        }
        tResult := now.AddDate(0, int(-diff), 0)

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error("relativeToAbsoluteTime result err")   
        }
    }
}
