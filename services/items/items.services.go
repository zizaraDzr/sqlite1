package items

import (
	"mvc/domain/items"
	"mvc/utils/errors"
)

func GetItems() (*items.Item, *errors.RestErr) {
	item := &items.Item{}
	item.Get()
	return item, nil
}
