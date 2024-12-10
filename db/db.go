package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pgxpool.Pool

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	Conn, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	_, err = Conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS route_calls (
        id SERIAL PRIMARY KEY,
        route_name TEXT NOT NULL,
        call_count INTEGER NOT NULL DEFAULT 0
    )`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
}

func IncrementRouteCall(routeName string) {
	_, err := Conn.Exec(context.Background(), `
        INSERT INTO route_calls (route_name, call_count)
        VALUES ($1, 1)
        ON CONFLICT (route_name)
        DO UPDATE SET call_count = route_calls.call_count + 1;
    `, routeName)
	if err != nil {
		log.Printf("Unable to increment route call count: %v\n", err)
	}
}

func GetTotalRequests() int64 {
	var total int64
	rows, err := Conn.Query(context.Background(), "SELECT SUM(call_count) FROM route_calls")
	if err != nil {
		return 0
	}
	defer rows.Close()
	
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0
		}
	}
	return total
}
