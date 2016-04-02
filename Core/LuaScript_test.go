package KsanaDB
import(
        "testing" 
)

func Test_getTag(t *testing.T) {  
    setTag := getLuaScript("setTag")
    getTag := getLuaScript("getTag")
    getMetric := getLuaScript("getMetric")
    getMetricKeys := getLuaScript("getMetricKeys")
    noTag := getLuaScript("noTag")

    if setTag == "" {
        t.Error("get setTag script fail")    
    }

    if getTag == "" {
        t.Error("get getTag script fail")    
    }

    if getMetric == "" {
        t.Error("get getMetric script fail")    
    }

    if getMetricKeys == "" {
        t.Error("get getMetricKeys script fail")    
    }

    if noTag != "" {
        t.Error("default no script fail")    
    }

}
