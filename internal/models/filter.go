package models

type Filter struct {
	Lyear      int    `json:"lyear"`
	Tyear      int    `json:"tyear"`
	Regnum     string `json:"regNum"`
	Mark       string `json:"mark"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
