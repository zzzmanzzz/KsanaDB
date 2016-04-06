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
    Unit *string
}

type Sample struct {
    Value *json.Number
    Unit *string
}

type AggregatorType struct {
    Name string
    Sampling Sample
}

type MetricType struct {
   Tags map[string]string
   Name *string
   Limit *json.Number
   GroupBy  []string
   Aggregator AggregatorType
}

type Query struct {
    StartAbsolute *json.Number 
    EndAbsolute *json.Number
    StartRelative *RelativeTime
    EndRelative *RelativeTime
    TimeZone string
    Metric MetricType
}

type AllTagSeqType struct {
    Val map[string]string
    Seq map[string][]string
}


type GroupType struct {
    Tags map[string]string    
    Values [][] interface{}
}

type ResultType struct {
    Name string
    Start int64//json.Number
    End int64//json.Number
    Filter map[string]string `json:"Filter,omitempty"`
    AggregateFunction string
    TimeRange int
    TimeUnit string
    GroupBy []string `json:"GroupBy,omitempty"`
    Group []GroupType 
}

type OutputForm struct {
     
    Result ResultType
}
