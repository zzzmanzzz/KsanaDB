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

func Test_CheckGeoDataFail(t *testing.T) {  
     shouldFalse := checkGeoData(nil) 
     if shouldFalse == true {
        t.Error("check input nil fail")    
     }

    var g = GeoData{ 
        Longitude: 121.5560405,
        Latitude: 24.9997971,
    }

     shouldFalse = checkGeoData(&g) 
     if shouldFalse == true {
        t.Error("check input  name nil fail")    
     }
}

func Test_CheckGeoDataSuccess(t *testing.T) {  

    var g = GeoData{ 
        Name : "test",
        Longitude: 121.5560405,
        Latitude: 24.9997971,
    }

     shouldTrue := checkGeoData(&g) 
     if shouldTrue == false {
        t.Error("check geo data fail")    
     }
}
