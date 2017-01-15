



























// MIGRATED TO SAVECASEDRIVER.GO




















package main

import (
	//"html/template"
	"net/http"
	"time"
	
	"strconv"
	
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"appengine/user"
	
	//"appengine/blobstore"
	"google.golang.org/appengine/blobstore"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
//func saveCase(r *http.Request, ctx appengine.Context) (string) {
func saveCase(r *http.Request) (string) {
	output := ""
	
	ctx := appengine.NewContext(r)
	
	// ========== ========== ========== ========== ==========
	output += "<div>PRE: saveCaseBlobstore()</div>"
	blobkey := saveCaseBlobstore(r)
	//blobkey := ""
	output += "<h1>blobkey = ["+blobkey+"]</h1>"
	output += "<div>POST: saveCaseBlobstore()</div>"
	// ========== ========== ========== ========== ==========
	
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
	caseFrequencyLow, _		:= strconv.Atoi(r.FormValue("casefrequencylow"))
	caseFrequencyHigh, _	:= strconv.Atoi(r.FormValue("casefrequencyhigh"))
	
	caseLength, _			:= strconv.Atoi(r.FormValue("caselength"))
	casewidth, _			:= strconv.Atoi(r.FormValue("casewidth"))
	caseheight, _			:= strconv.Atoi(r.FormValue("caseheight"))
	caseweight, _			:= strconv.Atoi(r.FormValue("caseweight"))
	casebattery, _			:= strconv.Atoi(r.FormValue("casebattery"))
	caseprice, _			:= strconv.Atoi(r.FormValue("caseprice"))
	
	casesold, _				:= strconv.ParseBool(r.FormValue("casesold"))
	
	// Create a new Case from Struct.go
	caseData := Case {
		Name:				r.FormValue("casename"),
		Overview:			r.FormValue("caseoverview"),
		Featuring:			r.FormValue("casefeaturing"),
		
		//FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		FrequencyLow:		int32(caseFrequencyLow), // int
		FrequencyHigh:		int32(caseFrequencyHigh), // int
		
		Length:				int8(caseLength), // int
		Width:				int8(casewidth), // int
		Height:				int8(caseheight), // int
		
		Weight:				int8(caseweight), // int
		Battery:			int8(casebattery), // int
		Notes:				r.FormValue("casenotes"),
		
		Price:				int32(caseprice), // int
		
		Sold:				casesold, // bool
		
		BlobKey:			blobkey,
		
		DateAdded:			time.Now(),
	}
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	output += "<h1>r.FormValue(\"casename\") = ["+r.FormValue("casename")+"]</h1>"
	output += "<h1>caseData.Name = ["+caseData.Name+"]</h1>"
	output += saveCaseDatastore(r, caseData)
	if caseData.Name != "" {
		//output += "<div>PRE: saveCaseDatastore()</div>"
		output += saveCaseDatastore(r, caseData)
		//output += string(caseData.Name)
		//output += "<div>POST: saveCaseDatastore()</div>"
	} else if caseData.BlobKey != "" {
		// ========== ========== ========== ========== ==========
		// Delete blobstore entry
		//deleteBlobKey := "EUG76sbkgL8CDsNUokKcRQ=="
		output += "<div>Prepare to delete from blobstore: "+caseData.BlobKey+"</div>"
		blobstore.Delete(ctx, appengine.BlobKey(caseData.BlobKey)) // https://cloud.google.com/appengine/docs/go/blobstore/reference#Delete
		output += "<div>Finished deleting from blobstore</div>"
		output += `<div><a href="/">Return Home</a></div>`
		// ========== ========== ========== ========== ==========
	} else {
		output += "<h1>Form was submitted blank</h1>"
	}
	// ========== ========== ========== ========== ==========
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveCaseBlobstore(r *http.Request) (string) {
	output := ""
	
	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		/*
		serveError(ctx, w, err)
		return
		*/
		//output += "<h1>ERROR: "+err.Error()+"</h1>"
	}
	file := blobs["file"]
	
	//output += "casename: "+blobs["casename"]
	
	//output += "<div>LEN(FILE): "+string(len(file))+"</div>"
	//output += "<div>LEN(FILE): "+string(file[0].BlobKey)+"</div>"
	/*
	for f := 0; f < 10; f++ {
		output += "<div>LEN(FILE): "+string(file[f].BlobKey)+"</div>"
	}
	*/
	//output += "<div>LEN(FILE): ["+string(len(file))+"]</div>"
	
	if len(file) == 0 {
		/*
		log.Errorf(ctx, "no file uploaded")
		http.Redirect(w, r, "/", http.StatusFound)
		return
		*/
		//output += "<h1>WARNING: No image file uploaded to blobstore</h1>"
		output = ""
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
//func saveCaseDatastore(r *http.Request, ctx appengine.Context, caseData Case) (string) {
func saveCaseDatastore(r *http.Request, caseData Case) (string) {
	output := ""
	
	ctx := appengine.NewContext(r)
	
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
		if caseData.BlobKey != "" {
			output += "<img src=\"/serve/?blobKey="+caseData.BlobKey+"\" />"
		} else {
			output += "<h1>No image was uploaded</h1>"
		}
	}
	// ========== ========== ========== ========== ==========
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========






