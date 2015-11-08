package date

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQL(t *testing.T) {
	db, err := sql.Open("sqlite3", "date-test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("date-test.db")
	exec := func(q string, args ...interface{}) {
		if _, err := db.Exec(q, args...); err != nil {
			t.Error(err)
		}
	}

	exec(`CREATE TABLE IF NOT EXISTS info (deadline DATE);`)
	exec(`INSERT INTO info (deadline) VALUES("2012-12-12");`)

	row := db.QueryRow("select deadline from info")
	var d Date
	if err := row.Scan(&d); err != nil {
		t.Error(err)
	}
	log.Println(d.String())

	d = New(2012, 12, 11, time.UTC)
	exec(`INSERT INTO info (deadline) VALUES(?);`, d)

	row = db.QueryRow("select deadline from info where deadline = '2012-12-11'")
	if err := row.Scan(&d); err != nil {
		t.Error(err)
	}
	log.Println(d.String())
}
