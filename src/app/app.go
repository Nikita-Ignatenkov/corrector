package app

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
)

type Server struct {
	DB *pgx.Conn
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("SHOWCASE_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	s.DB = conn
}

func (s *Server) Close() {
	if err := s.DB.Close(context.Background()); err != nil {
		log.Fatal(err)
	}
}
