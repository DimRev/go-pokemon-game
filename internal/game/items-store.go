package game

import "fmt"

type User struct {
	Name  string
	Score int
}

type ItemStore struct {
	Items []Item
}

type Item struct {
	ID          string
	Name        string
	Description string
}

func (is *ItemStore) GetItemById(itemID string) (*Item, error) {
	for _, currItem := range is.Items {
		if currItem.ID == itemID {
			return &currItem, nil
		}
	}
	return nil, fmt.Errorf("no item with id %s found in itemStore", itemID)
}

func (is *ItemStore) AddItemToItemStore(newItem Item) {
	is.Items = append(is.Items, newItem)
}

func (is *ItemStore) RemoveItemFromItemStore(itemID string) error {
	newItems := []Item{}
	removedItem := false
	for _, currItem := range is.Items {
		if currItem.ID != itemID {
			newItems = append(newItems, currItem)
		} else {
			if !removedItem {
				removedItem = true
			} else {
				return fmt.Errorf("2 items of the id %s found in itemStore", itemID)
			}
		}
	}
	if !removedItem {
		return fmt.Errorf("no item with id %s located in itemStore", itemID)
	}
	is.Items = newItems
	return nil
}
