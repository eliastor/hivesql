package hivesql_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/beltran/gohive"
	_ "github.com/eliastor/hivesql"
)

func Test_Workflow(t *testing.T) {
	Example_Underline_workflow()
	Example_Workflow()

}

func Example_Underline_workflow() {

	cfg := gohive.NewConnectConfiguration()
	conn, err := gohive.Connect("localhost", 10000, "NONE", cfg)
	if err != nil {
		log.Fatal(err)
	}
	cursor := conn.Cursor()
	defer conn.Close()
	ctx := context.TODO()
	cursor.Exec(ctx, "SHOW DATABASES")
	if cursor.Err != nil {
		log.Fatal(cursor.Err)
	}
	log.Println(cursor.Description())
	var s string
	for cursor.HasMore(ctx) {
		//cursor.FetchOne(ctx, &s)
		rowMap := cursor.RowMap(ctx)
		if cursor.Err != nil {
			log.Fatal(cursor.Err)
		}
		log.Println(rowMap, s)
	}
}

func Example_Workflow() {
	db, err := sql.Open("hivesql", "localhost:10000/default")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SHOW DATABASES;")
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cols)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			rows.Close()
			log.Fatal(err)
		}
		log.Println(name)
	}
	rows.Close()

	rows, err = db.Query("SHOW TABLE IN default")
	if err != nil {
		log.Fatal(err)
	}
	cols, err = rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cols)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			rows.Close()
			log.Fatal(err)
		}
		log.Println(name)
	}
	rows.Close()

}
