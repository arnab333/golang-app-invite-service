# Golang Invite Only Service

- Users will only able to login if they have the invitation code
- For database `MySQL` has been used
- For storing the auth tokens `Redis` has been used
- For storing the invitation tokens in memory storage has been used
- For database connectivity `GORM` has been used
- For authentication `JWT` has been used
- CRON has been used to auto clean up the expired invitation tokens
- OpenAPI Version 3 has been used for documenting the API

## Install

- Clone this repository
- Run `go mod tidy` to install all necessary packages
- Run `go run main.go` to start the application

## API Documentation

- https://documenter.getpostman.com/view/8548410/Uz5FHbpg
