package ChronosDB
import "menteslibres.net/gosexy/redis"
import "log"

var client *redis.Client


func GetLink(host string, port uint) {

    client = redis.New()
    err := client.Connect(host, port)
    if err != nil {
        log.Fatalf("Connect failed: %s\n", err.Error())
        return
    }
}

func SetData(key string, value string, time uint32) {
    key = key
    value = value
    time = time
    log.Printf("SET hello 1\n")
    client.Set("hello", 1)
} 

func Close() {
    client.Quit()    
}
