## Overview:
This is a Go program that emulates a RESTful API for money transfers between accounts. The application downloads account information from a JSON file, ingests it into an in-memory datastore, and provides endpoints to view all accounts and perform transfers between them.

### Requirements:
1. Go 1.16 or higher 

### Installation and Setup:

1. Clone the repository using the command git clone https://github.com/karthaj/backend-engineering-challenge
2. Navigate to the root directory of the project. 
3. Run go mod download to download all dependencies.
4. To run the program, execute go run main.go.
5. By default, the application will run on port 8085. You can specify a different port by setting the environment variable _PORT_ 

### API Usage
