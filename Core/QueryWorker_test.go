package KsanaDB
import(
        "testing" 
        "fmt"
)

func Test_rangeAggreator(t *testing.T) {  
   rangeStartTime := int64(0)
   rangeEndTime := int64(0)
   aggResult := float64(0)
   currentElementTime := int64(6)
   currentElementValue := float64(1)
   aF := getFuncMap("sum")
   startTimestamp := int64(5)
   timeRange := int64(10)
   ret := []map[string]interface{}{}

   rangeStartTime, rangeEndTime, aggResult, ret =  rangeAggreator(rangeStartTime, rangeEndTime, aggResult, currentElementTime, currentElementValue, aF, startTimestamp, timeRange, ret)

    if rangeStartTime != 5 {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        t.Error("rangeStartTime fail")    
    }

    if rangeEndTime != 15 {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        t.Error("rangeEndTime fail")    
    }

    if len(ret) != 0 {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        t.Error("ret fail")    
    }

    currentElementTime = int64(16)

    rangeStartTime, rangeEndTime, aggResult, ret =  rangeAggreator(rangeStartTime, rangeEndTime, aggResult, currentElementTime, currentElementValue, aF, startTimestamp, timeRange, ret)

    if rangeStartTime != 15 {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        t.Error("rangeStartTime fail")    
    }

    if rangeEndTime != 25 {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        t.Error("rangeEndTime fail")    
    }

    if ret[0]["timestamp"] != int64(5) {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        fmt.Println(ret[0]["timestamp"])
        t.Error("accumulate timestamp fail")    
    }

    if ret[0]["value"] != float64(1) {
        fmt.Printf("%d %d %f\n",rangeStartTime, rangeEndTime, aggResult)
        fmt.Println(ret[0]["value"])
        t.Error("accumulate value fail")    
    }
}
