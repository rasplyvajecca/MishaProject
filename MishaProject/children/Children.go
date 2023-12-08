package children

type Children struct {
	Name string
	Age  int
}

func NewChildren() []Children {
	return []Children{
		{Name: "Вася", Age: 5},
		{Name: "Ира", Age: 2},
		{Name: "Лиза", Age: 2},
		{Name: "Маша", Age: 11},
		{Name: "Марина", Age: 13},
	}
}
