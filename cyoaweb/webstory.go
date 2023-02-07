package cyoaweb

import (
	"fmt"
	"net/http"
)

// func bookHandler(bk *Book, fallback http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		switch r.URL.Path {
// 		case "/intro":
// 			fmt.Fprintln(w, bk.Intro)
// 			log.Default().Println("Redirecting to intro")
// 		case "/favicon.ico":
// 			log.Default().Println("No favicon")
// 		default:
// 			log.Default().Println("Redirecting to default clause")
// 			fallback.ServeHTTP(w, r)
// 		}
// 	}
// }

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func RunStoryWeb() {
	http.ListenAndServe(":8080", http.HandlerFunc(hello))
}
