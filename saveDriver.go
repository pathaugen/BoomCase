
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
	"appengine/blobstore"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveDriver(r *http.Request) (string) {
	output := ""
	
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
	
	output += saveDriverDatastore(r, driverData)
	
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
		output += "<h1>ERROR: "+err.Error()+"</h1>"
	}
	file := blobs["file"]
	
	if len(file) == 0 {
		output += "<h1>WARNING: No image file uploaded to blobstore</h1>"
	} else {
		output = string(file[0].BlobKey)
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveDriverDatastore(r *http.Request, driverData Driver) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// Store Golang struct in the datastore
	key := datastore.NewIncompleteKey(ctx, "Driver", driverKey(ctx))
	_, err := datastore.Put(ctx, key, &driverData)
	if err != nil {
		output += "<h1>ERROR: datastore failed</h1>"
	} else {
		output += "<h1>SUCCESS: Created datastore entry for new driver</h1>"
		output += "<img src=\"/serve/?blobKey="+driverData.BlobKey+"\" />"
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========





