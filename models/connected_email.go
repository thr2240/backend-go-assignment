package models

import (
	"github.com/thr2240/backend-go-assignment/go-pop3"
	"github.com/emersion/go-message"
	"github.com/jordan-wright/email"
	"log"
)
type ConnectedEmail struct {
	ID           int
	Email        string
	Pop3pHost    string
	Pop3Port     int
	Pop3Username string
	Pop3Password string
}

// AuthPop3 checks if login credentials of POP3 are valid
func (c *ConnectedEmail) AuthPop3() error {
	// Create a new POP3 client
	client := pop3.New(pop3.Opt{
		Host:          c.Pop3Host,
		Port:          c.Pop3Port,
		DialTimeout:   3, // 3 seconds
		TLSEnabled:    true,
		TLSSkipVerify: false,
	})

	// Establish a connection to the POP3 server
	conn, err := client.NewConn()
	if err != nil {
		return err
	}
	defer conn.Quit()

	// Authenticate with the POP3 server
	if err := conn.Auth(c.Pop3Username, c.Pop3Password); err != nil {
		return err
	}

	return nil
}

// ReadEmailsPop3 reads emails from the inbox using POP3
func (c *ConnectedEmail) ReadEmailsPop3() ([]*message.Entity, error) {
	// Create a new POP3 client
	client := pop3.New(pop3.Opt{
		Host:          c.Pop3Host,
		Port:          c.Pop3Port,
		DialTimeout:   3, // 3 seconds
		TLSEnabled:    true,
		TLSSkipVerify: false,
	})

	// Establish a connection to the POP3 server
	conn, err := client.NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	// Authenticate with the POP3 server
	if err := conn.Auth(c.Pop3Username, c.Pop3Password); err != nil {
		return nil, err
	}

	// Get the list of message IDs
	messageIDs, err := conn.List(0)
	if err != nil {
		return nil, err
	}

	// Fetch each email message
	var emails []*message.Entity
	for _, msgID := range messageIDs {
		msg, err := conn.Retr(msgID.ID)
		if err != nil {
			log.Printf("Failed to retrieve email with ID %d: %v", msgID.ID, err)
			continue
		}

		emails = append(emails, msg)
	}

	return emails, nil
}