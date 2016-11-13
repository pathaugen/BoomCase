
package main

import (
    "fmt"
    "net/http"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	
	//"appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
// GAE handler - handlerSaveCaseDriver
func handlerSaveCaseDriver(w http.ResponseWriter, r *http.Request) {
	output := ""
	
	output += "<h1>handlerSaveCaseDriver()</h1>"
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Can send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
    
    output += "<h1>START: saveCaseDriverBlobstore()</h1>"
    //blobkey := saveCaseDriverBlobstore(r)
    blobkey := ""
    
    output += "<h1>START: saveCaseDriverDatastore()</h1>"
    output += saveCaseDriverDatastore(r, ctx, blobkey)
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
