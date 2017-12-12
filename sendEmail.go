
package main

import (
	"net/http"
	//"google.golang.org/appengine/blobstore" // https://cloud.google.com/appengine/docs/go/blobstore/reference

  //"gopkg.in/mailgun/mailgun-go.v1"
  "mailgun"
)


// ========== START: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
func sendEmail(r *http.Request) (string) {
	output := ""

	// ========== ========== ========== ========== ==========
	// Store the image in the blobstore
	//blobs, _, err := blobstore.ParseUpload(r)
	//if err != nil {  }
	//file := blobs["file"]

	//if len(file) == 0 { output = ""
	//} else { output = string(file[0].BlobKey) }
	// ========== ========== ========== ========== ==========


	// ========== ========== ========== ========== ==========
  mg := mailgun.NewMailgun(yourdomain, ApiKey, publicApiKey)
  message := mailgun.NewMessage(
    "sender@example.com",
    "Fancy subject!",
    "Hello from Mailgun Go!",
    "recipient@example.com")
  resp, id, err := mg.Send(message)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("ID: %s Resp: %s\n", id, resp)
	// ========== ========== ========== ========== ==========


  output += "sendEmail() CONTENT"

  return output
}
// ========== END: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
