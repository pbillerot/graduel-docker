// Le package qui m'a permis de valider la chîne de fabrication :
// - compiler
// - test en local
// - build du container Docker
// - test dans le container

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Docker</h1>"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Listen http://localhost:8008/hello")
	http.ListenAndServe(":8008", nil)
}
