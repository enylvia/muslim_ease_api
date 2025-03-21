package model

type Surah struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	NumberOfAyah  int64  `json:"number_of_ayah"`
	NumberOfSurah int64  `json:"number_of_surah"`
	Place         string `json:"place"`
	Type          string `json:"type"`
}
type ParamGetSurah struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Search string `form:"search"`
}
