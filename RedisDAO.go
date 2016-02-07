package ChronosDB
import (
    "crypto/sha1"
    "menteslibres.net/gosexy/redis"
    "log"
)
var client *redis.Client
var hasher = sha1.New()


func GetLink(host string, port uint) {

    client = redis.New()
    err := client.Connect(host, port)
    if err != nil {
        log.Fatalf("Connect failed: %s\n", err.Error())
        return
    }
}

func SetTimeSeries(metrics string, key []string, value []string, tag []string, time uint32) {
    log.Printf("metrics : %s, key : %s, value : %s, tag : %s, time : %d \n", metrics, key, value, tag, time)
    client.Set("hello", 1)
}

/*
func SetTagName(metrics string, key string, tag string) {
    log.Printf("SET hello 1\n")
    client.Set("hello", 1)
} 
*/

func SetTagID(metrics string, tag string, tagID uint64) (int64, error) {
    Mkey := metrics + "_TagID"
    ID, err := client.HLen(Mkey)
    if err != nil {
        log.Fatalf("HLen failed: %s\n", err.Error())
        return -1, err
    }
    ID += 1
    client.HSetNX(Mkey, tag, ID)
    return ID, nil
} 


func Close() {
    client.Quit()    
}
