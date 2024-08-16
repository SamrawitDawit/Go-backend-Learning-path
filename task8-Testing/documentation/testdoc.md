# Unit Test Suite Documentation
## Overview
This document provides instructions for running the unit tests, measuring test coverage, and addressing any issues encountered during testing for the Go project.

## Running Tests
To run the unit tests locally, follow these steps:

* Ensure Go is Installed:

Make sure you have Go installed on your machine. You can download it from golang.org.

* Navigate to the Project Directory:

Open a terminal and navigate to the root directory of your project.

* Run the Tests:

Execute the following command to run all tests:
 `go test ./... -v`

* Run Specific Tests
Execute the following command to run a specific test file:
`go test ./path/to/testfile_test.go`

Execute the following command to run a specific test function:
`go test -run TestFunctionName`

Test Coverage
To measure test coverage, you can use the built-in Go tool. Follow these steps:

Run Tests with Coverage:

Execute the following command to run tests and generate a coverage report:
`go test ./... -coverprofile=coverage.out`

View Coverage Report:

To view the coverage report in a human-readable format, run:
`go tool cover -html=coverage.out`

## Issues Encountered During Testing
### Pointer to Interface Error:

An error was encountered where a pointer to an interface was used, which is not allowed in Go. This was resolved by ensuring that the TaskUsecase interface was implemented by a concrete type and that the concrete type was used in the tests.

### Dependency Management:

Ensure that all dependencies are correctly managed using `go mod tidy` to avoid issues with missing packages.

### Environment Configuration:

Make sure the correct Go version is specified in the CI workflow to avoid compatibility issues.