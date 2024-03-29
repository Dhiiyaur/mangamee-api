package mangaservice

import (
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

func MangareadIndex(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML(".page-item-detail.manga", func(e *colly.HTMLElement) {

		var coverImage string
		checkImage := strings.Split(e.ChildAttr("a > img", "data-src"), " ")
		if len(checkImage) < 2 {
			coverImage = e.ChildAttr("a > img", "data-src")
		} else {
			coverImage = checkImage[len(checkImage)-2]
		}
		returnData = append(returnData, entity.MangaData{

			Cover:       coverImage,
			Title:       e.ChildAttr("a", "title"),
			LastChapter: strings.Split(e.ChildText("span.chapter.font-meta > a"), " ")[1],
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/?m_orderby=new-manga&page=" + params.PageNumber)

	if err != nil {
		logger.Info("MangareadIndex Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MangareadSearch(ctx context.Context, params entity.RequestParams) ([]entity.MangaData, error) {

	var returnData []entity.MangaData

	c := colly.NewCollector()
	c.OnHTML(".row.c-tabs-item__content", func(e *colly.HTMLElement) {

		var lastChapter string
		checkChapter := strings.Split(e.ChildText("span.font-meta.chapter > a"), " ")

		if len(checkChapter) > 2 {
			lastChapter = checkChapter[len(checkChapter)-2]
		} else {
			lastChapter = checkChapter[len(checkChapter)-1]
		}

		returnData = append(returnData, entity.MangaData{
			Cover:       e.ChildAttr("a > img", "data-src"),
			Title:       e.ChildAttr("a", "title"),
			LastChapter: lastChapter,
			Id:          strings.Split(e.ChildAttr("a", "href"), "/")[4],
		})

	})

	err := c.Visit("https://www.mangaread.org/?s=" + params.Search + "&post_type=wp-manga")
	if err != nil {
		logger.Info("MangareadSearch Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}

func MangareadDetail(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter
	limit := 0

	c := colly.NewCollector()

	c.OnHTML(".post-title", func(e *colly.HTMLElement) {

		if limit == 0 {
			returnData.Title = strings.Split(e.ChildText("h1"), "  ")[0]
		}
		limit++
	})

	c.OnHTML(".summary_image", func(e *colly.HTMLElement) {
		returnData.Cover = e.ChildAttr("img", "data-src")
	})

	c.OnHTML(".summary__content", func(e *colly.HTMLElement) {
		returnData.Summary = e.ChildText("p")
	})

	c.OnHTML(".wp-manga-chapter", func(e *colly.HTMLElement) {

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		tempName := strings.ReplaceAll(re.FindAllString(e.ChildText("a"), -1)[0], "-", "")

		chapters = append(chapters, entity.Chapter{
			Name: tempName,
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[5],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/" + params.MangaId)

	returnData.Chapters = chapters
	returnData.OriginalServer = "https://www.mangaread.org/manga/" + params.MangaId

	if err != nil {
		logger.Info("MangareadDetail Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}

	return returnData, nil
}

func MangareadImage(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var dataImages []entity.Image
	c := colly.NewCollector()

	c.OnHTML(".wp-manga-chapter-img", func(e *colly.HTMLElement) {

		dataImages = append(dataImages, entity.Image{
			Image: "https://" + strings.Split(e.Attr("data-src"), "//")[1],
		})

	})
	err := c.Visit("https://www.mangaread.org/manga/" + params.MangaId + "/" + params.ChapterId)

	re := regexp.MustCompile(`\d+`)
	returnData.Images = entity.DataChapters{
		ChapterName: re.FindAllString(params.ChapterId, -1)[0],
		Images:      dataImages,
	}

	returnData.OriginalServer = "https://www.mangaread.org/manga/" + params.MangaId + "/" + params.ChapterId
	if err != nil {
		logger.Info("MangareadImage Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}

func MangareadChapter(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	var chapters []entity.Chapter
	c := colly.NewCollector()

	c.OnHTML(".wp-manga-chapter", func(e *colly.HTMLElement) {

		re := regexp.MustCompile(`\d+`)
		tempName := re.FindAllString(e.ChildText("a"), -1)[0]

		chapters = append(chapters, entity.Chapter{
			Name: tempName,
			Id:   strings.Split(e.ChildAttr("a", "href"), "/")[5],
		})
	})

	err := c.Visit("https://www.mangaread.org/manga/" + params.MangaId)

	returnData.Chapters = chapters

	if err != nil {
		logger.Info("MangareadChapter Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}

func MangareadMetaTag(ctx context.Context, params entity.RequestParams) (entity.MangaData, error) {

	var returnData entity.MangaData
	limit := 0

	c := colly.NewCollector()
	c.OnHTML(".post-title", func(e *colly.HTMLElement) {

		if limit == 0 {
			returnData.Title = strings.Split(e.ChildText("h1"), "  ")[0]
		}
		limit++
	})

	c.OnHTML(".summary_image", func(e *colly.HTMLElement) {
		returnData.Cover = e.ChildAttr("img", "data-src")
	})
	err := c.Visit("https://www.mangaread.org/manga/" + params.MangaId)
	if err != nil {
		logger.Info("MangareadMetaTag Error", zap.Any("request_id", helper.GetRequestId(ctx)), zap.Error(err))
		return returnData, exception.NewInternal()
	}
	return returnData, nil
}
