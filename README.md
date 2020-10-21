GO application that implements chat functionality.
Messages are persisted in MySQL database,
chat functions are exposed through the REST API,
with following methods:
- /sendMsg - 
- /getMsg - 
- /getGroups - 
- /createGroup -

Run it with:

`go run main.go`

Application exposes following REST URL: `http://localhost:8080/chat`
