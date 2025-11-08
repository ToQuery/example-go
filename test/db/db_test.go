package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func TestDBpgx(t *testing.T) {
	dsn := "" // &TimeZone=Asia/Shanghai"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// 验证时区
	var now time.Time
	err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB time:", now.Local())
}

func TestDbLibpq(t *testing.T) {
	dsn := "" // &TimeZone=Asia/Shanghai"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var now1 time.Time
	if err := db.QueryRow("SELECT NOW()").Scan(&now1); err != nil {
		log.Fatal(err)
	}

	log.Println("DB Time1:", now1.Local())

	query2 := db.QueryRowContext(context.Background(), "SELECT NOW()")

	var now2 time.Time
	err = query2.Scan(&now2)
	if err != nil {
		return
	}
	log.Println("DB Time2:", now2.Local())

}
