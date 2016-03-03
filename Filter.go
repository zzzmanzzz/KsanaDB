package KsanaDB
import(
    "fmt"
    "sort"
)

func filter(tagFilter []int64, tags []string) bool {
    dataTagLen := len(tags)
    for _, i := range(tagFilter) {
         elem := fmt.Sprintf("%d", i)
         pos := sort.SearchStrings(tags, elem)
         if pos == dataTagLen || tags[pos] != elem {
              return false   
         }
    }
    return true
}
