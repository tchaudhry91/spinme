package spin_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tchaudhry91/spinme/spin"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func ExamplePostgres() {
	out, err := spin.Postgres(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give postgres a couple of seconds to boot-up, sadly there is no "ready" check yet
	time.Sleep(5 * time.Second)
	var hostEp string
	var ok bool
	// Grab the host endpoint mapping for the container
	if hostEp, ok = out.Endpoints["5432/tcp"]; !ok {
		fmt.Println("Didn't find expected port mapping")
	}
	// pq nees an independent port, not the entire endpoint
	ep := strings.Split(hostEp, ":")
	connStr := fmt.Sprintf("user=postgres password=password dbname=testdb port=%s sslmode=disable", ep[len(ep)-1])
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!")
	//Output: Connected!
}
