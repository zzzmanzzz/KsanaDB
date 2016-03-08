package KsanaDB
import(
        "testing" 
        "fmt"
)

func Test_groupBy(t *testing.T) {  
    group := map[string][]string{
       "type":{"2","6","11","18","27"}, 
       "host":{"3","4","7","9","12","14","16","19","21","23"},
    }
    HitTagsA := []string{"12", "6"}
    HitTagsB := []string{"12", "6", "9999"}
    UnHitTagsA := []string{"1", "6"}
    UnHitTagsB := []string{"6"}
    UnHitTagsC := []string{"100", "9999"}
    
    r:=groupBy(group, HitTagsA)
    if len(r) != 2 {
        fmt.Println(group) 
        fmt.Println(HitTagsA)
        t.Error("err")
    }
    r=groupBy(group, HitTagsB)
    if len(r) != 2 {
        fmt.Println(group) 
        fmt.Println(HitTagsB)
        t.Error("err")
    }
    r=groupBy(group, UnHitTagsA)
    if len(r) != 0 {
        fmt.Println(group) 
        fmt.Println(UnHitTagsA)
        t.Error("err")
    }
    r=groupBy(group, UnHitTagsB)
    if len(r) != 0 {
        fmt.Println(group) 
        fmt.Println(UnHitTagsB)
        t.Error("err")
    }
    r=groupBy(group, UnHitTagsC)
    if len(r) != 0 {
        fmt.Println(group) 
        fmt.Println(UnHitTagsC)
        t.Error("err")
    }
}
