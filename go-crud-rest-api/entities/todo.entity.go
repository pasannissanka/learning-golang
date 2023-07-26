package entities

type Todo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}
