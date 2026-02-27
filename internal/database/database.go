package database // Define database module
import (
	"context" // Manages request lifetimes and timeouts
	"fmt"
	"log"
	"os" // access environment variables

	"github.com/JorgeJola/indnratebackend/internal/models" // Load data structure for simulations
	"github.com/jackc/pgx/v5/pgxpool"                      //Postgres SQL connection pool
	"github.com/joho/godotenv"                             // Load .env file into environmental variables (Database credentials)
)

var DB *pgxpool.Pool // Global variable that use the pool


// Postgres connection
func Connect() {
	err := godotenv.Load() // Load environmental variables (.env file)
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}

func Query(cellID int, nitroPrice float64, grainPrice float64)([]models.Simulation, error){
	rows,err := DB.Query(context.Background(),
	"SELECT id_sim, id_cell, nitro_kg_ha, yield_kg_ha FROM simulations WHERE id_cell=$1",cellID)
	if err != nil {
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
		s.NitroLbAc = s.NitroKgHa / 0.892
		s.YieldBsAc = s.YieldKgHa * 15.9
		s.Profit_dol = (s.YieldBsAc * s.GrainPrice) - (s.NitroLbAc * s.NitroPrice)
        sims = append(sims, s)
    }

    return sims, nil
}
