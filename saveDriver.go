
package main

import (
	//"html/template"
	"net/http"
	"time"
	
	"strconv"
	
	"appengine"
	"appengine/datastore"
	//"appengine/user"
	
	//"google.golang.org/appengine/blobstore"
	"appengine/blobstore" // https://cloud.google.com/appengine/docs/go/blobstore/reference
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveDriver(r *http.Request, ctx appengine.Context) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// Delete blobstore entry
	output += "<div>Prepare to delete from blobstore</div>"
	deleteBlobKey := "EUG76sbkgL8CDsNUokKcRQ=="
	blobstore.Delete(ctx, appengine.BlobKey(deleteBlobKey)) // https://cloud.google.com/appengine/docs/go/blobstore/reference#Delete
	output += "<div>Finished deleting from blobstore</div>"
	// ========== ========== ========== ========== ==========
	
	blobkey := saveDriverBlobstore(r)
	
	// ========== ========== ========== ========== ==========
	// Pull the POST form fields into a Golang struct
	/*
	type Driver struct {
		Name				string
		FrequencyResponse	string
		Width				int
		Price				int
		
		BlobKey				string
		
		DateAdded			time.Time
	}
	*/
	casewidth, _	:= strconv.Atoi(r.FormValue("casewidth"))
	caseprice, _	:= strconv.Atoi(r.FormValue("caseprice"))
	
	driverData := Driver {
		Name:				r.FormValue("casename"),
		FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		
		Width:				casewidth, // int
		Price:				caseprice, // int
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	// ========== ========== ========== ========== ==========
	
	output += saveDriverDatastore(r, ctx, driverData)
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveDriverBlobstore(r *http.Request) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		//output += "<h1>ERROR: "+err.Error()+"</h1>"
	}
	file := blobs["file"]
	
	if len(file) == 0 {
		//output += "<h1>WARNING: No image file uploaded to blobstore</h1>"
		output = ""
	} else {
		output = string(file[0].BlobKey)
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveDriverDatastore(r *http.Request, ctx appengine.Context, driverData Driver) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// Store Golang struct in the datastore
	key := datastore.NewIncompleteKey(ctx, "Driver", driverKey(ctx))
	_, err := datastore.Put(ctx, key, &driverData)
	if err != nil {
		output += "<h1>ERROR: datastore failed</h1>"
	} else {
		output += "<h1>SUCCESS: Created datastore entry for new driver</h1>"
		if driverData.BlobKey != "" {
			output += "<img src=\"/serve/?blobKey="+driverData.BlobKey+"\" />"
		} else {
			output += "<h1>No image was uploaded</h1>"
		}
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========





