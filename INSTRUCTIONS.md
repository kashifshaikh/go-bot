# Instructions

To execute: `go run cmd/bot.go`

This will launch an HTTP server by default on `http://localhost:8888`. 
A boltdb file will be created under /tmp/bot.db. This program was designed under MacOS, but should run under windows and linux as well. 

You can change port configuration in the `.env.dev` file.

# Testing
You can test against the `/profiles` endpoint using `tests/tests.http` file and using VSCode extension https://marketplace.visualstudio.com/items?itemName=mkloubert.vscode-http-client

The http-client extension works similar postman, but you can execute HTTP requests within vscode.
