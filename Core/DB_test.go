package KsanaDB
import(
        "testing" 
        "fmt"
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
    result, err := QueryData(q)

    if err != nil {
        t.Error(err)    
    }

    fmt.Println(result)
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
