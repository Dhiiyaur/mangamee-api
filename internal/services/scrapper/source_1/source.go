package source_1

import (
	"strings"

	"mangamee-api/internal/models"

	"github.com/gocolly/colly"
)

func MangaIndex(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()
	c.OnHTML(".page-item-detail.manga", func(e *colly.HTMLElement) {

		tempDataCover := strings.Split(e.ChildAttr("a > img", "data-srcset"), " ")

		dataMangas = append(dataMangas, models.MangaData{

			Cover:       tempDataCover[len(tempDataCover)-2],
			Title:       e.ChildAttr("a", "title"),
			LastChapter: strings.Split(e.ChildText("span.chapter.font-meta > a"), " ")[1],
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/?m_orderby=new-manga&page=" + queryParams.Page)
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaSearch(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()
	c.OnHTML(".row.c-tabs-item__content", func(e *colly.HTMLElement) {

		dataMangas = append(dataMangas, models.MangaData{
			Cover:       e.ChildText("span.font-meta.chapter > a"),
			Title:       e.ChildAttr("a", "title"),
			LastChapter: e.ChildText("span.font-meta.chapter > a"),
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
		})

	})

	err := c.Visit("https://www.mangaread.org/?s=" + queryParams.Search + "&post_type=wp-manga")
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaDetail(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var chapters []models.Chapter
	limit := 0

	c := colly.NewCollector()

	c.OnHTML(".post-title", func(e *colly.HTMLElement) {

		if limit == 0 {
			dataMangas.Title = strings.Split(e.ChildText("h1"), "  ")[0]
		}
		limit++
	})

	c.OnHTML(".summary_image", func(e *colly.HTMLElement) {
		dataMangas.Cover = e.ChildAttr("img", "data-src")
	})

	c.OnHTML(".summary__content", func(e *colly.HTMLElement) {
		dataMangas.Summary = e.ChildText("p")
	})

	c.OnHTML(".wp-manga-chapter", func(e *colly.HTMLElement) {

		chapters = append(chapters, models.Chapter{
			Name: e.ChildText("a"),
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[5],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/" + queryParams.Id)

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaImage(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var dataImages []models.Image
	c := colly.NewCollector()

	c.OnHTML(".wp-manga-chapter-img", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, models.Image{
			Image: "https://" + strings.Split(e.Attr("data-src"), "//")[1],
		})

	})
	err := c.Visit("https://www.mangaread.org/manga/" + queryParams.Id + "/" + queryParams.ChapterId)

	dataMangas.Images = dataImages

	if err != nil {
		return dataMangas, err
	}
	return dataMangas, nil

}

func MangaChapter(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var chapters []models.Chapter
	c := colly.NewCollector()

	c.OnHTML(".wp-manga-chapter", func(e *colly.HTMLElement) {

		chapters = append(chapters, models.Chapter{
			Name: e.ChildText("a"),
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[5],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/" + queryParams.Id)

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}
