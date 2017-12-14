
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


	// ========== ========== ========== ========== ==========
	textEmail := `
BoomCase Email Order
user@domain.com
Thank you for your order!
Our expert team will review your custom design and contact you shortly.
Unless there is a major design issue that should be addressed you will receive an invoice for the total along with additional info on the delivery of your order.`
	// ========== ========== ========== ========== ==========

	// ========== ========== ========== ========== ==========
	htmlEmail := `
<h1>BoomCase Email Order</h1>
<div>
	user@domain.com
</div>
<div>
	Thank you for your order!
	Our expert team will review your custom design and contact you shortly.
	Unless there is a major design issue that should be addressed you will receive an invoice for the total along with additional info on the delivery of your order.
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
	e.To = []string{"pathaugen@gmail.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	e.Cc = []string{"info@productionmediadesign.com"}
	e.Subject = "BoomCase Email Order for user@domain.com"
	e.Text = []byte(textEmail)
	e.HTML = []byte(htmlEmail)
	//e.AttachFile("test.txt")
	e.AttachFile("/boomcase-logo.png")
	e.Send("smtp.mailgun.org:587", smtp.PlainAuth("", "postmaster@mg.productionmediadesign.com", "bbcfbbc6173a7dc7e8f7c22108c47145", "smtp.mailgun.org"))
	// ========== ========== ========== ========== ==========


  output += "sendEmail() CONTENT"

  return output
}
// ========== END: saveImage ========== ========== ========== ========== ========== ========== ========== ========== ==========
