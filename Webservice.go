package main

import (
    "strconv"                                                                                                              
        "fmt"                                                                                                                  
            "math/rand"         
            "time"
    "github.com/zzzmanzzz/KsanaDB/Core"
    "github.com/go-martini/martini"
)


func init() {

}

func main() {
      m := martini.Classic()
KsanaDB.InitRedis("tcp", "127.0.0.1:6379")

 i:= 1
          d := fmt.Sprintf(`[{"name": "wyatt_test", "timestamp": %d, "value": %s, "tags": {"host": "server%d", "speed":"%d", "type":"tp%d"}}]`, (1389024000000 + 100  * int64(i) + rand.Int63()%10) , strconv.Itoa(i%10), i%10, rand.Int63()%100, rand.Int63()%5)

      m.Get("/", func() string {

       KsanaDB.SetData(d)
                return "Hello 世界!"
                })
      m.RunOnAddr(fmt.Sprintf(":%d", 3000 + time.Now().Unix()%100))
     // m.Run()
}
