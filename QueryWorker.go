package KsanaDB
import(
//    "fmt"
//     "strconv" 
     "log"
//    "time"
//     "encoding/json"
)


func queryWorker(dataList []string, startTimestamp int64, tagFilter []int64, aggregateFunction string, sampleUnit string, sampleRange int) ([]map[string]interface{}, error){
    ret := []map[string]interface{}{}

    //check this and next one time
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
     
        //fmt.Println(tags)
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

