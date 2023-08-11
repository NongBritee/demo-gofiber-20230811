package mockup

import (
	"demo-gofiber/query"
	"math/rand"
)

type mockRepository struct {
	items []query.QueryItem
}

func NewMockRepository() *mockRepository {
	// random item 100 rows
	var items []query.QueryItem
	for i := 0; i < 100; i++ {
		// randome active power 1-1000
		// randome power input 1-1000
		items = append(items, query.QueryItem{
			ActivePower: rand.Intn(1000),
			PowerInput:  rand.Intn(1000),
		})
	}

	return &mockRepository{
		items: items,
	}
}

func (m mockRepository) QueryAll() ([]query.QueryItem, error) {
	return m.items, nil
}
