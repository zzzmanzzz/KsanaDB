package ChronosDB
import (                                                                        
//        "fmt" 
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
        return dateZeroOclock, timestamp - dateZeroOclock 
} 

func relativeToAbsoluteTime(tNow time.Time, diff int, unit string) (int64, error) {
    var tResult time.Time 
    if unit == "ms" || unit == "s" || unit == "m" || unit == "h" {
         d := strconv.Itoa(-diff)           
         dur, err := time.ParseDuration(d + unit)
         if err == nil {
             tResult = tNow.Add(dur)
             return tResult.UTC().Unix() * 1000, err
         } else {
             return 0, err    
         }
    } 
        
    if unit == "d" {
        tResult = tNow.AddDate(0, 0, -diff)
    } else if unit == "w" {
        tResult = tNow.AddDate(0, 0, 7 * -diff)
    } else if unit == "M" {
        tResult = tNow.AddDate(0, -diff, 0)
    } else if unit == "y" {
        tResult = tNow.AddDate(-diff, 0, 0)
    }
    return tResult.UTC().Unix() * 1000, nil 
}

