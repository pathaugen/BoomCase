
package main

import (
    "net/http"
	"appengine"
	"appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func handlerServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
