package htmlparser

import (
	"errors"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetDocument(lnk string) (doc *goquery.Document, err error) {

	// Читаем страницу
	resp, err := http.Get(lnk)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Что-то пошло не так
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", resp.Status, resp.StatusCode)
		return
	}

	// Парсим документ
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}
