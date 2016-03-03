package KsanaDB
import(
    "fmt"
    "sort"
)

func filter(tagFilter []int64, tags []string) bool {
    dataTagLen := len(tags)
    for _, i := range(tagFilter) {
         elem := fmt.Sprintf("%d", i)
         if sort.SearchStrings(tags, elem) == dataTagLen {
              return false   
         }
    }
    return true
}
