package database

import (
	"fmt"
	"go-blog-api/internal/types"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	CreateUser(c *gin.Context, user types.AuthRequest)
	GetUser(c *gin.Context, username string) *Users
}

type service struct {
	db *gorm.DB
}

var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DB")
)

func New() Service {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	dbInstance := &service{
		db: db,
	}
	PushSchema(db)

	log.Println("Connected to database")
	return dbInstance
}
