package dtype

type Email struct {
	Value string
}

func NewEmail(email string) (*Email, error) {
	e := &Email{
		Value: email,
	}

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Email) IsValid() error {
	return nil
}
