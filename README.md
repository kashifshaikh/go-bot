# go-bot

I wanted to build my own bot that can scrape web sites and auto checkout game consoles like PS5 that are always out of stock on various retail websites like Walmart, Bestbuy, etc.

So I created a new project built with Electron, React, and Golang, where the React app would call the embedded go server to store/manage application config/state. Go-bot is the core component of my Gameover Bot project, that would perform the web-scraping and auto-checkout logic.

The golang project is using Echo web framework, Gorilla Websockets, BoltDB, and Storm ORM.

This is unfinished project - currently you can add/remove/update/get Billing/Credit card profiles and we perform validation checks on the CC and Addresses. You can also send websocket ping messages to `/ws` endpoint, and with pong back with the payload. 

# To run
go run cmd/bot.go
