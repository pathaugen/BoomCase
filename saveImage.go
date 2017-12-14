
package main

import (
	//"html/template"
	"net/http"
	//"time"

	//"strconv"

	//"appengine"
	//"appengine/datastore"
	//"appengine/user"

	//"appengine/blobstore"
	"google.golang.org/appengine/blobstore" // https://cloud.google.com/appengine/docs/go/blobstore/reference
)


// ========== START: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveImage(r *http.Request) (string) {
	output := ""

	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {  }
	file := blobs["file"]

	if len(file) == 0 { output = ""
	} else { output = string(file[0].BlobKey) }
	// ========== ========== ========== ========== ==========

  return output
}
// ========== END: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
