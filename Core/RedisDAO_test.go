package KsanaDB
import(
    "testing" 
     mock "github.com/rafaeljusto/redigomock"
     redis "github.com/garyburd/redigo/redis"
)

func getMock() redis.Conn {
    c := mock.NewConn()
   
    c.Command("EVALSHA").Expect("ok")
    c.Command("ZADD").Expect("ok")

    return c    
}

func init() {
    clientFunction = getMock
}

func Test_BulkSetTimeSeries(t *testing.T) {  
    var input = []interface{}{`{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[1458790110000,0],[1458790110001,1],[1458790110002,2]]}`,`{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458790110003,0],[1458790110103,1],[1458790110203,2]]}`}
    BulkSetTimeSeries("test", input)
}

func Test_SetTimeSeries(t *testing.T) {  
    data := `{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"value":1.000000}`
    SetTimeSeries("test", data, 1234567890000)
}
