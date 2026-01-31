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
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
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
		log.Fatal("Failed to initialize database:", err)
		return
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	//localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:" + port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
