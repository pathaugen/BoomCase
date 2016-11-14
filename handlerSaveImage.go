
package main

import (
    //"fmt"
    "net/http"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	
	//"appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func handlerSaveImage(w http.ResponseWriter, r *http.Request) {
	output := ""
	
	output += "<h1>handlerSaveImage()</h1>"
	
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
    
    output += "<h1>START: saveImage()</h1>"
    blobkey := saveImage(r)
    output += "<h1>blobkey: "+blobkey+"</h1>"
    //blobkey := ""
    
    //output += "<h1>START: saveCaseDriver()</h1>"
    //output += saveCaseDriver(r, ctx, blobkey)
    
    //w.Header().Set("Content-Type", "text/html; charset=utf-8")
    //fmt.Fprint(w, output)
    
    http.Redirect(w, r, "/dashboard?blobkey="+blobkey, http.StatusFound)
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
