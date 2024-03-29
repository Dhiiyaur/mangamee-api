package mangaservice

import (
	"fmt"
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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

func MangatownIndex(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML("body > section > article > div > div.manga_pic_content > ul.manga_pic_list > li", func(e *colly.HTMLElement) {
		var mangaId, lastChapter string

		mangaIdCheck := strings.Split(e.ChildAttr("a", "href"), "/")
		mangaId = mangaIdCheck[len(mangaIdCheck)-2]

		lastChapterCheck := strings.Split(e.ChildText("p.new_chapter"), " ")
		lastChapter = lastChapterCheck[len(lastChapterCheck)-1]

		mangaCoverCheck := strings.Replace(e.ChildAttr("a > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		returnData = append(returnData, entity.MangaData{
			Id:          mangaId,
			Title:       e.ChildAttr("a", "title"),
			Cover:       mangaCoverCheck,
			LastChapter: lastChapter,
		})

	})
	err := c.Visit("https://www.mangatown.com/hot/" + params.PageNumber + ".htm?wviews.za")

	if err != nil {
		logger.Info("MangatownIndex Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil

}

func MangatownSearch(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML(".manga_pic_list > li", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("a.manga_cover > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		returnData = append(returnData, entity.MangaData{
			Cover: mangaCoverCheck,
			Title: e.ChildAttr("a.manga_cover", "title"),
			Id:    strings.Split(e.ChildAttr("a.manga_cover", "href"), "/")[2],
		})

	})

	err := c.Visit("https://www.mangatown.com/search?name=" + params.Search)

	if err != nil {
		logger.Info("MangatownSearch Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MangatownDetail(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".article_content", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("div.detail_info.clearfix > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		returnData.Title = e.ChildText("h1.title-top")
		returnData.Cover = mangaCoverCheck
		returnData.Summary = e.ChildText("div.detail_info.clearfix > ul > li > span")

	})

	c.OnHTML(".chapter_list > li", func(e *colly.HTMLElement) {

		var chapterName string

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		arr := re.FindAllString(e.ChildText("a"), -1)
		if len(arr) != 0 {
			chapterName = arr[len(arr)-1]
		} else {
			chapterName = "0"
		}

		chapters = append(chapters, entity.Chapter{
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[3],
			Name: chapterName,
		})
	})

	err := c.Visit("https://www.mangatown.com/manga/" + params.MangaId)

	returnData.Chapters = chapters
	returnData.OriginalServer = "https://www.mangatown.com/manga/" + params.MangaId

	if err != nil {
		logger.Info("MangatownDetail Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MangatownImage(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var dataImages []entity.Image
	var link string

	c := colly.NewCollector()

	c.OnHTML(".read_img", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("img", "src"), "zjcdn.mangahere.org", "fmcdn.mangatown.com", -1)

		link = "https:" + mangaCoverCheck

	})

	err := c.Visit("https://www.mangatown.com/manga/" + params.MangaId + "/" + params.ChapterId + "/")

	baseLink, imageLink := ReturnLastSliceAndJoinLink(link)
	imageExtension, frontRawData, loopData := GetRawImageData(imageLink)

	for i := 0; i < 100; i++ {
		tempNumber := loopData + i
		if tempNumber < 10 {
			a := fmt.Sprintf("%v00%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, entity.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 100 && tempNumber > 9 {
			a := fmt.Sprintf("%v0%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, entity.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 1000 && tempNumber > 99 {
			a := fmt.Sprintf("%v%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, entity.Image{
				Image: joinImageLink,
			})

		} else if tempNumber < 10000 && tempNumber > 999 {
			a := fmt.Sprintf("%v%v.%v", frontRawData, strconv.Itoa(tempNumber), imageExtension)
			joinImageLink := fmt.Sprintf("%v/%v", baseLink, a)
			dataImages = append(dataImages, entity.Image{
				Image: joinImageLink,
			})

		}
	}

	re := regexp.MustCompile(`\d+`)
	returnData.Images = entity.DataChapters{
		ChapterName: re.FindAllString(params.ChapterId, -1)[0],
		Images:      dataImages,
	}
	returnData.OriginalServer = "https://www.mangatown.com/manga/" + params.MangaId + "/" + params.ChapterId + "/"

	if err != nil {
		logger.Info("MangatownImage Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}

func MangatownChapter(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".chapter_list > li", func(e *colly.HTMLElement) {

		var chapterName string

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		arr := re.FindAllString(e.ChildText("a"), -1)
		if len(arr) != 0 {
			chapterName = arr[len(arr)-1]
		} else {
			chapterName = "0"
		}

		chapters = append(chapters, entity.Chapter{
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[3],
			Name: chapterName,
		})
	})

	err := c.Visit("https://www.mangatown.com/manga/" + params.MangaId)

	returnData.Chapters = chapters

	if err != nil {
		logger.Info("MangatownChapter Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MangatownMetaTag(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData

	c := colly.NewCollector()

	c.OnHTML(".article_content", func(e *colly.HTMLElement) {

		mangaCoverCheck := strings.Replace(e.ChildAttr("div.detail_info.clearfix > img", "src"), "https://fmcdn.mangahere.com/", "http://fmcdn.mangatown.com/", -1)

		returnData.Title = e.ChildText("h1.title-top")
		returnData.Cover = mangaCoverCheck

	})

	err := c.Visit("https://www.mangatown.com/manga/" + params.MangaId)

	if err != nil {
		logger.Info("MangatownMetaTag Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}
