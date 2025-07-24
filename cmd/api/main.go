package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Projeto-Gobid/internal/api"
	"github.com/joaomarcosg/Projeto-Gobid/internal/services"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store/pgstore"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GO_BID_DATABASE_USER"),
		os.Getenv("GO_BID_DATABASE_PASSWORD"),
		os.Getenv("GO_BID_DATABASE_HOST"),
		os.Getenv("GO_BID_DATABASE_PORT"),
		os.Getenv("GO_BID_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	store := pgstore.NewPGUserStore(pool)

	s := scs.New()
	s.Store = pgxstore.New(pool)

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *services.NewUserService(store),
		Sessions:    s,
	}

	api.BindRoutes()

	fmt.Println("Starting server on port :3080")
	if err := http.ListenAndServe(":3080", api.Router); err != nil {
		panic(err)
	}

}
