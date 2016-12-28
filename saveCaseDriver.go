
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
	caseFrequencyLow, _		:= strconv.Atoi(r.FormValue("casefrequencylow"))
	caseFrequencyHigh, _	:= strconv.Atoi(r.FormValue("casefrequencyhigh"))
	
	caseLength, _			:= strconv.Atoi(r.FormValue("caselength"))
	caseWidth, _			:= strconv.Atoi(r.FormValue("casewidth"))
	caseHeight, _			:= strconv.Atoi(r.FormValue("caseheight"))
	caseWeight, _			:= strconv.Atoi(r.FormValue("caseweight"))
	caseBattery, _			:= strconv.Atoi(r.FormValue("casebattery"))
	caseWatts, _			:= strconv.Atoi(r.FormValue("casewatts"))
	casePrice, _			:= strconv.Atoi(r.FormValue("caseprice"))
	
	caseSold, _				:= strconv.ParseBool(r.FormValue("casesold"))
	
	caseData := Case {
		Name:				r.FormValue("casename"),
		Overview:			r.FormValue("caseoverview"),
		Featuring:			r.FormValue("casefeaturing"),
		
		//FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		FrequencyLow:		int32(caseFrequencyLow), // int32
		FrequencyHigh:		int32(caseFrequencyHigh), // int32
		
		Length:				int8(caseLength), // int8
		Width:				int8(caseWidth), // int8
		Height:				int8(caseHeight), // int8
		
		Weight:				int8(caseWeight), // int8
		Battery:			int8(caseBattery), // int8
		Notes:				r.FormValue("casenotes"),
		
		Price:				int32(casePrice), // int32
		Watts:				int16(caseWatts), // int16
		Sold:				caseSold, // bool
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	
	// Protection of blobkey from accidental manipulation
	//if blobkey != "" && blobkey != "<BLOBKEY>" { caseData.BlobKey = blobkey }
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// Create Struct - Driver
	driverCircle, _			:= strconv.ParseBool(r.FormValue("drivercircle"))
	
	driverFrequencyLow, _	:= strconv.Atoi(r.FormValue("driverfrequencylow"))
	driverFrequencyHigh, _	:= strconv.Atoi(r.FormValue("driverfrequencyhigh"))
	
	driverWeight, _			:= strconv.Atoi(r.FormValue("driverweight"))
	driverDiameter, _		:= strconv.Atoi(r.FormValue("driverdiameter"))
	driverPrice, _			:= strconv.Atoi(r.FormValue("driverprice"))
	
	driverData := Driver {
		Name:				r.FormValue("drivername"),
		Type:				r.FormValue("drivertype"),
		Circle:				driverCircle,
		
		//FrequencyResponse:	r.FormValue("driverfrequencyresponse"),
		FrequencyLow:		int32(driverFrequencyLow), // int32
		FrequencyHigh:		int32(driverFrequencyHigh), // int32
		
		Diameter:			int16(driverDiameter), // int16
		Weight:				int8(driverWeight), // int8
		Price:				int32(driverPrice), // int32
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	if caseData.Name != "" {
		output += "<h1>Proceeding with CASE creation in datastore</h1>"
		
		output += "<h1>r.FormValue(\"casename\") = ["+r.FormValue("casename")+"]</h1>"
		output += "<h1>caseData.Name = ["+caseData.Name+"]</h1>"
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		// Store Golang struct in the datastore
		
		// Use a new generated key, or replace with existing key to overwrite data
		newKey := datastore.NewIncompleteKey(ctx, "Case", caseKey(ctx)) // newKey.IntID() is 0 until .Put into the datastore
		
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
			//output += `<h1><a href="/case/`+strconv.Itoa(int(newKey.IntID()))+`">Go Directly to New Case</a></h1>` // Only works for existing cases..
			//output += `<h1><a href="/case/`+newKey.StringID()+`">Go Directly to New Case</a></h1>` // Only works for existing cases..
			output += `<h1><a href="/customize">Go to Case Selection</a></h1>`
			if caseData.BlobKey != "" {
				output += "<img src=\"/serve/?blobKey="+caseData.BlobKey+"\" />"
			} else { output += "<h1>No image was uploaded</h1>" }
		}
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	} else if driverData.Name != "" {
		output += "<h1>Proceeding with DRIVER creation in datastore</h1>"
		
		output += "<h1>r.FormValue(\"drivername\") = ["+r.FormValue("drivername")+"]</h1>"
		output += "<h1>driverData.Name = ["+driverData.Name+"]</h1>"
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		// Store Golang struct in the datastore
		
		// Use a new generated key, or replace with existing key to overwrite data
		newKey := datastore.NewIncompleteKey(ctx, "Driver", driverKey(ctx)) // newKey.IntID() is 0 until .Put into the datastore
		
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		if r.FormValue("existingdatastoreid") != "" {
			// Datastore query to get the key of existing entry to overwrite
			
			// Array to hold the results
			var driverArray []Case
			
			// Datastore query
			q := datastore.NewQuery("Driver").Ancestor(driverKey(ctx)) //.Filter("ID =", "5488762045857792") //.Filter("Featuring =", "featuring") //.Filter("ID=", pageRequestedVariables1) //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
			keys, err := q.GetAll(ctx, &driverArray)
			if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
			
			// ========== ========== ========== ========== ==========
			//outputCases := ""
			for i, _ := range driverArray {
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
		
		_, err := datastore.Put(ctx, newKey, &driverData)
		if err != nil {
			output += "<h1>ERROR: datastore failed: "+err.Error()+"</h1>"
		} else {
			output += "<h1>SUCCESS: Created datastore entry for new driver</h1>"
			output += `<h1><a href="/">Return Home</a></h1>`
			output += `<h1><a href="/customize">Go to Case Selection</a></h1>`
			if driverData.BlobKey != "" {
				output += "<img src=\"/serve/?blobKey="+driverData.BlobKey+"\" />"
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
		output += `<div><a href="/">Return Home</a></div>`
		// ========== ========== ========== ========== ==========
	} else { output += "<h1>Form was submitted blank</h1>" }
	// ========== ========== ========== ========== ==========
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========






