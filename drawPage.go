
package main

import (
    //"fmt"
    //"os"
    "io/ioutil"
    "strings"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPage(pageRequested string) (string) {
	
	
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
	
	/* Using the pageRequested, pull the .html and .css for the selected resource else send a 404 .html/.css */
	pageRequestedData, err := ioutil.ReadFile("resources/html/"+pageRequested+".html")
	if err != nil {
		/* panic(err) */ /* else it uses "ERROR LOADING PAGE" */
		pageRequested404, err := ioutil.ReadFile("resources/html/404.html")
		if err != nil {
			/*panic(err)*/
		} else {
			stylesheetLink = `<link rel="stylesheet" type="text/css" href="resources/stylesheets/404.css" />`
			htmlContent = string(pageRequested404)
		}
	} else {
		stylesheetLink = `<link rel="stylesheet" type="text/css" href="resources/stylesheets/`+pageRequested+`.css" />`
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
    output = strings.Replace(output, "<PAGETITLE>", pageRequested, -1)
    output = strings.Replace(output, "<URL>", "http://boomcase.productionmediadesign.com/"+pageRequested, -1)
    
    
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

