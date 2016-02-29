package KsanaDB

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

func BulkSetTimeSeries(metrics string, input []interface{}) (int, error) {
  //  log.Printf("metrics : %s\n", metrics)
  //  log.Println(input)
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}

func SetTimeSeries(metrics string, value string, time int64) (int, error) {
  //  log.Printf("metrics : %s, value : %s, time offset : %d\n", metrics, value, time)
  //  log.Println(value)
    input := []interface{}{}
    input = append(input,time)
    input = append(input,value)
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}

func queryTimeSeries(prefix string, name string, start int64, stop int64) ([]string) {
    //options := ""//"withscores"
    cmds := getTimeseriesQueryCmd(prefix, name, start, stop)
    for _, cmd := range cmds {
    //.AddFlat(options)
        client.Send("ZRANGEBYSCORE", redis.Args{cmd["keyName"], cmd["from"], cmd["to"]}...)
    }
    client.Flush()
    
    ret := []string{}
    for _,_ = range cmds {
        p, err := redis.Strings(client.Receive())
        if err != nil {
            continue
        }
        ret = append(ret, p...)
    }
    return ret
}


func setTags(prefix string, metrics string, tags []string) (string) {
    hashName := prefix + metrics + "\tTagHash"
    listName := prefix + metrics + "\tTagList"

    // args = eval key + eval args
    //here is keys 
    args := []string{}

    args = append(args, tags...)
    args = append(args, hashName)
    args = append(args, listName)

    scriptArgs := make([]interface{}, len(args))
    for i, v := range args {
            scriptArgs[i] = v
    }

    s := getLuaScript("setTag")
    script := redis.NewScript(len(tags), s)

    result, err := redis.String(script.Do(client, scriptArgs...))
    if err != nil {
        log.Println(err)    
    }
    return result
} 

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
