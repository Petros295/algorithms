package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	item := Product{
		Name:     "Смартфон",
		Price:    25000.50,
		Quantity: 3,
	}
	fmt.Println("Товар", item.Name, "цена:", item.Price, "Остаток:", item.Quantity)
	if item.Quantity < 5 {
		fmt.Println("Лимит исчерпан, нужно заказать еще.")
	} else {
		fmt.Println("Товара на складе достаточно.")
	}
}
