package ChronosDB

import (
    redis "github.com/garyburd/redigo/redis"
    "log"
    "fmt"
)

var MAX_POOL_SIZE = 20
var redisPoll chan redis.Conn
var client redis.Conn

func putRedis(conn redis.Conn) {
    if redisPoll == nil {
        redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
    }
    if len(redisPoll) >= MAX_POOL_SIZE {
        conn.Close()
            return
    }
    redisPoll <- conn
}

func InitRedis(network, address string) redis.Conn {
    if len(redisPoll) == 0 {
        redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
            go func() {
                for i := 0; i < MAX_POOL_SIZE/2; i++ {
                    c, err := redis.Dial(network, address)
                        if err != nil {
                            panic(err)
                        }
                    putRedis(c)
                }
            } ()
    }
    return <-redisPoll
}



func GetLink(host string, port uint) {
     client = InitRedis("tcp",host+":"+fmt.Sprint(port))
}

func BulkSetTimeSeries(metrics string, input []interface{}, tags []string) (int, error) {
    log.Printf("metrics : %s\n", metrics)
    log.Println(input)
    log.Println(tags)
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}

func SetTimeSeries(metrics string, value string, time int64, tags []string) (int, error) {
    log.Printf("metrics : %s, value : %s, time offset : %d, tag : %s\n", metrics, value, time, tags)
    input := []interface{}{}
    input = append(input,time)
    input = append(input,value)
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}
/*
func QueryTimeSeries() {
return redis.Strings(s.Do("ZREVRANGEBYSCORE", redigo.Args{key, start, stop}.AddFlat(options)...))    
}
*/
/*
func SetTagName(metrics string, key string, tag string) {
    log.Printf("SET hello 1\n")
    client.Set("hello", 1)
} 
*/
/*
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
*/

func Close() {
    client.Close()   
}
