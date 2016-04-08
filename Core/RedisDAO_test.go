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

    var retHMGET []interface{}
    d := []byte("test")
    retHMGET = append(retHMGET, d)
    c.Command("HMGET").Expect(retHMGET)
    return c    
}

func init() {
    clientFunction = getMock
}
func Test_init(t *testing.T) { 
    InitRedis("tcp", "127.0.0.1:1234")
}

func Test_BulkSetTimeSeries(t *testing.T) {  
    var input = []interface{}{`{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[1458790110000,0],[1458790110001,1],[1458790110002,2]]}`,`{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458790110003,0],[1458790110103,1],[1458790110203,2]]}`}
    BulkSetTimeSeries("test", input)
}

func Test_SetTimeSeries(t *testing.T) {  
    data := `{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"value":1.000000}`
    SetTimeSeries("test", data, 1234567890000)
}

func Test_getMetric(t *testing.T) {
    prefix = "KSANADBv1\t"
    getMetric(prefix)    
}

func Test_getMetricKeys(t *testing.T) {
    prefix = "KSANADBv1\t"
    getMetricKeys(prefix, "wyatt_test")    
}

func Test_deleteKeys(t *testing.T) {
    data := []string{
             "KSANADBv1\twyatt_new\tTagList",
             "KSANADBv1\twyatt_new\t1459555200000",
             "KSANADBv1\twyatt_new\tTagHash",
        }
    deleteKeys(data)    
}
