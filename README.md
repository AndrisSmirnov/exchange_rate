# Genesis Software Engineering School 2023 API

## Description

This repository contains the code for a simple API developed for the Genesis Software Engineering School 2023.

## Installation

1. Clone the repository: `git clone https://github.com/AndrisSmirnov/exchange_rate.git`
2. Navigate to the project directory: `cd exchange_rate`

## Usage

### Run locally

1. Run `go run ./cmd/main.go` to start the server locally.

### Run in Docker

1. Run `docker build -t exchange_rate .` to build the docker image.
2. Run `docker run -p 8080:8080 exchange_rate` to start the server in docker container.

### List of endpoints:

- `/api/rate` (GET): get current bitcoin rate in UAH
- `/api/subscribe` (POST): subscribe to mailing list
- `/api/sendEmails` (POST): send emails with current currency rate to all subscribers

### Features list:

- Getting exchange rates directly from the Binance
- Added the possibility to expand the service to receive the exchange rate of other currency
- Building the e-mail using HTML templates for better viewing
- The e-mail contains an "Unsubscribe" button.
- The e-mail contains links to all social networks
