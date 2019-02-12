package main

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

//Request - request from AWS Lambda
type Request struct {
	HerokuURL []string `json:"herokuapp"`
}

//Response - response log to AWS Lambda
type Response struct {
	Log string `json:"log"`
}

func main() {
	lambda.Start(wakeUpHeroku)
}

func wakeUpHeroku(req Request) (Response, error) {
	var log string
	for i := range req.HerokuURL {
		resp, err := http.Get(req.HerokuURL[i])
		if err != nil {
			log += "\n" + req.HerokuURL[i] + " - Failed!"
		} else {
			log += "\n" + req.HerokuURL[i] + " - Respond: " + getBodyContent(resp)
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
