package KsanaDB
import(
    "fmt"
//     "strconv" 
     "log"
     "strings"
//    "time"
//     "encoding/json"
)

func queryWorker(dataList []string, startTimestamp int64, tagFilter []string, groupByTag map[string][]string, aggregateFunction string, sampleUnit string, sampleRange int) ([]map[string]interface{}, error){

    ret := []map[string]interface{}{}
    hasGroupBy := len(groupByTag) > 0

    if hasGroupBy == false {
        ret, err := nonConcurrentQuery(dataList, startTimestamp, tagFilter, aggregateFunction, sampleUnit, sampleRange)
        return ret, err
    } else {
        if hasGroupBy == true  {
            concurrentQuery(dataList, startTimestamp, tagFilter, groupByTag, aggregateFunction, sampleUnit, sampleRange)
        }
    }
    return ret, nil
}

func concurrentPart(key string, startTimestamp int64, timeRange int64, aggregateFunction string, in chan map[string]interface{}, out chan map[string][]map[string]interface{}) {
    aggResult := float64(0)
    rangeStartTime := int64(0)
    rangeEndTime := int64(0)
    ret := map[string][]map[string]interface{}{}

    aF := getFuncMap(aggregateFunction)

    for {
        d, more := <-in
        if more == false {
            break    
        }
        tc := d["timestamp"].(int64)
        vc := d["value"].(float64)

        if rangeStartTime == 0 {
            rangeStartTime = tc - ( tc - startTimestamp ) % timeRange
            rangeEndTime = rangeStartTime + timeRange
            //fmt.Printf("###\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
            //vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)
        } 
            //fmt.Printf("###\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
            //vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)

        if tc > rangeEndTime {
            ele := make(map[string]interface{})
            ele["timestamp"] = rangeStartTime
            ele["value"] = aggResult
            ret[key] = append(ret[key], ele) 
            rangeStartTime = tc - ( tc - startTimestamp ) % timeRange
            rangeEndTime = rangeStartTime + timeRange
            aggResult = aF(0,vc) 
        //fmt.Printf("@@@\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
        //    vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)
        } else {
            aggResult = aF(aggResult, vc)
        }

    }
    out <- ret
}

func concurrentQuery(dataList []string, startTimestamp int64, tagFilter []string, groupByTag map[string][]string, aggregateFunction string, sampleUnit string, sampleRange int) ([]map[string]interface{}, error){
    ret := []map[string]interface{}{}
    end := len(dataList)  

    timeRange, err := getTimeRange(startTimestamp, sampleRange, sampleUnit )

    if err != nil {
         return nil, err    
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
        fmt.Println(d)
        if counter == len(groupByInputChan) {
            close(out)
            break
        }

    }
    return ret, nil
}

func nonConcurrentQuery(dataList []string, startTimestamp int64, tagFilter []string, aggregateFunction string, sampleUnit string, sampleRange int) ([]map[string]interface{}, error){
    
    ret := []map[string]interface{}{}
    end := len(dataList)  

    timeRange, err := getTimeRange(startTimestamp, sampleRange, sampleUnit )

    if err != nil {
         return nil, err    
    }

    aggResult := float64(0)
    rangeStartTime := int64(0)
    rangeEndTime := int64(0)

    aF := getFuncMap(aggregateFunction)

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

        if rangeStartTime == 0 {
            rangeStartTime = tc - ( tc - startTimestamp ) % timeRange
            rangeEndTime = rangeStartTime + timeRange
//        fmt.Printf("###\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
//            vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)
        } 
 
//        fmt.Printf("###\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
//            vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)

        if tc > rangeEndTime {
            ele := make(map[string]interface{})

            ele["timestamp"] = rangeStartTime
            ele["value"] = aggResult
            if hasTagFilter == true {
                ele["tags"] = tagFilter
            }
            ret = append(ret, ele) 
            rangeStartTime = tc - ( tc - startTimestamp ) % timeRange
            rangeEndTime = rangeStartTime + timeRange
            aggResult = aF(0,vc) 
//        fmt.Printf("@@@\nvc   \t%f\nstart\t%d\ntc   \t%d\nend  \t%d\ntime Range %d\nstartTimestamp  %d\n", 
//            vc, rangeStartTime, tc, rangeEndTime, timeRange,startTimestamp)
        } else {
            aggResult = aF(aggResult, vc)
        }
    } 
    return ret, nil
}

