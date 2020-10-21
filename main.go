package chatapp

import (
	"database/sql"
	"github.com/cavke/go-chat-app/http"
	"github.com/cavke/go-chat-app/mysql"
)

func main() {
	// Connect to database
	// TODO get DB properties from ENV: os.Getenv("DB")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/chatapp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Create services.
	us := &mysql.UserService{DB: db}

	// Attach to HTTP handler
	var h http.Handler
	h.UserService = us

	// TODO start http server
	//port := os.Getenv("PORT")

}
