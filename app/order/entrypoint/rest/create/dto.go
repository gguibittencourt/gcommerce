package create

type (
	Payload struct {
		CPF   string `validate:"validCPF"`
		Items []Item `validate:"required"`
	}

	Item struct {
		Amount      uint32
		Description string
		Price       float64
	}
)
