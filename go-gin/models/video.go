package models

type person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

//Video struct
type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"description" binding:"min=2,max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      person `json:"author" binding:"required"`
}
