package family

type Family struct {
	Member string
	Name   string
	Age    int
}

func NewFamily() []Family {
	return []Family{
		{Member: "Папа", Name: "Яша", Age: 30},
		{Member: "Мама", Name: "Оля", Age: 30},
		{Member: "Бабушка", Name: "Яна", Age: 68},
		{Member: "Дедушка", Name: "Ваня", Age: 72},
		{Member: "Дедушка", Name: "Паша", Age: 88},
	}
}
