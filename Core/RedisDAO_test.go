package KsanaDB
import(
    "testing" 
     mock "github.com/rafaeljusto/redigomock"
     redis "github.com/garyburd/redigo/redis"
)

func getMock() redis.Conn {
    return mock.NewConn()    
}

func init() {
    clientFunction = getMock
}

func Test_BulkSetTimeSeries(t *testing.T) {  
    var input = []interface{}{"1", "2 "}
    BulkSetTimeSeries("test", input)
}

func Test_SetTimeSeries(t *testing.T) {  
    data := `{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"value":1.000000}`
    SetTimeSeries("test", data, 1234567890000)
}
