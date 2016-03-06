package KsanaDB
import(
    "sort"
)

func filter(tagFilter []string, tags []string) bool {
    dataTagLen := len(tags)
    for _, elem := range(tagFilter) {
         pos := sort.SearchStrings(tags, elem)
         if pos == dataTagLen || tags[pos] != elem {
              return false   
         }
    }
    return true
}
