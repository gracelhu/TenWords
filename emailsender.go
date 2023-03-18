package main

import (

	"fmt"
	"net/smtp"
	"time"

)

	func main() {
		// Set up authentication information.
		auth := smtp.PlainAuth(
			"",
			"aayeshasns3113@gmail.com", // Your email address
			"zurajanai",       // Your email password
			"smtp.gmail.com",           // Your email provider's SMTP server address
		)

		// Set up email message
		from := "aayeshasns3113@gmail.com"                                    // Your email address
		to := []string{"aayeshamislam@gmail.com", "abdullahnj2004@gmail.com"} // Recipient email addresses
		subject := "Auto Email Notifier by Aayesha"
		body := "This is an automated email notification sent at " + time.Now().Format("2006-01-02 15:04:05")

		message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body)

		// Send email

		err := smtp.SendMail("smtp.gmail.com:465", auth, from, to, []byte(message))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Email sent!")
	}
