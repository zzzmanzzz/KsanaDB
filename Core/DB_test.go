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
