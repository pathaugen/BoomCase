
package main

import (
    "fmt"
    "net/http"
	"appengine"
	"appengine/user"
)

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
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	//if u := user.Current(ctx); u != nil {
	u := user.Current(ctx)
	if u == nil {
		//g.Author = u.String()
		url, _ := user.LoginURL(ctx, "/")
		output += `
			<div style="font-size:2.0em;text-align:center;padding:2%;">
				<a href="`+url+`">Sign into BoomCase</a>
			</div>`
	} else {
		url, _ := user.LogoutURL(ctx, "/")
		output += `Welcome `+u.String()+`! (<a href="`+url+`">sign out</a>)`
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ==========
