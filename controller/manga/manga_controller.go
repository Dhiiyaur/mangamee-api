package mangacontroller

import "github.com/gin-gonic/gin"

type MangaController interface {
	GetIndex(c *gin.Context)
	GetSearch(c *gin.Context)
	GetDetail(c *gin.Context)
	GetImage(c *gin.Context)
	GetChapter(c *gin.Context)
	GetSource(c *gin.Context)
	GetMetaTag(c *gin.Context)
	GetMangaProxy(c *gin.Context)
}
