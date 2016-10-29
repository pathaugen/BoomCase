
package main

import (
    "fmt"
    "net/http"
)


// Golang BUILD output
/*
func main() {
	fmt.Printf("Boomcase Console App Output\n")
}
*/


// ========== ========== ========== ========== ==========
// GAE config handlers for URIs
func init() {
    http.HandleFunc("/", handlerRoot)
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    fmt.Fprint(w, drawPage(r.URL.Path[1:]))
}
// ========== ========== ========== ========== ==========

