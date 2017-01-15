
package main

import (
    "fmt"
    "net/http"
    
	"google.golang.org/appengine"
	//"appengine/datastore"
	"google.golang.org/appengine/user"
	
	//"appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
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
		//g.Author = u.Email
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
    
    //output += "<h1>START: saveImage()</h1>"
    //blobkey := saveImage(r)
    //blobkey := ""
    
    output += "<h1>START: saveCaseDriver()</h1>"
    //output += saveCaseDriver(r, ctx, blobkey)
    output += saveCaseDriver(r, ctx)
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
