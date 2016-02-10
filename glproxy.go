package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func msg(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	log_str := "Status: %d, " + msg + "\n"
	log.Printf(log_str, status)
	fmt.Fprintln(w, msg)
}

var u_proto_regex = regexp.MustCompile(`(http(|s))-`)

func proxyGLWebHook(w http.ResponseWriter, r *http.Request) {
	t_url_parts := strings.Split(r.URL.Path, "/")
	domain := u_proto_regex.ReplaceAllString(t_url_parts[1], "$1://")
	log.Println(domain)
	t_url := domain + "/" + strings.Join(t_url_parts[2:], "/")
	log.Println("Target URL: " + t_url)
	procr, er := http.NewRequest("POST", t_url, r.Body)
	if er != nil {
		msg(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	procr.Header.Add("X-Gitlab-Event", r.Header.Get("X-Gitlab-Event"))
	client := &http.Client{}
	procrsp, er := client.Do(procr)
	if er != nil {
		msg(w, "Failed to post", http.StatusBadRequest)
		log.Println(er)
		return
	}
	msg(w, "Successfully transfered", procrsp.StatusCode)
}

func main() {
	http.HandleFunc("/", proxyGLWebHook)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
