package KsanaDB

import (
    redis "github.com/garyburd/redigo/redis"
    "log"
    "fmt"
    "time"
    "errors"
)

var pool *redis.Pool 
var maxPipeline int

func InitRedis(network, address string)  {
    maxPipeline = 8000 // too many  pipeline will get fewer data
    pool = &redis.Pool{                                                                                                    
        MaxIdle:     80,
        MaxActive: 12000,
        IdleTimeout: 240 * time.Second,                          
        Dial: func() (redis.Conn, error) {                                                                        
            c, err := redis.Dial(network, address)                                                                   
                if err != nil { 
                    log.Fatalf(err.Error())
                    return nil, err                                                                               
                }                                                                                                 
                return c, err                                                                                         
        },                                                                                                        
        TestOnBorrow: func(c redis.Conn, t time.Time) error { 
                  _, err := c.Do("PING") 
                  if err != nil { 
                      log.Fatalf(err.Error())
                      return err 
                  }
                  return nil 
        },                                                                                                           
    } 
}  

func BulkSetTimeSeries(metrics string, input []interface{}) (int, error) {
    client := pool.Get()
    defer client.Close()
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}

func SetTimeSeries(metrics string, value string, time int64) (int, error) {
    client := pool.Get()
    defer client.Close()
    input := []interface{}{}
    input = append(input,time)
    input = append(input,value)
    return redis.Int(client.Do("ZADD", redis.Args{metrics}.AddFlat(input)...))
}

func queryTimeSeries(prefix string, name string, start int64, stop int64) ([]string, error) {
    client := pool.Get()
    defer client.Close()
    cmds := getTimeseriesQueryCmd(prefix, name, start, stop)

    if len(cmds) > maxPipeline {
        return []string{}, errors.New(fmt.Sprintf("time %d - %d over upper limit duration days %d", start, stop, maxPipeline))    
    }

    for _, cmd := range cmds {

        client.Send("ZRANGEBYSCORE", redis.Args{cmd["keyName"], cmd["from"], cmd["to"]}...)
    }
    client.Flush()
    
    ret := []string{}
    for _,_ = range cmds {
        p, err := redis.Strings(client.Receive())
        if err != nil || len(p) == 0{
            continue
        }

        ret = append(ret, p...)
    }
    return ret, nil
}


func setTags(prefix string, metrics string, tags []string) (string) {
    //TODO: call function
    client := pool.Get()
    defer client.Close()
    hashName := prefix + metrics + "\tTagHash"
    listName := prefix + metrics + "\tTagList"

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

func getTags(prefix string, metrics string, target string, keyName string) (string) {
    client := pool.Get()
    defer client.Close()
    listName := prefix + metrics + "\tTagList"
    s := getLuaScript("getTag")
    script := redis.NewScript(0, s)

    result, err := redis.String(script.Do(client, listName, target, keyName))
    if err != nil {
        log.Println(err)    
    }
    return result
} 

func getSeqByKV(prefix string, metrics string, filterKeyValue []string) ([]string, error) {
    client := pool.Get()
    defer client.Close()
    hashName := prefix + metrics + "\tTagHash"
    return redis.Strings(client.Do("HMGET", redis.Args{hashName}.AddFlat(filterKeyValue)...))
}

func Close() {
    pool.Close()
}
