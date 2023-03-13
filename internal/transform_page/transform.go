package transform_page

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news_nosu/internal/utils"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func GetNews() [][]string {
	var wg sync.WaitGroup
	ch := make(chan []string, 32)
	for i := 1; i < 60; i++ {
		idx := i
		go GetTitles(idx, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var news [][]string

	for r := range ch {
		news = append(news, r)
	}
	return news
}

func GetTitles(pageNumber int, c chan []string, wg *sync.WaitGroup) {
	wg.Add(1)
	url := "https://www.nosu.ru/category/news/#"
	formData := make(map[string]string)
	formData["paged"] = fmt.Sprint(pageNumber)
	haders := make(map[string]string)
	haders["Content-Type"] = "application/x-www-form-urlencoded"
	page, _, err := utils.GetPage(context.Background(), http.MethodPost, url, nil, haders, formData, 0)
	if err != nil {
		log.Println(err)
	}
	page.Find(".content-block").Each(func(i int, s *goquery.Selection) {
		date := s.Find(".date").Text()
		title := s.Find("a").Text()
		description := s.Find("p").Text()
		//result := fmt.Sprint(date,": ", title)
		var result []string
		result = append(result, date, title, description)
		c <- result
	})
	wg.Done()
}
