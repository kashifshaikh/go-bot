# Instructions

To execute: `go run cmd/bot.go`

This will launch an HTTP server by default on `http://localhost:8888`

You can change configuration in the `.env.dev` file.

# Testing
You can test against the `/profiles` endpoint using `tests/tests.http` file and using VSCode extension https://marketplace.visualstudio.com/items?itemName=mkloubert.vscode-http-client

The http-client extension works similar postman, but you can execute HTTP requests within vscode.
