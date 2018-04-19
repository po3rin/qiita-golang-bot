package main

import (
	"fmt"
	"time"
)

func main() {
	var items *[]Item
	var r = RawGet{
		endpoint: "https://qiita.com/api/v2/items",
		time:     time.Now().Format("2006-01-02"),
		tag:      "golang",
	}

	items = r.getQiitaItem()
	client := TwitterClient()

	for _, i := range *items {
		t := i.Title + i.URL
		fmt.Println(t)
		if 1 != 1 {
			t := i.Title + i.URL
			fmt.Println(t)
			_, _, err := client.Statuses.Update(t, nil)
			if err != nil {
				fmt.Println("投稿に失敗しました: ", err)
			}
		}
	}
}
