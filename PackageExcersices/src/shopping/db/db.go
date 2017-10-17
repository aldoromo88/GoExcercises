package db

type Item struct {
	Price float64
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 19.99,
	}
}
