package source_4

import (
	"fmt"
	"mangamee-api/internal/models"
	"strings"

	"github.com/gocolly/colly"
)

func MangaIndex(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()

	c.OnHTML(".list-story-item", func(e *colly.HTMLElement) {

		tempLastChapter := strings.Split(e.ChildAttr("div > a:nth-child(2)", "href"), "-")
		tempMangaID := strings.Split(e.ChildAttr("a.item-img", "href"), "/")

		dataMangas = append(dataMangas, models.MangaData{
			Title:       e.ChildAttr("img", "alt"),
			Id:          tempMangaID[len(tempMangaID)-1],
			Cover:       e.ChildAttr("a > img", "src"),
			LastChapter: tempLastChapter[len(tempLastChapter)-1],
		})
	})

	err := c.Visit("https://m.mangabat.com/manga-list-all/" + queryParams.Page + "/")
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaSearch(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()

	c.OnHTML(".list-story-item", func(e *colly.HTMLElement) {

		tempLastChapter := strings.Split(e.ChildAttr("div > a:nth-child(2)", "href"), "-")
		tempMangaID := strings.Split(e.ChildAttr("a.item-img", "href"), "/")

		dataMangas = append(dataMangas, models.MangaData{
			Title:       e.ChildAttr("img", "alt"),
			Id:          tempMangaID[len(tempMangaID)-1],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: tempLastChapter[len(tempLastChapter)-1],
		})
	})

	err := c.Visit("https://m.mangabat.com/search/manga/" + queryParams.Search)

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaDetail(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var chapters []models.Chapter

	c := colly.NewCollector()

	c.OnHTML(".panel-story-info", func(e *colly.HTMLElement) {

		dataMangas.Cover = e.ChildAttr("span > img", "src")
		dataMangas.Title = e.ChildText("div.story-info-right > h1")
		dataMangas.Summary = e.ChildText("div.panel-story-info-description")

	})

	c.OnHTML(".chapter-name", func(e *colly.HTMLElement) {

		tempMangaID := strings.Split(e.Attr("href"), "/")
		tempMangaName := strings.Split(e.Attr("href"), "-")

		chapters = append(chapters, models.Chapter{
			Name: tempMangaName[len(tempMangaName)-1],
			Id:   tempMangaID[len(tempMangaID)-1],
		})

	})

	err := c.Visit("https://readmangabat.com/" + queryParams.Id + "/")

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

	c.OnHTML(".img-content", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, models.Image{
			Image: fmt.Sprintf("%vproxy/4?id=%v", "https://mangamee-api.herokuapp.com/manga/", e.Attr("src")),
		})

	})

	err := c.Visit("https://readmangabat.com/" + queryParams.ChapterId + "/")

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

	c.OnHTML(".chapter-name", func(e *colly.HTMLElement) {

		tempMangaID := strings.Split(e.Attr("href"), "/")
		tempMangaName := strings.Split(e.Attr("href"), "-")

		chapters = append(chapters, models.Chapter{
			Name: tempMangaName[len(tempMangaName)-1],
			Id:   tempMangaID[len(tempMangaID)-1],
		})

	})

	err := c.Visit("https://readmangabat.com/" + queryParams.Id + "/")

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaMetaTag(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData

	c := colly.NewCollector()

	c.OnHTML(".panel-story-info", func(e *colly.HTMLElement) {

		dataMangas.Cover = e.ChildAttr("span > img", "src")
		dataMangas.Title = e.ChildText("div.story-info-right > h1")

	})

	err := c.Visit("https://readmangabat.com/" + queryParams.Id + "/")

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil

}
