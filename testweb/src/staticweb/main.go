package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler1 struct {
	message string
}

func (m *messageHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

type messageHandler2 struct {
	message string
}

func (m *messageHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func messageHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "messageHandler3 as a function")
}

//Handler logic into a Closure
func messageHandler4(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { 
			fmt.Fprintf(w, message)    
		}) 
}

func makeHref(text string) string {
	href := "<A href="
	href += "\"/"
	href += text
	href += "\">"
	href += text
	href += "</A></br>"
	return href
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	menu := makeHref("handler1")
	menu += makeHref("handler2")
	menu += makeHref("handler3")
	menu += makeHref("handler33")
	menu += makeHref("handler4")
	menu += makeHref("handler5")
	fmt.Fprintf(w, menu)
}

func main() {
	mux := http.NewServeMux()

	mh1 := &messageHandler1{"messageHandler1 as struct"}
	mux.Handle("/handler1", mh1)

	mh2 := &messageHandler2{"messageHandler2 as struct"}
	mux.Handle("/handler2", mh2)

	mh3 := http.HandlerFunc(messageHandler3)
	mux.Handle("/handler3", mh3)
	
	mux.HandleFunc("/handler33", messageHandler3)

	mux.Handle("/handler4", messageHandler4("messageHandler4 as a Closure"))

	//fs := http.FileServer(http.Dir("public"))
	//mux.Handle("/", fs)
	mhMenu := http.HandlerFunc(menuHandler)
	mux.Handle("/", mhMenu)

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", mux) 
}