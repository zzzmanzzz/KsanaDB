package KsanaDB
import(
        "testing" 
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
