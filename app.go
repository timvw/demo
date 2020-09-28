package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubMessage struct {
	Ref string `json:"ref"`
	Repository struct {
		SshURL string `json:"ssh_url"`
	} `json:"repository"`
}

func Github(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	fmt.Println("need to handle a request...")
	var payload GithubMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Printf("failed to decode %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("ref: %v\n", payload.Ref)
	fmt.Printf("sshurl: %v\n", payload.Repository.SshURL)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/github", Github)
	http.ListenAndServe(":80", nil)
}
