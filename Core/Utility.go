package KsanaDB
import (                                                                        
     "fmt" 
//   "log" 
     "time" 
     "strconv"
     "errors"
     "encoding/json"
) 

func getDateStartSec(timestamp int64) ( int64, int64 ) { 
        offset := timestamp % 86400000
        return timestamp - offset, offset
}

func relativeToAbsoluteTime(tNow time.Time, diff int, unit string) (int64, error) {
    var tResult time.Time 
    if unit == "ms" {
        return tNow.UTC().Unix() * 1000 - int64(diff), nil
    }

    if unit == "s" || unit == "m" || unit == "h" {
         d := strconv.Itoa(-diff)           
         dur, err := time.ParseDuration(d + unit)
         if err == nil {
             tResult = tNow.Add(dur)
             ret := tResult.UTC().Unix() * 1000 
             return ret, err
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
    } else {
        return 0, errors.New("unfeasible unit " + unit)    
    }

    return tResult.UTC().Unix() * 1000, nil 
}

func getTimeRange(reference int64, diff int, unit string) (int64, error) {
    ret, err := relativeToAbsoluteTime(time.Unix(reference/1000, 1000000 * (reference%1000)),  -diff, unit)
    return ret - reference, err
}

func getTimeseriesQueryCmd(prefix string, metricName string, from int64, to int64) ([]map[string]string )  {
       keyPrefix := prefix + metricName + "\t"
     
       from0, fromOffset := getDateStartSec(from)
       to0, toOffset := getDateStartSec(to)
       begin := time.Unix(from0/1000, 1000000 * (from0%1000))
       end := time.Unix(to0/1000, 1000000 * (to0%1000))
       
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

func getTagSeq (tags map[string]interface{}, prefix string, name string) ([]interface{}) {
    var jTagSeq string

    jTagSeq = "[]"
    if tags != nil {
        kvArray := []string{}
        for k, v := range tags {
            data := v
            switch v.(type) {
                case float64 :
                    data = fmt.Sprintf("%f",v)
            }

            kvArray = append(kvArray, fmt.Sprintf("%s\t%s",removeTab(k),removeTab(data.(string)))) 
        }
        jTagSeq = setTags(prefix, name, kvArray)
    }
    jarray,_ := ParseJson(jTagSeq)
    return jarray
}

func getQueryTime(tNow time.Time, unit *string, value *json.Number) (int64, error) {
    var absTime int64

    if unit == nil || value == nil {
        return -1, errors.New("Need set relative time")    
    } else {
        relativeTime, err := value.Int64()
        if err != nil {
             return -2, err 
        }
        absTime, err = relativeToAbsoluteTime(tNow, int(relativeTime), *unit)
        if err != nil {
            return -3, err 
        }
    }
    return absTime, nil
}
