package models

type ApiParam struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required string `json:"required"`
	Remark   string `json:"remark"`
}
