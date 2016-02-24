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

func getTimeseriesQueryCmd(prefix string, metricName string, from int64, to int64) ([]map[string]string )  {
       keyPrefix := prefix + metricName + "\t"
     
       from0, fromOffset := getDateStartSec(from)
       to0, toOffset := getDateStartSec(to)
       begin := time.Unix(from0/1000, from0%1000)
       end := time.Unix(to0/1000, to0%1000)
       
       ret := []map[string]string{}
       element := make( map[string]string)

       element["keyName"] = keyPrefix + strconv.FormatInt(from0, 10)
       element["from"] = strconv.FormatInt(fromOffset, 10) 

       if from0 == to0 {
           element["to"] = strconv.FormatInt(toOffset, 10)
           ret = append(ret, element)
           return ret
       } else {
           element["to"] = "inf"
       }

       ret = append(ret, element)

       for i := begin.AddDate(0, 0, 1) ;  i.Before(end) ; i = i.AddDate(0, 0, 1) {
           element := make( map[string]string)
           element["keyName"] = keyPrefix + strconv.FormatInt(i.UTC().Unix()*1000, 10)
           element["from"] = "-inf"
           element["to"] = "inf"
           ret = append(ret, element)
       }

       element = make( map[string]string)
       element["keyName"] = keyPrefix + strconv.FormatInt(to0, 10)
       element["from"] = "-inf"
       element["to"] = strconv.FormatInt(toOffset, 10)

       ret = append(ret, element)
       return ret
}

func generateTimeSeriesData(prefix string, name string, timestamp int64) (string, int64 ) {
     zeroOclock , offset := getDateStartSec(timestamp)
     keyname := prefix + name + "\t" + strconv.FormatInt(zeroOclock, 10)
     return keyname, offset
}
