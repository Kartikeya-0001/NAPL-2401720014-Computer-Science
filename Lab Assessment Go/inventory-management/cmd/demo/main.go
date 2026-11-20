package main

import (
	"fmt"
	"log"

	inventory "example.com/inventory-management"
)

func main() {
	m := inventory.NewManager()

	// Populate items
	err := m.Add(inventory.Item{ID: 1, Name: "Notebook", Quantity: 10, Price: 50.00})
	if err != nil {
		log.Fatalf("failed to add item: %v", err)
	}

	err = m.Add(inventory.Item{ID: 2, Name: "Pen", Quantity: 25, Price: 12.00})
	if err != nil {
		log.Fatalf("failed to add item: %v", err)
	}

	// Print Current inventory matching output format
	fmt.Println("Current inventory:")
	for _, item := range m.ItemsSnapshot() {
		fmt.Printf("%d | %s | quantity: %d | unit price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	// Compute statistics concurrently via goroutine & channel
	summary := m.SummaryAsync()

	// Print summary results matching output format
	fmt.Printf("Products: %d\n", summary.Products)
	fmt.Printf("Total units: %d\n", summary.TotalUnits)
	fmt.Printf("Total value: %.2f\n", summary.TotalValue)
}