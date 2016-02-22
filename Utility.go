package ChronosDB
import (                                                                        
        "fmt" 
//        "log" 
        "time" 
        "strconv" 
) 

func getDateStartSec(timestamp int64) ( int64, int64 ) { 
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

func relativeToAbsoluteTime(input string, unit string) (int64, error) {
    tNow := time.Now()
    var tResult time.Time 
    diff, err := strconv.ParseInt(input, 10, 64)

    if err == nil {
        if unit == "ms" {
        } else if unit == "s" {
        } else if unit == "m" {
        } else if unit == "h" {
        } else if unit == "d" {
            tResult = tNow.AddDate(0, 0, int(-diff))
        } else if unit == "w" {
            tResult = tNow.AddDate(0, 0, 7 * int(-diff))
        } else if unit == "M" {
            tResult = tNow.AddDate(0, int(-diff), 0)
        } else if unit == "y" {
            tResult = tNow.AddDate(int(-diff), 0, 0)
        }
    }

    fmt.Println(tNow)
    fmt.Println(tResult)
    return tResult.UTC().Unix() * 1000, err
}

func addMonth(t time.Time, d_month int) time.Time {
    year :=  t.Year()
    month := t.Month()
    day := t.Day()
    newMonth := int(month) + d_month
    newLastDay := getLastDay(year, newMonth)
    var newDay int
    if day > newLastDay {
        newDay = newLastDay    
    } else {
        newDay = day    
    }
    fmt.Println(year)
    fmt.Println(month)
    fmt.Println(newMonth)
    fmt.Println(newLastDay)
    return time.Date(year, time.Month(month), newDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func getLastDay(year int, month int)  int {
    t := time.Date(year, time.Month(month + 1), 1, 0, 0, 0, 0, time.Local) 
    t = t.AddDate(0, 0, -1)
    return t.Day()
}
