package source_2

import (
	"fmt"
	"mangamee-api/internal/models"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func ReturnLastSliceAndJoinLink(s string) (string, string) {
	slice := strings.Split(s, "/")
	return strings.Join(slice[:len(slice)-1], "/"), slice[len(slice)-1]
}

func GetRawImageData(s string) (string, string, int) {

	var imageExtension, frontRawData string
	var loopData int

	a := strings.Split(s, ".")
	imageExtension = a[len(a)-1]

	if strings.Contains(s, "_") {
		b := strings.Split(a[0], "_")
		loopData, _ = strconv.Atoi(b[len(b)-1])

		if len(b) > 2 {
			frontRawData = fmt.Sprintf("%v_%v_", b[0], b[1])
		} else {
			frontRawData = fmt.Sprintf("%v_", b[0])
		}

	} else {
		frontRawData = a[0][0:1]
		loopData, _ = strconv.Atoi(a[0][1:])
	}

	return imageExtension, frontRawData, loopData

}

//

func MangaIndex(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()
	c.OnHTML("body > section > article > div > div.manga_pic_content > ul.manga_pic_list > li", func(e *colly.HTMLElement) {
		var mangaId, lastChapter string

		mangaIdCheck := strings.Split(e.ChildAttr("a", "href"), "/")
		mangaId = mangaIdCheck[len(mangaIdCheck)-2]

		lastChapterCheck := strings.Split(e.ChildText("p.new_chapter"), " ")
		lastChapter = lastChapterCheck[len(lastChapterCheck)-1]

		mangaCoverCheck := strings.Replace(e.ChildAttr("a > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		dataMangas = append(dataMangas, models.MangaData{
			Id:          mangaId,
			Title:       e.ChildAttr("a", "title"),
			Cover:       mangaCoverCheck,
			LastChapter: lastChapter,
		})

	})
	err := c.Visit("https://www.mangatown.com/hot/" + queryParams.Page + ".htm?wviews.za")
	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil

}

func MangaSearch(queryParams models.QueryParams) ([]models.MangaData, error) {

	var dataMangas []models.MangaData

	c := colly.NewCollector()
	c.OnHTML(".manga_pic_list > li", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("a.manga_cover > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		dataMangas = append(dataMangas, models.MangaData{
			Cover: mangaCoverCheck,
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

func MangaDetail(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var chapters []models.Chapter

	c := colly.NewCollector()

	c.OnHTML(".article_content", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("div.detail_info.clearfix > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		dataMangas.Title = e.ChildText("h1.title-top")
		dataMangas.Cover = mangaCoverCheck
		dataMangas.Summary = e.ChildText("div.detail_info.clearfix > ul > li > span")

	})

	c.OnHTML(".chapter_list > li", func(e *colly.HTMLElement) {

		chapters = append(chapters, models.Chapter{
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[3],
			Name: strings.Split(e.ChildAttr("a", "href"), "/")[3],
		})
	})

	err := c.Visit("https://www.mangatown.com/manga/" + queryParams.Id)

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}

func MangaImage(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData
	var dataImages []models.Image
	var link string

	c := colly.NewCollector()

	c.OnHTML(".read_img", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("img", "src"), "zjcdn.mangahere.org", "fmcdn.mangatown.com", -1)

		link = "https:" + mangaCoverCheck

	})

	err := c.Visit("https://www.mangatown.com/manga/" + queryParams.Id + "/" + queryParams.ChapterId + "/")

	baseLink, imageLink := ReturnLastSliceAndJoinLink(link)
	imageExtension, frontRawData, loopData := GetRawImageData(imageLink)

	for i := 0; i < 100; i++ {
		tempNumber := loopData + i
		if tempNumber < 10 {
			a := fmt.Sprintf("%v00%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, models.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 100 && tempNumber > 9 {
			a := fmt.Sprintf("%v0%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, models.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 1000 && tempNumber > 99 {
			a := fmt.Sprintf("%v%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, models.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 10000 && tempNumber > 999 {
			a := fmt.Sprintf("%v%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, models.Image{
				Image: joinImageLink,
			})

		}
	}

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

	c.OnHTML(".chapter_list > li", func(e *colly.HTMLElement) {

		chapters = append(chapters, models.Chapter{
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[3],
			Name: strings.Split(e.ChildAttr("a", "href"), "/")[3],
		})
	})

	err := c.Visit("https://www.mangatown.com/manga/" + queryParams.Id)

	dataMangas.Chapters = chapters

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil

}

func MangaMetaTag(queryParams models.QueryParams) (models.MangaData, error) {

	var dataMangas models.MangaData

	c := colly.NewCollector()

	c.OnHTML(".article_content", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("div.detail_info.clearfix > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		dataMangas.Title = e.ChildText("h1.title-top")
		dataMangas.Cover = mangaCoverCheck

	})

	err := c.Visit("https://www.mangatown.com/manga/" + queryParams.Id)

	if err != nil {
		return dataMangas, err
	}

	return dataMangas, nil
}
