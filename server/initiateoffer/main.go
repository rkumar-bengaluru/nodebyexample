package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// req url
	reqUrl, _ := url.Parse("https://fcm.googleapis.com/v1/projects/zuellingpharma/messages:send")
	// create request body , replace device token
	reqBody := ioutil.NopCloser(strings.NewReader(`
	{
		"message": {
			"notification": {
				"title": "Zuelling Pharma September Offer",
				"body": "Take 10% off any [product or service] from now until the end of the month!",
				"image"  : "https://www.giving.sg/image/logo?img_id=35892708g"
			},
			"token" : "dnGZsLe9xwQKp1VhzDHvvR:APA91bHuKVsYBleR5QGlt2kQsaAavLc2eFFfGsBu-2-r_ly6OHrjaU4QUb61pJ2-ieDBoY__ao9PKSqpdaD7AVZGx4ItrS5QbmUH2lXVPtWkNMZ2TDbV-nqX9NJAfc0qDFDVSK8UWr7N",
			"webpush":{
       			"headers":{
         			"image":"https://www.giving.sg/image/logo?img_id=35892708g"
       			}
     		},
			"data": {
				"title" : "Awesome title",
				"body"  : "Your awesome push notification body",
				"image"  : "your_image_url"
			}
		} 
	}`))
	// create request object & replace oath2 token from https://developers.google.com/oauthplayground/
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {" Bearer ya29.a0AVA9y1vSI7qQevpm9LRBpX1sP5-87rBsrF2lC5UHWl7xe7kgXAdQBQTyIV_3cno34xHBELXSaPCOgPcD7ElQ6qZeaLLxmfv_Sx3woCZvUdO1FdFyGGIl5chLhhCFfJ8Ns4OTHl--TjttAqS7h7tIfgoyZlOxaCgYKATASARISFQE65dr8r3aQ2jzm3UYwbpFNcyQOHw0163"},
		},
		Body: reqBody,
	}
	// send the request
	res, err := http.DefaultClient.Do(req)
	// check for error
	if err != nil {
		log.Fatal("Error:", err)
	}
	// read response body
	data, _ := ioutil.ReadAll(res.Body)
	// close response body
	res.Body.Close()
	// print response status
	fmt.Printf("status: %d\n", res.StatusCode)
	fmt.Printf("body : %s\n", data)
	fmt.Println("initiating campaign...")
}
