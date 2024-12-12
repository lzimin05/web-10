package usecase

func (u *Usecase) FetchHelloUserMessage() (string, error) {
	msg, err := u.p.SelectRandomHelloUser()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetHelloUserMessage(msg string) error {
	err := u.p.InsertUser(msg)
	if err != nil {
		return err
	}
	return nil
}
