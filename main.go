
package main

import (
    "fmt"
    "net/http"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
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
	
	http.HandleFunc("/savecase", handlerSaveCase)
	
    http.HandleFunc("/", handlerRoot)
    
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerSaveCase
func handlerSaveCase(w http.ResponseWriter, r *http.Request) {
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	c := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(c); u != nil {
		//g.Author = u.String()
	}
	// ========== ========== ========== ========== ==========
	
	
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    //fmt.Fprint(w, drawPage(r.URL.Path[1:]))
    fmt.Fprint(w, saveCase(r))
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerRoot
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	c := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(c); u != nil {
		//g.Author = u.String()
	}
	// ========== ========== ========== ========== ==========
	
	
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    //fmt.Fprint(w, drawPage(r.URL.Path[1:]))
    fmt.Fprint(w, drawPage(r))
}
// ========== ========== ========== ========== ==========

