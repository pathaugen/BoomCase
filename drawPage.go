
package main

import (
    //"fmt"
    "net/http"

    "golang.org/x/net/context"

	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"

	"google.golang.org/appengine/blobstore"

    //"os"
    "io/ioutil"
    "strings"
	"strconv"
)

// ========== START: drawPage ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPage(r *http.Request, ctx context.Context) (string) { // context.Context vs appengine.Context

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
		if adminEmails[u.Email] {
			adminAuthorization += "<div>Authorized Administrator</div>"
		} else { adminAuthorization += "<div>standard user</div>" }
		// ========== ========== ========== ========== ==========
		url, _ := user.LogoutURL(ctx, "/")
		adminBar := `
			<div style="text-align:center;background-color:yellow;font-weight:bold;padding:2%;">
				<div>Welcome <a href="/login">`+u.Email+`</a>! (<a href="`+url+`">sign out</a>)</div>
				`+adminAuthorization+`
			</div>`
	    templateContent = strings.Replace(templateContent, "<ADMINBAR>", adminBar, -1)
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========




	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	stylesheetLink := ""
	htmlContent := `<div style="text-align:center;"><h1>ERROR LOADING PAGE</h1></div>`

	//if pageRequested == "" { pageRequested = "home" }
  if pageRequested == "admin" { pageRequested = "home" } // "home" is the original dashboard page of selections
  if pageRequested == "" { pageRequested = "customize" }
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
  output = strings.Replace(output, "<CONTENT>",   htmlContent, -1)
  output = strings.Replace(output, "<PAGETITLE>", pageRequestedTitle, -1)
  output = strings.Replace(output, "<URL>",       "http://boomcase.productionmediadesign.com/"+pageRequested, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========




	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	pageRequestedVariables1 := ""
	pageRequestedVariables2 := ""
	blobkey := ""

	if pageRequested == "case" || pageRequested == "404" {
		if len(s) > 1 { pageRequestedVariables1 = string(s[1]) }
		if len(s) > 2 { pageRequestedVariables2 = string(s[2]) }

    output = strings.Replace(output, "<PAGEVARIABLES1>", pageRequestedVariables1, -1)
    output = strings.Replace(output, "<PAGEVARIABLES2>", pageRequestedVariables2, -1)

    // drawPageCase.go
		if pageRequested == "case" { output, blobkey = drawPageCase(ctx, output, pageRequestedVariables1) }
	}
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========




	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	if pageRequested == "dashboard" {
		// ========== ========== ========== ========== ==========
		// If blobkey query string is set, display the image from the datastore
		blobkey = r.FormValue("blobkey")
		blobkeyHtml := ""

		if blobkey != "" {
			blobkeyHtml = `
				<div style="">
					<h1>Image Successfully Uploaded!</h1>
					<div><img src="/serve/?blobKey=`+blobkey+`" style="width:500px;" /></div>
					Now add this image as a Custom Case or Driver Option by selecting an option.
				</div>
			`
		} else {
			// if blobkey doesn't exist, replace editing forms with image upload forms
  		u := user.Current(ctx)
      if u != nil && adminEmails[u.Email] {
  			output = strings.Replace(output, "<FORMCASE>", "<FORMIMAGE>", -1)
  			output = strings.Replace(output, "<FORMDRIVER>", "<FORMIMAGE>", -1)
      } else if u != nil && !adminEmails[u.Email] {
        notAuthorized := `<div><h2>WARNING:<br />Your email address<br />(<a href="/login">`+u.Email+`</a>)<br />is NOT authorized to add/edit cases or drivers!</h2></div>`
  			output = strings.Replace(output, "<FORMCASE>", notAuthorized+"<FORMCASE>", -1)
  			output = strings.Replace(output, "<FORMDRIVER>", notAuthorized+"<FORMDRIVER>", -1)
      } else {}
		}

		output = strings.Replace(output, "<BLOBKEYIMAGE>", blobkeyHtml, -1)
		// ========== ========== ========== ========== ==========
	} else if pageRequested == "customize" {
		// drawPageCustomize.go
		output = drawPageCustomize(ctx, output)
  }
	if pageRequested == "dashboard" || pageRequested == "customize" || pageRequested == "case" {
		// ========== ========== ========== ========== ==========
		// [START if_user]
		u := user.Current(ctx)
		if u != nil && adminEmails[u.Email] {
			// ========== ========== ========== ========== ==========
			// Load in modular forms - usable throughout the webapp
			formCase, _ := ioutil.ReadFile("resources/html/formcase.html")
	    output = strings.Replace(output, "<FORMCASE>", string(formCase), -1)

			formDriver, _ := ioutil.ReadFile("resources/html/formdriver.html")
	    output = strings.Replace(output, "<FORMDRIVER>", string(formDriver), -1)

			formImage, _ := ioutil.ReadFile("resources/html/formimage.html")
	    output = strings.Replace(output, "<FORMIMAGE>", string(formImage), -1)
			// ========== ========== ========== ========== ==========

			// ========== ========== ========== ========== ==========
			// Generation of URL to save Case or Driver
			uploadURLCaseDriver, err := blobstore.UploadURL(ctx, "/savecasedriver", nil)
			if err != nil { /* serveError(ctx, w, err); return */ } else { output = strings.Replace(output, "<FORMACTIONCASEDRIVER>", uploadURLCaseDriver.String(), -1) }

			uploadURLImage, err := blobstore.UploadURL(ctx, "/saveimage", nil)
			if err != nil { /* serveError(ctx, w, err); return */ } else { output = strings.Replace(output, "<FORMACTIONIMAGE>", uploadURLImage.String(), -1) }
			// ========== ========== ========== ========== ==========

			// Slip the blobkey into the forms
			//if r.FormValue("blobkey") != "" { output = strings.Replace(output, "<BLOBKEY>", r.FormValue("blobkey"), -1) }
			if blobkey != "" { output = strings.Replace(output, "<CASEBLOBKEY>", blobkey, -1) }
			if blobkey != "" { output = strings.Replace(output, "<DRIVERBLOBKEY>", blobkey, -1) }

			// Stylesheet for the case/driver form
			stylesheetLink += `<link rel="stylesheet" type="text/css" href="/resources/stylesheets/formcasedriverimage.css" />`

			// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
			// Fill in the form with values for editing, else eliminate the values
			valueExistingDataStoreId   := "" // Image to use - hidden value


			// Case Specific Variables
			valueCaseName              := ""
			valueCaseOverview          := ""
			valueCaseFeaturing         := ""

			//valueCaseFrequencyresponse	:= "" // Old variable before low/high
			valueCaseFrequencyLow      := ""
			valueCaseFrequencyHigh     := ""

			valueCaseLength            := ""
			valueCaseWidth             := ""
			valueCaseHeight            := ""
			valueCaseWeight            := ""
			valueCaseBattery           := ""
			valueCaseNotes             := ""
			valueCasePrice             := ""
			valueCaseSold              := ""
			valueCaseDriverMultiplier  := ""


			// Driver Specific Variables
			valueDriverName            := ""

			valueDriverTypeLow         := ""
			valueDriverTypeMid         := ""
			valueDriverTypeHigh        := ""

			valueDriverDiameter        := "" // Multiply these inches x 100 for base, then by case multiplier to be exact

			//valueDriverFrequencyresponse	:= "" // Old variable before low/high
			valueDriverFrequencyLow    := ""
			valueDriverFrequencyHigh   := ""

			valueDriverWeight          := ""
			valueDriverPrice           := ""
			valueDriverCircle          := ""

			// ========== ========== ========== ========== ==========
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
					id := int64(key.IntID())

					if strconv.Itoa(int(id)) == pageRequestedVariables1 {
						// c is the variable with data from the datastore (check the struct.go file)

						// Populate the form with current values from datastore
						valueExistingDataStoreId	= strconv.Itoa(int(id))
						valueCaseName           = c.Name
						valueCaseOverview       = c.Overview
						valueCaseFeaturing      = c.Featuring

						//valueCaseFrequencyresponse	= c.FrequencyResponse
						valueCaseFrequencyLow		= strconv.Itoa(int(c.FrequencyLow))
						valueCaseFrequencyHigh  = strconv.Itoa(int(c.FrequencyHigh))

						valueCaseLength         = strconv.Itoa(int(c.Length))
						valueCaseWidth          = strconv.Itoa(int(c.Width))
						valueCaseHeight         = strconv.Itoa(int(c.Height))
						valueCaseWeight         = strconv.Itoa(int(c.Weight))
						valueCaseBattery        = strconv.Itoa(int(c.Battery))
						valueCaseNotes          = c.Notes
						valueCasePrice          = strconv.Itoa(int(c.Price))
						if c.Sold { valueCaseSold = "checked" } else { valueCaseSold = "" }

						//valueCaseDriverMultiplier	= strconv.Itoa(int(c.DriverMultiplier))
						valueCaseDriverMultiplier	= c.DriverMultiplier
					}
				}
				// ========== ========== ========== ========== ==========
			}
			// ========== ========== ========== ========== ==========

			// Replace the HTML placeholders with blank values, or ones from datastore
	    output = strings.Replace(output, "<EXISTINGDATASTOREID>",        valueExistingDataStoreId, -1)


	    // Case Specific Values
	    output = strings.Replace(output, "<VALUECASENAME>",              valueCaseName, -1)
	    output = strings.Replace(output, "<VALUECASEOVERVIEW>",          valueCaseOverview, -1)
	    output = strings.Replace(output, "<VALUECASEFEATURING>",         valueCaseFeaturing, -1)

	    //output = strings.Replace(output, "<VALUECASEFREQUENCYRESPONSE>",	valueCaseFrequencyresponse, -1)
	    output = strings.Replace(output, "<VALUECASEFREQUENCYLOW>",      valueCaseFrequencyLow, -1)
	    output = strings.Replace(output, "<VALUECASEFREQUENCYHIGH>",     valueCaseFrequencyHigh, -1)

	    output = strings.Replace(output, "<VALUECASELENGTH>",            valueCaseLength, -1)
	    output = strings.Replace(output, "<VALUECASEWIDTH>",             valueCaseWidth, -1)
	    output = strings.Replace(output, "<VALUECASEHEIGHT>",            valueCaseHeight, -1)
	    output = strings.Replace(output, "<VALUECASEWEIGHT>",            valueCaseWeight, -1)
	    output = strings.Replace(output, "<VALUECASEBATTERY>",           valueCaseBattery, -1)
	    output = strings.Replace(output, "<VALUECASENOTES>",				     valueCaseNotes, -1)
	    output = strings.Replace(output, "<VALUECASEPRICE>",				     valueCasePrice, -1)
	    output = strings.Replace(output, "<VALUECASESOLD>",					     valueCaseSold, -1)
	    output = strings.Replace(output, "<VALUECASEDRIVERMULTIPLIER>",  valueCaseDriverMultiplier, -1)


	    // Driver Specific Values
	    output = strings.Replace(output, "<VALUEDRIVERNAME>",            valueDriverName, -1)

	    output = strings.Replace(output, "<VALUEDRIVERTYPELOW>",         valueDriverTypeLow, -1)
	    output = strings.Replace(output, "<VALUEDRIVERTYPEMID>",         valueDriverTypeMid, -1)
	    output = strings.Replace(output, "<VALUEDRIVERTYPEHIGH>",			   valueDriverTypeHigh, -1)

	    output = strings.Replace(output, "<VALUEDRIVERDIAMETER>",        valueDriverDiameter, -1)

	    //output = strings.Replace(output, "<VALUEDRIVERFREQUENCYRESPONSE>",	valueDriverFrequencyresponse, -1)
	    output = strings.Replace(output, "<VALUEDRIVERFREQUENCYLOW>",    valueDriverFrequencyLow, -1)
	    output = strings.Replace(output, "<VALUEDRIVERFREQUENCYHIGH>",   valueDriverFrequencyHigh, -1)

	    output = strings.Replace(output, "<VALUEDRIVERWEIGHT>",          valueDriverWeight, -1)
	    output = strings.Replace(output, "<VALUEDRIVERPRICE>",           valueDriverPrice, -1)
	    output = strings.Replace(output, "<VALUEDRIVERCIRCLE>",          valueDriverCircle, -1)
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
// ========== END: drawPage ========== ========== ========== ========== ========== ========== ========== ========== ==========
