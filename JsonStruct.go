package KsanaDB
import (
    "encoding/json"
)


type DataPoint struct {
    Name string
    Value *json.Number
    Tags  map[string]interface{}
    Timestamp *json.Number
    Datapoints [][]json.Number
}

