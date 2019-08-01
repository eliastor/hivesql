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
	//Example_underlineWorkflow()
	Example_workflow()

}

func Example_underlineWorkflow() {

	cfg := gohive.NewConnectConfiguration()
	conn, err := gohive.Connect("localhost", 10000, "NONE", cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cursor := conn.Cursor()
	defer cursor.Close()
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

func Example_workflow() {
	db, err := sql.Open("hivesql", "localhost:10000/default")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("DESCRIBE DATABASE EXTENDED default")
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cols)

	resp := make([]*sql.NullString, len(cols))
	for k := range resp {
		resp[k] = new(sql.NullString)
	}
	respi := make([]interface{}, len(cols))
	for k := range respi {
		respi[k] = (interface{})(resp[k])
	}
	for rows.Next() {
		if err := rows.Scan(respi...); err != nil {
			rows.Close()
			log.Fatal(err)
		}
		for k := range resp {
			log.Print(*resp[k], ", ")
		}
		log.Println()
	}
	rows.Close()

	rows, err = db.Query("SHOW TABLES IN default")
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
		log.Println("    ", name)
	}
	rows.Close()

	rows, err = db.Query("SHOW TABLE EXTENDED IN ? LIKE ?", "default", "pokes")
	if err != nil {
		log.Fatal(err)
	}
	cols, err = rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cols)

	resp = make([]*sql.NullString, len(cols))
	for k := range resp {
		resp[k] = new(sql.NullString)
	}
	respi = make([]interface{}, len(cols))
	for k := range respi {
		respi[k] = (interface{})(resp[k])
	}

	for rows.Next() {
		if err := rows.Scan(respi...); err != nil {
			rows.Close()
			log.Fatal(err)
		}
		for k := range resp {
			log.Print(*resp[k], ", ")
		}
		log.Println()
	}
	rows.Close()

	//pokes
}
