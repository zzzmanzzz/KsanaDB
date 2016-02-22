package ChronosDB

import (
    "github.com/go-martini/martini"
   // "github.com/martini-contrib/render"
)

func main() {
      m := martini.Classic()
        m.Get("/", func() string {
                return "Hello 世界!"
                  })
          m.Run()
}
