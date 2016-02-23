package ChronosDB 
import (
    "testing"
    "time"
    "strconv"
)

func Test_relativeToAbsoluteTimeInDay(t *testing.T) {
    input := "256"
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "d")
    if err != nil {
        t.Error(err)    
    } else {
        diff, err := strconv.ParseInt(input, 10, 64)
        if err != nil {
            t.Error( err)   
        }
        tResult := now.AddDate(0, 0,int(-diff))

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
        }
    }
}
func Test_relativeToAbsoluteTimeInMonth(t *testing.T) {
    input := "5"
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "M")
    if err != nil {
        t.Error(err)    
    } else {
        diff, err := strconv.ParseInt(input, 10, 64)
        if err != nil {
            t.Error( err)   
        }
        tResult := now.AddDate(0, int(-diff), 0)

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
        }
    }
}

func Test_relativeToAbsoluteTimeInYear(t *testing.T) {
    input := "5"
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "y")
    if err != nil {
        t.Error(err)    
    } else {
        diff, err := strconv.ParseInt(input, 10, 64)
        if err != nil {
            t.Error(err)   
        }
        tResult := now.AddDate(int(-diff), 0, 0)

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
            t.Error(tResult.UTC().Unix() * 1000 )   
        }
    }
}
