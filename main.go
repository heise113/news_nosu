package main

import (
	"fmt"
	"news_nosu/internal/browser"
	"news_nosu/internal/telegram"
	"news_nosu/internal/transform_page"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	news := transform_page.GetNews()
	fmt.Println(news)

	go telegram.Start(wg, news)
	go browser.Start(wg, news)

	wg.Wait()
}
