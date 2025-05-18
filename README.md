# Credit Card Validator

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A command-line tool that validates credit card numbers using the Luhn algorithm, implemented in Go.

## Features

- ✅ Validates 16-digit credit card numbers
- 🛡️ Sanitizes input (handles spaces/dashes)
- 📊 Supports all major card brands (Visa, Mastercard, etc.)
- 🔍 Detailed validation feedback
- 🚀 Single executable - no dependencies

## Installation

### Pre-built Binaries
Download the latest executable for your OS:

- [Windows (.exe)](https://example.com/download/ccvalidator.exe)
- [MacOS](https://example.com/download/ccvalidator-mac)
- [Linux](https://example.com/download/ccvalidator-linux)

# Interactive mode
./ccvalidator

# Single-check mode
./ccvalidator "4539 1488 0343 6467"

# Pipe input
echo "4111 1111 1111 1111" | ./ccvalidator

$ ./ccvalidator
Enter card number: 4539 1488 0343 6467
✅ Valid (Visa)

$ ./ccvalidator "1234-5678-9012-3456"
❌ Invalid (Failed Luhn check)

# Run tests
go test -v

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o ccvalidator.exe

# Generate test coverage
go test -coverprofile=coverage.out
go tool cover -html=coverage.out


### Key Sections Included:
1. **Badges** - Visual indicators for version/license
2. **Features** - Quick overview of capabilities
3. **Installation** - Both pre-built and source options
4. **Usage** - Multiple running examples
5. **Algorithm** - Brief technical explanation
6. **Development** - Common dev commands
7. **Contributing** - Open source guidelines

### From Source
```bash
git clone [https://github.com/Leul0M/CreditCard_Validator.git]
cd credit-card-validator
go build -o ccvalidator
