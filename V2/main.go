package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	deptNo   string
	deptName string
	sqlDb    = os.Getenv("MYSQL_DATABASE")
	sqlHost  = os.Getenv("MYSQL_HOST")
	sqlPort  = os.Getenv("MYSQL_PORT")
)

func main() {
	errChan := make(chan error)
	sqlUser, sqlPass, err := getVaultCreds()
	if err != nil {
		log.Fatal(err)
	}
	dbconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqlUser, sqlPass, sqlHost, sqlPort, sqlDb)
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select dept_no, dept_name from departments where dept_no = ?",
		`d005`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&deptNo, &deptName)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// start http server
	log.Println("Starting Web App server.")
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/health", healthHandler)
	http.Handle("/", r)
	go func() {
		errChan <- http.ListenAndServe("0.0.0.0:8081", nil)
	}()
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// healthHandler returns http status ok.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alive!"))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Please add /health to the URL")
}

func parseStringEnv(flagName string) string {
	if os.Getenv(flagName) != "" {
		return os.Getenv(flagName)
	}
	return os.Getenv(strings.ToUpper(flagName))
}
