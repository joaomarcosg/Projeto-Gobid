package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Projeto-Gobid/internal/api"
	"github.com/joaomarcosg/Projeto-Gobid/internal/services"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store/pgstore"
	"github.com/joho/godotenv"
)

func main() {

	gob.Register(uuid.UUID{})

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBID_DATABASE_USER"),
		os.Getenv("GOBID_DATABASE_PASSWORD"),
		os.Getenv("GOBID_DATABASE_HOST"),
		os.Getenv("GOBID_DATABASE_PORT"),
		os.Getenv("GOBID_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	userStore := pgstore.NewPGUserStore(pool)
	productStore := pgstore.NewPGProductStore(pool)
	bidStore := pgstore.NewPGBidStore(pool)

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.Api{
		Router:         chi.NewMux(),
		UserService:    *services.NewUserService(userStore),
		ProductService: *services.NewProductService(productStore),
		Sessions:       s,
		WsUpgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		BidService: *services.NewBidService(bidStore),
		AuctionLobby: services.AuctionLobby{
			Rooms: make(map[uuid.UUID]*services.AuctionRoom),
		},
	}

	api.BindRoutes()

	port := os.Getenv("GOBID_APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port :%s\n", port)
	if err := http.ListenAndServe(":"+port, api.Router); err != nil {
		panic(err)
	}

}
