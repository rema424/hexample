package service2

import "context"

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// RegisterPerson ...
func (p *Provider) RegisterPerson(ctx context.Context, name, email string) (Person, error) {
	psn := Person{
		Name:  name,
		Email: email,
	}

	psn, err := p.r.RegisterPerson(ctx, psn)
	if err != nil {
		return Person{}, err
	}

	return psn, nil
}

// GetPersonByID ...
func (p *Provider) GetPersonByID(ctx context.Context, id int64) (Person, error) {
	psn, err := p.r.GetPersonByID(ctx, id)
	if err != nil {
		return Person{}, err
	}

	return psn, nil
}
