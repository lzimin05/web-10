package usecase

import "errors"

func (u *Usecase) FetchCount() (string, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return "", err
	}
	if msg == "" {
		return "", errors.New("count not found in DB")
	}
	return msg, nil
}

func (u *Usecase) SetCount(msg int) error {
	err := u.p.InsertCount(msg)
	if err != nil {
		return err
	}
	return nil
}
