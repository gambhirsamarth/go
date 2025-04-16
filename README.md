# Go Programming Learning Path: A Structured Guide to Learning Go Fundamentals

This repository provides a comprehensive, hands-on approach to learning Go programming language fundamentals through practical examples and progressive exercises. It offers a structured learning path from basic syntax to more complex programming concepts, making it ideal for beginners starting with Go.

The repository contains carefully organized examples and practice questions that build upon each other, covering essential Go concepts like variables, data types, functions, control structures, and data structures. Each topic includes practical examples and exercises that reinforce learning through hands-on coding experience. The practice questions are organized by days, allowing for a systematic approach to learning Go programming concepts.

## Repository Structure
```
.
├── go.mod                  # Module definition and Go version specification
└── main/                   # Main source directory containing all Go files
    ├── 1HelloWorld.go      # Entry point demonstrating basic Go program structure
    ├── 2Variables.go       # Examples of variable declaration and initialization
    ├── 3DataTypes.go       # Comprehensive overview of Go data types
    └── practice-questions/ # Organized daily practice exercises
        ├── Day1/          # Basic output and string manipulation
        ├── Day2/          # Numbers and type checking
        ├── Day3/          # Functions and parameter passing
        ├── Day4/          # Control structures and loops
        └── Day5/          # Arrays, slices, and basic algorithms
```

## Usage Instructions
### Prerequisites
- Go 1.24.0 or later installed on your system
- Basic understanding of programming concepts
- Text editor or IDE (VS Code, GoLand, etc.)

### Installation
1. Install Go:
```bash
# macOS (using Homebrew)
brew install go

# Linux
sudo apt-get update
sudo apt-get install golang-go

# Windows
# Download installer from https://golang.org/dl/
```

2. Clone the repository:
```bash
git clone <repository-url>
cd <repository-name>
```

3. Verify Go installation:
```bash
go version
```

### Quick Start
1. Run the Hello World program:
```bash
cd main
go run 1HelloWorld.go
```

2. Explore variable declarations:
```bash
go run 2Variables.go
```

3. Study data types:
```bash
go run 3DataTypes.go
```

### More Detailed Examples
1. Working with strings:
```go
package main
import "fmt"

func main() {
    // Basic string declaration
    var name string = "Go Developer"
    fmt.Println(name)
    
    // Multiline string
    multiline := `Line 1
    Line 2
    Line 3`
    fmt.Println(multiline)
}
```

2. Basic function example:
```go
package main
import "fmt"

func sum(a, b int) int {
    return a + b
}

func main() {
    result := sum(5, 3)
    fmt.Printf("Sum: %d\n", result)
}
```

### Troubleshooting
1. Common GOPATH issues:
   - Error: `go: cannot find main module`
   - Solution: Ensure you're in the correct directory with go.mod file
   ```bash
   echo $GOPATH
   go env GOPATH
   ```

2. Build errors:
   - Error: `undefined: fmt.Println`
   - Solution: Check import statement: `import "fmt"`

## Data Flow
The repository follows a progressive learning path, starting with basic concepts and building towards more complex topics.

```ascii
Basic Concepts → Variables → Data Types → Functions → Control Structures → Data Structures
[HelloWorld] → [Variables] → [DataTypes] → [Functions] → [Loops/Conditionals] → [Arrays/Slices]
```

Key component interactions:
1. Each concept builds upon previous lessons
2. Practice questions reinforce theoretical concepts
3. Examples demonstrate practical usage
4. Daily exercises provide structured learning progression
5. Code files include detailed comments for better understanding
6. Each topic includes multiple examples for comprehensive coverage
7. Practice questions increase in complexity as concepts advance