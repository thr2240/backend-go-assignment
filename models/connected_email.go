package models

type ConnectedEmail struct {
	ID           int
	Email        string
	Pop3pHost    string
	Pop3Port     int
	Pop3Username string
	Pop3Password string
}

func (c *ConnectedEmail) AuthPop3() error {
	client, err := pop3.DialTLS(c.Pop3Host, c.Pop3Port, nil)
	if err != nil {
		return err
	}
	defer client.Quit()

	if err := client.User(c.Pop3Username); err != nil {
		return err
	}

	if err := client.Pass(c.Pop3Password); err != nil {
		return err
	}

	return nil
}

func (c *ConnectedEmail) ReadEmailsPop3() error {
	client, err := pop3.DialTLS(c.Pop3Host, c.Pop3Port, nil)
	if err != nil {
		return nil, err
	}
	defer client.Quit()

	if err := client.User(c.Pop3Username); err != nil {
		return nil, err
	}

	if err := client.Pass(c.Pop3Password); err != nil {
		return nil, err
	}

	count, size, err := client.Stat()
	if err != nil {
		return nil, err
	}

	var bodies []enmime.MIMEBody

	for i := 1; i <= count; i++ {
		msg, err := client.Retr(i)
		if err != nil {
			return nil, err
		}

		body, err := enmime.ParseMIMEBody(msg)
		if err != nil {
			return nil, err
		}

		bodies = append(bodies, *body)

		if err := client.Dele(i); err != nil {
			return nil, err
		}
	}

	if err := client.Noop(); err != nil {
		return nil, err
	}

	if err := client.Quit(); err != nil {
		return nil, err
	}

	return bodies, nil
}

