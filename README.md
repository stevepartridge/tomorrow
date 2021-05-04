# tomorrow
A basic Go client for Tomorrow API


### Tests
In order for tests to run tests it does require an API key.  The tests leverage a namespaced environment variable `TEST_TOMORROW_API_KEY`.  To run the tests use the following:
```
TEST_TOMORROW_API_KEY=<your api key> go test -v -cover .
```
