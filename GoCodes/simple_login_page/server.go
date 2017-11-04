package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type loginStruct struct {
	Username string `json:username`
	Password string `json:password`
}

func LoginCheck(w http.ResponseWriter, r *http.Request) {
	userName := "admin"
	passWord := "pass"

	var creds loginStruct
	credsBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad Request: Error while reading")
		w.WriteHeader(400)
	}
	err = json.Unmarshal(credsBody, &creds)

	if err != nil {
		log.Printf("Internal Error: Error while unmarshalling")
		w.WriteHeader(500)
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte("username or password is required"))
	}
 	if creds.Username != userName || creds.Password != passWord {

		w.WriteHeader(401)
	} else {

		w.WriteHeader(200)
	}

}

func main() {
	fmt.Printf("Starting the server.. localhost:8085")

	mux := http.NewServeMux()

	mux.HandleFunc("/login", LoginCheck)

	log.Fatal(http.ListenAndServe(":8085", mux))
}
