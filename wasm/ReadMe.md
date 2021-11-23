## Go example in Wasm

Follow this document to compile and run in browser
<https://github.com/golang/go/wiki/WebAssembly#getting-started>

To run pre-compiled files 
Install goexec:  go get -u github.com/shurcooL/goexec
export go path : export PATH="$PATH:$(go env GOPATH)/bin"

goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

Then go to 
http://localhost:8080/httpclient/ for http call example

http://localhost:8080/prettyjson/ for pretty printing json
