## Library Management Console Application 

## Features 
 - Adding user every time the open the application 
 - Adding books to the library 
 - Removing Books from the library 
## Library Management System

This project is a library management system that allows users to borrow and return books. The system ensures that borrowed books are marked as unavailable until they are returned.

### Features
- Users can borrow books from the library, making the borrowed books unavailable.
- Users can return borrowed books, making them available again.

### File Structure 

├── controllers
│   └── library_controller.go
├── docs
│   └── documentation.md
├── go.mod
├── main.go
├── models
│   ├── book.go
│   └── member.go
├── README.md
├── services
│   └── library_service.go
└── test
	├── library_controller_test.go
	└── library_service_test.go

## Controller
The controller is responsible for user interface and communicating with methods inside the logic of the application in services folder 

## Docs
The docs is responsible for containing the documentation of the library 

## The main.go file 
The main is the entry point of the application running the application and calling the controller right away 

## Model
The models are responsible for creating the models of the application. For example, the Book struct and the Member struct.

## Service
The service contains the logic of the application, including methods for adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Test 
The test is responsible for testing the service of the application.
