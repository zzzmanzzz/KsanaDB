package main

import (
    "strconv"
    "fmt" 
    "os"
    "bytes"
    "net/http"
    "github.com/zzzmanzzz/KsanaDB/Core"
    "github.com/go-martini/martini"
    "gopkg.in/yaml.v2"
)

var redisServer string
var redisPort string
var ksanaDBPort int 

type Conf struct {
    RedisServer string `yaml:"RedisServer"`
    RedisPort int `yaml:"RedisPort"`
    KsanaDBPort int `yaml:"KsanaDBPort"`
}


func init() {
    //default
    redisServer = "127.0.0.1"
    redisPort = "6379"
    ksanaDBPort = 13000

    file, err := os.Open("./KsanaDb.conf")
    if err != nil {
        fmt.Println(err) 
        fmt.Println("Open config KsanaDB.conf fail")
        fmt.Println("Use default Config")
        printConfig()
        return
    }

    data := make([]byte, 100)
    _, err = file.Read(data)
    if err != nil {
        fmt.Println(err) 
        fmt.Println("Read config KsanaDB.conf fail")
        fmt.Println("Use default Config")
        printConfig()
        return
    }

    c := Conf{}
    err = yaml.Unmarshal( bytes.Trim(data, "\x00"), &c)
    if err != nil {
        fmt.Println(err) 
        fmt.Println("Parse config KsanaDB.conf fail")
        fmt.Println("Use default Config")
        printConfig()
        return
    }

    redisServer = c.RedisServer
    redisPort = strconv.Itoa(c.RedisPort)
    ksanaDBPort = c.KsanaDBPort
    printConfig()
}

func printConfig () {
    fmt.Printf("redis server ip : %s\nredis port : %s\nKsanaDB port : %d\n" , redisServer, redisPort, ksanaDBPort) 
}

func main() {
      m := martini.Classic()
      KsanaDB.InitRedis("tcp", redisServer + ":" + redisPort )

      m.Post("/api/v1/datapoints", func(w http.ResponseWriter, r *http.Request) {
           var err *error
           err = KsanaDB.SetData(r.FormValue("data"))
           if err == nil {
               w.WriteHeader(http.StatusOK)
           } else {
               w.WriteHeader(400)
           }
      })

      m.Post("/api/v1/tagvalues", func(w http.ResponseWriter, r *http.Request) {
           ret, err := KsanaDB.GetMetricsTag(r.FormValue("metric"), "TagValue", r.FormValue("tag"))
           if err == nil {
               w.Write([]byte(ret))
           } else {
               w.WriteHeader(400)
           }
      })

      m.Post("/api/v1/tagnames", func(w http.ResponseWriter, r *http.Request) {
           ret, err := KsanaDB.GetMetricsTag(r.FormValue("metric"), "TagKey", "")
           if err == nil {
               w.Write([]byte(ret))
           } else {
               w.WriteHeader(400)
           }
      })

      m.Post("/api/v1/metricnames", func(w http.ResponseWriter, r *http.Request) {
           ret, err := KsanaDB.GetMetric()
           if err == nil {
               w.Write([]byte(ret))
           } else {
               w.WriteHeader(400)
           }
      })

      m.Delete("/api/v1/metric", func(w http.ResponseWriter, r *http.Request) {
           ret, err := KsanaDB.DeleteMetric(r.FormValue("metric"))
           if err == nil {
               w.Write([]byte(ret))
           } else {
               w.WriteHeader(400)
           }
      })

      m.Post("/api/v1/query", func(w http.ResponseWriter, r *http.Request) {
           q, err := KsanaDB.ParseQueryJson(r.FormValue("data"))

           if err != nil {
               w.WriteHeader(400)
               return
           }

           ret, err := KsanaDB.QueryData(q)
           if err == nil {
               //w.WriteHeader(http.StatusOK)
               w.Write([]byte(ret))
           } else {
               fmt.Println(err)
               w.WriteHeader(400)
               w.Write([]byte(err.Error()))
           }

      })

      m.RunOnAddr(fmt.Sprintf(":%d", ksanaDBPort ))
}
