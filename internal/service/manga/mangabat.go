package mangaservice

import (
	"fmt"
	"mangamee-api/internal/entity"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func MangabatIndex(params entity.MangaParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML(".list-story-item", func(e *colly.HTMLElement) {

		tempLastChapter := strings.Split(e.ChildAttr("div > a:nth-child(2)", "href"), "-")
		tempMangaID := strings.Split(e.ChildAttr("a.item-img", "href"), "/")

		returnData = append(returnData, entity.MangaData{
			Title:       e.ChildAttr("img", "alt"),
			Id:          tempMangaID[len(tempMangaID)-1],
			Cover:       e.ChildAttr("a > img", "src"),
			LastChapter: tempLastChapter[len(tempLastChapter)-1],
		})
	})

	err := c.Visit("https://m.mangabat.com/manga-list-all/" + params.PageNumber + "/")
	if err != nil {
		return returnData, err
	}

	return returnData, nil
}

func MangabatSearch(params entity.MangaParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()

	c.OnHTML(".list-story-item", func(e *colly.HTMLElement) {

		tempLastChapter := strings.Split(e.ChildAttr("div > a:nth-child(2)", "href"), "-")
		tempMangaID := strings.Split(e.ChildAttr("a.item-img", "href"), "/")

		returnData = append(returnData, entity.MangaData{
			Title:       e.ChildAttr("img", "alt"),
			Id:          tempMangaID[len(tempMangaID)-1],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: tempLastChapter[len(tempLastChapter)-1],
		})
	})

	err := c.Visit("https://m.mangabat.com/search/manga/" + params.Search)

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}

func MangabatDetail(params entity.MangaParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".panel-story-info", func(e *colly.HTMLElement) {

		returnData.Cover = e.ChildAttr("span > img", "src")
		returnData.Title = e.ChildText("div.story-info-right > h1")
		returnData.Summary = e.ChildText("div.panel-story-info-description")

	})

	c.OnHTML(".chapter-name", func(e *colly.HTMLElement) {

		tempMangaID := strings.Split(e.Attr("href"), "/")
		tempMangaName := strings.Split(e.Attr("href"), "-")

		chapters = append(chapters, entity.Chapter{
			Name: tempMangaName[len(tempMangaName)-1],
			Id:   tempMangaID[len(tempMangaID)-1],
		})

	})

	err := c.Visit("https://readmangabat.com/" + params.MangaId + "/")

	returnData.Chapters = chapters
	returnData.OriginalServer = "https://readmangabat.com/" + params.MangaId + "/"

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}

func MangabatImage(params entity.MangaParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var dataImages []entity.Image
	var name string

	c := colly.NewCollector()

	c.OnHTML(".img-content", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, entity.Image{
			Image: fmt.Sprintf("%vproxy?id=%v", "https://api.mangamee.space/manga/", e.Attr("src")),
		})

	})

	err := c.Visit("https://readmangabat.com/" + params.ChapterId + "/")

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	if strings.Contains(params.ChapterId, "chap") {
		tmp := strings.Split(params.ChapterId, "chap")
		name = re.FindAllString(tmp[len(tmp)-1], -1)[0]
	} else {
		name = re.FindAllString(params.ChapterId, -1)[0]
	}

	returnData.Images = entity.DataChapters{
		ChapterName: name,
		Images:      dataImages,
	}

	returnData.OriginalServer = "https://readmangabat.com/" + params.ChapterId + "/"

	if err != nil {
		return returnData, err
	}
	return returnData, nil
}

func MangabatChapter(params entity.MangaParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".chapter-name", func(e *colly.HTMLElement) {

		tempMangaID := strings.Split(e.Attr("href"), "/")
		tempMangaName := strings.Split(e.Attr("href"), "-")

		chapters = append(chapters, entity.Chapter{
			Name: tempMangaName[len(tempMangaName)-1],
			Id:   tempMangaID[len(tempMangaID)-1],
		})

	})

	err := c.Visit("https://readmangabat.com/" + params.MangaId + "/")

	returnData.Chapters = chapters

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}

func MangabatMetaTag(params entity.MangaParams) (entity.MangaData, error) {

	var returnData entity.MangaData

	c := colly.NewCollector()

	c.OnHTML(".panel-story-info", func(e *colly.HTMLElement) {

		returnData.Cover = e.ChildAttr("span > img", "src")
		returnData.Title = e.ChildText("div.story-info-right > h1")

	})

	err := c.Visit("https://readmangabat.com/" + params.MangaId + "/")

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}
