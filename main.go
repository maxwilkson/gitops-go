package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(`
			<div style="display:flex; justify-content:center">
				<h1><strong><span style="font-size:24px">Hello Argo CD&nbsp;</span></strong>🐙</h1>
			</div>
			<div style="display:flex; justify-content:center">
				<p><span style="font-size:18px">A GitOps Project&nbsp;</span>😎</p>
			</div>`))
	})

	http.ListenAndServe(":8080", nil)
}
