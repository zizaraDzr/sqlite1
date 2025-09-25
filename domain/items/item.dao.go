package items

import (
	"mvc/utils/errors"
	"strconv"
)

var (
	itemsDB = []Item{}
)

func generateItems() []Item {
	items := make([]Item, 500)
	for i := 0; i < 500; i++ {
		items[i] = Item{
			Name: "Item " + strconv.Itoa(i+1),
			ID:   i + 1,
		}
	}
	return items
}

func (items *Item) Get() *errors.RestErr {
	itemsDB = generateItems()

	return nil
}
