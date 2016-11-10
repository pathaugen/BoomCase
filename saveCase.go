
package main

import (
	//"html/template"
	"net/http"
	"time"
	
	"appengine"
	"appengine/datastore"
	//"appengine/user"
)




// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func saveCase(r *http.Request) (string) {
	output := ""
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	//output += "TEST111"
	//output += "TEST222"
	
	
	// ========== ========== ========== ========== ==========
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
	caseData := Case {
		Name:				r.FormValue("casename"),
		Overview:			r.FormValue("caseoverview"),
		Featuring:			r.FormValue("casefeaturing"),
		FrequencyResponse:	r.FormValue("casefrequencyresponse"),
		
		Length:				r.FormValue("caselength"),
		Width:				r.FormValue("casewidth"),
		Height:				r.FormValue("caseheight"),
		
		Weight:				r.FormValue("caseweight"),
		Battery:			r.FormValue("casebattery"),
		Notes:				r.FormValue("casenotes"),
		
		Price:				r.FormValue("caseprice"),
		
		Sold:				r.FormValue("casesold"),
		
		DateAdded:			time.Now(),
	}
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	key := datastore.NewIncompleteKey(ctx, "Case", caseKey(ctx))
	_, err := datastore.Put(ctx, key, &caseData)
	if err != nil {
		/*
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		*/
	}
	// ========== ========== ========== ========== ==========
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

