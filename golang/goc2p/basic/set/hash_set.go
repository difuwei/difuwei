package set

type HashSet struct {
    m map[interface{}]bool
}

func (set *HashSet) Add(e interface{}) bool {
    if !set.m[e]{
        set.m[e] = true
        return true
    }
    return false
}