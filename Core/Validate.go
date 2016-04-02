package KsanaDB                                                                                                            
import (
    "strings"
)

func isMetricNameValidate(metric string) bool {
     return !strings.Contains(metric, "\t",)
}
