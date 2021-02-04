package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// db constants
const (
	DBHost     = "DBHost"
	DBUser     = "DBUser"
	DBPassword = "DBPassword"
	DBName     = "DBName"
	Migration  = `CREAT TABLE IF NOT EXISTS Quote (
		id serial PRIMARY KEY,
		author text NOT NULL,
		content text NOT NULL,
		created_at timestamp with time zone DEFAULT current_timestamp
		)`
)

// Quote - struct definition of quote
type Quote struct {
	Author    string    `json:"author" binding:"required"`
	Content   string    `json:"content" bidning:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

// GetQuotes - get list of quotes
func GetQuotes() ([]Quote, error) {
	const q = `SELECT author, content, created_at FROM Quote ORDER BY created_at DESC LIMIT 100`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	results := make([]Quote, 0)

	for rows.Next() {
		var author string
		var content string
		var createdAt time.Time
		if err := rows.Scan(&author, &content, &createdAt); err != nil {
			return nil, err
		}
		results = append(results, Quote{Author: author, Content: content, CreatedAt: createdAt})
	}

	return nil, nil
}

// AddQuote - add quote to db
func AddQuote(quote Quote) error {
	const q = `INSERT INTO Quote(author, content, created_at) VALUES($1, $2 ,$3)`
	_, err := db.Exec(q, quote.Author, quote.Content, quote.CreatedAt)
	return err
}

func main() {

	var err error
	r := gin.Default()
	r.GET("/quotes", func(ctx *gin.Context) {
		results, err := GetQuotes()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error: " + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, results)
	})

	r.POST("/quote", func(ctx *gin.Context) {
		var q Quote
		if ctx.Bind(&q) == nil {
			q.CreatedAt = time.Now()
			if err := AddQuote(q); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error: " + err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv(DBHost), os.Getenv(DBUser), os.Getenv(DBPassword), os.Getenv(DBName))
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Query(Migration)
	if err != nil {
		log.Println("Failed to run migrations", err.Error())
		return
	}
	log.Println("Running quote server...")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
