package email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
)

func Email() {
	// sender data
	godotenv.Load(".env")
	from := os.Getenv("FROM_EMAIL")     //ex: "John.Doe@gmail.com"
	password := os.Getenv("EMAIL_PASS") // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := os.Getenv("fadel.lukman.dev@gmail.com") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	fmt.Println(from, password)
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Our Golang Email\n"
	body := "our first email!"
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

func EmailV2(emailAddress []byte) {
	godotenv.Load(".env")
	from := os.Getenv("FROM_EMAIL")     //ex: "John.Doe@gmail.com"
	password := os.Getenv("EMAIL_PASS") // ex: "ieiemcjdkejspqz"
	fmt.Println(from, password)

	email := mail.NewMessage()

	email.SetHeader("From", from)
	email.SetHeader("To", string(emailAddress))
	email.SetHeader("Subject", "Welcome to FAPA-STORE!")
	email.SetBody("text/plain", "This verification token")

	send := mail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := send.DialAndSend(email); err != nil {
		fmt.Println(err)
		panic(err)
	}

	//Bentuknya Publisher
	//Publisher dulu

}

// func SendEmail(send *mail.Dialer, email *mail.Message) {

// }

// func PublisherEmail(emailAddress string) {
// 	queue.PubliserEmail(emailAddress)
// }

// func ConsumerEmail(send *mail.Dialer, email *mail.Message) {

// }
