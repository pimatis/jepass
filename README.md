# JePass - Password Management Tool

JePass is a command-line password management tool written in Go that provides password generation, validation, and updating capabilities with strength analysis.

## Features

- **Password Generation**: Create secure passwords with customizable length
- **Password Validation**: Check password strength and identify security issues
- **Password Updates**: Enhance existing passwords by appending secure characters
- **Strength Analysis**: Evaluate passwords with detailed strength ratings
- **Interactive CLI**: User-friendly command-line interface

## Installation

### Prerequisites
- Go 1.16 or higher
- Git

### Setup
```bash
git clone https://github.com/pimatis/jepass.git
cd jepass
```

## Usage

Run the application:
```bash
`./jepass` or `go run main.go`
```

### Available Operations

1. **Password Creation** - Generate new secure passwords
2. **Password Update** - Enhance existing passwords
3. **Password Check** - Validate and analyze password strength
4. **Exit** - Close the application

## API Documentation

### Password Generation (`functions.Create`)

Generates a secure password with mixed character types.

```go
func Create(length int) string
```

**Parameters:**
- `length`: Desired password length (minimum 8 characters)

**Returns:**
- Generated password string containing lowercase, uppercase, numbers, and symbols

**Example:**
```go
password := functions.Create(12)
// Returns: "aB3#fG7@kL9$"
```

### Password Validation (`functions.Check`)

Analyzes password strength and identifies security issues.

```go
func Check(password string) (bool, PasswordStrength, []string)
```

**Parameters:**
- `password`: Password string to validate

**Returns:**
- `bool`: Whether password is valid (no critical issues)
- `PasswordStrength`: Strength level (Weak, Medium, Strong, VeryStrong)
- `[]string`: List of identified issues

**Validation Criteria:**
- Minimum 8 characters length
- Contains lowercase letters
- Contains uppercase letters
- Contains numbers
- Contains symbols
- Avoids common patterns ("password", "123456")

**Example:**
```go
isValid, strength, issues := functions.Check("MyPass123!")
// Returns: true, Strong, []
```

### Password Update (`functions.Update`)

Enhances existing passwords by appending secure characters.

```go
func Update(oldPass string) string
```

**Parameters:**
- `oldPass`: Existing password to enhance

**Returns:**
- Enhanced password with additional secure characters

**Example:**
```go
newPassword := functions.Update("oldpass")
// Returns: "oldpassaB3#"
```

## Password Strength Levels

| Level | Score Range | Description |
|-------|-------------|-------------|
| Weak | ≤ 2 | Missing multiple security criteria |
| Medium | 3-4 | Meets basic requirements |
| Strong | 5-6 | Good security with most criteria met |
| VeryStrong | ≥ 7 | Excellent security with all criteria exceeded |

## Security Features

- **Character Diversity**: Ensures mix of character types
- **Length Validation**: Enforces minimum security standards
- **Pattern Detection**: Identifies common weak patterns
- **Strength Scoring**: Quantitative security assessment
- **Visual Feedback**: Color-coded issue reporting

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source. Please check the license file for details.

<div align="center" style="display: flex; align-items: center; justify-content: space-between;">
   <p style="margin-left: 25rem; margin-top: 1.2rem;">Created by <a href="https://github.com/pimatis">Pimatis Labs</a></p>
   <img src="https://www.upload.ee/image/17796243/logo.png" alt="PiContent Logo" width="30" style="opacity: 0.2; position: absolute;">
</div>
