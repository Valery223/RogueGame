package main

import "fmt"

type HasID interface {
	GetID() int
}

type Item struct {
	id   int
	name string
}

type AxeItem struct {
	Item
	damage float64
}

func (i Item) testMetod() {
	fmt.Println(" testMetod", i.id, i.name)
}
func (i Item) GetID() int {
	return i.id
}

func test(i any) {
	if hasID, ok := i.(HasID); ok {
		fmt.Println("ID:", hasID.GetID())
	}
}

func main() {
	item := Item{id: 1, name: "First"}
	item.testMetod()

	axeItem := AxeItem{Item: Item{id: 2, name: "Second"}, damage: 20.3}
	axeItem.testMetod()

	test(item)
	test(axeItem.Item)

}
