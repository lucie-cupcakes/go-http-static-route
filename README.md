# go-http-static-route

Simple Go module for making net/http static routes.

This module preloads all static files on the memory!


# usage

Grab the module

	go get github.com/lucie-cupcakes/go-http-static-route

Add it to your code:

```go
package main

import (
	"log"
	"net/http"
	httpStaticRoute "github.com/lucie-cupcakes/go-http-static-route"
)

func main() {
	staticFiles, err := httpStaticRoute.LoadStaticFiles("./www",
	func(filePath string) bool {
		return true
	})
	if err != nil {
		panic(err)
	}
	httpStaticRoute.AddStaticRoutes(staticFiles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```
Enjoy!

# beta software

This project is in heavy development, please use with caution on production.

