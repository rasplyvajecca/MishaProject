package furniture

type Furniture struct {
	Object string
	Size   float64
	Count  int
}

func NewFurniture() []Furniture {
	return []Furniture{
		{Object: "Диван", Size: 6.55, Count: 3},
		{Object: "Кровать", Size: 7.77, Count: 2},
		{Object: "Шкаф", Size: 2, Count: 2},
		{Object: "Тумба", Size: 1.5, Count: 5},
		{Object: "Ковер", Size: 3.333, Count: 3},
	}
}
