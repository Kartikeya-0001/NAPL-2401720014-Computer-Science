package inventory

import (
	"errors"
	"fmt"
	"sync"
)

// Item struct representing an inventory item
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Summary stores calculated inventory statistics
type Summary struct {
	Products   int
	TotalUnits int
	TotalValue float64
}

// Manager handles inventory operations
type Manager struct {
	mu    sync.RWMutex
	items []Item
}

// NewManager initializes a new inventory manager
func NewManager() *Manager {
	return &Manager{
		items: make([]Item, 0),
	}
}

// Add adds a new item after validation
func (m *Manager) Add(item Item) error {
	if item.ID <= 0 {
		return errors.New("invalid ID: must be greater than 0")
	}
	if item.Name == "" {
		return errors.New("invalid Name: cannot be empty")
	}
	if item.Quantity < 0 {
		return errors.New("invalid Quantity: cannot be negative")
	}
	if item.Price < 0 {
		return errors.New("invalid Price: cannot be negative")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, existing := range m.items {
		if existing.ID == item.ID {
			return fmt.Errorf("item with ID %d already exists", item.ID)
		}
	}

	m.items = append(m.items, item)
	return nil
}

// Update updates an existing item's quantity and price
func (m *Manager) Update(id int, quantity int, price float64) error {
	if quantity < 0 {
		return errors.New("invalid Quantity: cannot be negative")
	}
	if price < 0 {
		return errors.New("invalid Price: cannot be negative")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	for i, item := range m.items {
		if item.ID == id {
			m.items[i].Quantity = quantity
			m.items[i].Price = price
			return nil
		}
	}

	return fmt.Errorf("item with ID %d not found", id)
}

// Remove removes an item by ID
func (m *Manager) Remove(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, item := range m.items {
		if item.ID == id {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("item with ID %d not found", id)
}

// ItemsSnapshot returns a copy of current items to prevent race conditions
func (m *Manager) ItemsSnapshot() []Item {
	m.mu.RLock()
	defer m.mu.RUnlock()

	snapshot := make([]Item, len(m.items))
	copy(snapshot, m.items)
	return snapshot
}

// SummaryAsync computes inventory statistics concurrently using a goroutine and channel
func (m *Manager) SummaryAsync() Summary {
	m.mu.RLock()
	itemsCopy := make([]Item, len(m.items))
	copy(itemsCopy, m.items)
	m.mu.RUnlock()

	ch := make(chan Summary)

	go func(items []Item) {
		var totalUnits int
		var totalValue float64

		for _, item := range items {
			totalUnits += item.Quantity
			totalValue += float64(item.Quantity) * item.Price
		}

		ch <- Summary{
			Products:   len(items),
			TotalUnits: totalUnits,
			TotalValue: totalValue,
		}
	}(itemsCopy)

	return <-ch
}