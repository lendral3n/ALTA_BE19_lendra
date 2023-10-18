package main

import (
    "fmt"
    "sort"
)

func MostAppearItem(items []string) string {
    itemCounts := make(map[string]int)

    for _, item := range items {
        itemCounts[item]++
    }

    type itemWithCount struct {
        item  string
        count int
    }

    var sortedItems []itemWithCount
    for item, count := range itemCounts {
        sortedItems = append(sortedItems, itemWithCount{item, count})
    }

    sort.Slice(sortedItems, func(i, j int) bool {
        return sortedItems[i].item < sortedItems[j].item
    })

    result := ""
    for _, item := range sortedItems {
        result += item.item + "->" + fmt.Sprint(item.count) + " "
    }

    return result
}

func main() {
    fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // golang->1 ruby->2 js->4
    fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // C->1 D->2 B->3 A->4
    fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"})) // football->1 basketball->1 tenis->1
}
