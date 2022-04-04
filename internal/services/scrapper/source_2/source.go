package source_2

import (
	"mangamee-api/internal/models"
	"strings"

	"github.com/gocolly/colly"
)

func MangaSearch(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()
	c.OnHTML(".manga_pic_list > li", func(e *colly.HTMLElement) {

		dataMangas = append(dataMangas, models.MangaData{
			Cover: e.ChildAttr("a.manga_cover > img", "src"),
			Title: e.ChildAttr("a.manga_cover", "title"),
			Id:    strings.Split(e.ChildAttr("a.manga_cover", "href"), "/")[2],
		})

	})

	err := c.Visit("https://www.mangatown.com/search?name=" + queryParams.Search)
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}
