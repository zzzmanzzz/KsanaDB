package ChronosDB

import (
    "menteslibres.net/gosexy/redis"
    "log"
)
var client *redis.Client

func GetLink(host string, port uint) {
    if client == nil {
        client = redis.New()
        err := client.Connect(host, port)
        if err != nil {
            log.Fatalf("Connect failed: %s\n", err.Error())
                return
        }
    }
}

func BulkSetTimeSeries(metrics string, input []interface{}, tags []string) (int64, error) {
    log.Printf("metrics : %s\n", metrics)
    log.Println(input)
    log.Println(tags)
    ret, err := client.ZAdd(metrics,  input...)
    return ret, err
}

func SetTimeSeries(metrics string, value string, time int64, tags []string) (int64, error) {
    log.Printf("metrics : %s, value : %s, time offset : %d, tag : %s\n", metrics, value, time, tags)
    input := []interface{}{}
    input = append(input,time)
    input = append(input,value)
    ret, err := client.ZAdd(metrics,  input...)
    return ret, err
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
