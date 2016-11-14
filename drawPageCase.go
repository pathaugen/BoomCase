
package main

import (
    //"fmt"
    //"net/http"
    
    //"x/net/context"
    
	"appengine"
	"appengine/datastore"
	"appengine/user"
	
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
			
			// ========== ========== ========== ========== ==========
			// Utilizing the requested case id, pull the image key for display via blobstore
			//output = strings.Replace(output, "<CASEIMAGE>", "<img src=\"/caseuploads/blankcase"+pageRequestedVariables1+".jpg\" />", -1)
			output = strings.Replace(output, "<CASEIMAGE>", `<img src="/serve/?blobKey=`+blobkey+`" />`, -1)
			// ========== ========== ========== ========== ==========
			
			output = strings.Replace(output, "<CASENAME>",				c.Name, -1)
			output = strings.Replace(output, "<CASEOVERVIEW>",			c.Overview, -1)
			
			output = strings.Replace(output, "<CASEFEATURING>",			c.Featuring, -1)
			output = strings.Replace(output, "<CASEFREQUENCYRESPONSE>",	c.FrequencyResponse, -1)
			
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
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	u := user.Current(ctx)
	if u != nil && adminEmails[u.String()] {
		
		formDriverButton := `
			<div id="page-formdriver-button" style="clear:both;">
				<a href="" class="btn btn-primary" id="admin-edit-driver" aria-label="ADMIN: Add Driver">
					<i class="fa fa-plus-circle" aria-hidden="true"></i> ADMIN: Add Driver
				</a>
			</div>`
		output = strings.Replace(output, "<FORMDRIVER>", formDriverButton+"<FORMDRIVER>", -1)
		
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

