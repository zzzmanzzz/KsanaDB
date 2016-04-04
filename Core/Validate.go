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
