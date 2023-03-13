package browser

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

//func date(c echo.Context) error {
//for i, n := range news {
//	if n[0] == update.Message.Text {
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n\n%s", n[0], n[1]))
//		bot.Send(msg)
//		break
//	}
//	if i == len(news)-1 {
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Новостей по данной дате не найдено!")
//		bot.Send(msg)
//	}
//}
//return c.String(http.StatusOK, fmt.Sprintf("%s", news[0][0]))
//}

func Start(wg sync.WaitGroup, news [][]string) {
	defer wg.Done()
	e := echo.New()
	//e.GET("/", date)
	e.GET("/", func(c echo.Context) error {
		for i, n := range news {
			if n[0] == c.QueryParam("data") {
				return c.HTML(http.StatusOK, fmt.Sprintf("<p>%s</p><h2>%s</h2><p>%s</p>", n[0], n[1], n[2]))
			}
			if i == len(news)-1 {
				return c.String(http.StatusOK, "Новостей по данной дате не найдено.")
			}
		}
		return c.String(http.StatusOK, "lalala")
	})
	e.Logger.Fatal(e.Start(":7777"))
}
