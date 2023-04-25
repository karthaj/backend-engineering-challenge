### About Me
### CV : [Ashfak-Khajudeen.pdf](https://github.com/karthaj/backend-engineering-challenge/files/11316110/Ashfak-Khajudeen.pdf)

#### Summary
As a Full-Stack Engineer with 2.8 years of experience at PickMe - Sri Lanka, I have successfully led the development of multiple projects with
varying sizes and complexities. My experience ranges from small-scale applications to large-scale systems with millions of users. As a key
member of a cross-functional passenger team, including designers, product owners, and developers, I have ensured the timely and successful
delivery of high-quality products. My focus on implementing industry best practices, such as unit testing, code reviews, and CI/CD pipelines, has
consistently resulted in high code quality and maintainability.
#### I am willing to explore job opportunities that require me to relocate to a new location.

---

# Backend Engineering Challenge

## Overview:
This is a Go program that emulates a RESTful API for money transfers between accounts. The application downloads account information from a JSON file, ingests it into an in-memory datastore, and provides endpoints to view all accounts and perform transfers between them.

## High Level Architecture
![image](https://user-images.githubusercontent.com/48028155/234168699-d142256b-6094-4986-9c75-ade53f96e509.png)



### features 
  1. Authenticated with JWT - ( Not implemented 100% )
  2. Custom Logs on server console 
  3. Validation of endpoint and method
  4. Custom error handling
  5. Im memory DB - BadgerDB ( new to the library )
    

# REST API example application
## Install

    go mod downlaod && go mod tidy

## Run the app

    go run main.go 

# REST API

The REST API to the backend-engineering-challenge is given below. Tester may use the curl to execute the API calls.

## Get Account by ID

#### Request - GET

   `{BASE-API}/v1.0/account/get/id/23477b82-84a1-41fe-b259-c41179182451`

    curl --location --request GET 'localhost:8085/v1.0/account/get/id/23477b82-84a1-41fe-b259-c41179182451' \
    --header 'correlation-id: ASHFAk-c0d2-45f2-84b2-0349a4af1b4e' \
    --header 'Authorization: Bearer test-token'

#### Success Response - 200

    {
        "data": {
            "accounts": [
                {
                    "id": "23477b82-84a1-41fe-b259-c41179182451",
                    "name": "Dynabox",
                    "balance": 2260.45
                }
            ]
        },
        "meta": {
            "code": 200,
            "message": "Success"
    }


#### Error Response - 400

    {
        "errors": [
            {
                "correlationId": "ASHFAk-c0d2-45f2-84b2-0349a4af1b4e",
                "code": "API-100105",
                "message": "Account not found"
            }
        ]
    }
---

## Get list of all accounts

#### Request - GET
`{BASE-API}/v1.0/account/get/all`

    curl --location --request GET 'localhost:8085/v1.0/account/get/all' \
    --header 'correlation-id: ASHFAk-c0d2-45f2-84b2-0349a4af1b4e' \
    --header 'Authorization: Bearer test-token'

#### Success Response - 200

    {
        "data": {
            "accounts": [
                {
                    "id": "03eb9399-e526-431f-812f-2fda01659022",
                    "name": "Browsecat",
                    "balance": 3172.14
                },
                {
                    "id": "04943793-8f35-4d73-aa93-0ef2da57d22e",
                    "name": "Flashset",
                    "balance": 3724.11
                },
                {
                    "id": "054f9801-f070-4f16-bde7-155430417d43",
                    "name": "Quimba",
                    "balance": 1011.56
                },
                ....
            ]
        {
    {


#### Error Response - 400

    {
        "errors": [
            {
                "correlationId": "ASHFAk-c0d2-45f2-84b2-0349a4af1b4e",
                "code": "API-100105",
                "message": "No accounts available"
            }
        ]
    }

---

## Perform transfers

#### Request - POST
`{BASE-API}/v1.0/account/get/all`

    curl --location --request GET 'localhost:8085/v1.0/account/get/all' \
    --header 'correlation-id: ASHFAk-c0d2-45f2-84b2-0349a4af1b4e' \
    --header 'Authorization: Bearer test-token'

#### Success Response - 200

    {
        "data": {
            "account": {
                "id": "2d9f59e6-3896-406c-b035-9a2fc99adbf8",
                "name": "Tagcat",
                "balance": 4779.34,
                "reference": "TRX-20232504081244"
            }
        },
        "meta": {
            "code": 200,
            "message": "Success"
        }
    }


#### Error Response - 400

    {
        "errors": [
            {
                "correlationId": "ASHFAk-c0d2-45f2-84b2-0349a4af1b4e",
                "code": "API-100105",
                "message": "Invalid sender account to perform transaction"
            }
        ]
    }

 
    {
        "errors": [
            {
                "correlationId": "ASHFAk-c0d2-45f2-84b2-0349a4af1b4e",
                "code": "API-100105",
                "message": "Invalid sender account to perform transaction"
            }
        ]
    }
    

    {
        "errors": [
            {
                "correlationId": "ASHFAQ-c0d2-45f2-84b2-0349a4af1b4e",
                "code": "API-100106",
                "message": "Insufficient account balance"
            }
        ]
    }


---
 
