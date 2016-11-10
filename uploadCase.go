
package main

import (
	//"html/template"
	//"net/http"
	"time"
	
	//"appengine"
	//"appengine/datastore"
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
	
	Length				string
	Width				string
	Height				string
	
	Weight				string
	Battery				string
	Notes				string
	
	Price				string
	
	Sold				string // Mark as sold
	
	// image
	
	DateAdded			time.Time
}
// [END case_struct]
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// [START driver_struct]
type Driver struct {
	Name				string
	FrequencyResponse	string
	Width				string
	Price				string
	
	// image
	
	DateAdded			time.Time
}
// [END driver_struct]
// ========== ========== ========== ========== ==========




// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func uploadCase() (string) {
	output := ""
	
	
	output = `
		<a href="/case/1" class="case-block">
			<img src="/caseuploads/blankcase1.jpg" />
			<span class="name-container">Case Name</span>
			<span class="watt-container"><span class="watts">123</span> Watt BoomCase</span>
			<span class="price-container">$<span class="price">1500</span></span>
		</a>
	`
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	
	//c := appengine.NewContext(r)
	
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	
	
	
	// ========== ========== ========== ========== ==========
	// Getting current user
	/*
	// If the user is already signed in to your application, user.Current returns a pointer to a user.User value. Otherwise, it returns nil:
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting is in the same entity group.
	// Queries across the single entity group will be consistent.
	// However, the write rate to a single entity group should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	*/
	// ========== ========== ========== ========== ==========
	
	
	
	
	// ========== ========== ========== ========== ==========
	// func root
	/*
	// Ancestor queries, as shown here, are strongly consistent with the High Replication Datastore.
	// Queries that span entity groups are eventually consistent.
	// If we omitted the .Ancestor from this query there would be a slight chance that Greeting that had just been written would not show up in a query.
	
	// The following code constructs a Query value that requests the 10 most recent Greeting objects that are descendants of the root guestbook key in Date-descending order
	// [START query]
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	// [END query]
	
	// The following code calls q.GetAll(c, &greetings) to run the query and append the results to the greetings slice:
	// [START getall]
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// [END getall]
	
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	*/
	// ========== ========== ========== ========== ==========
	
	
	
	
	// ========== ========== ========== ========== ==========
	// func sign
	/*
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	
	// [START if_user]
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting is in the same entity group.
	// Queries across the single entity group will be consistent.
	// However, the write rate to a single entity group should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	// [END if_user]
	*/
	// ========== ========== ========== ========== ==========
	
	
	
	
	
	
	
	
	
	
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

