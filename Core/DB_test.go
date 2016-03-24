package KsanaDB
import(
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
