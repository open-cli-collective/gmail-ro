# Contributing to gmail-ro

Thank you for your interest in contributing to gmail-ro!

## Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/piekstra/gmail-ro.git
   cd gmail-ro
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

3. Build:
   ```bash
   make build
   ```

4. Run tests:
   ```bash
   make test
   ```

## Development Workflow

### Before Submitting

Run all checks:
```bash
make verify
```

This runs:
- `go fmt` - Code formatting
- `golangci-lint` - Linting
- `go test` - Tests with race detection

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# Run quick tests (skip slow ones)
make test-short
```

### Linting

```bash
make lint
```

Requires [golangci-lint](https://golangci-lint.run/welcome/install/).

## Code Style

- Follow standard Go conventions
- Use `go fmt` for formatting
- Write table-driven tests where appropriate
- Add comments for exported functions

## Important: Read-Only Design

This CLI is intentionally **read-only**. Do not add features that:
- Send emails
- Delete emails
- Modify emails or labels
- Perform any write operations

The `gmail.readonly` scope is fundamental to this project's design.

## Pull Request Process

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `make verify`
5. Submit a pull request

## Reporting Issues

Please include:
- Go version (`go version`)
- OS and architecture
- Steps to reproduce
- Expected vs actual behavior
