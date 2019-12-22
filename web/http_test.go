package web

import "net/http"

func Example() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, world!"))
	})
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
