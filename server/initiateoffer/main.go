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
			"token" : "ff2PwqPTRyKgKHEEC714g2:APA91bH5l9VynJ1CF9fr0OcQec48U2L8ZQrHlVQI28DQ6Bo7ahXcH_mWalj71HUHHdsBmOP-9XlpWPAqMBD9WTAHglikZRcOxHx5Pbh23_q8-4oAlIEztohUFDOvh-csu5Vznw4NWU8T",
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
