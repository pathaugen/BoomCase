
package main

import (
	"net/http"
	"net/smtp"

  //"gopkg.in/mailgun/mailgun-go.v1"
  //"github.com/mailgun/mailgun-go/mailgun"

	"github.com/jordan-wright/email"
)


// ========== START: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
func sendEmail(r *http.Request) (string) {
	output := ""


	// ========== ========== ========== ========== ==========
	//thumbOpts := image.ServingURLOptions { Size: 800, }
	//thumbKey := appengine.BlobKey(caseblobkey)
	//thumbnail, _ := image.ServingURL(ctx, thumbKey, &thumbOpts) // (*url.URL, error)
	// ========== ========== ========== ========== ==========


	userEmail := "pathaugen@gmail.com"


	// ========== ========== ========== ========== ==========
	textEmail := `
BoomCase Email Order
`+userEmail+`
Thank you for your order!
Our expert team will review your custom design and contact you shortly.
Unless there is a major design issue that should be addressed you will receive an invoice for the total along with additional info on the delivery of your order.`
	// ========== ========== ========== ========== ==========

	// ========== ========== ========== ========== ==========
	htmlEmail := `
<div style="font-size:1.2em;">
	<div style="background-color:black;color:white;text-align:center;padding:50px 0;">
		<h1>BoomCase Email Order</h1>
	</div>
	<div>
		`+userEmail+`
	</div>
	<div style="padding:50px 0;">
		Thank you for your order!
		Our expert team will review your custom design and contact you shortly.
		Unless there is a major design issue that should be addressed you will receive an invoice for the total along with additional info on the delivery of your order.
	</div>
	<div style="background-color:black;color:white;text-align:center;padding:50px 0;">
		The BoomCase
	</div>
</div>`
	// ========== ========== ========== ========== ==========


	// ========== ========== ========== ========== ==========
	// MAILGUN API
  //mg := mailgun.NewMailgun(yourdomain, ApiKey, publicApiKey)
  //message := mailgun.NewMessage(
    //"sender@example.com",
    //"Fancy subject!",
    //"Hello from Mailgun Go!",
    //"recipient@example.com")
  //resp, id, err := mg.Send(message)
  //if err != nil {
    //log.Fatal(err)
  //}
  //fmt.Printf("ID: %s Resp: %s\n", id, resp)
	// ========== ========== ========== ========== ==========


	// ========== ========== ========== ========== ==========
	// SMTP

	//e := email.NewEmail()
	//e.From = "Your Name <foo@YOUR_DOMAIN_NAME>"
	//e.To = []string{"bar@example.com"}
	//e.Subject = "Hello"
	//e.Text = []byte("Testing some Mailgun awesomeness")
	//err := e.Send("smtp.mailgun.com:587", smtp.PlainAuth("", "YOUR_USERNAME", "YOUR_PASSWORD", "smtp.mailgun.com"))
	//if err != nil {
		//panic(err)
	//}

	e := email.NewEmail()
	e.From = "Production Media Design <postmaster@mg.productionmediadesign.com>"
	e.To = []string{userEmail}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"info@productionmediadesign.com"}
	e.Subject = "BoomCase Email Order for user@domain.com"
	e.Text = []byte(textEmail)
	e.HTML = []byte(htmlEmail)
	//e.AttachFile("test.txt")
	//e.AttachFile("/boomcase-logo.png")
	e.Send("smtp.mailgun.org:587", smtp.PlainAuth("", "postmaster@mg.productionmediadesign.com", "bbcKEYGOESHERE", "smtp.mailgun.org"))
	// ========== ========== ========== ========== ==========


  output += "sendEmail() CONTENT"

  return output
}
// ========== END: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
