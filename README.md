# ses-go
A simple wrapper for AWS SES in golang to send email


### Sending an email


```` go
package main

import (
	"fmt"

	"github.com/srajelli/ses-go"
)

func main() {
	ses.SetConfiguration(aws_key_id, aws_secret_key, aws_region)
	
	emailData := ses.Email{
		To:      "someone@example.com",
		From:    "me@example.com",
		Text:    "Hi this is the text message body",
		Subject: "Sending email from aws ses api",
		ReplyTo: "me@example.com",
	}

	resp := ses.SendEmail(emailData)

	fmt.Println(resp)
}
````