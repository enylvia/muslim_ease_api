package repository

import (
	"go-api-quran/model"
	"gorm.io/gorm"
)

type SurahRepository interface {
	GetDB() *gorm.DB
	GetListSurah(param model.ParamGetSurah) ([]model.Surah, int64, error)
}
type SurahRepositoryImplement struct {
	db *gorm.DB
}

func NewSurahRepository(db *gorm.DB) SurahRepository {
	return &SurahRepositoryImplement{db: db}
}
func (sr *SurahRepositoryImplement) GetDB() *gorm.DB {
	return sr.db
}

func (sr *SurahRepositoryImplement) GetListSurah(param model.ParamGetSurah) ([]model.Surah, int64, error) {
	var surahList []model.Surah
	var count int64
	baseQuery := sr.db.Table("surahs")

	if param.Search != "" {
		baseQuery = baseQuery.Where("name ILIKE ?", "%"+param.Search+"%")
	}
	if err := baseQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// Salin ulang baseQuery untuk ambil data paginated
	query := baseQuery
	if param.Limit != 0 {
		query = query.Limit(param.Limit)
	}
	if param.Offset != 0 {
		query = query.Offset(param.Offset)
	}
	if err := query.Find(&surahList).Error; err != nil {
		return nil, 0, err
	}
	return surahList, count, nil
}
