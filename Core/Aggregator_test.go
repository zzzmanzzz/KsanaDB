package KsanaDB
import(
        "testing" 
)

func Test_getFuncMap(t *testing.T) {  
    sum := getFuncMap("sum")
    sumRet := sum(100,200)
    if sumRet != 300 {
        t.Error("sum fial")    
    }

    max := getFuncMap("max")
    maxRet := max(100,200)
    if maxRet != 200 {
        t.Error("max fial")    
    }

    min := getFuncMap("min")
    minRet := min(100,200)
    if minRet != 100 {
        t.Error("min fial")    
    }

}
