package api

type Usecase interface {
	FetchCount() (string, error)
	SetCount(msg int) error
}
