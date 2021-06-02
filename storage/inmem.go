package storage

import (
	"fmt"
	"math/rand"
)

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func GetItems(maxItems int) []Item {
	items := make([]Item, maxItems)
	for idx := 0; idx < maxItems; idx++ {
		items[idx] = Item{
			Title:       fmt.Sprintf("Some item #%d", idx),
			Description: "Some description",
			Price:       rand.Intn(maxItems),
		}
	}

	return items
}
