package web

type CategoriesUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=100" json:"name"`
	//Name string `json:"name" validate:"required"`
}
