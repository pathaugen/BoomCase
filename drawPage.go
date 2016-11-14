
package main

import (
    //"fmt"
    "net/http"
    
    //"x/net/context"
    
	"appengine"
	"appengine/datastore"
	"appengine/user"
	
	"appengine/blobstore"
    
    //"os"
    "io/ioutil"
    "strings"
	"strconv"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPage(r *http.Request, ctx appengine.Context) (string) { //context.Context
	
	// Calculate the page that was requested
	pageRequestedString := r.URL.Path[1:]
	s := strings.Split(pageRequestedString, "/")
	pageRequested := s[0]
	
	// Load the HTML framework template
	loadedTemplate, err := ioutil.ReadFile("resources/html/framework.html")
	templateContent := "[CONTENT]"
	if err != nil { panic(err) }
	templateContent = string(loadedTemplate)
	
	// ========== ========== ========== ========== ==========
	// Draw Admin Bar
	// [START if_user]
	u := user.Current(ctx)
	if u != nil {
		// ========== ========== ========== ========== ==========
		// Check if user is an admin
		adminAuthorization := ""
		if adminEmails[u.String()] {
			adminAuthorization += "<div>Authorized Administrator</div>"
		} else { adminAuthorization += "<div>standard user</div>" }
		// ========== ========== ========== ========== ==========
		url, _ := user.LogoutURL(ctx, "/")
		adminBar := `
			<div style="text-align:center;background-color:yellow;font-weight:bold;padding:2%;">
				<div>Welcome <a href="/login">`+u.String()+`</a>! (<a href="`+url+`">sign out</a>)</div>
				`+adminAuthorization+`
			</div>`
	    templateContent = strings.Replace(templateContent, "<ADMINBAR>", adminBar, -1)
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	stylesheetLink := ""
	htmlContent := `<div style="text-align:center;"><h1>ERROR LOADING PAGE</h1></div>`
	
	if pageRequested == "" { pageRequested = "home" }
	pageRequestedTitle := pageRequested
	
	// Using the pageRequested, pull the .html and .css for the selected resource else send a 404 .html/.css
	pageRequestedData, err := ioutil.ReadFile("resources/html/"+pageRequested+".html")
	if err != nil {
		// panic(err) // else it uses "ERROR LOADING PAGE"
		pageRequested404, err := ioutil.ReadFile("resources/html/404.html")
		if err != nil { /* panic(err) */ } else {
			stylesheetLink += `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/404.css" />`
			htmlContent = string(pageRequested404)
			pageRequested = "404"
		}
	} else {
		stylesheetLink += `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/`+pageRequested+`.css" />`
		htmlContent = string(pageRequestedData)
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    output := templateContent
    output = strings.Replace(output, "<CONTENT>", htmlContent, -1)
    output = strings.Replace(output, "<PAGETITLE>", pageRequestedTitle, -1)
    output = strings.Replace(output, "<URL>", "http://boomcase.productionmediadesign.com/"+pageRequested, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	pageRequestedVariables1 := ""
	pageRequestedVariables2 := ""
	
	if pageRequested == "case" || pageRequested == "404" {
		
		if len(s) > 1 { pageRequestedVariables1 = string(s[1]) }
		if len(s) > 2 { pageRequestedVariables2 = string(s[2]) }
		
	    output = strings.Replace(output, "<PAGEVARIABLES1>", pageRequestedVariables1, -1)
	    output = strings.Replace(output, "<PAGEVARIABLES2>", pageRequestedVariables2, -1)
	    
	    // drawPageCase.go
		if pageRequested == "case" { output = drawPageCase(ctx, output, pageRequestedVariables1) }
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	if pageRequested == "dashboard" {
	} else if pageRequested == "customize" {
		// drawPageCustomize.go
		output = drawPageCustomize(ctx, output)
    }
	if pageRequested == "dashboard" || pageRequested == "customize" || pageRequested == "case" {
		// ========== ========== ========== ========== ==========
		// [START if_user]
		u := user.Current(ctx)
		if u != nil && adminEmails[u.String()] {
			// ========== ========== ========== ========== ==========
			// Load in modular forms - usable throughout the webapp
			formCase, _ := ioutil.ReadFile("resources/html/formcase.html")
		    output = strings.Replace(output, "<FORMCASE>", string(formCase), -1)
		    
			formDriver, _ := ioutil.ReadFile("resources/html/formdriver.html")
		    output = strings.Replace(output, "<FORMDRIVER>", string(formDriver), -1)
			// ========== ========== ========== ========== ==========
			
			// ========== ========== ========== ========== ==========
			// Generation of URL to save Case or Driver
			uploadURLCase, err := blobstore.UploadURL(ctx, "/savecasedriver", nil)
			if err != nil { /* serveError(ctx, w, err); return */ } else { output = strings.Replace(output, "<FORMACTIONCASEDRIVER>", uploadURLCase.String(), -1) }
			// ========== ========== ========== ========== ==========
			
			// Stylesheet for the case/driver form
			stylesheetLink += `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/formcasedriver.css" />`
			
			// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
			// Fill in the form with values for editing, else eliminate the values
			valueCaseName				:= ""
			valueCaseOverview			:= ""
			valueCaseFeaturing			:= ""
			valueCaseFrequencyresponse	:= ""
			valueCaseLength				:= ""
			valueCaseWidth				:= ""
			valueCaseHeight				:= ""
			valueCaseWeight				:= ""
			valueCaseBattery			:= ""
			valueCaseNotes				:= ""
			valueCasePrice				:= ""
			valueCaseSold				:= ""
			
			if pageRequested == "case" {
				
				// Array to hold the results
				var caseArray []Case
				
				// Datastore query
				q := datastore.NewQuery("Case") //.Filter("Featuring =", "featuring") //.Filter("ID=", pageRequestedVariables1) //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
				keys, err := q.GetAll(ctx, &caseArray)
				if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
				
				// ========== ========== ========== ========== ==========
				for i, c := range caseArray {
					key := keys[i]
					id := uint64(key.IntID())
					
					if strconv.Itoa(int(id)) == pageRequestedVariables1 {
						// Populate the form with current values from datastore
						valueCaseName				= c.Name
						valueCaseOverview			= c.Overview
						valueCaseFeaturing			= c.Featuring
						valueCaseFrequencyresponse	= c.FrequencyResponse
						valueCaseLength				= strconv.Itoa(int(c.Length))
						valueCaseWidth				= strconv.Itoa(int(c.Width))
						valueCaseHeight				= strconv.Itoa(int(c.Height))
						valueCaseWeight				= strconv.Itoa(int(c.Weight))
						valueCaseBattery			= strconv.Itoa(int(c.Battery))
						valueCaseNotes				= c.Notes
						valueCasePrice				= strconv.Itoa(int(c.Price))
						if c.Sold { valueCaseSold = "checked" } else { valueCaseSold = "" }
					}
				}
				// ========== ========== ========== ========== ==========
			}
			// Replace the HTML placeholders with blank values, or ones from datastore
		    output = strings.Replace(output, "<VALUECASENAME>",					valueCaseName, -1)
		    output = strings.Replace(output, "<VALUECASEOVERVIEW>",				valueCaseOverview, -1)
		    output = strings.Replace(output, "<VALUECASEFEATURING>",			valueCaseFeaturing, -1)
		    output = strings.Replace(output, "<VALUECASEFREQUENCYRESPONSE>",	valueCaseFrequencyresponse, -1)
		    output = strings.Replace(output, "<VALUECASELENGTH>",				valueCaseLength, -1)
		    output = strings.Replace(output, "<VALUECASEWIDTH>",				valueCaseWidth, -1)
		    output = strings.Replace(output, "<VALUECASEHEIGHT>",				valueCaseHeight, -1)
		    output = strings.Replace(output, "<VALUECASEWEIGHT>",				valueCaseWeight, -1)
		    output = strings.Replace(output, "<VALUECASEBATTERY>",				valueCaseBattery, -1)
		    output = strings.Replace(output, "<VALUECASENOTES>",				valueCaseNotes, -1)
		    output = strings.Replace(output, "<VALUECASEPRICE>",				valueCasePrice, -1)
		    output = strings.Replace(output, "<VALUECASESOLD>",					valueCaseSold, -1)
			// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		} else {
			url, _ := user.LoginURL(ctx, "/")
			loginLink := `
				<div style="font-size:2.0em;text-align:center;padding:2%;">
					<a href="`+url+`">Sign into BoomCase</a>
				</div>`
		    output = strings.Replace(output, "<FORMCASE>", loginLink, -1)
		    output = strings.Replace(output, "<FORMDRIVER>", loginLink, -1)
		}
		// [END if_user]
		// ========== ========== ========== ========== ==========
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    output = strings.Replace(output, "<STYLESHEET>", stylesheetLink, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

