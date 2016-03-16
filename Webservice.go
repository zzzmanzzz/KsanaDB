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

      m.Post("/api/v1/query", func(w http.ResponseWriter, r *http.Request) {
           q, err := KsanaDB.ParseQueryJson(r.FormValue("data"))

           if err != nil {
               w.WriteHeader(400)
               return
           }

           ret, err := KsanaDB.QueryData(q)
           if err == nil {
               for k, v := range(ret) {
               fmt.Print(k)
               fmt.Println(v)
               }
               w.WriteHeader(http.StatusOK)
           } else {
               fmt.Println(err)
               w.WriteHeader(400)
           }

      })

      m.RunOnAddr(fmt.Sprintf(":%d", 13000 ))//+ time.Now().Unix()%100))
}
