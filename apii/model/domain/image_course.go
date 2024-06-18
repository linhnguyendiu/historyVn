package domain

type ImageCourse struct {
	Id          int `gorm:"primaryKey"`
	CourseId    int `gorm:"not null"`
	ImageType   int `gorm:"not null"`
	Description string
	ImageAlt    string
}
