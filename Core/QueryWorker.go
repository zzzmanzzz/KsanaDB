package KsanaDB
import(
     "log"
     "strings"
)

func queryWorker(dataList []string, startTimestamp int64, tagFilter []string, groupByTag map[string][]string, aggregateFunction string, sampleUnit string, sampleRange int) (map[string][]map[string]interface{}, error){

    var ret map[string][]map[string]interface{}
    hasGroupBy := len(groupByTag) > 0
    var err error
    if hasGroupBy == false {
        ret, err = nonConcurrentQuery(dataList, startTimestamp, tagFilter, aggregateFunction, sampleUnit, sampleRange)
    } else {
        ret, err = concurrentQuery(dataList, startTimestamp, tagFilter, groupByTag, aggregateFunction, sampleUnit, sampleRange)
    }
    return ret, err
}

func concurrentPart(key string, startTimestamp int64, timeRange int64, aggregateFunction string, in chan map[string]interface{}, out chan map[string][]map[string]interface{}) {
    aggResult := float64(0)
    rangeStartTime := int64(0)
    rangeEndTime := int64(0)
    ret := map[string][]map[string]interface{}{}
    localResult := []map[string]interface{}{}

    aggFun := aggreatorFactory(aggregateFunction) 

    for {
        d, more := <-in
        if more == false {
            break    
        }
        tc := d["timestamp"].(int64)
        vc := d["value"].(float64)

        rangeStartTime, rangeEndTime, aggResult, localResult = aggFun(rangeStartTime, rangeEndTime, aggResult, tc, vc, startTimestamp, timeRange, localResult)
    }
    if len(localResult) > 0 {
         ret[key] = localResult
    }
    out <- ret
}

func concurrentQuery(dataList []string, startTimestamp int64, tagFilter []string, groupByTag map[string][]string, aggregateFunction string, sampleUnit string, sampleRange int) (map[string][]map[string]interface{}, error){
    ret := map[string][]map[string]interface{}{}
    end := len(dataList)  

    var timeRange int64
    if isTimeRangeFunction(aggregateFunction) {
        var err error
        timeRange, err = getTimeRange(startTimestamp, sampleRange, sampleUnit )
            if err != nil {
                return nil, err    
            }
    }

    hasTagFilter := len(tagFilter) > 0
    groupByInputChan := map[string]chan map[string]interface{}{}
    out := make(chan map[string][]map[string]interface{}, 1000)

    for i := 0; i< end; i ++ {
        tc, vc, tags, err := ParseJsonHash(dataList[i])
     
        if err != nil {
            log.Println(err)
            continue    
        }
        
        if hasTagFilter && filter(tagFilter, tags) == false {
            continue
        }
        
        //unique and sort hit group tags and filter tags
        gkey := groupBy(groupByTag, tags)
        if len(gkey) > 0 {
            key := strings.Join(gkey,"\t")
            if groupByInputChan[key] == nil {
                groupByInputChan[key] = make(chan map[string]interface{}, 1000) 
                go concurrentPart(key, startTimestamp, timeRange, aggregateFunction, groupByInputChan[key], out)
            }
            ele := map[string]interface{}{}
            ele["value"] = vc
            ele["timestamp"] = tc
            groupByInputChan[key] <- ele
        } else {
            continue    
        }
    }

    for _, v := range(groupByInputChan) {
        close(v)
    }
    
    counter := 0
    for d := range(out) {
        counter++
        
        for k, v := range(d) {
            ret[k] = v
        }
        
        if counter == len(groupByInputChan) {
            close(out)
            break
        }

    }
    return ret, nil
}

func nonConcurrentQuery(dataList []string, startTimestamp int64, tagFilter []string, aggregateFunction string, sampleUnit string, sampleRange int) (map[string][]map[string]interface{}, error){
    
    ret := []map[string]interface{}{}
    result := map[string][]map[string]interface{}{}
    end := len(dataList)  

    var timeRange int64
    if isTimeRangeFunction(aggregateFunction) {
        var err error
        timeRange, err = getTimeRange(startTimestamp, sampleRange, sampleUnit )
            if err != nil {
                return nil, err    
            }
    }

    aggResult := float64(0)
    rangeStartTime := int64(0)
    rangeEndTime := int64(0)

    aggfun := aggreatorFactory(aggregateFunction) 

    hasTagFilter := len(tagFilter) > 0

    for i := 0; i< end; i ++ {
        tc, vc, tags, err := ParseJsonHash(dataList[i])
     
        if err != nil {
            log.Println(err)
            continue    
        }
        
        if hasTagFilter && filter(tagFilter, tags) == false {
            continue
        }
        rangeStartTime, rangeEndTime, aggResult, ret = aggfun(rangeStartTime, rangeEndTime, aggResult, tc, vc, startTimestamp, timeRange, ret)
    } 

    result["single"] = ret
    return result, nil
}

func aggreatorFactory(aggFunction string) func(rangeStartTime int64, rangeEndTime int64, aggResult float64, currentElementTime int64, currentElementValue float64, startTimestamp int64, timeRange int64, ret []map[string]interface{}) (int64, int64, float64, []map[string]interface{}) {

    aF := getFuncMap(aggFunction)
   var aggreator func (rangeStartTime int64, rangeEndTime int64, aggResult float64, currentElementTime int64, currentElementValue float64, startTimestamp int64, timeRange int64, ret []map[string]interface{}) (int64, int64, float64, []map[string]interface{})

   switch aggFunction {
       case "count", "min", "max", "sum", "avg", "std" :
           aggreator = func (rangeStartTime int64, rangeEndTime int64, aggResult float64, currentElementTime int64, currentElementValue float64, startTimestamp int64, timeRange int64, ret []map[string]interface{}) (int64, int64, float64, []map[string]interface{}) {
        
               if rangeStartTime == 0 {
                   rangeStartTime = currentElementTime - ( currentElementTime - startTimestamp ) % timeRange
                   rangeEndTime = rangeStartTime + timeRange
               } 
 
               if currentElementTime > rangeEndTime {
                   ele := make(map[string]interface{})
                   ele["timestamp"] = rangeStartTime
                   ele["value"] = aggResult
                   ret = append(ret, ele) 
                   rangeStartTime = currentElementTime - ( currentElementTime - startTimestamp ) % timeRange
                   rangeEndTime = rangeStartTime + timeRange
                   aggResult = aF(0,currentElementValue) 
                   aF = getFuncMap(aggFunction)
               } else {
                   aggResult = aF(aggResult, currentElementValue)
               }
               return rangeStartTime, rangeEndTime, aggResult, ret
           }
   
       case "raw":
           aggreator = func (rangeStartTime int64, rangeEndTime int64, aggResult float64, currentElementTime int64, currentElementValue float64, startTimestamp int64, timeRange int64, ret []map[string]interface{}) (int64, int64, float64, []map[string]interface{}) {
               aggResult = aF(0,currentElementValue) 
               ele := make(map[string]interface{})
               ele["timestamp"] = currentElementTime
               ele["value"] = aggResult
               ret = append(ret, ele) 
               aF = getFuncMap(aggFunction)
               return rangeStartTime, rangeEndTime, aggResult, ret
           }
  
   }
   return aggreator

}

