package KsanaDB
import(
        "fmt"
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

    max = getFuncMap("max")
    maxRet = max(200,100)
    if maxRet != 200 {
        t.Error("max fial")    
    }

    min := getFuncMap("min")
    minRet := min(100,200)
    if minRet != 100 {
        t.Error("min fial")    
    }

    min = getFuncMap("min")
    minRet = min(200,100)
    if minRet != 100 {
        t.Error("min fial")    
    }

    count := getFuncMap("count")
    countRet := count(0,100)
    if countRet != 1 {
        t.Error("count fial")    
    }

    avg := getFuncMap("avg")
    data := []float64 {10, 20, 30, 40, 50, 60}
    answer := []float64 {10, 15, 20, 25, 30, 35}
    average := float64(0)
    for i, e := range(data) {
        average = avg(average ,e)
        if average != answer[i] {
            t.Error("average fial")
            fmt.Println(average)
            fmt.Println(answer[i])
        }
    }


}
