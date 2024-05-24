package structs

type UpdateClientReadsBody struct {
	Client int `json:"client" binding:"required"`
	Reads  int `json:"reads" binding:"required"`
}
