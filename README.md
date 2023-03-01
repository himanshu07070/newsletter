[![License](https://img.shields.io/badge/License-GPL%20v3-yellow&?style=for-the-badge)](LICENSE)
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

```
If found error regarding credentials please follow this article: 
        ***https://help.saleshandy.com/article/135-smtp-could-not-authenticate***
```

# API Endpoints
- /subscribe (POST): Adds a new user to the database and subscribes them to the newsletter. Expects an email address in the request body and tag which shows from which topic user is subscribing for newsletter.
- /unsubscribe (POST): Removes a user from the database and unsubscribes them from the newsletter. Expects an email address in the request body.
- /newsletter (POST): Sends a newsletter with the latest published article to all subscribed users. Expects a metadata like:
 
     - link of that new article   
	   - tag  of that new article
	   - article name 
	   - mail subject

# Future Work
We plan to implement an algorithm that will recommend articles to users based on their interests and reading history. This will improve user engagement and make the newsletter more personalized.
