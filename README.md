# Backend With Go

## Learning Backend with Go

### Project Structure

```
/home/chera/All/Projects/go/
├── README.md
├── go_basics/
│   ├── array/
│   │   └── main.go
│   ├── array_two/
│   │   └── main.go
│   ├── defer/
│   │   └── main.go
│   ├── Errors/
│   │   └── main.go
│   ├── functions/
│   │   └── another.go
│   ├── goio/
│   │   ├── main.go
│   │   └── newMain/
│   │       └── task.go
│   ├── hash_map/
│   │   └── main.go
│   ├── interface/
│   │   └── main.go
│   ├── loop_examples/
│   │   └── loop_ex.go
│   ├── loops/
│   │   └── main.go
│   ├── methods/
│   │   └── main.go
│   ├── numberics_const/
│   │   └── main.go
│   ├── pointers/
│   │   └── pointers.go
│   ├── printinh/
│   │   └── printing.go
│   ├── set/
│   │   └── main.go
│   ├── simple_math/
│   │   └── main.go
│   ├── type_assertion/
│   │   └── main.go
│   ├── type_conversion/
│   │   └── main.go
│   └── variables/
│       └── variable.go
├── a2sv_tasks/
│   ├── task1/
│   │   ├── main.go
│   │   └── main_test.go
│   └── task2/
│       ├── frequency_count.go
│       ├── frequency_test.go
│       ├── main.go
│       ├── palindrome_check.go
│       └── palindrom_test.go
├── go.mod
└── README.md
```

### Description

This repository contains my learning path for backend development with Go, structured as follows:

- **go_basics/**: Contains the basic Go projects and examples.
    - **array/**: Examples of array usage.
    - **array_two/**: More examples of array usage.
    - **defer/**: Examples of using defer in Go.
    - **Errors/**: Error handling examples.
    - **functions/**: Function usage examples.
    - **goio/**: Examples of Go I/O operations.
    - **hash_map/**: Examples of hash map usage.
    - **interface/**: Examples of interface usage.
    - **loop_examples/**: Loop examples.
    - **loops/**: More loop examples.
    - **methods/**: Method usage examples.
    - **numberics_const/**: Numeric constants examples.
    - **pointers/**: Pointer usage examples.
    - **printinh/**: Printing examples.
    - **set/**: Set usage examples.
    - **simple_math/**: Simple math operations.
    - **type_assertion/**: Type assertion examples.
    - **type_conversion/**: Type conversion examples.
    - **variables/**: Variable usage examples.

- **a2sv_tasks/**: Contains tasks and exercises from the A2SV backend learning path.
    - **task1/**: Contains the main.go and main_test.go for task 1.
    - **task2/**: Contains various Go files for task 2 including frequency count and palindrome check.

### Getting Started

To get started with this project, clone the repository and navigate to the project directory:

```sh
git clone /home/chera/All/Projects/go/
cd go
```

### Running the Application

To run the application, use the following command:

```sh
go run go_basics/main.go
```

### Running Tests

To run the tests, use the following command:

```sh
go test ./go_basics/tests/...
```

### Contributing

Feel free to contribute to this project by opening issues or submitting pull requests.