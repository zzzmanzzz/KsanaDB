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
    var input = []interface{}{"1", "2"}
    BulkSetTimeSeries("test", input)
}

func Test_SetTimeSeries(t *testing.T) {  
    SetTimeSeries("test", "1.0", 123456789)
}
