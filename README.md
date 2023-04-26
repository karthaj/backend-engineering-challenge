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
![image](https://user-images.githubusercontent.com/48028155/234176587-cf36e7ef-5cc5-488e-8eb3-eabf7cb904f5.png)

### Private API Request Flow with JWT Authentication Middleware
![go-arch-private-api-request-flow](https://user-images.githubusercontent.com/48028155/234178020-fb46f851-278c-4c55-9c4c-818ff91a7c16.png)

---

### features 
  1. Authenticated with JWT - required to process call but Not implemented payload validation ( expiery, invalid, blacklisted ) 
  2. Custom Logs on server console 
  3. Validation of endpoint and method
  4. Custom error handling
  5. Im memory DB - BadgerDB ( new to the library )
    


## Install

    go mod downlaod && go mod tidy

## Run the app

    go run main.go 

# REST API - [Postman Collection with examples](https://app.getpostman.com/join-team?invite_code=3328a1003a191d886a27b9af9c96fdd7&target_code=c1ef287fe4da05a5e9365e73e6383fe9)


The REST API to the backend-engineering-challenge is given below. Tester may use the curl to execute the API calls.
 

## Ping
#### Request - GET

   `{ BASE-API }/v1.0/ping`

    curl --location --request GET 'localhost:8085/v1.0/ping'

#### Success Response - 200
      {
          "Data": {
              "Status": "P I N G - Wed, 26 Apr 2023 11:11:44 +0530"
          },
          "Meta": {
              "Code": 200,
              "Message": "Success"
          }
      }

## Get Account by ID

#### Request - GET

   `{ BASE-API }/v1.0/account/get/id/{ ACCOUNT-ID }`

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
`{ BASE-API }/v1.0/account/get/all`

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
`{ BASE-API }/v1.0/account/transaction`

    curl --location --request POST 'localhost:8085/v1.0/account/transaction' \
    --header 'correlation-id: ASHFAK-c0d2-45f2-84b2-0349a4af1b4e' \
    --header 'Authorization: Bearer test-token' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "amount": 100,
        "fromAccountId": "4c0b1e82-2df3-48a1-a297-094cddde5546"
    }'

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

#### Error Response - 422

    {
        "errors": {
            "message": "Validation Error",
            "code": "API-100101",
            "correlationId": "ASHFAK-c0d2-45f2-84b2-0349a4af1b4e",
            "fields": {
                "ToAccountId": "required"
            }
        }
    }

---
 
