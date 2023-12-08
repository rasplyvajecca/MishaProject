package rooms

type Rooms struct {
	Room  string
	Size  float64
	Count int
}

func NewRooms() []Rooms {
	return []Rooms{
		{Room: "Кухня", Size: 50, Count: 1},
		{Room: "Туалет", Size: 3.33, Count: 2},
		{Room: "Ванная", Size: 5.6, Count: 2},
		{Room: "Прихожая", Size: 12, Count: 1},
		{Room: "Спальня", Size: 25.5, Count: 3},
	}
}
