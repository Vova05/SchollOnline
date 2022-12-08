package models

type Application struct {
	IdRequest  int    `json:"id"`
	DateStart  string `json:"date_start"`
	DateEnd    string `json:"date_end"`
	IdWorkshop int    `json:"id_workshop"`
	NameAdmin  string `json:"name_admin"`
	NameClient string `json:"name_client"`
}
