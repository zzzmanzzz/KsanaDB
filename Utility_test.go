package KsanaDB 
import (
    "testing"
    "time"
    "strconv"
    "fmt"
)

func Test_relativeToAbsoluteTimeInMilliSecond(t *testing.T) {
    input := 256
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "ms")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        if now.UTC().Unix() * 1000 - int64(diff) != ret {
            t.Error(ret)    
        }
    }
}

func Test_relativeToAbsoluteTimeInSecond(t *testing.T) {
    input := 256
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "s")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        d := strconv.Itoa(-diff)
        dur,_  := time.ParseDuration(d + "s")
        tResult := now.Add(dur) 
        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)    
        }
    }
}

func Test_relativeToAbsoluteTimeInMinute(t *testing.T) {
    input := 256
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "m")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        d := strconv.Itoa(-diff)
        dur,_  := time.ParseDuration(d + "m")
        tResult := now.Add(dur) 
        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)    
        }
    }
}

func Test_relativeToAbsoluteTimeInHour(t *testing.T) {
    input := 256
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "h")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        d := strconv.Itoa(-diff)
        dur,_  := time.ParseDuration(d + "h")
        tResult := now.Add(dur) 
        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)    
        }
    }
}

func Test_relativeToAbsoluteTimeInDay(t *testing.T) {
    input := 256
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "d")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        tResult := now.AddDate(0, 0,int(-diff))

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
        }
    }
}

func Test_relativeToAbsoluteTimeInMonth(t *testing.T) {
    input := 5
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "M")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        tResult := now.AddDate(0, int(-diff), 0)

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
        }
    }
}

func Test_relativeToAbsoluteTimeInYear(t *testing.T) {
    input := 5
    now := time.Now()

    ret, err := relativeToAbsoluteTime(now, input, "y")
    if err != nil {
        t.Error(err)    
    } else {
        diff := input
        tResult := now.AddDate(int(-diff), 0, 0)

        if tResult.UTC().Unix() * 1000 != ret {
            t.Error(ret)   
        }
    }
}

func Test_getTimeseriesQueryCmd(t *testing.T) {
    now := time.Now().UTC().Unix() * 1000
    then := time.Now().UTC().Unix() * 1000 + 86400000 * 5 + 1000
    ret := getTimeseriesQueryCmd("CHRONOSDBv1\t", "test", now, then)
    if len(ret) != 6 {
       t.Error(ret)
    }
    fmt.Println(ret)
}

func Test_getTimeRange(t *testing.T) {
    reference := time.Now().Unix() * 1000
    Onemillisec, err := getTimeRange(reference, 1, "ms")
    Onesec, err := getTimeRange(reference, 1, "s")
    Onemin, err := getTimeRange(reference, 1, "m")
    Oneday, err := getTimeRange(reference, 1, "d")
    Oneweek, err := getTimeRange(reference, 1, "w")

    if err != nil {
        t.Error(err)    
    }
    if Onemillisec != 1 {
        t.Error("1 millisec range fail")
        fmt.Println(Onemillisec)
    }
    if Onesec != 1000 {
        t.Error("1 sec range fail")
        fmt.Println(Onesec)
    }
    if Onemin != 60000 {
        t.Error("1 min range fail")
        fmt.Println(Onemin)
    }
    if Oneday != 86400000 {
        t.Error("1 day range fail")
        fmt.Println(Oneday)
    }
    if Oneweek != 7 * 86400000 {
        t.Error("1 week range fail")
        fmt.Println(Oneday)
    }


}

