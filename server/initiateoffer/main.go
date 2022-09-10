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
	// create request body
	reqBody := ioutil.NopCloser(strings.NewReader(`
	{
		"message": {
			"notification": {
				"title": "Zuelling Pharma September Offer",
				"body": "Take 10 off any [product or service] from now until the end of the month!",
				"image"  : "https://www.giving.sg/image/logo?img_id=35892708g"
			},
			"token" : "fW8Ol04Nkh3b0EVgOw3fJg:APA91bF0iNX37FDEVi_ZSBAn6rGbuwgvKMKdgyWjB5e0x02je8RjL6STCD55mmwvNqhF_8oJOQB48YaGajZdyTL-QrIx4_WodZyAyfQWq7rQOHPhdzUO9BuO9PF3Z1PfRslHShmpidW5",
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
	// create request object
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {" Bearer ya29.a0AVA9y1uldvZ-kZLVa9imbC2OYTMo6WHMyOIdtLhwKZUG9c7gzFP92N012twTZux2sIIGVjJJDxLG6wd_a7gGpDo4eVpgNIVdZf-I1DXG1iKk0DIO4qYG1drCOQWYS9PHqkePtAu02QQhfQx2uyQcuClSzE2BaCgYKATASARMSFQE65dr8IV_9Oo9BpoBBgNqXASyWmQ0163"},
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
