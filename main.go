
package main

import (
    "net/http"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
// GAE config handlers for URIs
func init() {
	http.HandleFunc("/savecasedriver",	handlerSaveCaseDriver)
	http.HandleFunc("/saveimage",		handlerSaveImage)
	http.HandleFunc("/serve/",			handlerServe)
	http.HandleFunc("/login",			handlerLogin)
	
	// API Versions
	http.HandleFunc("/api/1.0/",		handlerAPI10) // API version 1.0
	
    http.HandleFunc("/",				handlerRoot)
}

// if true, running in production environment
// os.getenv('SERVER_SOFTWARE', '').startswith('Google App Engine/')
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

// Dummy Main() as required - Used for executable when running after build
func main() {
	
}
