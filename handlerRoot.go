
package main

import (
    "fmt"
    "net/http"
	"appengine"
	"appengine/user"
)

// ========== ========== ========== ========== ==========
// GAE handler - handlerRoot
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, drawPage(r, ctx))
}
// ========== ========== ========== ========== ==========
