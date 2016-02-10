package main

import (
	"flag"
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

func parseURL(raw string) (t_url string) {
	t_url_parts := strings.Split(raw, "/")
	domain := u_proto_regex.ReplaceAllString(t_url_parts[1], "$1://")
	t_url = domain + "/" + strings.Join(t_url_parts[2:], "/")
	log.Println("Target URL: " + t_url)
	return
}

func parseArgs() (p int) {
	p = *(flag.Int("port", 8080, "Listening port number: 1024 <= port <= 65535"))
	if p < 1024 || p > 65535 {
		log.Fatalln("Invalid Argument!")
	}

	return
}

var u_proto_regex = regexp.MustCompile(`(http(|s))-`)

func proxyGLWebHook(w http.ResponseWriter, r *http.Request) {
	t_url := parseURL(r.URL.Path)
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
	p := parseArgs()
	http.HandleFunc("/", proxyGLWebHook)
	lurl := fmt.Sprintf(":%d", p)
	log.Println("Trying to Listening on: " + lurl)
	log.Fatalln(http.ListenAndServe(lurl, nil))
}
