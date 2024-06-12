# Simple JSON-Based Database in Go

This repository contains a simple JSON-based database implementation written in Go. It provides basic functionalities to create, read, update, and delete records stored in JSON files.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Code Explanation](#code-explanation)


## Features

- Create and write records to JSON files
- Read single records or all records from a collection
- Delete records or entire collections
- Uses mutexes for thread-safe operations
- Configurable logging

## Installation

### Prerequisites

- Go (version 1.22.3 or later)

## Clone the Repository

git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name

Install Dependencies
The required dependencies are listed in the go.mod and go.sum files. To install them, run:
go mod tidy

## Usage

Running the Application
To run the application, execute the following command:
go run main.go
I have commented out the delete function call. If you want to use the function just pass the name of resourse which you wanted to delete a particular file, other wise pass it empty (like this "" )

Example Usage

The main.go file contains an example usage of the JSON-based database with a set of predefined users.

## Code Explanation

Logger Interface: Defines logging methods like Fatal, Error, Warn, Info, Debug, and Trace.
Driver Struct: Contains the main database logic with a directory for storing data and a logger.
Options Struct: Allows passing optional configurations like a custom logger.
New Function: Initializes a new database driver, creating the directory if it doesn't exist.
Write Method: Saves a record to a specified collection and resource.
Read Method: Reads a specific record from a collection.
ReadAll Method: Reads all records from a collection.
Delete Method: Deletes a specific record or entire collection.
stat Function: Helper function to check file existence with a .json extension fallback.


