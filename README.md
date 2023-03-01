# Newsletter Subscription Service

## Introduction
This is a newsletter subscription service that allows users to subscribe and receive updates on new articles published on our website. Users can subscribe by providing their email address and can unsubscribe at any time. The service will send out newsletters to all subscribed users whenever a new article is published.

## Technologies and Framework
- Go
- MongoDB
- Gin

## Installation
- Clone the repository.
- Install dependencies with "go mod tidy"
- Create a .env file and add the following variables
``` 
from = "senderemail@domain-name"
password = "(Application-Specific Passwords)"
```
- Start the server with "go run main.go"

*https://help.saleshandy.com/article/135-smtp-could-not-authenticate*
