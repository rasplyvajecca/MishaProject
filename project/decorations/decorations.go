package decorations

type Decorations struct {
	Object string
	Color  string
	Count  int
}

func NewDecorations() []Decorations {
	return []Decorations{
		{Object: "Ковер", Color: "Красный", Count: 1},
		{Object: "Картины", Color: "Черный", Count: 9},
		{Object: "Рисунки", Color: "Белый", Count: 100},
		{Object: "Плакаты", Color: "Красный", Count: 5},
		{Object: "Вазы", Color: "Розовый", Count: 3},
	}
}
