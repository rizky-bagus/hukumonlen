// Package main is the entrypoint for the API server.
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"hukum-onlen-go/common/middlewares"
	"hukum-onlen-go/internal/config"
	"hukum-onlen-go/internal/http/routes"
)

func main() {
	cfg, err := config.NewConfig(".env")
	panicOnError(err)

	bunDB, err := initDB(&cfg.MySQL)
	panicOnError(err)

	ginEngine := gin.Default()
	middlewares.NewCorsMiddleware(cfg, ginEngine)
	registerRoutes(bunDB, ginEngine)

	if err = ginEngine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}

func initDB(cfg *config.MySQLConfig) (*bun.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.User, cfg.Password, cfg.Protocol, cfg.Host, cfg.Database)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())

	return db, nil
}

func registerRoutes(db *bun.DB, engine *gin.Engine) {
	routes.NewGatheringRoutes(db, engine)
	routes.NewMemberRoutes(db, engine)
	routes.NewPublicRoutes(engine)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
