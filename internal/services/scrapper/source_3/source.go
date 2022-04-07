package source_3

import (
	"fmt"
	"mangamee-api/internal/models"
	"strings"

	"github.com/gocolly/colly"
)

func MangaIndex(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()

	c.OnHTML("body > main > div > div.container > div.flexbox4 > div.flexbox4-item", func(e *colly.HTMLElement) {

		var checkLastChapter string

		tempLastChapter := strings.Split(e.ChildText("li > a"), " ")[1]
		if strings.Contains(tempLastChapter, "Ch.") {
			checkLastChapter = strings.Split(tempLastChapter, "C")[0]
		} else {
			checkLastChapter = tempLastChapter
		}
		dataMangas = append(dataMangas, models.MangaData{
			Title:       e.ChildAttr("a", "title"),
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: checkLastChapter,
		})
	})

	err := c.Visit("https://www.maid.my.id/page/" + queryParams.Page + "/")
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaSearch(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()

	c.OnHTML("body > main > div > div > div.flexbox2 > div.flexbox2-item", func(e *colly.HTMLElement) {

		var checkLastChapter string

		tempLastChapter := strings.Split(e.ChildText("div.season"), " ")

		if len(tempLastChapter) > 1 {
			checkLastChapter = tempLastChapter[1]
		}

		dataMangas = append(dataMangas, models.MangaData{
			Title:       e.ChildAttr("a", "title"),
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: checkLastChapter,
		})
	})

	err := c.Visit("https://www.maid.my.id/?s=" + queryParams.Search)

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaDetail(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var chapters []models.Chapter

	c := colly.NewCollector()

	c.OnHTML(".series-thumb", func(e *colly.HTMLElement) {
		dataMangas.Cover = e.ChildAttr(`img`, "src")
	})

	c.OnHTML(".series-title", func(e *colly.HTMLElement) {

		dataMangas.Title = e.ChildText(`h2`)
	})

	c.OnHTML(".series-synops", func(e *colly.HTMLElement) {

		dataMangas.Summary = e.Text

	})

	c.OnHTML(".flexch-infoz", func(e *colly.HTMLElement) {

		var chapterName string
		tempChapterName := e.ChildAttr(`a`, "title")

		if strings.Contains(tempChapterName, "Bahasa Indonesia") {
			a := strings.Split(tempChapterName, "Bahasa Indonesia")
			b := strings.Split(a[len(a)-2], " ")
			chapterName = fmt.Sprintf("%v %v", b[len(b)-3], b[len(b)-2])

		} else {
			a := strings.Split(tempChapterName, " ")
			chapterName = fmt.Sprintf("%v %v", a[len(a)-2], a[len(a)-1])
		}

		chapters = append(chapters, models.Chapter{
			Name: chapterName,
			Id:   strings.Split(e.ChildAttr(`a`, "href"), "/")[3],
		})

	})

	err := c.Visit("https://www.maid.my.id/manga/" + queryParams.Id + "/")

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

	c.OnHTML(".reader-area img", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, models.Image{
			Image: e.Attr("src"),
		})

	})

	err := c.Visit("https://www.maid.my.id/" + queryParams.ChapterId + "/")
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

	c.OnHTML(".flexch-infoz", func(e *colly.HTMLElement) {

		var chapterName string
		tempChapterName := e.ChildAttr(`a`, "title")

		if strings.Contains(tempChapterName, "Bahasa Indonesia") {
			a := strings.Split(tempChapterName, "Bahasa Indonesia")
			b := strings.Split(a[len(a)-2], " ")
			chapterName = fmt.Sprintf("%v %v", b[len(b)-3], b[len(b)-2])

		} else {
			a := strings.Split(tempChapterName, " ")
			chapterName = fmt.Sprintf("%v %v", a[len(a)-2], a[len(a)-1])
		}

		chapters = append(chapters, models.Chapter{
			Name: chapterName,
			Id:   strings.Split(e.ChildAttr(`a`, "href"), "/")[3],
		})

	})

	err := c.Visit("https://www.maid.my.id/manga/" + queryParams.Id + "/")

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}
