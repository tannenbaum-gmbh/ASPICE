# ASPICE Calculator Demo

This project demonstrates CI capabilities of GitHub Actions for ASPICE software lifecycle. It includes a simple Golang calculator application with unit tests and a CI pipeline that runs on various triggers.

## Quick Start

```bash
# Build the calculator
make build

# Run the calculator
make run ARGS="add 5 3"
```

## Project Structure

```
├── calculator/
│   ├── calculator.go         # Calculator implementation
│   └── calculator_test.go    # Unit tests for calculator
├── cmd/
│   └── calc/
│       └── main.go           # Main application
├── .github/
│   └── workflows/
│       ├── ci.yml            # CI pipeline configuration
│       └── issue-retest.yml  # Issue retest workflow
├── .devcontainer/            # Development container configuration
│   ├── devcontainer.json
│   └── Dockerfile
├── Makefile                  # Build automation
└── go.mod                    # Go module definition
```

## Features

The calculator provides the following operations:
- Addition
- Subtraction
- Multiplication
- Division (with error handling for division by zero)

## ASPICE Compliance Features

1. **Traceability**: 
   - CI pipeline enforces code reviews through pull requests
   - Pull requests can be mapped to requirements

2. **Verification and Validation**: 
   - Automated testing with coverage reporting
   - Static code analysis through linting
   - Security scanning

3. **Documentation**: 
   - Enforced documentation standards through automated checks

4. **Version Control**: 
   - Special handling for release branches
   - Artifact generation and storage

5. **Quality Assurance**:
   - Code format checking
   - Dependency verification
   - Multiple verification steps

## CI Pipeline

The CI pipeline is triggered on:
- Push to main/master branches
- Push to release/* branches
- Pull requests to main/master and release/* branches
- When an issue comment contains '#retest'

The pipeline includes the following jobs:
1. **Lint**: Code quality and formatting checks
2. **Build**: Compilation and artifact generation
3. **Test**: Unit testing with coverage reporting
4. **Documentation**: Documentation standards verification
5. **Security Scan**: Security vulnerability detection

## How to Use

### Building the Application

Using Make:
```bash
make build
```

Or directly with Go:
```bash
go build -o calc ./cmd/calc
```

### Running the Application

Using Make:
```bash
make run ARGS="add 5 3"
make run ARGS="subtract 10 4"
make run ARGS="multiply 6 7"
make run ARGS="divide 20 5"
```

Or directly:
```bash
./calc add 5 3
./calc subtract 10 4
./calc multiply 6 7
./calc divide 20 5
```

### Running Tests

Using Make:
```bash
make test         # Run tests
make test-coverage # Run tests with coverage
```

Or directly:
```bash
go test ./...
```

## Issue Retest Feature

This project includes a special feature to rerun the CI pipeline directly from GitHub issues:

1. When someone comments on an issue with the text '#retest', it automatically triggers the CI pipeline
2. The system will comment on the issue to acknowledge the retest request
3. When the tests complete, the results are automatically posted back to the issue

This feature is useful for:
- Verifying bug fixes without requiring code changes
- Reproducing intermittent test failures
- Validating the current state of the codebase against reported issues

## License

This project is licensed under the MIT License - see the LICENSE file for details.