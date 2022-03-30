package source_2

import (
	"strings"

	"github.com/dhiiyaur/go-mangamee-2/internal/models"
	"github.com/gocolly/colly"
)

func MangaIndex(queryParams models.QueryParams) ([]models.MangaData, error) {

	dataMangas := []models.MangaData{}

	c := colly.NewCollector()
	c.OnHTML(".page-item-detail.manga", func(e *colly.HTMLElement) {

		result := models.MangaData{
			Cover:       e.ChildAttr("a > img", "data-src"),
			Title:       e.ChildAttr("a", "title"),
			LastChapter: strings.Split(e.ChildText("span.chapter.font-meta > a"), " ")[1],
			Link:        strings.Split(e.ChildAttr("a", "href"), "/")[4],
		}

		dataMangas = append(dataMangas, result)
	})

	err := c.Visit("https://www.mangaread.org/manga/?m_orderby=new-manga&page=" + queryParams.SourcePage)
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil

}
