package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/JorgeJola/indnratebackend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

// Connect opens the pool using DATABASE_URL (preferred on Render) or DB_* vars.
func Connect() error {
	fmt.Fprintln(os.Stderr, "database: loading configuration")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := dsnFromEnv()
	if dsn == "" {
		return fmt.Errorf("set DATABASE_URL or DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD (and ensure the Postgres instance is linked or vars are set in Render)")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("create pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf(`database ping failed: %w

Fix on Render: (1) Postgres → Connections → copy the Internal Database URL into DATABASE_URL on this Web Service (same region as the DB). External hostnames often fail from the private network. (2) If you still see TLS errors on an older Postgres version, append sslnegotiation=postgres to DATABASE_URL to override direct TLS.`, err)
	}
	DB = pool
	log.Println("Connected to PostgreSQL (ping OK)")
	return nil
}

func dsnFromEnv() string {
	if u := strings.TrimSpace(os.Getenv("DATABASE_URL")); u != "" {
		return ensureSSLMode(u)
	}
	return buildDSNFromParts()
}

// ensureSSLMode forces sslmode=require. Render URLs often use sslmode=prefer, which can
// trigger plaintext startup and "unexpected EOF" against servers that require TLS.
func ensureSSLMode(conn string) string {
	u, err := url.Parse(conn)
	if err != nil {
		return conn + "?sslmode=require"
	}
	if u.Scheme != "postgres" && u.Scheme != "postgresql" {
		return conn
	}
	q := u.Query()
	if q.Get("sslmode") == "disable" {
		return conn
	}
	q.Set("sslmode", "require")
	q.Set("connect_timeout", "15")
	// Render and some cloud Postgres endpoints close with "unexpected EOF" when using the
	// legacy SSLRequest handshake; direct TLS (PostgreSQL 17+ / pgx sslnegotiation) avoids it.
	if q.Get("sslnegotiation") == "" {
		q.Set("sslnegotiation", "direct")
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func buildDSNFromParts() string {
	host := strings.TrimSpace(os.Getenv("DB_HOST"))
	port := strings.TrimSpace(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := strings.TrimSpace(os.Getenv("DB_NAME"))
	if host == "" || port == "" || name == "" || user == "" {
		return ""
	}
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, pass),
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   "/" + name,
	}
	q := u.Query()
	q.Set("sslmode", "require")
	q.Set("connect_timeout", "15")
	if q.Get("sslnegotiation") == "" {
		q.Set("sslnegotiation", "direct")
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func Query_sim(cellID int, nitroPrice float64, grainPrice float64) ([]models.Simulation, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id_sim, id_cell, nitro_kg_ha, yield_kg_ha FROM simulations WHERE id_cell=$1", cellID)
	if err != nil {
		log.Printf("database query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var sims []models.Simulation
	for rows.Next() {
		var s models.Simulation
		if err := rows.Scan(&s.IDSim, &s.IDCell, &s.NitroKgHa, &s.YieldKgHa); err != nil {
			log.Println("Row scan error:", err)
			continue
		}
		s.NitroPrice = nitroPrice
		s.GrainPrice = grainPrice
		s.NitroLbAc = s.NitroKgHa * 0.892
		s.YieldBsAc = s.YieldKgHa / 62.77
		s.Profit_dol = (s.YieldBsAc * s.GrainPrice) - (s.NitroLbAc * s.NitroPrice)
		sims = append(sims, s)
	}
	if err := rows.Err(); err != nil {
		log.Printf("database rows error: %v", err)
		return nil, err
	}
	return sims, nil
}

func Query_trials(regionID string) ([]models.Trials, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id_region, aonr FROM on_farm WHERE id_region=$1", regionID)
	if err != nil {
		log.Printf("database query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var trials []models.Trials

    for rows.Next() {
        var t models.Trials
        // Match the columns in your SELECT statement: id_region, aonr
        if err := rows.Scan(&t.IDRegion, &t.AONR); err != nil {
            log.Println("Row scan error:", err)
            continue
        }
        trials = append(trials, t)
    }


	if err := rows.Err(); err != nil {
		log.Printf("database rows error: %v", err)
		return nil, err
	}
	return trials, nil
}