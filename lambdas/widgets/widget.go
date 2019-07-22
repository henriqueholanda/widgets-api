package widgets

type Widget struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     string `json:"price"`
	Melts     bool   `json:"melts"`
	Inventory string `json:"inventory"`
}

type Widgets []Widget
