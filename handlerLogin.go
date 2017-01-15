
package main

import (
    "fmt"
    "net/http"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// ========== ========== ========== ========== ==========
// Define admin authorized emails in a map
var adminEmails = map[string]bool {
	"pathaugen": true,					// Patrick Haugen
	"pathaugen@gmail.com": true,		// Patrick Haugen
	"DominicOdbert": true,				// Dominic Odbert
	"dominicodbert@gmail.com": true,	// Dominic Odbert
	"rex00x@yahoo.com": true,			// JP Odbert
}
// ========== ========== ========== ========== ==========

// ========== ========== ========== ========== ==========
func handlerLogin(w http.ResponseWriter, r *http.Request) {
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	output := ""
	output += `<div style="font-size:2.0em;text-align:center;padding:2%;">`
	output += `<h1>BoomCase Login</h1>`
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	//if u := user.Current(ctx); u != nil {
	u := user.Current(ctx)
	if u == nil {
		//g.Author = u.Email
		url, _ := user.LoginURL(ctx, "/")
		output += `<div><a href="`+url+`">Sign into BoomCase</a></div>`
	} else {
		url, _ := user.LogoutURL(ctx, "/")
		output += `<div>Welcome `+u.Email+`! (<a href="`+url+`">sign out</a>)</div>`
		
		// ========== ========== ========== ========== ==========
		// Check if user is an admin
		if adminEmails[u.Email] {
			output += `<div><h1>You are an admin!</h1></div>`
		} else { output += `<div>standard user - NOT an admin</div>` }
		// ========== ========== ========== ========== ==========
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
	output += `<div><a href="/">BoomCase Home</a></div>`
	output += "</div>"
	
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ==========
