package main

import (
    "github.com/zzzmanzzz/KsanaDB/Core"
    "github.com/go-martini/martini"
)


func init() {

}

func main() {
      m := martini.Classic()

      KsanaDB.GetLink("127.0.0.1", 6379)
      KsanaDB.Close()  

      m.Get("/", func() string {
                return "Hello 世界!"
                })
      m.Run()
}
