package requests

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"strings"
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

type Product struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Url      string `json:"url"`
	Shipping bool   `json:"shipping"`
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

func GetGoodsFromHtml(word string, pageIndex int) []Product {
	targetUrl := "https://www.dhgate.com/wholesale/search.do?act=search&sus=&searchkey=%s&catalog=#hpsearch1806"
	word = strings.Replace(word, " ", "+", -1)
	if pageIndex == 0 {
		targetUrl = fmt.Sprintf(targetUrl, word)
	} else {
		targetUrl = fmt.Sprintf("https://www.dhgate.com/w/%s/%d.html", word, pageIndex)
	}
	fmt.Println(targetUrl)
	doc, _ := goquery.NewDocument(targetUrl)
	result := make([]Product, 0)
	doc.Find("#proList .listitem").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find(".subject").Attr("href")
		var isShip bool
		if shipping := selection.Find(".free").Text(); shipping == "Free shipping" {
			isShip = true
		} else {
			isShip = false
		}
		title := selection.Find(".subject").Text()
		if strings.TrimSpace(title) != "" {
			fmt.Println(title)
			result = append(result, Product{
				Title:    title,
				Url:      url,
				Price:    selection.Find(".price").Text(),
				Shipping: isShip,
			})
		}
	})
	return result

}

func main() {
	//GetGoods(1, 10, "thir")
	fmt.Println(GetGoodsFromHtml("t shirt", 1))
}
