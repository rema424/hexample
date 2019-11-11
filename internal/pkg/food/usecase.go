package food

// Interactor .
type Interactor struct {
	repo Repository
}

// NewInteractor .
func NewInteractor(r Repository) *Interactor {
	return &Interactor{r}
}
