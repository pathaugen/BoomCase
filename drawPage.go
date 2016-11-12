
package main

import (
    //"fmt"
    "net/http"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	
	"appengine/blobstore"
    
    //"os"
    "io/ioutil"
    "strings"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
//func drawPage(pageRequestedString string) (string) {
func drawPage(r *http.Request) (string) {
	pageRequestedString := r.URL.Path[1:]
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	s := strings.Split(pageRequestedString, "/")
	//pageRequested, pageRequestedData := s[0], s[1]
	
	/*
	pageRequested := ""
	if string(s[0]) != "" {
		pageRequested = s[0]
	} else {
		pageRequested = pageRequestedString
	}
	*/
	pageRequested := s[0]
	
	// Load template
	//loadedTemplate := ReadFile("resources/html/framework.html", string) ([]byte, error)
	loadedTemplate, err := ioutil.ReadFile("resources/html/framework.html")
	templateContent := "[CONTENT]"
	if err != nil {
		panic(err)
	}
	templateContent = string(loadedTemplate)
	
	
	// ========== ========== ========== ========== ==========
	stylesheetLink := ""
	htmlContent := `<div style="text-align:center;"><h1>ERROR LOADING PAGE</h1></div>`
	
	if pageRequested == "" {
		pageRequested = "home"
	}
	
	pageRequestedTitle := pageRequested
	//pageRequested404 := ""
	/* Using the pageRequested, pull the .html and .css for the selected resource else send a 404 .html/.css */
	pageRequestedData, err := ioutil.ReadFile("resources/html/"+pageRequested+".html")
	if err != nil {
		/* panic(err) */ /* else it uses "ERROR LOADING PAGE" */
		pageRequested404, err := ioutil.ReadFile("resources/html/404.html")
		if err != nil {
			/*panic(err)*/
		} else {
			stylesheetLink = `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/404.css" />`
			htmlContent = string(pageRequested404)
			//pageRequestedTitle = "404"
			pageRequested = "404"
		}
	} else {
		stylesheetLink = `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/`+pageRequested+`.css" />`
		htmlContent = string(pageRequestedData)
	}
	
	/*
	if pageRequested == "" {
		htmlContent = `
			<div style="text-align:center;">
				<h1>TESTING 222</h1>
				<h1><a href="customize">Boomcase Customize</a></h1>
				<h1><a href="boombarrel">Boom Barrel</a></h1>
			</div>
		`
	    //fmt.Fprint(htmlContent)
	} else if pageRequested == "customize" {
		loadedCustomize, err := ioutil.ReadFile("resources/html/customize.html")
		if err != nil {
			panic(err)
		}
		stylesheetLink = `<link rel="stylesheet" type="text/css" href="resources/stylesheets/customize.css" />`
		htmlContent = string(loadedCustomize)
	} else if pageRequested == "boombarrel" {
		loadedBoombarrel, err := ioutil.ReadFile("resources/html/boombarrel.html")
		if err != nil {
			panic(err)
		}
		stylesheetLink = `<link rel="stylesheet" type="text/css" href="resources/stylesheets/boombarrel.css" />`
		htmlContent = string(loadedBoombarrel)
	} else {
		htmlContent = "PAGE COULD NOT BE FOUND: "+pageRequested
	}
	*/
	
    // ========== ========== ========== ========== ==========
    
    
    output := templateContent
    //output = strings.Replace(htmlContent, "Home", "TEST", -1)
    output = strings.Replace(output, "<STYLESHEET>", stylesheetLink, -1)
    output = strings.Replace(output, "<CONTENT>", htmlContent, -1)
    //output = strings.Replace(output, "<DELETECASE>", `<span>Delete Item</span>`, -1)
    output = strings.Replace(output, "<PAGETITLE>", pageRequestedTitle, -1)
    output = strings.Replace(output, "<URL>", "http://boomcase.productionmediadesign.com/"+pageRequested, -1)
    
    
	// ========== ========== ========== ========== ==========
	if pageRequested == "dashboard" {
		// [START if_user]
		if u := user.Current(ctx); u != nil {
			//g.Author = u.String()
			output = "TEST"
		} else {
			//output = "TEST"
		}
		// [END if_user]
		
		uploadURL, err := blobstore.UploadURL(ctx, "/savecase", nil)
		if err != nil {
			/*
			serveError(ctx, w, err)
			return
			*/
		} else {
		    output = strings.Replace(output, "<FORMACTION>", uploadURL.String(), -1)
		}
	}
	// ========== ========== ========== ========== ==========
	
    
    if pageRequested == "customize" {
		output = strings.Replace(output, "<CASE>", drawCase(r), -1)
    }
    
    
	if pageRequested == "case" || pageRequested == "404" {
		pageRequestedVariables1 := ""
		pageRequestedVariables2 := ""
		//if string(s[1]) != "" {
		if len(s) > 1 {
			pageRequestedVariables1 = string(s[1])
		}
		if len(s) > 2 {
			pageRequestedVariables2 = string(s[2])
		}
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
    
    
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

