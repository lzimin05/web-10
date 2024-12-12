package usecase

type Provider interface {
	SelectRandomHelloUser() (string, error)
	InsertUser(msg string) error
}
