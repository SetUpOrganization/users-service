package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func cleanCSVValue(value string) string {
	value = strings.ReplaceAll(value, "&nbsp;", " ")
	value = strings.ReplaceAll(value, "\u00A0", " ")
	return strings.TrimSpace(value)
}

func ReplaceFirstSubTag(html string) string {
	html = strings.ReplaceAll(html, "<sub>", "_")

	return strings.ReplaceAll(html, "</sub>", "")
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("URL не указан в .env файле")
	}

	c := colly.NewCollector()
	file, err := os.Create("source/dynamic_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Название", "Формула", "H", "G", "S", "C"}
	writer.Write(headers)

	c.OnHTML(".tablepress-id-3", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, row *colly.HTMLElement) {
			rowForWrite := []string{}

			for col := 1; col <= len(headers); col++ {
				if col == 2 {
					htmlContent, _ := row.DOM.Find(fmt.Sprintf("td.column-%d", col)).Html()
					data := ReplaceFirstSubTag(htmlContent)
					rowForWrite = append(rowForWrite, cleanCSVValue(data))
					continue
				}

				data := row.ChildText(fmt.Sprintf("td.column-%d", col))

				rowForWrite = append(rowForWrite, cleanCSVValue(data))
			}

			writer.Write(rowForWrite)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Посещаю:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Ошибка:", err)
	})

	c.Visit(url)
}
