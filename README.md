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
│       ├── ci.yml                  # Main CI pipeline configuration
│       ├── issue-retest.yml        # Issue retest workflow
│       └── release-validation.yml  # Release validation workflow
├── .devcontainer/            # Development container configuration
│   ├── devcontainer.json
│   └── Dockerfile
├── CHANGELOG.md              # Record of changes
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
   - Special release validation for dev to release branch PRs

2. **Verification and Validation**: 
   - Automated testing with coverage reporting
   - Static code analysis through linting
   - Security scanning
   - Enforced coverage thresholds for release candidates

3. **Documentation**: 
   - Enforced documentation standards through automated checks

4. **Version Control**: 
   - Special handling for release branches
   - Artifact generation and storage

5. **Quality Assurance**:
   - Code format checking
   - Dependency verification
   - Multiple verification steps

## CI Pipelines

### Main CI Pipeline

The main CI pipeline is triggered on:
- Push to main/master branches
- Push to release/* branches
- Pull requests to main/master and release/* branches
- When an issue comment contains '#retest'

The pipeline includes the following jobs:
1. **Lint**: Code quality and formatting checks
2. **Build**: Compilation and artifact generation
3. **Test**: Unit testing with coverage reporting
4. **Security Scan**: Security vulnerability detection

### Release Validation Pipeline

A specialized validation pipeline triggers only for:
- Pull requests from `dev/*` branches to `release/*` branches

This pipeline includes stricter validations:
1. **Source Branch Verification**: Ensures PR is from a dev branch
2. **Lint**: Enhanced code quality and formatting checks
3. **Build**: Artifact generation with release-specific naming
4. **Test**: Enforced minimum test coverage (80%)
5. **Security Scan**: Blocks release if high-severity issues exist
6. **Release Readiness**: Documentation and changelog verification

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

## Branch Structure

This project follows a structured branching model to support ASPICE compliance:

- `main` / `master`: Primary production branch
- `release/*`: Release branches (e.g., `release/v1.0`, `release/v2.3`)
- `dev/*`: Development branches (e.g., `dev/feature-x`, `dev/bugfix-y`)

The CI pipelines enforce certain rules:
- All code must pass basic CI before merging to any branch
- Only `dev/*` branches can be merged to `release/*` branches
- Special validation rules apply when promoting from dev to release

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