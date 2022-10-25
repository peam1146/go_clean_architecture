package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	book_datasource "github.com/peam1146/go_clean_architecture/src/book/data/datasource/remote_datasource"
	book_repo_impl "github.com/peam1146/go_clean_architecture/src/book/data/repository"
	book_controller "github.com/peam1146/go_clean_architecture/src/book/domain/controllers"
	book_router "github.com/peam1146/go_clean_architecture/src/book/presentation/router"
	"github.com/peam1146/go_clean_architecture/src/common/config"
	router "github.com/peam1146/go_clean_architecture/src/common/presentation/router/fiber"
	"log"
)

func setUpDB(dbConf *config.Database) (*sql.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.SSL)
	db, err := sql.Open("postgres", postgresInfo)
	return db, err
}

func main() {
	conf, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := setUpDB(&conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	bookDataSource := book_datasource.New(db)
	bookRepo := book_repo_impl.NewBookRepositoryImpl(bookDataSource)
	bookController := book_controller.NewBookController(bookRepo)

	r := router.NewFiber()

	g := r.Group("/book")

	bookRouter := book_router.NewBookRouter(bookController, g)
	bookRouter.Register()
	r.Listen(fmt.Sprintf(":%d", conf.App.Port))
}
