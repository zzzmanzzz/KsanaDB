package main

import (
    //"strconv"           
    "fmt" 
    //"math/rand"         
//    "time"
    "net/http"
    "github.com/zzzmanzzz/KsanaDB/Core"
    "github.com/go-martini/martini"
)

func init() {

}

func main() {
      m := martini.Classic()
      KsanaDB.InitRedis("tcp", "127.0.0.1:6379")

      m.Post("/api/v1/datapoints", func(w http.ResponseWriter, r *http.Request) {
           var err *error
           err = KsanaDB.SetData(r.FormValue("data"))
           if err == nil {
               w.WriteHeader(http.StatusOK)
           } else {
               w.WriteHeader(400)
           }
      })

      m.Post("/api/v1/metricnames", func(w http.ResponseWriter, r *http.Request) {
           fmt.Println(r.FormValue("data"))
           ret, err := KsanaDB.GetMetricsTag(r.FormValue("data"), "TagKey", "")
           if err == nil {
               w.Write([]byte(ret))
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

      m.RunOnAddr(fmt.Sprintf(":%d", 13000 ))//+ time.Now().Unix()%100))
}
