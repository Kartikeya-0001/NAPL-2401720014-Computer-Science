package inventory

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewManager()
	err := m.Add(Item{ID: 1, Name: "Notebook", Quantity: 10, Price: 50.00})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	items := m.ItemsSnapshot()
	if len(items) != 1 || items[0].Name != "Notebook" {
		t.Errorf("failed to add item correctly")
	}
}

func TestValidation(t *testing.T) {
	m := NewManager()

	// Invalid ID
	if err := m.Add(Item{ID: 0, Name: "Item", Quantity: 5, Price: 10}); err == nil {
		t.Error("expected error for invalid ID, got nil")
	}

	// Invalid Name
	if err := m.Add(Item{ID: 1, Name: "", Quantity: 5, Price: 10}); err == nil {
		t.Error("expected error for empty name, got nil")
	}

	// Invalid Quantity
	if err := m.Add(Item{ID: 2, Name: "Item", Quantity: -1, Price: 10}); err == nil {
		t.Error("expected error for negative quantity, got nil")
	}

	// Invalid Price
	if err := m.Add(Item{ID: 3, Name: "Item", Quantity: 5, Price: -5.0}); err == nil {
		t.Error("expected error for negative price, got nil")
	}
}

func TestUpdate(t *testing.T) {
	m := NewManager()
	_ = m.Add(Item{ID: 1, Name: "Pen", Quantity: 10, Price: 5.0})

	err := m.Update(1, 20, 7.50)
	if err != nil {
		t.Fatalf("expected no error on update, got %v", err)
	}

	items := m.ItemsSnapshot()
	if items[0].Quantity != 20 || items[0].Price != 7.50 {
		t.Errorf("update failed: expected qty 20 and price 7.50, got qty %d and price %.2f", items[0].Quantity, items[0].Price)
	}
}

func TestRemove(t *testing.T) {
	m := NewManager()
	_ = m.Add(Item{ID: 1, Name: "Pen", Quantity: 10, Price: 5.0})

	err := m.Remove(1)
	if err != nil {
		t.Fatalf("expected no error on remove, got %v", err)
	}

	if len(m.ItemsSnapshot()) != 0 {
		t.Errorf("expected empty inventory after removal")
	}
}

func TestSummaryAsync(t *testing.T) {
	m := NewManager()
	_ = m.Add(Item{ID: 1, Name: "Notebook", Quantity: 10, Price: 50.00})
	_ = m.Add(Item{ID: 2, Name: "Pen", Quantity: 25, Price: 12.00})

	summary := m.SummaryAsync()

	if summary.Products != 2 {
		t.Errorf("expected 2 products, got %d", summary.Products)
	}
	if summary.TotalUnits != 35 {
		t.Errorf("expected 35 total units, got %d", summary.TotalUnits)
	}
	if summary.TotalValue != 800.00 {
		t.Errorf("expected 800.00 total value, got %.2f", summary.TotalValue)
	}
}

func TestItemsSnapshot(t *testing.T) {
	m := NewManager()
	_ = m.Add(Item{ID: 1, Name: "Notebook", Quantity: 10, Price: 50.00})

	snapshot := m.ItemsSnapshot()
	snapshot[0].Name = "Modified Externally"

	original := m.ItemsSnapshot()
	if original[0].Name == "Modified Externally" {
		t.Errorf("snapshot returned direct pointer to internal state, should be isolated copy")
	}
}