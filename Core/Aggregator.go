package KsanaDB  
import(
    "math"
    "errors"
)

type aggFunc func(float64, float64, ...interface{}) float64


var fnRegistry = map[string] interface{} {
    "sum": func(sum float64, val float64, others ...interface{}) float64 { 
             return sum + val
         },
    "max": func() func(dummy float64, val float64, others ...interface{}) float64 { 
            max := math.SmallestNonzeroFloat64
            return func(dummy float64, val float64, others ...interface{}) float64 {
                if max < val { 
                   max = val
                   return val 
                } else { 
                   return max
                }
            }
         },
    "min": func() func(dummy float64, val float64, others ...interface{}) float64 { 
            min := math.MaxFloat64 
            return func(dummy float64, val float64, others ...interface{}) float64 {
                if min > val { 
                    min = val
                    return val 
                } else { 
                    return min
                }
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
    "std": func() func(dummy float64, val float64, others ...interface{}) float64 {
            i := 0
            mean := float64(0)
            m2 := float64(0)

            // Welford's algorithm
            return func(dummy float64, val float64, others ...interface{}) float64 { 
                i = i + 1
                delta := val - mean
                mean = mean + delta / float64(i)
                m2 = m2 + delta * (val - mean)
                if i < 2 {
                    return math.SmallestNonzeroFloat64
                } 
                return math.Sqrt(m2/float64((i-1)))
            }
         },
    "first": func() func(dummy float64, val float64, others ...interface{}) float64 { 
            var first *float64
            return func(dummy float64, val float64, others ...interface{}) float64 {
                if first == nil {
                    first = &val
                }
                return *first
            }
         },
    "raw": func(dummy float64, val float64, others ...interface{}) float64 { 
             return val
         },
}

func getFuncMap(funName string) func(float64, float64, ...interface{}) float64 {
    var aggf aggFunc
    switch funName {
        case "sum":
            aggf = fnRegistry["sum"].(func(float64, float64, ...interface{}) float64)
        case "max":
            f := fnRegistry["max"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
        case "min":
            f := fnRegistry["min"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
        case "count":
            aggf = fnRegistry["count"].(func(float64, float64, ...interface{}) float64)
        case "avg":
            f := fnRegistry["avg"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
        case "std":
            f := fnRegistry["std"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
        case "first":
            f := fnRegistry["first"].(func() func(float64, float64, ...interface{}) float64)
            aggf = f()
        case "raw", "last":
            aggf = fnRegistry["raw"].(func(float64, float64, ...interface{}) float64)
    }
    return aggf
}

func isTimeRangeFunction(f string) (bool, error) {
    ret := false
    var err error
    switch f {
        case "sum", "max", "min", "count", "avg", "std", "first", "last":
            ret = true
        case "raw":
            ret = false
        default:
            err = errors.New("no such aggreate function")
    }
    return ret, err
}

