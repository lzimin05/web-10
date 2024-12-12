package api

type Usecase interface {
	FetchHelloUserMessage() (string, error)
	SetHelloUserMessage(msg string) error
}
