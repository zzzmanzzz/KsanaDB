package KsanaDB
import(
        "testing" 
        "fmt"
)

func Test_queryWorkerConcurrent(t *testing.T) {
    time := int64(1389024000000)
    tagFilter := []string{}
    ret, err := queryWorker(testinput,time, tagFilter,groupByTest,"sum","ms",1) 
    if err != nil {
        t.Error("queryWorkerConcurren fail") 
        fmt.Println(err)
    }
    fmt.Println(ret)
}

func Test_queryWorkerConcurrentWithFilter(t *testing.T) {
    time := int64(1389024000000)
    tagFilter := []string{"12"}
    ret, err := queryWorker(testinput,time, tagFilter,groupByTest,"sum","ms",1) 
    if err != nil {
        t.Error("queryWorkerConcurren fail") 
        fmt.Println(err)
    }
    fmt.Println(ret)
}

func Test_queryWorkerNonCurrent(t *testing.T) {
    time := int64(1389024000000)
    tagFilter := []string{}
    noGroupby := map[string][]string{}
    ret, err :=   queryWorker(testinput,time, tagFilter,noGroupby,"sum","s",1) 
    if err != nil {
        t.Error("queryWorkerConcurren fail") 
        fmt.Println(err)
    }
    fmt.Println(ret)
}

func Test_queryWorkerNonCurrentWithFilter(t *testing.T) {
    time := int64(1389024000000)
    tagFilter := []string{"5"}
    noGroupby := map[string][]string{}
    ret, err :=   queryWorker(testinput,time, tagFilter,noGroupby,"sum","s",1) 
    if err != nil {
        t.Error("queryWorkerConcurren fail") 
        fmt.Println(err)
    }
    fmt.Println(ret)
}

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

var testinput = []string{
                   `{"tags":["1","2","3"],"timestamp":"1389024000000","value":"0.000000"}`,
                   `{"tags":["4","5","6"],"timestamp":"1389024000101","value":"1.000000"}`,
                   `{"tags":["3","7","8"],"timestamp":"1389024000208","value":"2.000000"}`,
                   `{"tags":["10","11","9"],"timestamp":"1389024000309","value":"3.000000"}`,
                   `{"tags":["12","13","4"],"timestamp":"1389024000404","value":"4.000000"}`,
                   `{"tags":["14","15","3"],"timestamp":"1389024000503","value":"5.000000"}`,
                   `{"tags":["16","17","18"],"timestamp":"1389024000600","value":"6.000000"}`,
                   `{"tags":["11","19","20"],"timestamp":"1389024000706","value":"7.000000"}`,
                   `{"tags":["11","21","22"],"timestamp":"1389024000808","value":"8.000000"}`,
                   `{"tags":["23","24","3"],"timestamp":"1389024000900","value":"9.000000"}`, 
                   `{"tags":["1","18","25"],"timestamp":"1389024001004","value":"0.000000"}`,
                   `{"tags":["26","27","5"],"timestamp":"1389024001109","value":"1.000000"}`, 
                   `{"tags":["11","28","7"],"timestamp":"1389024001203","value":"2.000000"}`,
                   `{"tags":["29","3","9"],"timestamp":"1389024001300","value":"3.000000"}`, 
                   `{"tags":["11","12","30"],"timestamp":"1389024001400","value":"4.000000"}`,
                   `{"tags":["11","14","31"],"timestamp":"1389024001507","value":"5.000000"}`,
                   `{"tags":["16","32","4"],"timestamp":"1389024001600","value":"6.000000"}`,
                   `{"tags":["11","19","33"],"timestamp":"1389024001701","value":"7.000000"}`,
                   `{"tags":["18","21","34"],"timestamp":"1389024001809","value":"8.000000"}`,
                   `{"tags":["24","27","35"],"timestamp":"1389024001905","value":"9.000000"}`,
                   `{"tags":["1","27","36"],"timestamp":"1389024002008","value":"0.000000"}`,
                   `{"tags":["11","37","5"],"timestamp":"1389024002100","value":"1.000000"}`,
                   `{"tags":["27","34","7"],"timestamp":"1389024002209","value":"2.000000"}`,
                   `{"tags":["18","38","9"],"timestamp":"1389024002300","value":"3.000000"}`,
                   `{"tags":["12","39","4"],"timestamp":"1389024002402","value":"4.000000"}`,
                   `{"tags":["14","25","3"],"timestamp":"1389024002507","value":"5.000000"}`,
                   `{"tags":["16","27","40"],"timestamp":"1389024002609","value":"6.000000"}`,
                   `{"tags":["19","23","3"],"timestamp":"1389024002701","value":"7.000000"}`,
                   `{"tags":["18","21","41"],"timestamp":"1389024002801","value":"8.000000"}`,
                   `{"tags":["18","24","42"],"timestamp":"1389024002905","value":"9.000000"}`,
}

var groupByTest = map[string][]string{
   "host":{"1","5","7","9","12","14","16","19","21","24"},
   "type":{"3","4","11","18","27"},
}
