# TESTSUITE V2

End to end test runner for QA Engineer to run test cases & scenario sequentially

## Purpose of this repo

1. Entry level QA automation
2. Complete documentation of our test scenario in `Github`
3. Run test scenario in CI/CD pipeline

## Developer Reference

### Initialize

1. Install `packr` using `go get -u github.com/gobuffalo/packr/packr`
2. Checkout this repo
3. `cd mississippi`
4. `dep ensure -v`

### Quick run

1. `cd cmd`
2. `go run main.go -domain=apartment`

### Building the executables

1. `cd cmd`
2. `packr build -o testsuite`
3. `./testsuite -domain=apartment`

### About domain

The `domain` flag determine which folder inside `pkg` we want to run.
So `-domain=apartment` means we are running all tests scenario under `pkg/apartment`

---

## QA Reference

### Working with testsuite

* QA works inside `pkg` folder
* Folders under `pkg` are considered as test `scenario`
* `json` files under `scenario` is considered as individual test `case`
* `scenario.json` file is a special file we use to describe our test `scenario`

### Test Scenario Data Structure

```json
{
    "name" : "I am scenario name",
    "description" : "Describe your test goal / acceptance criteria"
}
```

### Test Case Date Structure

```json
{
    "order": 1,
    "name": "Firebase Auth",
    "description": "Authenticate to firebase",
    "request": {
        "url": "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=MY_SECRET_APIKEY",
        "method": "POST",
        "headers": null,
        "payload": {
            "email": "someone@spacestock.com",
            "password": "mypassword",
            "returnSecureToken": true
        }
    },
    "expect": {
        "status_code": 200,
        "headers": null,
        "evaluate": null
    },
    "pipeline": {
        "cache": true,
        "cache_as": "auth",
        "on_failure": "exit"
    }
}
```

## TODO

1. Dockerize
2. Deploy
3. Integrate
