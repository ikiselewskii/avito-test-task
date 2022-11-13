package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"net/http"
	"os"
	"time"
	"log"
)

func main() {
	router := gin.Default()
	InitializeEndpoints(router)
	db := InitializeDBConnection(os.Getenv("DSN"))
	CreateUserTable(db)
	CreateReservationsTable(db)
	router.Run()

}

func InitializeEndpoints(r *gin.Engine) {
	r.GET("/", rootRoute)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func InitializeDBConnection(dsn string) (db *bun.DB) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	err:= sqldb.Ping()
	if err != nil{
		log.Println("can`t connect bc of ", err)
	}else {
		log.Println("everything`s fine")
	}
	db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return
}

// endpoint handlers
func rootRoute(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world", "status": http.StatusOK})
}

// database models
type MyBaseModel struct {
	bun.BaseModel
	ID        int       `bun:",pk"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type User struct {
	MyBaseModel `bun:"table:users,alias:u"`
	Balance     float32
}
type Reservation struct {
	MyBaseModel `bun:"table:reservations,alias:r"`
	UserID      int
	User        User `bun:"rel:belongs-to,join:user_id=id"`
	ProductID   string
	OrderID     string
	Price       string
	Status      string
}

//database tables

func CreateUserTable(base *bun.DB) {
	_, err := base.NewCreateTable().
		Model((*User)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		panic(err)
	}
}

func CreateReservationsTable(base *bun.DB) {
	_, err := base.NewCreateTable().
		Model((*Reservation)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		panic(err)
	}
}
