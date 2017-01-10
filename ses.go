package ses

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Email options.
// credits : https://github.com/tj/go-ses
type Email struct {
	From    string // From source email
	To      string // To destination email(s)
	Subject string // Subject text to send
	Text    string // Text is the text body representation
	HTML    string // HTMLBody is the HTML body representation
	ReplyTo string // Reply-To email(s)
}

// *********************************************************************
//	create a new aws session and returns the session var
//	@returns sess *session.Session
//
func startNewSession() *session.Session {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
	}
	return sess
}

// *********************************************************************
//	aws configuration
//
func SetConfiguration(aws_key_id string, aws_secret_key string, aws_region string) {
	os.Setenv("AWS_REGION", aws_region)
	os.Setenv("AWS_ACCESS_KEY_ID", aws_key_id)
	os.Setenv("AWS_SECRET_ACCESS_KEY", aws_secret_key)
}

// *********************************************************************
//	create and send text or html email to single receipents.
//	@returns resp *ses.SendEmailOutput
//
func SendEmail(emailData Email) *ses.SendEmailOutput {
	// start a new aws session
	sess := startNewSession()
	// start a new ses session
	svc := ses.New(sess)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(emailData.To), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Text: &ses.Content{
					Data:    aws.String(emailData.Text), // Required
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(emailData.Subject), // Required
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(emailData.From), // Required

		ReplyToAddresses: []*string{
			aws.String(emailData.ReplyTo), // Required
			// More values...
		},
	}

	// send email
	resp, err := svc.SendEmail(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	return resp
}
