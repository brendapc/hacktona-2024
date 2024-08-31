package initializerDB

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Importa o driver do Postgres para o package database/sql
)

const (
	host     = "db"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// StartDB inicializa a conexão com o banco de dados e retorna um ponteiro para a instância *sql.DB
func StartDB() *sql.DB {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Configura o pool de conexões
	db.SetMaxOpenConns(25)                 // Número máximo de conexões abertas
	db.SetMaxIdleConns(25)                 // Número máximo de conexões ociosas
	db.SetConnMaxLifetime(5 * time.Minute) // Tempo máximo de vida para uma conexão

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db
}
