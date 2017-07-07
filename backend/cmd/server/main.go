package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bqluan/bookbook/handler"
)

var (
	address         = flag.String("address", "0.0.0.0", "The IP address for the backend server to serve on. (set to 0.0.0.0 for all interfaces)")
	mariadbHostname = flag.String("mariadb-hostname", "mariadb", "The hostname or IP address on which the mariadb server serves.")
	mariadbPort     = flag.Int("mariadb-port", 3306, "The port on which the mariadb server serves.")
	mariadbUser     = flag.String("mariadb-user", "root", "The user of the mariadb server.")
	mariadbPwd      = flag.String("mariadb-pwd", "root", "The password of the mariadb user.")
	port            = flag.Int("port", 8081, "The port for the backend server to serve on.")
)

func main() {
	flag.Parse()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/bookbook?charset=utf8&parseTime=True&loc=Local",
		*mariadbUser,
		*mariadbPwd,
		*mariadbHostname,
		*mariadbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	apiHandler, err := handler.CreateAPIHandler(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server: ready to serve on %s:%d", *address, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), apiHandler))
}
