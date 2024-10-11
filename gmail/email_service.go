package gmail

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/api/gmail/v1"
)

func ForwardEmail(service *gmail.Service, emailID, toEmail string) error {
	msg, err := service.Users.Messages.Get("me", emailID).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve the message: %v", err)
	}

	var fromEmail, subject string
	for _, header := range msg.Payload.Headers {
		if header.Name == "From" {
			fromEmail = header.Value
		}
		if header.Name == "Subject" {
			subject = header.Value
		}
	}

	encodedMessage, err := encodeMessage(msg.Snippet, toEmail, fromEmail, subject) 
	if err != nil {
		return fmt.Errorf("unable to encode the message: %v", err)
	}

	newMsg := &gmail.Message{
		Raw: encodedMessage, 
	}

	_, err = service.Users.Messages.Send("me", newMsg).Do()
	if err != nil {
		return fmt.Errorf("unable to forward the message: %v", err)
	}

	return nil
}

func encodeMessage(rawMessage, toEmail, fromEmail, subject string) (string, error) {
	// Create the email header and body
	emailHeader := fmt.Sprintf("From: %s\r\n", fromEmail)
	emailHeader += fmt.Sprintf("To: %s\r\n", toEmail)
	emailHeader += fmt.Sprintf("Subject: Fwd: %s\r\n", subject) // Prefix with "Fwd:" to indicate forwarding
	emailHeader += "\r\n" // Blank line separating headers from the body

	// Combine header and body
	fullMessage := emailHeader + rawMessage

	// Encode the message to base64
	encodedMessage := base64.URLEncoding.EncodeToString([]byte(fullMessage))
	return encodedMessage, nil
}
