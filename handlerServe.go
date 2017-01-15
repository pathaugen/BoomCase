
package main

import (
    "net/http"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func handlerServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
