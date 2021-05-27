package main

import (
    "./hashtable"
    "fmt"
)

func main() {
    var snowflakes = []hashtable.Snowflake{
        // identical right [0,1]
        //{4,5,6,1,2,3},
        //{1,2,3,4,5,6},

        // identical left [2,3]
        //{1,2,3,4,5,6},
        //{6,5,4,3,2,1},

        // unidentical
        {1,2,3,4,5,6},
        {1,2,3,4,5,8},

        {1,8,3,4,5,8},
        {1,9,3,0,0,8},
        {1,2,1,1,1,8},
        {1,2,1,11,5,9},

        // identical with one of the above
        {4,5,8,1,2,3},
    }

    table := hashtable.Snowflakes_hashtable{}

    for _, v := range snowflakes {
        table.Add(v)
    }

    fmt.Println("repeated hashes: ", table.Repeated_snowflakes_hashes)
    fmt.Println("contains identical: ", table.ContainsIdentical())
}