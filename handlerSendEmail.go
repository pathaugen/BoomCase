
package main

import (
  "fmt"
  "net/http"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func handlerSendEmail(w http.ResponseWriter, r *http.Request) {
	output := ""

	output += "<h1>handlerSendEmail()</h1>"

	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	//ctx := appengine.NewContext(r) // c or ctx
	// Can send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========

    output += "<h1>START: sendEmail()</h1>"
    output += sendEmail(r)
    //blobkey := sendEmail(r)
    //output += "<h1>blobkey: "+blobkey+"</h1>"
    //blobkey := ""

    //output += "<h1>START: saveCaseDriver()</h1>"
    //output += saveCaseDriver(r, ctx, blobkey)

    //w.Header().Set("Content-Type", "text/html; charset=utf-8")
    //fmt.Fprint(w, output)

    //http.Redirect(w, r, "/dashboard?blobkey="+blobkey, http.StatusFound)

    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
