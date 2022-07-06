# Practice

###задание см. тут https://rhetorical-pewter-3f8.notion.site/2-Waters-c490b766d4744da8920e31914f7ecb2f

###Что сделал:
- Нашёл апи, которое отдаёт сразу все айтемы (сразу все категории), но айтемы разбросаны по страницам. Сам запрос на это апи выглядит так: https://prodservices.waters.com/api/waters/search/category_facet$shop:Shop?isocode=en_US&page=%d&rows=99, где
page - номер страницы с товарами, rows - количество товаров отображаемое на странице (максимум, который может отдать сайт waters - это 99 товаров, поэтому число 99 захардкожено)
