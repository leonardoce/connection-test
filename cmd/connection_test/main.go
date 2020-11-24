//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=../../internal/server/views

package main

import (
	"flag"
	"github.com/leonardoce/connection-test/internal/logging"
	"github.com/leonardoce/connection-test/internal/server"

	_ "github.com/lib/pq"
)

var (
	log = logging.New("main")
)

func main() {
	port := flag.String("-addr", ":8080", "the network address to be used for listening")
	flag.Parse()

	err := server.Run(*port)
	if err != nil {
		log.Error(err, "Error while starting server")
	}
}
