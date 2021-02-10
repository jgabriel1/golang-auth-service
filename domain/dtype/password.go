package dtype

import "errors"

type Password struct {
	Value string
}

func NewPassword(password string) (*Password, error) {
	p := &Password{
		Value: password,
	}

	err := p.IsValid()
	if err != nil {
		return nil, errors.New("Invalid password.")
	}

	return p, nil
}

func (p *Password) IsValid() error {
	return nil
}

type PasswordHash struct {
	Value []byte
}

func NewPasswordHash(hash []byte) *PasswordHash {
	return &PasswordHash{
		Value: hash,
	}
}
