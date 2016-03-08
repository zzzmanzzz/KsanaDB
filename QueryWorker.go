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
    hasGroupBy := len(groupByTag) > 0

    var groupByAggreateIndex map[string][]int
    if hasGroupBy == true  {
         groupByAggreateIndex = map[string][]int{}
    }

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

        if hasGroupBy == true {
            //unique and sort hit group tags and filter tags
             gkey := groupBy(groupByTag, tags)
             if len(gkey) > 0 {
                 key := strings.Join(gkey,"\t")
                 groupByAggreateIndex[key] = append(groupByAggreateIndex[key], i)
             } else {
                 continue    
             }
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
    fmt.Println(groupByAggreateIndex)
    return ret, nil
}

