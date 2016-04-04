package KsanaDB
import(
        "testing" 
        "strings"
)

func Test_isMetricNameValidate(t *testing.T) {  
    ret := isMetricNameValidate("A\tB")
    if ret == true {
        t.Error("test isMetricNameValidate fail")    
    }

    ret = isMetricNameValidate("AB")
    if ret != true {
        t.Error("test isMetricNameValidate fail")    
    }
}

func Test_removeTab(t *testing.T) { 
    ret := removeTab("A\tB")
    if strings.Contains(ret, "\t") == true {
        t.Error("remove tab fail")    
    }
}
