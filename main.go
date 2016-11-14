
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
    http.HandleFunc("/",				handlerRoot)
}

// if true, running in production environment
// os.getenv('SERVER_SOFTWARE', '').startswith('Google App Engine/')
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
