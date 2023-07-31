package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte("<div><h1 style=\"text-align:center\"><strong><span style=\"font-size:24px\">Hello Argo CD&nbsp;</span></strong>🐙</h1></div>
			<div><p style=\"text-align:center\"><span style=\"font-size:18px\">A GitOps Project&nbsp;</span>😎</p></div>"
			)
		)
	})

	http.ListenAndServe(":8080", nil)
}
