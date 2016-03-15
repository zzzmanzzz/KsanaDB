package KsanaDB  

type aggFunc func(float64, float64) float64


var fnRegistry = map[string] interface{} {
    "sum": func(sum float64, val float64) float64 { return sum + val},
    "max": func(max float64, val float64) float64 { if max < val { 
                                                        return val 
                                                    } else { 
                                                        return max
                                                    }
                                                  },
    "min": func(min float64, val float64) float64 { if min > val { 
                                                        return val 
                                                    } else { 
                                                        return min
                                                    }
                                                  },
}

func getFuncMap(funName string) aggFunc {
    var aggf aggFunc
    switch funName {
        case "sum":
            aggf = fnRegistry["sum"].(func(float64,float64)float64)
        case "max":
            aggf = fnRegistry["max"].(func(float64,float64)float64)
        case "min":
            aggf = fnRegistry["min"].(func(float64,float64)float64)
    }
    return aggf
}

