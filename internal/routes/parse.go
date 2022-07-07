package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	parser "webScraper/pkg/data_parser"
)

func Parse(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "", 0)
	products, err := parser.GetAllData(logger, []string{"US", "GB", "CH"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	parser.Print(os.Stdout, products)
	w.WriteHeader(http.StatusOK)
	duration := time.Since(start)
	logger.Printf("Время работы: %v", duration.Seconds())
	w.Write([]byte("Работа окончена"))
}
