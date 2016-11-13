
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
func saveCaseDriverBlobstore(r *http.Request) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
	}
	file := blobs["file"]
	
	if len(file) == 0 {
		output = ""
	} else {
		output = string(file[0].BlobKey)
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
//func saveCaseDatastore(r *http.Request, ctx appengine.Context, caseData Case) (string) {
func saveCaseDriverDatastore(r *http.Request, ctx appengine.Context, blobkey string) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	//ctx := appengine.NewContext(r) // c or ctx
	// Can send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// Create Struct - Case
	caseLength, _	:= strconv.Atoi(r.FormValue("caselength"))
	casewidth, _	:= strconv.Atoi(r.FormValue("casewidth"))
	caseheight, _	:= strconv.Atoi(r.FormValue("caseheight"))
	caseweight, _	:= strconv.Atoi(r.FormValue("caseweight"))
	casebattery, _	:= strconv.Atoi(r.FormValue("casebattery"))
	caseprice, _	:= strconv.Atoi(r.FormValue("caseprice"))
	
	casesold, _		:= strconv.ParseBool(r.FormValue("casesold"))
	
	caseData := Case {
		Name:				r.FormValue("casename"),
		Overview:			r.FormValue("caseoverview"),
		Featuring:			r.FormValue("casefeaturing"),
		FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		
		Length:				caseLength, // int
		Width:				casewidth, // int
		Height:				caseheight, // int
		
		Weight:				caseweight, // int
		Battery:			casebattery, // int
		Notes:				r.FormValue("casenotes"),
		
		Price:				caseprice, // int
		
		Sold:				casesold, // bool
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// Create Struct - Driver
	/*
	driverdiameter, _		:= strconv.Atoi(r.FormValue("driverdiameter"))
	driverprice, _			:= strconv.Atoi(r.FormValue("driverprice"))
	
	driverData := Driver {
		Name:				r.FormValue("drivername"),
		FrequencyResponse:	r.FormValue("driverfrequencyresponse"),
		
		Diameter:			driverdiameter, // int
		Price:				driverprice, // int
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	*/
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	output += "<h1>r.FormValue(\"casename\") = ["+r.FormValue("casename")+"]</h1>"
	output += "<h1>caseData.Name = ["+caseData.Name+"]</h1>"
	if caseData.Name != "" {
		// ========== ========== ========== ========== ==========
		// Store Golang struct in the datastore
		key := datastore.NewIncompleteKey(ctx, "Case", caseKey(ctx))
		_, err := datastore.Put(ctx, key, &caseData)
		if err != nil {
			output += "<h1>ERROR: datastore failed</h1>"
		} else {
			output += "<h1>SUCCESS: Created datastore entry for new case</h1>"
			if caseData.BlobKey != "" {
				output += "<img src=\"/serve/?blobKey="+caseData.BlobKey+"\" />"
			} else {
				output += "<h1>No image was uploaded</h1>"
			}
		}
		// ========== ========== ========== ========== ==========
	} else if caseData.BlobKey != "" {
		// ========== ========== ========== ========== ==========
		// Delete blobstore entry
		//deleteBlobKey := "EUG76sbkgL8CDsNUokKcRQ=="
		output += "<div>Prepare to delete from blobstore: "+caseData.BlobKey+"</div>"
		blobstore.Delete(ctx, appengine.BlobKey(caseData.BlobKey)) // https://cloud.google.com/appengine/docs/go/blobstore/reference#Delete
		output += "<div>Finished deleting from blobstore</div>"
		// ========== ========== ========== ========== ==========
	} else {
		output += "<h1>Form was submitted blank</h1>"
	}
	// ========== ========== ========== ========== ==========
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========






