package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-quran/model"
	"go-api-quran/service"
	"net/http"
)

type SurahHandler interface {
	ListSurah(c *gin.Context)
}

type SurahHandlerImplement struct {
	surahService service.SurahService
}

func NewSurahHandler(surahService service.SurahService) SurahHandler {
	return &SurahHandlerImplement{surahService: surahService}
}

// ListSurah godoc
// @Summary Get list of surah
// @Description Retrieve a list of surah from the Quran with optional search and pagination
// @Tags Surah
// @Accept  json
// @Produce  json
// @Param search query string false "Search by surah name"
// @Param limit query int false "Limit number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} model.ResponseWithMeta[model.Surah]
// @Failure 400 {object} model.Response[any]
// @Router /surah [get]
func (s SurahHandlerImplement) ListSurah(c *gin.Context) {
	var param model.ParamGetSurah
	if err := c.ShouldBindQuery(&param); err != nil {
		model.Error(c, http.StatusBadRequest, "Invalid query parameter")
		return
	}
	data, count, err := s.surahService.ListSurah(param)
	if err != nil {
		model.Error(c, http.StatusBadRequest, fmt.Sprintf("Failed to retrieve surah data: %s", err.Error()))
		return
	}
	model.PaginatedSuccess(c, count, data, "Successfully retrieved surah data")
	return
}
