## Go example in Wasm

### Follow this document to compile and run in browser
<https://github.com/golang/go/wiki/WebAssembly#getting-started>

### Run pre-compiled files 
Install goexec and run the server: 
```shell 
go get -u github.com/shurcooL/goexec

export go path : export PATH="$PATH:$(go env GOPATH)/bin"

goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```
### Below the url's
HTTP client : http://localhost:8080/httpclient/

Pretty Print Json : http://localhost:8080/prettyjson/ for pretty printing json
