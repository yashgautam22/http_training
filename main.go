package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/welcome/", WelcomeHandler)
	http.HandleFunc("/hello", DummyHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Server")
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func DummyHandler(w http.ResponseWriter, r *http.Request) {

	u := user{}

	fmt.Println(u)

	// client := http.Client{}
	// _, err := http.NewRequest("POST", "http://127.0.0.1:8080/hello", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// resp, err := client.Do(r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)

	//fmt.Println(r.Method)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, "Welcome to the Server")
}
