package KsanaDB  

type aggFunc func(float64, float64, ...interface{}) float64


var fnRegistry = map[string] interface{} {
    "sum": func(sum float64, val float64, others ...interface{}) float64 { 
             return sum + val
         },
    "max": func(max float64, val float64, others ...interface{}) float64 { 
             if max < val { 
                 return val 
             } else { 
                 return max
             }
         },
    "min": func(min float64, val float64, others ...interface{}) float64 { 
             if min > val { 
                 return val 
             } else { 
                 return min
             }
         },
    "count": func(sum float64, val float64, others ...interface{}) float64 { 
             return sum + 1
         },
    "avg": func() func(dummy float64, val float64, others ...interface{}) float64 {
            i := 0
            sum := float64(0)
            return func(dummy float64, val float64, others ...interface{}) float64 {
                  i = i + 1
                  sum = sum + val
                  return sum/float64(i)
              }
         },
}

func getFuncMap(funName string) func(float64, float64, ...interface{}) float64 {
    var aggf aggFunc
    switch funName {
        case "sum":
            aggf = fnRegistry["sum"].(func(float64, float64, ...interface{}) float64)
        case "max":
            aggf = fnRegistry["max"].(func(float64, float64, ...interface{}) float64)
        case "min":
            aggf = fnRegistry["min"].(func(float64, float64, ...interface{}) float64)
        case "count":
            aggf = fnRegistry["count"].(func(float64, float64, ...interface{}) float64)
        case "avg":
            f := fnRegistry["avg"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
    }
    return aggf
}

