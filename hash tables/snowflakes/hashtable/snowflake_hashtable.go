package hashtable

const TABLE_SIZE = 1000

type snowflake_list struct {
    head *snowflake_list_item
}
func (sl *snowflake_list) IsEmpty() bool {
    return sl.head == nil
}
func (sl *snowflake_list) HasOne() bool {
    return sl.head != nil && sl.head.next == nil
}
func (sl *snowflake_list) Add(s Snowflake) {
    list_item := snowflake_list_item{
        value: s,
        next: sl.head,
    }
    sl.head = &list_item
}
func (sl *snowflake_list) GetValues() []Snowflake {
    snowflakes := []Snowflake{}
    currentListItem := sl.head
    for currentListItem != nil {
        snowflakes = append(snowflakes, currentListItem.value)
        currentListItem = currentListItem.next
    }
    return snowflakes
}

type Snowflake [6]int
func (s Snowflake) Sum() int {
    sum := 0
    for _, v := range s {
        sum += v
    }
    return sum
}
func (s Snowflake) GetHash() int {
    return s.Sum() % TABLE_SIZE
}
func (s Snowflake) SameAs(s2 Snowflake) bool {
    return identical(s, s2)
}

type snowflake_list_item struct {
    value Snowflake
    next  *snowflake_list_item
}

type Snowflakes_hashtable struct {
    Repeated_snowflakes_hashes []int
    table [TABLE_SIZE]snowflake_list
}
func (sht *Snowflakes_hashtable) Add(s Snowflake) {
    hash := s.GetHash()
    if sht.table[hash].HasOne() {
        sht.Repeated_snowflakes_hashes = append(sht.Repeated_snowflakes_hashes, hash)
    }
    sht.table[hash].Add(s)
}
func (sht *Snowflakes_hashtable) ContainsIdentical() bool {
    comparisonFn := func(s Snowflake, s2 Snowflake) bool {return s.SameAs(s2)};
    for _, v := range sht.Repeated_snowflakes_hashes {
        same_sum_snowflakes := sht.table[v].GetValues()
        set_is_unique := unique_set(same_sum_snowflakes, comparisonFn)
        if !set_is_unique {
            return true
        }
    }
    return false
}

func identical_right(arr1, arr2 Snowflake) bool {
    var identical bool = false
    outer:
    for i := 0; i < 6; i++ {
        inner:
        for j := 0; j < 6; j++ {
            arr1_index := j
            arr2_index := (j + i)%6

            if arr1[arr1_index] != arr2[arr2_index] {
                break inner
            }
            if j == 5 {
                identical = true
                break outer
            }
        }
    }

    return identical
}
func identical_left(arr1, arr2 Snowflake) bool {
    var identical bool = false
outer:
    for i := 0; i < 6; i++ {
    inner:
        for j := 0; j < 6; j++ {
            arr1_index := j
            arr2_index := (i - j)%6 // we need to move left
            if arr2_index < 0 {
                arr2_index = 6 + arr2_index
            }

            if arr1[arr1_index] != arr2[arr2_index] {
                break inner
            }
            if j == 5 {
                identical = true
                break outer
            }
        }
    }

    return identical
}

func identical(arr1, arr2 Snowflake) bool {
    return identical_left(arr1, arr2) || identical_right(arr1, arr2)
}

func unique_set(
    arr []Snowflake,
    equalFn func(s1, s2 Snowflake) bool,
) bool {
    for i, v := range arr {
        for j, v2 := range arr {
            if (i == j) {
                continue
            }
            if equalFn(v, v2) {
                return false
            }
        }
    }

    return true
}
