
package main

import (
	//"html/template"
	//"net/http"
	"time"
	
	"appengine"
	"appengine/datastore"
	//"appengine/user"
)




// ========== ========== ========== ========== ==========
// [START greeting_struct]
/*
type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}
*/
// [END greeting_struct]
// ========== ========== ========== ========== ==========



// ========== ========== ========== ========== ==========
// [START case_struct]
type Case struct {
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
}
// [END case_struct]
// ========== ========== ========== ========== ==========
// ========== ========== ========== ========== ==========
// caseKey returns the key used for all case entries.
func caseKey(ctx appengine.Context) *datastore.Key {
	// The string "default_case" here could be varied to have multiple guestbooks.
	return datastore.NewKey(ctx, "Case", "default_case", 0, nil)
}
// ========== ========== ========== ========== ==========



// ========== ========== ========== ========== ==========
// [START driver_struct]
type Driver struct {
	Name				string
	FrequencyResponse	string
	Width				string
	Price				string
	
	// image
	
	DateAdded			time.Time
}
// [END driver_struct]
// ========== ========== ========== ========== ==========





