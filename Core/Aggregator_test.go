package KsanaDB
import(
        "fmt"
        "testing" 
        "math"
)

func Test_getFuncMap(t *testing.T) {  
    sum := getFuncMap("sum")
    sumRet := sum(100,200)
    if sumRet != 300 {
        t.Error("sum fial")    
    }

    max := getFuncMap("max")
    maxRet := max(100,200) //100 is dummy
    if maxRet != 200 {
        fmt.Println(maxRet)
        t.Error("max fial")    
    }
    maxRet = max(100,2000) //100 is dummy
    if maxRet != 2000 {
        fmt.Println(maxRet)
        t.Error("max fial")    
    }

    //reset max and test negative data
    max = getFuncMap("max")
    maxRet = max(100,-700) //100 is dummy
    if maxRet != -700 {
        fmt.Println(maxRet)
        t.Error("max negative fial")    
    }
    maxRet = max(100,-1100) //100 is dummy
    if maxRet != -700 {
        fmt.Println(maxRet)
        t.Error("max negative fial")    
    }

    min := getFuncMap("min")
    minRet := min(100,200) //100 is dummy
    if minRet != 200 {
        fmt.Println(minRet)
        t.Error("min fial")    
    }
    minRet = min(100,50) // 100 is dummy
    if minRet != 50 {
        fmt.Println(minRet)
        t.Error("min fial")    
    }
   
    //reset test
    min = getFuncMap("min")
    minRet = min(200,100)
    if minRet != 100 {
        fmt.Println(minRet)
        t.Error("min reset fial")    
    }

    count := getFuncMap("count")
    countRet := count(0,100)
    if countRet != 1 {
        t.Error("count fial")    
    }

    firstF := getFuncMap("first")
    firstRet := firstF(100,200) // 100 is dummy
    firstRet = firstF(100,700)
    if firstRet != 200 {
        fmt.Println(firstRet)
        t.Error("first fial")    
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

    std := getFuncMap("std")
    data = []float64 {20, 20, 20, 20, 20, 20}
    answer = []float64 {math.SmallestNonzeroFloat64, 0, 0, 0, 0, 0, 0}
    stand := float64(0)
    for i, e := range(data) {
        stand = std(stand ,e)
        if stand != answer[i] {
            fmt.Println(i)
            fmt.Println(stand)
            fmt.Println(answer[i])
            t.Error("stand derivation fial")
        }
    }

    raw := getFuncMap("raw")
    rawVal := raw(100,200)
    if rawVal != 200 {
        t.Error("raw fial")    
    }
}


func Test_isTimeRangeFunction(t *testing.T) {  
    rangeFunction := []string {"sum", "max", "min", "count", "avg", "std"}
    nonRangeFunction := []string{"raw"}
    for _,d := range(rangeFunction) {
        ret, err := isTimeRangeFunction(d)
        if ret == false  {
             fmt.Println(d)
             t.Error("isTimeRangeFunction fail")    
        }
        if err != nil {
             fmt.Println(d)
             t.Error(err)    
        }
    }

    for _,d := range(nonRangeFunction) {
        ret, err := isTimeRangeFunction(d) 
        if ret  == true  {
             fmt.Println(d)
             t.Error("isTimeRangeFunction fail")    
        }
        if err != nil {
             fmt.Println(d)
             t.Error(err)    
        }
    }
   
    _, err := isTimeRangeFunction("noSuchFunction")
    if err == nil {
         t.Error("Not detect non-exist function")    
    }

}
