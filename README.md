# tomorrow
[![GoDoc](https://godoc.org/github.com/stevepartridge/tomorrow?status.svg)](https://godoc.org/github.com/stevepartridge/tomorrow)
[![Go Report Card](https://goreportcard.com/badge/github.com/stevepartridge/tomorrow)](https://goreportcard.com/report/github.com/stevepartridge/tomorrow)
[![Coverage](http://gocover.io/_badge/github.com/stevepartridge/service)](http://gocover.io/github.com/stevepartridge/tomorrow)

A basic Go client for Tomorrow API


### Tests
In order for tests to run tests it does require an API key.  The tests leverage a namespaced environment variable `TEST_TOMORROW_API_KEY`.  To run the tests use the following:
```
TEST_TOMORROW_API_KEY=<your api key> go test -v -cover .
```
