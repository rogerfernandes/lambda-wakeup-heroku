# lambda-wakeup-heroku
Lambda to call Heroku App

## Usage
Make shure you have [Go](https://golang.org/doc/install), [Make](https://www.gnu.org/software/make/) and [Dep](https://github.com/golang/dep) installed.

```bash
$ go get -u github.com/rogerfernandes/lambda-wakeup-heroku
$ cd $GOPATH/src/github.com/rogerfernandes/lambda-wakeup-heroku
$ make build
```

- `make build` will build **main.zip**
- Upload _main.zip_ at **AWS Lambda Function**
- In **AWS CloudWatch** configure a Rule calling the **AWS Lambda Function** with send JSON `{"heroku_urls" : ["https://YOURAPP.herokuapp.com", "http://anotherurl.com"]}` and etc
