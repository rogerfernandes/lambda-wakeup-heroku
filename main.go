package main

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

//Request - request from AWS Lambda
type Request struct {
	HerokuURLs []string `json:"heroku_urls"`
}

//Response - response log to AWS Lambda
type Response struct {
	Log string `json:"log"`
}

func main() {
	lambda.Start(WakeUpHeroku)
}

// WakeUpHeroku is a function will make a simple GET request to your Heroku app through its url to wake it up
func WakeUpHeroku(req Request) (Response, error) {
	var log string
	for i := range req.HerokuURLs {
		resp, err := http.Get(req.HerokuURLs[i])
		if err != nil {
			log += "\n" + req.HerokuURLs[i] + " - Failed!"
		} else {
			log += "\n" + req.HerokuURLs[i] + " - Respond: " + getBodyContent(resp)
		}
	}
	return Response{Log: log}, nil
}

func getBodyContent(resp *http.Response) string {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(content)
}
