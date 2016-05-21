package KsanaDB                                                                                                            
import (
    "strings"
)

func isMetricNameValidate(metric string) bool {
     return !strings.Contains(metric, "\t",)
}

func removeTab(tag string) string {
    return strings.Replace(tag, "\t", "", -1)
}

func checkGeoData(d *GeoData) bool {
     if d == nil {
         return false
     }
     if d.Name == "" {
         return false
     }
     return true 
}
