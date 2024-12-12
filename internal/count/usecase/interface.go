package usecase

type Provider interface {
	SelectCount() (string, error)
	InsertCount(int) error
}
