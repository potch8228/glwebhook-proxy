GitLab WebHook Forward Proxy
====

What's this?
----
This program redirects GitLab WebHook to target URL which can not reach normally because of networking structure and/or etc.

Usage
----
Setup target URL as following format
```
http[s]://<url@this program>/<target url>/<remainder of url path>
```

Target url format:  
Protocol part of "://" need to be replaced with "-" (single minus sign); for other parts, follow usual URL format

e.g.
https://thisproxy.example.com/https-targetciserver.example.com/job/test-ci-job

Command Flags
----
-port (default: 8080): proxy listen port
-github=true (default: false): proxy for GitHub/GitHub Enterprise; Can not work side-by-side with GitLab CE mode

Compile Environment
----
go 1.5.3 darwin/amd64

File Structure
----
glproxy.go -> main program  
test-echoserver.go -> test echo server

Author
----
Tatsuya Kobayashi <pikopiko28@gmail.com>

