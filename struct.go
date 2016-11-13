
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
	
	Length				int
	Width				int
	Height				int
	
	Weight				int
	Battery				int
	Notes				string
	
	Price				int
	
	Sold				bool // Mark as sold
	
	BlobKey				string
	
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
	Diameter			int
	Price				int
	
	BlobKey				string
	
	DateAdded			time.Time
}
// [END driver_struct]
// ========== ========== ========== ========== ==========
// ========== ========== ========== ========== ==========
// driverKey returns the key used for all case entries.
func driverKey(ctx appengine.Context) *datastore.Key {
	// The string "default_case" here could be varied to have multiple guestbooks.
	return datastore.NewKey(ctx, "Driver", "default_driver", 0, nil)
}
// ========== ========== ========== ========== ==========





