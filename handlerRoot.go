
package main

import (
    "fmt"
    "net/http"

	"google.golang.org/appengine"
	//"appengine/user"
)

// ========== START: handlerRoot ========== ========== ========== ==========
func handlerRoot(w http.ResponseWriter, r *http.Request) {

	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, drawPage(r, ctx))
}
// ========== END: handlerRoot ========== ========== ========== ==========
