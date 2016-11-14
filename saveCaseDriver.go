
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
//func saveCaseDatastore(r *http.Request, ctx appengine.Context, caseData Case) (string) {
//func saveCaseDriverDatastore(r *http.Request, ctx appengine.Context, blobkey string) (string) {
//func saveCaseDriver(r *http.Request, ctx appengine.Context, blobkey string) (string) {
func saveCaseDriver(r *http.Request, ctx appengine.Context) (string) {
	output := ""
	
	// Capture the blobkey from query string
	blobkey := r.FormValue("blobkey")
	output += "<h1>blobkey: "+blobkey+"</h1>"
	
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
	casewatts, _	:= strconv.Atoi(r.FormValue("casewatts"))
	caseprice, _	:= strconv.Atoi(r.FormValue("caseprice"))
	
	casesold, _		:= strconv.ParseBool(r.FormValue("casesold"))
	
	caseData := Case {
		Name:				r.FormValue("casename"),
		Overview:			r.FormValue("caseoverview"),
		Featuring:			r.FormValue("casefeaturing"),
		FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		
		Length:				int8(caseLength), // int8
		Width:				int8(casewidth), // int8
		Height:				int8(caseheight), // int8
		
		Weight:				int8(caseweight), // int8
		Battery:			int8(casebattery), // int8
		Notes:				r.FormValue("casenotes"),
		
		Price:				int32(caseprice), // int16
		Watts:				int16(casewatts), // int16
		Sold:				casesold, // bool
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	
	// Protection of blobkey from accidental manipulation
	//if blobkey != "" && blobkey != "<BLOBKEY>" { caseData.BlobKey = blobkey }
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
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		// Store Golang struct in the datastore
		
		// Use a new generated key, or replace with existing key to overwrite data
		newKey := datastore.NewIncompleteKey(ctx, "Case", caseKey(ctx))
		
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		if r.FormValue("existingdatastoreid") != "" {
			// Datastore query to get the key of existing entry to overwrite
			
			// Array to hold the results
			var caseArray []Case
			
			// Datastore query
			q := datastore.NewQuery("Case").Ancestor(caseKey(ctx)) //.Filter("ID =", "5488762045857792") //.Filter("Featuring =", "featuring") //.Filter("ID=", pageRequestedVariables1) //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
			keys, err := q.GetAll(ctx, &caseArray)
			if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
			
			// ========== ========== ========== ========== ==========
			//outputCases := ""
			for i, _ := range caseArray {
				key := keys[i]
				id := int64(key.IntID())
				
				if strconv.Itoa(int(id)) == r.FormValue("existingdatastoreid") {
					// Store the key to use
					newKey = key
				}
			}
			// ========== ========== ========== ========== ==========
			//if queryKey { newKey = queryKey }
		}
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		
		_, err := datastore.Put(ctx, newKey, &caseData)
		if err != nil {
			output += "<h1>ERROR: datastore failed: "+err.Error()+"</h1>"
		} else {
			output += "<h1>SUCCESS: Created datastore entry for new case</h1>"
			output += `<h1><a href="/">Return Home</a></h1>`
			if caseData.BlobKey != "" {
				output += "<img src=\"/serve/?blobKey="+caseData.BlobKey+"\" />"
			} else { output += "<h1>No image was uploaded</h1>" }
		}
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	} else if caseData.BlobKey != "" {
		// ========== ========== ========== ========== ==========
		// Delete blobstore entry
		//deleteBlobKey := "EUG76sbkgL8CDsNUokKcRQ=="
		output += "<div>Prepare to delete from blobstore: "+caseData.BlobKey+"</div>"
		blobstore.Delete(ctx, appengine.BlobKey(caseData.BlobKey)) // https://cloud.google.com/appengine/docs/go/blobstore/reference#Delete
		output += "<div>Finished deleting from blobstore</div>"
		// ========== ========== ========== ========== ==========
	} else { output += "<h1>Form was submitted blank</h1>" }
	// ========== ========== ========== ========== ==========
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========






