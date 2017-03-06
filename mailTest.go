package main

import (
//    "log"
//   "net/smtp"
    "github.com/go-gomail/gomail"
)


func main() {
     m := gomail.NewMessage()
     m.SetHeader("From", "john@gardenway.org")
     m.SetHeader("To", "jwmurray@alum.mit.edu")
     //m.SetAddressHeader("Cc", "dan@example.com", "Dan")
     m.SetHeader("Subject", "Hello from Go to MIT!")
     m.SetBody("text/html", "Hello <b>Jim</b> and <i>Cora</i>!")
     //m.Attach("/home/Alex/lolcat.jpg")
     
     d := gomail.NewDialer("smtp.gmail.com", 587, "john@gardenway.org", "lancer83")
     
     // Send the email to Bob, Cora and Dan.
     if err := d.DialAndSend(m); err != nil {
         panic(err)
	 }
}