
package main

import (
    //"fmt"
    "net/http"
    
    //"x/net/context"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	
	"appengine/blobstore"
    
    //"os"
    "io/ioutil"
    "strings"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPage(r *http.Request, ctx appengine.Context) (string) { //context.Context
	
	// ========== ========== ========== ========== ==========
	// Calculate the page that was requested
	pageRequestedString := r.URL.Path[1:]
	s := strings.Split(pageRequestedString, "/")
	pageRequested := s[0]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// Load the HTML framework template
	loadedTemplate, err := ioutil.ReadFile("resources/html/framework.html")
	templateContent := "[CONTENT]"
	if err != nil { panic(err) }
	templateContent = string(loadedTemplate)
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	stylesheetLink := ""
	htmlContent := `<div style="text-align:center;"><h1>ERROR LOADING PAGE</h1></div>`
	
	// ========== ========== ========== ========== ==========
	if pageRequested == "" { pageRequested = "home" }
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
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
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    output := templateContent
    output = strings.Replace(output, "<CONTENT>", htmlContent, -1)
    output = strings.Replace(output, "<PAGETITLE>", pageRequestedTitle, -1)
    output = strings.Replace(output, "<URL>", "http://boomcase.productionmediadesign.com/"+pageRequested, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	if pageRequested == "dashboard" {
		// [START if_user]
		if u := user.Current(ctx); u != nil {
			// g.Author = u.String()
			output = "TEST"
		} else {
			// output = "TEST"
		}
		// [END if_user]
		
	} else if pageRequested == "customize" {
		output = strings.Replace(output, "<CASE>", drawCase(r), -1)
		
    }
	if pageRequested == "dashboard" || pageRequested == "customize" {
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
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	if pageRequested == "case" || pageRequested == "404" {
		pageRequestedVariables1 := ""
		pageRequestedVariables2 := ""
		
		if len(s) > 1 { pageRequestedVariables1 = string(s[1]) }
		if len(s) > 2 { pageRequestedVariables2 = string(s[2]) }
		
	    output = strings.Replace(output, "<PAGEVARIABLES1>", pageRequestedVariables1, -1)
	    output = strings.Replace(output, "<PAGEVARIABLES2>", pageRequestedVariables2, -1)
	    
		output = strings.Replace(output, "<CASEIMAGE>", "<img src=\"/caseuploads/blankcase"+pageRequestedVariables1+".jpg\" />", -1)
		
		output = strings.Replace(output, "<CASENAME>", "Case Name", -1)
		output = strings.Replace(output, "<CASEOVERVIEW>", "Hard to find Gator Samsonite suitcase featuring a gold rimmed 15\" White Woofer and a Wide Excursion Horn.", -1)
		
		output = strings.Replace(output, "<CASELENGTH>", "18", -1)
		output = strings.Replace(output, "<CASEWIDTH>", "6", -1)
		output = strings.Replace(output, "<CASEHEIGHT>", "12", -1)
		
		output = strings.Replace(output, "<CASEWEIGHT>", "20", -1)
		output = strings.Replace(output, "<CASEBATTERY>", "18", -1)
		output = strings.Replace(output, "<CASEPRICE>", "550", -1)
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    output = strings.Replace(output, "<STYLESHEET>", stylesheetLink, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
    
    
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

