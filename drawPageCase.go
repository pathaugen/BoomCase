
package main

import (
    //"fmt"
    //"net/http"
    
    //"x/net/context"
    
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"appengine/image"
	
	//"appengine/blobstore"
    
    //"os"
    //"io/ioutil"
    "strings"
	"strconv"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPageCase(ctx appengine.Context, output string, pageRequestedVariables1 string) (string, blobkey string) {
	
	//blobkey := ""
	
	// ========== ========== ========== ========== ==========
	// Utilizing the requested case number, display the case values from database
	
	// Array to hold the results
	var caseArray []Case
	
	// Datastore query
	q := datastore.NewQuery("Case").Ancestor(caseKey(ctx)) //.Filter("ID =", "5488762045857792") //.Filter("Featuring =", "featuring") //.Filter("ID=", pageRequestedVariables1) //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
	keys, err := q.GetAll(ctx, &caseArray)
	if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
	
	//k := datastore.NewKey(ctx, "Case", pageRequestedVariables1, 0, nil)
	//c := new(Case)
	//var c Case
	//if err := datastore.Get(ctx, k, c); err != nil { /* http.Error(w, err.Error(), 500); return */ }
	
	//caseKey := datastore.NewKey(ctx, "Case", pageRequestedVariables2, 0, nil)
	//addressKey := datastore.NewKey(ctx, "Address", "", 1, employeeKey)
	//var caseInfo Case
	//err = datastore.Get(ctx, caseKey, &caseInfo)
	//if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
	
	// ========== ========== ========== ========== ==========
	//outputCases := ""
	for i, c := range caseArray {
		key := keys[i]
		id := int64(key.IntID())
		
		if strconv.Itoa(int(id)) == pageRequestedVariables1 {
			
			blobkey = c.BlobKey
			
			// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
			// Utilizing the requested case id, pull the image key for display via blobstore
			//output = strings.Replace(output, "<CASEIMAGE>", "<img src=\"/caseuploads/blankcase"+pageRequestedVariables1+".jpg\" />", -1)
			// ========== ========== ========== ========== ==========
			// Get a thumbnail of a blobstore image
			// https://github.com/golang/appengine/blob/master/image/image.go
			thumbOpts := image.ServingURLOptions { Size: 800, }
			thumbKey := appengine.BlobKey(blobkey)
			thumbnail, _ := image.ServingURL(ctx, thumbKey, &thumbOpts) // (*url.URL, error)
			// ========== ========== ========== ========== ==========
			//output = strings.Replace(output, "<CASEIMAGE>", `<img src="/serve/?blobKey=`+blobkey+`" />`, -1)
			output = strings.Replace(output, "<CASEIMAGE>", `<img src="`+thumbnail.String()+`" />`, -1)
			// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
			
			output = strings.Replace(output, "<CASENAME>",				c.Name, -1)
			output = strings.Replace(output, "<CASEOVERVIEW>",			c.Overview, -1)
			
			output = strings.Replace(output, "<CASEFEATURING>",			c.Featuring, -1)
			
			//output = strings.Replace(output, "<CASEFREQUENCYRESPONSE>",	c.FrequencyResponse, -1)
			output = strings.Replace(output, "<CASEFREQUENCYLOW>",		strconv.Itoa(int(c.FrequencyLow)), -1)
			output = strings.Replace(output, "<CASEFREQUENCYHIGH>",		strconv.Itoa(int(c.FrequencyHigh)), -1)
			
			output = strings.Replace(output, "<CASELENGTH>",			strconv.Itoa(int(c.Length)), -1)
			output = strings.Replace(output, "<CASEWIDTH>",				strconv.Itoa(int(c.Width)), -1)
			output = strings.Replace(output, "<CASEHEIGHT>",			strconv.Itoa(int(c.Height)), -1)
			
			output = strings.Replace(output, "<CASEWEIGHT>",			strconv.Itoa(int(c.Weight)), -1)
			output = strings.Replace(output, "<CASEBATTERY>",			strconv.Itoa(int(c.Battery)), -1)
			
			output = strings.Replace(output, "<CASENOTES>",				c.Notes, -1)
			
			output = strings.Replace(output, "<CASEPRICE>",				strconv.Itoa(int(c.Price)), -1)
		}
	}
	// ========== ========== ========== ========== ==========
	
	
	
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	// Add all the drivers from datastore
	
	var driverTemplate = ""
	
	// Array to hold the results
	var driverArray []Driver
	
	// Datastore query
	qDriver := datastore.NewQuery("Driver").Ancestor(driverKey(ctx)) //.Filter("ID =", "5488762045857792") //.Filter("Featuring =", "featuring") //.Filter("ID=", pageRequestedVariables1) //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
	keysDriver, err := qDriver.GetAll(ctx, &driverArray)
	if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
	
	for i, c := range driverArray {
		key := keysDriver[i]
		id := int64(key.IntID())
		
		blobkey = c.BlobKey
		
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		// Utilizing the requested driver id, pull the image key for display via blobstore
		// ========== ========== ========== ========== ==========
		// Get a thumbnail of a blobstore images
		// https://github.com/golang/appengine/blob/master/image/image.go
		thumbOpts := image.ServingURLOptions { Size: 800, }
		thumbKey := appengine.BlobKey(blobkey)
		thumbnail, _ := image.ServingURL(ctx, thumbKey, &thumbOpts) // (*url.URL, error)
		// ========== ========== ========== ========== ==========
		// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
		
		// Placeholders until these variables are coded in from datastore
		//var driverType = "low"
		//var driverInches = "12"
		
		driverTemplate += `
			<a href="" class="driver-info driver-info-`+c.Type+`" data-size="`+strconv.Itoa(int(c.Diameter))+`" data-type="`+c.Type+`">
				<!-- `+strconv.Itoa(int(id))+` -->
				<img src="`+thumbnail.String()+`" />
				<span class="name-container">`+c.Name+`</span>
				<span class="inch-container"><span class="size">`+strconv.Itoa(int(c.Diameter))+`</span>"</span>
				<span class="frequency-container"><span class="frequency">`+strconv.Itoa(int(c.FrequencyLow))+`hz - `+strconv.Itoa(int(c.FrequencyHigh))+`hz</span></span>
				<span class="price-container">+ $<span class="price">`+strconv.Itoa(int(c.Price))+`</span>/each</span>
				<span class="add-container"><i class="fa fa-plus-circle" aria-hidden="true"></i></span>
			</a>
		`
	}
	// TODO: String replace out drivers
	// output = strings.Replace(output, "<DRIVERS>", driverTemplate, -1)
	// output += driverTemplate
	output = strings.Replace(output, "<DRIVERS>", driverTemplate, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	u := user.Current(ctx)
	if u != nil && adminEmails[u.Email] {
		
		formDriverAddButton := `
			<div id="page-formdriver-button" style="clear:both;">
				<a href="" class="btn btn-primary" id="admin-add-driver" aria-label="ADMIN: Add Driver">
					<i class="fa fa-plus-circle" aria-hidden="true"></i> ADMIN: Add Driver
				</a>
			</div>`
		output = strings.Replace(output, "<FORMIMAGE>", formDriverAddButton+"<FORMIMAGE>", -1)
		
		/*
		formDriverEditButton := `
			<span class="page-formdriver-edit-button">
				<a href="" class="btn btn-primary" id="admin-add-driver" aria-label="Edit Driver">
					<i class="fa fa-plus-circle" aria-hidden="true"></i> Edit Driver
				</a>
			</span>`
		output = strings.Replace(output, "<FORMDRIVER>", formDriverEditButton+"<FORMDRIVER>", -1)
		*/
		
		formCaseButton := `
			<div id="page-formcase-button" style="clear:both;">
				<a href="" class="btn btn-primary" id="admin-edit-case" aria-label="ADMIN: Edit This Case">
					<i class="fa fa-plus-circle" aria-hidden="true"></i> ADMIN: Edit This Case
				</a>
			</div>`
		output = strings.Replace(output, "<FORMCASE>", formCaseButton+"<FORMCASE>", -1)
		
	} else {
		output = strings.Replace(output, "<FORMDRIVER>", "", -1)
		output = strings.Replace(output, "<FORMCASE>", "", -1)
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
    
    
    return output, blobkey
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

