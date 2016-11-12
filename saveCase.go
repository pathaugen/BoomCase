
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
func saveCase(r *http.Request) (string) {
	output := ""
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	//ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	//output += "<div>PRE: saveCaseBlobstore()</div>"
	blobkey := saveCaseBlobstore(r)
	//output += "<div>POST: saveCaseBlobstore()</div>"
	
	
	// ========== ========== ========== ========== ==========
	// Pull the POST form fields into a Golang struct
	/*
	Name				string
	Overview			string
	Featuring			string
	FrequencyResponse	string
	
	Length				string
	Width				string
	Height				string
	
	Weight				string
	Battery				string
	Notes				string
	
	Price				string
	
	Sold				string // Mark as sold
	
	// image
	
	DateAdded			time.Time
	*/
	
	// i, err := strconv.ParseInt("-42", 10, 64)
	// i, err := strconv.Atoi("-42")
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
	
	
	output += "<div>PRE: saveCaseDatastore()</div>"
	output += saveCaseDatastore(r, caseData)
	//output += string(caseData.Name)
	output += "<div>POST: saveCaseDatastore()</div>"
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveCaseBlobstore(r *http.Request) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	//ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		/*
		serveError(ctx, w, err)
		return
		*/
		output += "<h1>ERROR: "+err.Error()+"</h1>"
	}
	file := blobs["file"]
	
	//output += "<div>LEN(FILE): "+string(len(file))+"</div>"
	//output += "<div>LEN(FILE): "+string(file[0].BlobKey)+"</div>"
	/*
	for f := 0; f < 10; f++ {
		output += "<div>LEN(FILE): "+string(file[f].BlobKey)+"</div>"
	}
	*/
	//output += "<div>LEN(FILE): ["+string(len(file))+"]</div>"
	
	if len(file) == 0 {
		output += "<h1>WARNING: No image file uploaded to blobstore</h1>"
		/*
		log.Errorf(ctx, "no file uploaded")
		http.Redirect(w, r, "/", http.StatusFound)
		return
		*/
	} else {
		//http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
		//output += "BLOBKEY: "+string(file[0].BlobKey)
		
		//output += "<h1>SUCCESS: New image in blobstore<br />BLOBKEY: "+string(file[0].BlobKey)+"</h1>"
		output = string(file[0].BlobKey)
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveCaseDatastore(r *http.Request, caseData Case) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// Store Golang struct in the datastore
	key := datastore.NewIncompleteKey(ctx, "Case", caseKey(ctx))
	_, err := datastore.Put(ctx, key, &caseData)
	if err != nil {
		/*
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		*/
		output += "<h1>ERROR: datastore failed</h1>"
	} else {
		output += "<h1>SUCCESS: Created datastore entry for new case</h1>"
		output += "<img src=\"/serve/?blobKey="+caseData.BlobKey+"\" />"
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========






