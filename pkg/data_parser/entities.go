package data_parser

const (
	ItemsOnPage = 99
	url         = "https://prodservices.waters.com/api/waters/search/category_facet$shop:Shop?isocode=en_US&page=%d&rows=99"
)

type cntItemsFromOnePage struct {
	NumFound int `json:"num_found"`
}

//TODO: добавить поле с типом товара
type page struct {
	Items []struct {
		Title string `json:"title"`
		Id    string `json:"skucode"`
		Url   string `json:"url"`
	} `json:"documents"`
}

type PriceEntity struct {
	Currency string  `json:"currencyCode"`
	Value    float64 `json:"value"`
}

type price struct {
	Val PriceEntity `json:"basePrice"`
}

type availability struct {
	Status string `json:"productStatus"`
}

type product struct {
	Title        string
	Id           string
	Availability map[string]bool
	Url          string
	Price        PriceEntity
}

const StructFormat = `	ID:               %s
	TITLE:            %s
	PRICE:            %.2f %s
	URL:              %s`
