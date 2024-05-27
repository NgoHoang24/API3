package request

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required" json:"name"`
}
