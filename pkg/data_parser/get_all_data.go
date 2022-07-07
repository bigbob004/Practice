package data_parser

import (
	"fmt"
	"log"
	"math"
	"sync"
)

const maxRoutineNum = 1000

func GetAllData(lgr *log.Logger, countriesList []string) ([]product, error) {
	//ch := make(chan int, maxRoutineNum)

	itemsCnt, err := getItemsCnt(fmt.Sprintf(url, 1))
	if err != nil {
		return nil, err
	}
	products := make([]product, 0, itemsCnt)

	// TODO: выяснить, как работает арифметика при разных типах
	pageCnt := int(math.Ceil(float64(itemsCnt) / ItemsOnPage))
	lgr.Printf("%d items on %d pages found. Lets parse!", itemsCnt, pageCnt)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < pageCnt; i++ {
		lgr.Printf("Page %3.d parsing...", i+1)
		pg, err := getAllItemsFromOnePage(fmt.Sprintf(url, i+1))
		if err != nil {
			lgr.Printf("Error parsing page %d: %v", i, err)
			continue
		}
		for j := 0; j < len(pg.Items); j++ {
			//ch <- 1
			wg.Add(1)
			// Создам горутину для отдельного товара
			go func(j int, pg *page) {
				defer wg.Done()
				itemId := pg.Items[j].Id

				price, err := getPriceOfOneItem(itemId)
				if err != nil {
					lgr.Printf("Error parsing price of item %s: %v", itemId, err)
					return
				}
				avlb := getAvailabilityOfOneItemFromAllCountries(itemId, countriesList, lgr)

				mu.Lock()
				product := product{
					Id:           pg.Items[j].Id,
					Title:        pg.Items[j].Title,
					Url:          pg.Items[j].Url,
					Availability: avlb,
					Price:        price.Val,
				}
				if err != nil {
					lgr.Printf("Eror while marshaling: %v", err)
				} else {
					products = append(products, product)
				}
				mu.Unlock()
				//<-ch
			}(j, &pg)
		}
	}
	wg.Wait()
	return products, nil
}
