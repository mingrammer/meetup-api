[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Meetup API

> [Meetup web repository](https://github.com/qkraudghgh/meetup)

This repository is for API server for 9XD meetup management service. With integrated with Slack, users can use the meetup service bi-directionally between slack and web.

## Features

- Provides REST API for meetup web service
- Provides Slack bot RTM responder
- Hook the notification to slack when a meetup is made on web

## Setup

You must set up the following configurations in your `.bashrc`, `.zshrc` or other shells you are using

```bash
export API_SECRET_VALUE="any-secret-value"
export WEB_ENDPOINT="your-frontend-server-url"

export DB_USERNAME="db-username"
export DB_PASSWORD="db-pssword"

export GOOGLE_API_KEY="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export GOOGLE_MAP_API_KEY="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

export SLACK_CLIENT_ID="XXXXXXXXXXXX.XXXXXXXXXXXX"
export SLACK_CLIENT_SECRET="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
export SLACK_VERIFICATION_TOKEN="XXXXXXXXXXXXXXXXXXXXXXXX"

export SLACK_BOT_TOKEN="XXXX-XXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXX"
export SLACK_BOT_ID="XXXXXXXXX"
export SLACK_BOT_CHANNEL_ID="XXXXXXXXXX"
```

You should have SSL certificates to enable HTTPS for your API. It requires `certificate.crt` for certificate file and `private.key` for private key on project root directory.
If you don't need SSL, use `ListenAndServe` instead, in `main.go`

## Build & Run

In your `$GOPATH`, run this

```bash
go get github.com/mingrammer/meetup-api
```

Go to project directory, and run it

```bash
go build
./meetup-api

# 8080: REST API server for web
# 8081: Slack RTM server
```

## Maintainers

- [@qkraudghgh](https://github.com/qkraudghgh)
- [@mingrammer](https://github.com/mingrammer)
