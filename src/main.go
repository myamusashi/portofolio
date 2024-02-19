package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gopkg.in/gomail.v2"
)

func main() {
	http.HandleFunc("/confirm", Indexconfirm)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	port := ":9999"
	http.ListenAndServe(port, nil)
}

func Indexconfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tmpl := template.Must(
			template.New("result").ParseFiles("static/templates/confirmation.html"),
		)

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := Data{
			Email: r.FormValue("email"),
			Pesan: r.FormValue("subject"),
		}

		data.sendMail()

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

type Data struct {
	NamaDepan    string
	NamaBelakang string
	Email        string
	Pesan        string
}

func (d Data) sendMail() {
	m := gomail.NewMessage()

	m.SetHeader("From", d.Email)
	m.SetHeader("To", "jinxprogilang666@gmail.com")
	m.SetHeader("Subject", "New message")
	m.SetBody("text/html", d.Pesan)

	send := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"jinxprogilang666@gmail.com",
		"wgoozrudcqgpdvhd",
	)

	if err := send.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	} else {
		fmt.Println("EMail sent!")
		return
	}

}
