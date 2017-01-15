
package main

import (
    "fmt"
    "net/http"
)

// ========== ========== ========== ========== ==========
func handlerAPI10(w http.ResponseWriter, r *http.Request) {
	
	output := ""
	output += "API Response"
	
	// API can recieve HTTP verbs: GET or POST and is versioned for backwards compatability when large changes are made
	
	// URL: /api/
	// Behavior: Information about the versions of the API that are available
	
	// ========== ========== ========== ========== ==========
	// API 1.0 (/api10/)
	
	// URL: /api/10/
	// Behavior: Information about the 1.0 API
	
	
	// Case Management
	
	// URL: /api/10/cases/
	// Behavior: Collection of all cases in the datastore with limited data on each
	
	// URL: /api10/cases/[numericvalue]
	// Behavior: Element of the /cases/ collection (single case) returning more data
	
	
	// Driver Management
	
	// URL: /api/10/drivers/
	// Behavior: Collection of all drivers in the datastore with limited data on each
	
	// URL: /api/10/drivers/[numericvalue]
	// Behavior: Element of the /drivers/ collection (single driver) returning more data
	
	// ========== ========== ========== ========== ==========
	
    w.Header().Set("Content-Type", "text/html; charset=utf-8") // TODO: json output
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ==========
