package KsanaDB  

type aggFunc func(float64, float64) float64


var fnRegistry = map[string] interface{} {
    "sum": func(sum float64, val float64) float64 { return sum + val},
}

func getFunc(funName string) interface{} {
    return fnRegistry[funName]
}

