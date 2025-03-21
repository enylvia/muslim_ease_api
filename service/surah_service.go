package service

import (
	"go-api-quran/model"
	"go-api-quran/repository"
)

type SurahService interface {
	ListSurah(param model.ParamGetSurah) ([]model.Surah, int64, error)
}
type SurahServiceImplement struct {
	surahRepository repository.SurahRepository
}

func NewSurahService(surahRepository repository.SurahRepository) SurahService {
	return &SurahServiceImplement{surahRepository: surahRepository}
}

func (s SurahServiceImplement) ListSurah(param model.ParamGetSurah) ([]model.Surah, int64, error) {
	return s.surahRepository.GetListSurah(param)
}
