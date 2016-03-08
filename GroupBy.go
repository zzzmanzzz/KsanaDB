package KsanaDB
import(
    "sort"
)

func groupBy(group map[string][]string, tags []string) ([]string) {  
    ret := []string{}
    if len(tags) == 0 {
        return ret 
    }
    for _, tagSeq := range(group)  {
        sort.Strings(tagSeq)
        for _, tag := range(tags) {  
            groupTagLen := len(tagSeq)
            pos := sort.SearchStrings(tagSeq, tag) 
            if pos != groupTagLen && tagSeq[pos] == tag {
                ret = append(ret, tag)
            }
        }
    }

    retLen := len(ret)
    if retLen == len(group){
        sort.Strings(ret)
        return ret
    } else {
        ret = []string{}    
    } 
    return ret
}
