package models

type Patch struct {
	RegNum     string `json:"regNum,omitempty"`
	Mark       string `json:"mark,omitempty"`
	Model      string `json:"model,omitempty"`
	Year       int    `json:"year,omitempty"`
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
}
