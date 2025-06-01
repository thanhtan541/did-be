# writing-test

There are two types of tests:

- Unit tests: Whitebox testing
- Integration tests: Blackbox testing

## Unit tests

Unit tests are used to test individual components of the codebase. They are typically written using the `testing` package in Go.

> **_NOTE:_** All of Unit tests are written and reside in the `core` directory.

## Integration tests

Integration tests are used to test the interactions between different components of the codebase. They are typically written using the `testing` package in Go and have extra setup to create sandbox environment.

> **_NOTE:_** All of Integration tests are written and reside in the `api` directory.

## Test-Driven Development (TDD)

A software development approach where tests are written before the actual code is implemented

### How to write TDD tests

#### Core

Core is our Domain Model. All of the unit testings are written and reside in the `core` directory.
These tests are essential for ensuring the correctness and reliability of the domain model. The tests are covers:

1. Presentation: Parser between Business Logic and Data Access Layer
2. Business Logic: Use-cases that cover the expected behavior of the interaction such: verify, issue, update, delete, Did Document, etc.


Here are some steps to follow when writing TDD tests for Unit tests:

1. Identify the components that need to interact with each other: Unit test is not for only a file, but rather a set of intrucstion that achieve a task.
2. Write a test that covers the expected behavior of the interaction. Such as: verifier and Did Document
3. Implement the code to make the test pass.
4. Refactor the code as needed to improve its design and maintainability.

#### API

In the context of Integration tests, TDD involves writing tests that cover the interactions between different components of the codebase. For our `API`, we write tests to ensure that the exposed api are worked expected.
Such as: correct url, correct method, correct headers, correct body, etc.

Here are some steps to follow when writing TDD tests for Integration tests:

1. Think about the outer layer of the API, e.g: Endpoint, Body Struct, Query Params, Response data model
2. Write a test that make sure the enpoint is callable,
3. Then write enough code return desire output such StatusCode...
4. Refactor the code as needed to improve its design and maintainability.

## Example

### Unit Test

### Integration Test
- Where: All of integration tests are written and reside in the `api/tests` directory.
- Hows:
1. Create a file with meaningful name such as `ping_test.go` to focus on test `/ping` endpoint.
2. Use `tests/helper.go` to setup the test environment.
3. In some scenario, there is a chain of api call to achieve certain workflow.So, extending `calling endpoint` in `ping_test.go` to later re-use.
```go
func (ta *TestApp) getPing() *http.Response {
	apiUrl := fmt.Sprintf("%s/ping", ta.Url)
	res, err := ta.Client.Get(apiUrl)
	if err != nil {
		log.Fatalf("❌ Failed to call: %v", err)
	}

	return res
}

```
4. Follow Arrange-Act-Assert (AAA) pattern
```go
func TestPing(t *testing.T) {
	//Arrange
	app := SpawnApp()
	//Act
	resp := app.getPing()

	//Assert
	assert.Equal(t, 200, resp.StatusCode, "Request to Ping not sucessful")

	defer resp.Body.Close()

	var data PingResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("❌ Failed to json decode body: %v", err)
	}

	assert.Equal(t, "pong", data.Message, "Body has incorrect format")
}

5. Write more assertion to strengthen the test
```
