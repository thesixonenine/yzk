package mys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GoodsPage(game string, pageNo int, pageSize int) []Goods {
	if len(game) == 0 {
		game = "hk4e"
	}
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize < 20 {
		pageSize = 20
	}
	url := "https://api-takumi.mihoyo.com/mall/v1/web/goods/list?app_id=1&point_sn=myb&page_size=" + strconv.Itoa(pageSize) + "&page=" + strconv.Itoa(pageNo) + "&game=" + game
	resp, err := http.Get(url)
	if err != nil {
		return []Goods{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Goods{}
	}
	var mpJson map[string]interface{}
	err = json.Unmarshal(body, &mpJson)
	data := mpJson["data"]
	list := data.(map[string]interface{})["list"]
	for _, e := range list.([]interface{}) {
		//fmt.Println(e)
		item := e.(map[string]interface{})
		goodsName := item["goods_name"]
		goodsId := item["goods_id"]
		price := item["price"]

		fmt.Println("商品名称: " + goodsName.(string) + "\n商品ID: " + goodsId.(string) + "\n价格: " + strconv.FormatFloat(price.(float64), 'f', 0, 64) + " 米游币\n")
	}
	return []Goods{}
}

type Goods struct {
	GoodsId   string
	GoodsName string
	Price     int
}
