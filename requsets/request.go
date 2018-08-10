package requsets

import "github.com/parnurzeal/gorequest"

func GetGoods(pageNum, pageSize int) {
	apiUrl := "https://www.dhgate.com/dhrec/recommend_for_recentview_featured_left.action"
	request := gorequest.New()
	//page: 1
	//pagesize: 10
	//productnum: 10
	//pagetype: 4
	//keywords: t shirts
	//f:
	//ref:
	//_: 1533867778278
	request.Get(apiUrl).Query(`{
		page:}
}`)
}
