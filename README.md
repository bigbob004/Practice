# Practice

### задание см. тут https://rhetorical-pewter-3f8.notion.site/2-Waters-c490b766d4744da8920e31914f7ecb2f

### комментарий к заданию:
поговорили с Арсением насчёт того, какая информация должна присутствовать о конкретном товаре, он сказал следующее:
- каталожный номер товара
- название
- url
- цена
- доступность товара (проверяем доступноть в 3-ёх странах: Америка, Англия, Швейцария

### Что сделал:
- Нашёл апи, которое отдаёт сразу все айтемы (сразу все категории), а также сопутствующее апи для извлечения остальной информации о товаре
- Сделал вебсервер, который по вызову одного эндпоинта возвращает все товары (информацию о товарах)
- Извлекаю данные о товарах асинхронно для ускорения работы сервиса


### об апи сайта waters
У waters есть апи для извлечение всех товаров магазина, но товары разбросаны по страницам. Сам запрос на это апи выглядит так: https://prodservices.waters.com/api/waters/search/category_facet$shop:Shop?isocode=en_US&page=%d&rows=99, где
page - номер страницы с товарами, rows - количество товаров отображаемое на странице (максимум, который может отдать сайт waters - это 99 товаров, поэтому число 99 захардкожено). Это апи отдаёт много ненужной мне информации, я извлекаю только общее количество товаров (см. функцию getItemsCnt), каталожный номер (он пригодится далее), название и url товара. 

Часть информации мы извлекли, далее нужно извлечь цену и доступность. Тут я нашёл апи, которое выдаёт эту информации уже для конкретного товара (по каталожному номеру), а не пачкой по 99 товаров, как в предыдущем случае.

Апи для извлечения цены: https://api.waters.com/waters-product-exp-api-v1/api/products/prices?customerNumber=anonymous&productNumber= , где productNumber - каталожный номер товара. (см. функцию getPriceOfOneItem).

Апи для извлечения доступности: https://prodservices.waters.com/api/waters/product/v1/availability/productNumber/countryCode, где productNumber = каталожный номер товара, countryCode - код страны, для которой мы хотим узнать доступность (см. функцию getAvailabilityOfOneItem).

### Общая механика работы:
Сначала я извлекаю общее число товара магазина, затем узнаю количество страниц с товарами (общее кол-во / число товаров на странице (99 в нашем случае)).
Затем для данной страницы (для всех 99  товаров) извлекаю каталожный номер, url и название. Затем итерируюсь по всем товарам на странице и извлекаю для каждого из них цену и доступность (см. функцию GetAllData).

### Как это запускать?
- Склонируйте репозиторий в свою рабочую директорию
- Запустите сервер (точка входа в приложение файл cmd/main.go)
- В браузере набираем http://localhost:8080/parse (если у вас занят этот порт, можно выбрать другой, но тогда его придётся поменять и в коде)
- После отправки запроса на сервер, у вас появится файл info.log - туда будет записываться вся информация (сами товары, возможные ошибки), пока что это временно решение, скоро записи о товарах буду улетать в БД.
- Данные о всех товарах (а их на момент написания readme 10344) извлекаются примерно за 3 минуты, поэтому нужно дождаться, когда на страничке http://localhost:8080/parse появится надпись "Работа окончена". Также по окончании работы в конце файла info.log можно посмотреть общее время работы

### Ограничение нагрузки на сайт waters
- Для ограничение нагрузки (уменьшение общего кол-ва одновременных запросов) я создал фективный буферизированный канал, когда канал полностью заполнится, гланый поток заблокируется, пока го-рутины не вычитают из него инфорацию. Кол-во го-рутин регулируется константой maxRoutineNum.

### Информация об ошибках
- Информация об ошибках посылается в файл info.log.
- При отправке запросов на доступность для отдельных товаров в некоторых странах возникают ошибки с http-кодом 500. Это нормально, тут уже сервер waters так устроен, запросы на доступность в некоторых странах просто не продусмотрены. 
