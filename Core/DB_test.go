package KsanaDB
import(
        "encoding/json"
        "testing" 
)

func init() { 
    clientFunction = getMock                                                                                               
}

func Test_SetSingleData(t *testing.T) {  
    data := `[{"name": "wyatt_test", "timestamp": 1234567890, "value": 1.2, "tags": {"host": "server1", "speed":"55", "type":"tp0"}}]`
    err := SetData(data)

    if err != nil {
        t.Error(err)    
    }
}

func Test_SetDataPoints(t *testing.T) {  
    data := `[{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[1458791287000,0],[1458791287001,1],[1458791287002,2]]},{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458791287003,0],[1458791287103,1],[1458791287203,2]]}]`
    err := SetData(data)

    if err != nil {
        t.Error(err)    
    }
}

func Test_SetSingleDataNameEmpty(t *testing.T) {  
    data := `[{"name": "", "timestamp": 1234567890, "value": 1.2, "tags": {"host": "server1", "speed":"55", "type":"tp0"}}]`
    err := SetData(data)

    if err != nil {
        t.Error(err)    
    }
}

func Test_SetDataPointsTimestampFail(t *testing.T) {  
    data := `[{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[["ZZZ",0],[1458791287001,1],[1458791287002,2]]},{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458791287003,0],[1458791287103,1],[1458791287203,2]]}]`
    err := SetData(data)

    if err != nil {
        t.Error(err)    
    }
}

func Test_SetDataPointsJsonFial(t *testing.T) { 
    data := `[{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[xx1458791287000,0],[1458791287001,1],[1458791287002,2]]},{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458791287003,0],[1458791287103,1],[1458791287203,2]]}]`
    err := SetData(data) 

    if err == nil {
        t.Error("SetDataPointsJsonFial fial")    
    }
}

func Test_QueryOverMaxPipeline(t *testing.T) {  
    data := `{"startabsolute":1389024000000,"endabsolute":1389096010500,"metric":{"aggregator":{"name":"sum","sampling":{"unit":"h","value":1}},"tags":null,"name":"wyatt_test"}}`

    maxPipeline = 0
    q, err := ParseQueryJson(data)
    if err != nil {
        t.Error(err)    
    }
    _ , err = QueryData(q)
    
    if err == nil {
        t.Error(err)    
    }
}

func Test_QueryReturnNothing(t *testing.T) {  
    maxPipeline = 8000
    data := `{"startabsolute":1389024000000,"endabsolute":1389096010500,"metric":{"aggregator":{"name":"sum","sampling":{"unit":"h","value":1}},"tags":null,"name":"wyatt_test"}}`

    q, err := ParseQueryJson(data)
    if err != nil {
        t.Error(err)    
    }
    _, err = QueryData(q)

    if err != nil {
        t.Error(err)    
    }
}

func Test_QueryStartTimeFail(t *testing.T) {  
    maxPipeline = 8000
    data := `{"endabsolute":1389096010500,"metric":{"aggregator":{"name":"sum","sampling":{"unit":"h","value":1}},"tags":null,"name":"wyatt_test"}}`

    q, err := ParseQueryJson(data)
    if err != nil {
        t.Error(err)    
    }
    _, err = QueryData(q)

    if err == nil {
        t.Error("Not detect start time not input")    
    }

}

func Test_GetMetricsTag(t *testing.T) {
    name := "wyatt_test"
    target := "All"
    keyName := ""
    _, err := GetMetricsTag(name, target, keyName)   
    if err != nil {
        t.Error(err)    
    }
}

func Test_GetMetricsTag_WrongTargetName(t *testing.T) {
    name := "test"
    target := "all"
    keyName := ""
    _, err := GetMetricsTag(name, target, keyName)   
    if err == nil {
        t.Error(err)    
    }
}

func Test_GetMetricsTagSeq(t *testing.T) {
    name := "test"
    keyName := "speed"
    _ = GetMetricsTagSeq(name, keyName)  
}

func Test_GetFilterSeq_EmptyFilterList(t *testing.T) {
    name := "test"
    filterList := []string{}
    ret, err := GetFilterSeq(name, filterList) 
    if len(ret) != 0 {
        t.Error(err)    
    }
}

func Test_GetFilterSeq(t *testing.T) {
    name := "test"
    filterList := []string{"speed"}
    _, err := GetFilterSeq(name, filterList) 
    if err != nil {
        t.Error(err)    
    }
}


func Test_GetMetric(t *testing.T) {
    GetMetric()
}

func Test_deleteMetric(t *testing.T) {
    _, err := DeleteMetric("test")
    if err != nil && err.Error() == "Metric name can't content /t" {
        t.Error(err)    
    }
}

func Test_deleteMetricInvalidateName(t *testing.T) {
    _, err := DeleteMetric("test\t")
    if err == nil {
        t.Error("deleteMetricInvalidateName test fail")    
    }
}

func Test_generateOutputData(t *testing.T) {
    var result map[string][]map[string]interface{}
    json.Unmarshal([]byte(queryResultJ), &result) 

    var reverseHash map[string]string
    json.Unmarshal([]byte(reverseHashJ), &reverseHash) 

    var filterTag []string
    json.Unmarshal([]byte(filterTagJ), &filterTag) 

    var groupBy []string
    json.Unmarshal([]byte(groupByJ), &groupBy) 
 
    name := "wyatt_new"

    start := int64(1459900800000)
    end := int64(1459972810500)
 
    aggregationFunction := "sum"

    timeRange := 1
    unit := "m"

    _, err := generateOutputData(result, reverseHash, name, start, end, filterTag, groupBy, aggregationFunction, timeRange, unit)
    if err != nil {
        t.Error(err)    
    }

}

var queryResultJ = `
{"4\t5\t6":[{"timestamp":1459920000000,"value":24090},{"timestamp":1459920060000,"value":311700},{"timestamp":1459920120000,"value":671700},{"timestamp":1459920180000,"value":1.0317e+06},{"timestamp":1459920240000,"value":1.3917e+06}]}
`
var reverseHashJ = `
{"1":"host\tserver1","2":"speed\t10","3":"type\ttp2","4":"host\tserver11","5":"speed\t11","6":"type\ttp1"}
`
var filterTagJ = `
["type\ttp1","speed\t11"]
`

var groupByJ = `
["host","type","speed"]
`

var generateAnswer = `{"Name":"wyatt_new","Start":1459900800000,"End":1459972810500,"Filter":{"speed":"11","type":"tp1"},"AggregateFunction":"sum","TimeRange":1,"TimeUnit":"m","GroupBy":["host","type","speed"],"Group":[{"Tags":{"host":"server11","speed":"11","type":"tp1"},"Values":[[1459920000000,24090],[1459920060000,311700],[1459920120000,671700],[1459920180000,1.0317e+06],[1459920240000,1.3917e+06]]}]}`
