package server

import (
	"database/sql"
	"github.com/leonardoce/connection-test/internal/logging"
	"github.com/leonardoce/connection-test/internal/server/views"
	"net/http"
)

var (
	log = logging.New("server")
)

// Run start an HTTP server listening on that port
func Run(addr string) error {
	http.HandleFunc("/", handler)
	log.Info("Web server started", "addr", addr)
	return http.ListenAndServe(addr, http.DefaultServeMux)
}

func handler(resp http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.Path == "/":
		indexPage(resp, req)
	case req.URL.Path == "/connection_test":
		connectionTest(resp, req)
	}
}

func indexPage(resp http.ResponseWriter, req *http.Request) {
	views.WriteIndexPage(resp, "", "")
}

func connectionTest(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Error(err, "Error parsing form")
		resp.WriteHeader(http.StatusInternalServerError)
		views.WriteErrorPage(resp, err.Error())
		return
	}

	connectionString := req.FormValue("connection_string")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error(err, "Error while creating connection pool")
		views.WriteIndexPage(resp, connectionString, err.Error())
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error(err, "Error while pinging database")
		views.WriteIndexPage(resp, connectionString, err.Error())
		return
	}

	views.WriteIndexPage(resp, connectionString, "OK!")
}
