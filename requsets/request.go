package requsets

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type Reply struct {
	Msg    string `json:"msg"`
	Result []struct {
		Feedback  string `json:"feedback"`
		Sold      string `json:"sold"`
		Unit      string `json:"unit"`
		Highprice string `json:"highprice"`
		Lowprice  string `json:"lowprice"`
		Price     string `json:"price"`
		Link      string `json:"link"`
		Currency  string `json:"currency"`
		Title     string `json:"title"`
		ImgSrc    string `json:"imgSrc"`
	} `json:"result"`
	Code        int    `json:"code"`
	Ot          string `json:"ot"`
	Spot        string `json:"spot"`
	Pagetrack   string `json:"pagetrack"`
	Sppagetrack string `json:"sppagetrack"`
}

func GetGoods(pageNum, pageSize int, keyword string) Reply {
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
	var reply Reply
	request.Get(apiUrl).
		Query(fmt.Sprintf("pagesize=%d", pageSize)).
		Query(fmt.Sprintf("page=%d", pageNum)).
		Query(fmt.Sprintf("pagetype=%d", 4)).
		Query(fmt.Sprintf("keyword=%s", keyword)).
		Query(fmt.Sprintf("_=%d", time.Now().Unix())).
		EndStruct(&reply)
	return reply
}

func main() {
	GetGoods(1, 10, "thir")
}
