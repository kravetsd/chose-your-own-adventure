package cyoaweb

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

//Look towards refactoring this code. This webstrot.go file need to be renamed and may be moved to a cyoa folder.
