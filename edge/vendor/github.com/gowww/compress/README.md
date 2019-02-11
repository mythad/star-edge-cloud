# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) compress [![GoDoc](https://godoc.org/github.com/gowww/compress?status.svg)](https://godoc.org/github.com/gowww/compress) [![Build](https://travis-ci.org/gowww/compress.svg?branch=master)](https://travis-ci.org/gowww/compress) [![Coverage](https://coveralls.io/repos/github/gowww/compress/badge.svg?branch=master)](https://coveralls.io/github/gowww/compress?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/compress)](https://goreportcard.com/report/github.com/gowww/compress) ![Status Stable](https://img.shields.io/badge/status-stable-brightgreen.svg)

Package [compress](https://godoc.org/github.com/gowww/compress) provides a clever gzip compressing handler.

It takes care to not handle small contents, or contents that are already compressed (like JPEG, MPEG or PDF).  
Trying to gzip them not only wastes CPU but can potentially increase the response size.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/compress
	```

2. Import it in your code:

	```Go
	import "github.com/gowww/compress"
	```

## Usage

To wrap an [http.Handler](https://golang.org/pkg/net/http/#Handler), use [Handle](https://godoc.org/github.com/gowww/compress#Handle):

```Go
mux := http.NewServeMux()

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
})

http.ListenAndServe(":8080", compress.Handle(handler))
```

To wrap an [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc), use [HandleFunc](https://godoc.org/github.com/gowww/compress#HandleFunc):

```Go
http.Handle("/", compress.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}))

http.ListenAndServe(":8080", nil)
```

All in all, make sure to include this handler above any other handler that may write the response.
