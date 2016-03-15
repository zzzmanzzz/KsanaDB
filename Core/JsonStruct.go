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

type RelativeTime struct {
    Value *json.Number
    Unit string
}

type Sample struct {
    Value *json.Number
    Unit string
}

type AggregatorType struct {
    Name string
    Sampling Sample
}

type Metric struct {
   Tags map[string] []string
   Name string
   Limit *json.Number
   GroupBy []string
   Aggregator AggregatorType
}

type Query struct {
    StartAbsolute *json.Number 
    EndAbsoluate *json.Number
    StartRelative RelativeTime
    EndRelative RelativeTime
    TimeZone string
    Metrics []Metric
}
