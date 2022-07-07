package data_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// getItemsCnt для того, чтобы узнать общее количество товаров данного вида
func getItemsCnt(url string) (int, error) {
	body, err := MakeRequest(url, nil, 0, 0)
	if err != nil {
		return -1, err
	}
	cnt := cntItemsFromOnePage{}
	err = json.Unmarshal(body, &cnt)
	if err != nil {
		return -1, err
	}
	return cnt.NumFound, nil
}

func getAllItemsFromOnePage(request string) (page, error) {
	body, err := MakeRequest(request, nil, 0, 0)
	if err != nil {
		return page{}, err
	}
	var pg page
	if err := json.Unmarshal(body, &pg); err != nil {
		return page{}, err
	}
	for i := 0; i < len(pg.Items); i++ {
		sb := &strings.Builder{}
		sb.WriteString("https://www.waters.com")
		sb.WriteString(pg.Items[i].Url)
		pg.Items[i].Url = sb.String()
	}
	return pg, nil
}

func getPriceOfOneItem(itemId string) (price, error) {
	const baseUrl = "https://api.waters.com/waters-product-exp-api-v1/api/products/prices?customerNumber=anonymous&productNumber="
	siteUrl := baseUrl + itemId
	headers := map[string]string{
		"countryCode": "us",
		"channel":     "ECOMM",
		"language":    "en",
	}
	body, err := MakeRequest(siteUrl, headers, 0, 0)
	if err != nil {
		return price{}, err
	}
	var pr []price
	if err := json.Unmarshal(body, &pr); err != nil {
		return price{}, err
	}
	return pr[0], nil
}

func getAvailabilityOfOneItem(itemId string, countryCode string) (availability, error) {
	siteUrl := fmt.Sprintf("https://prodservices.waters.com/api/waters/product/v1/availability/%s/%s", itemId, countryCode)
	body, err := MakeRequest(siteUrl, nil, 0, 0)
	if err != nil {
		return availability{}, err
	}
	var av availability
	if err := json.Unmarshal(body, &av); err != nil {
		return availability{}, err
	}
	return av, nil
}

func getAvailabilityOfOneItemFromAllCountries(itemId string, countriesList []string, lgr *log.Logger) map[string]bool {
	avlb := make(map[string]bool)
	for _, country := range countriesList {
		availability, err := getAvailabilityOfOneItem(itemId, country)
		if err != nil {
			lgr.Printf("Error parsing availability in %s of item %s: %v", country, itemId, err)
		} else {
			avlb[country] = availability.Status == "IN_STOCK"
		}
	}
	return avlb
}

func Print(out io.Writer, prod []product) {
	for _, el := range prod {
		fmt.Fprintf(out, StructFormat+"\n", el.Id, el.Title, el.Price.Value, el.Price.Currency, el.Url)
		for countryCode, avlbty := range el.Availability {
			fmt.Fprintf(out, "\tAVAILABLE IN %s:  %v\n", countryCode, avlbty)
		}
		fmt.Println()
	}
}
