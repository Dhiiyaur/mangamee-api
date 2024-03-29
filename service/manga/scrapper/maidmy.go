package mangaservice

import (
	"fmt"
	"mangamee-api/entity"
	"mangamee-api/exception"
	"mangamee-api/helper"
	"mangamee-api/logger"

	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func MaidmyIndex(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML("body > main > div > div.container > div.flexbox4 > div.flexbox4-item", func(e *colly.HTMLElement) {

		var checkLastChapter string

		tempLastChapter := strings.Split(e.ChildText("li > a"), " ")[1]
		if strings.Contains(tempLastChapter, "Ch.") {
			checkLastChapter = strings.Split(tempLastChapter, "C")[0]
		} else {
			checkLastChapter = tempLastChapter
		}
		returnData = append(returnData, entity.MangaData{
			Title:       e.ChildAttr("a", "title"),
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: checkLastChapter,
		})
	})

	err := c.Visit("https://www.maid.my.id/page/" + params.PageNumber + "/")

	if err != nil {
		logger.Info("MaidMyIndex Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MaidmySearch(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()

	c.OnHTML("body > main > div > div > div.flexbox2 > div.flexbox2-item", func(e *colly.HTMLElement) {

		var checkLastChapter string

		tempLastChapter := strings.Split(e.ChildText("div.season"), " ")

		if len(tempLastChapter) > 1 {
			checkLastChapter = tempLastChapter[1]
		}

		returnData = append(returnData, entity.MangaData{
			Title:       e.ChildAttr("a", "title"),
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
			Cover:       e.ChildAttr("img", "src"),
			LastChapter: checkLastChapter,
		})
	})

	err := c.Visit("https://www.maid.my.id/?s=" + params.Search)

	if err != nil {
		logger.Info("MaidmySearch Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}

func MaidmyDetail(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".series-thumb", func(e *colly.HTMLElement) {
		returnData.Cover = e.ChildAttr(`img`, "src")
	})

	c.OnHTML(".series-title", func(e *colly.HTMLElement) {
		returnData.Title = e.ChildText(`h2`)
	})

	c.OnHTML(".series-synops", func(e *colly.HTMLElement) {
		returnData.Summary = e.Text

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

		chapters = append(chapters, entity.Chapter{
			Name: chapterName,
			Id:   strings.Split(e.ChildAttr(`a`, "href"), "/")[3],
		})

	})

	err := c.Visit("https://www.maid.my.id/manga/" + params.MangaId + "/")

	returnData.Chapters = chapters
	returnData.OriginalServer = "https://www.maid.my.id/manga/" + params.MangaId + "/"

	if err != nil {
		logger.Info("MaidmyDetail Detail", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MaidmyImage(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var dataImages []entity.Image
	c := colly.NewCollector()

	c.OnHTML(".reader-area img", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, entity.Image{
			Image: e.Attr("src"),
		})

	})

	err := c.Visit("https://www.maid.my.id/" + params.ChapterId + "/")

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	returnData.Images = entity.DataChapters{
		ChapterName: re.FindAllString(params.ChapterId, -1)[0],
		Images:      dataImages,
	}
	returnData.OriginalServer = "https://www.maid.my.id/" + params.ChapterId + "/"

	if err != nil {
		logger.Info("MaidmyImage Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MaidmyChapter(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter

	c := colly.NewCollector()

	c.OnHTML(".flexch-infoz", func(e *colly.HTMLElement) {

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		tempName := strings.ReplaceAll(re.FindAllString(e.ChildAttr(`a`, "title"), -1)[0], "-", "")
		chapters = append(chapters, entity.Chapter{
			Name: tempName,
			Id:   strings.Split(e.ChildAttr(`a`, "href"), "/")[3],
		})

	})

	err := c.Visit("https://www.maid.my.id/manga/" + params.MangaId + "/")

	returnData.Chapters = chapters

	if err != nil {
		logger.Info("MaidmyChapter Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MaidmyMetaTag(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData

	c := colly.NewCollector()

	c.OnHTML(".series-thumb", func(e *colly.HTMLElement) {
		returnData.Cover = e.ChildAttr(`img`, "src")
	})

	c.OnHTML(".series-title", func(e *colly.HTMLElement) {

		returnData.Title = e.ChildText(`h2`)
	})

	err := c.Visit("https://www.maid.my.id/manga/" + params.MangaId + "/")

	if err != nil {
		logger.Info("MaidmyMetaTag Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}
