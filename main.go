package main

import (
	"encoding/json"
	"fmt"
	"kasir-api-golang-2/database"
	"kasir-api-golang-2/handlers"
	"kasir-api-golang-2/repositories"
	"kasir-api-golang-2/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With, Accept",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	// config := Config{
	// 	Port:   viper.GetString("PORT"),
	// 	DBConn: viper.GetString("DB_CONN"),
	// }

	DBConn := viper.GetString("DB_CONN")
	fmt.Println("DB_CONN =", DBConn)

	db, err := database.InitDB(DBConn)
	if err != nil {
		// log.Fatal("Failed to initialize database:", err)
		log.Println("Failed to initialize database:", err)
		return
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	r := mux.NewRouter()
	r.HandleFunc("/api/produk", productHandler.HandleProducts).Methods("GET", "POST")
	r.HandleFunc("/api/produk/{id}", productHandler.HandleProductByID).Methods("GET", "PUT", "DELETE")

	//localhost:8080/health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	handler := enableCORS(r)
	// fmt.Println("Server running di localhost:" + port)
	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
