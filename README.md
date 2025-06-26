# ASPICE Calculator Demo

This project demonstrates CI capabilities of GitHub Actions for ASPICE software lifecycle. It includes a simple Golang calculator application with unit tests and a CI pipeline that runs on various triggers.

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
│       └── ci.yml            # CI pipeline configuration
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

The pipeline includes the following jobs:
1. **Lint**: Code quality and formatting checks
2. **Build**: Compilation and artifact generation
3. **Test**: Unit testing with coverage reporting
4. **Documentation**: Documentation standards verification
5. **Security Scan**: Security vulnerability detection

## How to Use

### Building the Application

```bash
go build -o calc ./cmd/calc
```

### Running the Application

```bash
./calc add 5 3
./calc subtract 10 4
./calc multiply 6 7
./calc divide 20 5
```

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.