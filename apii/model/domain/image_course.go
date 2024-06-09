package domain

type ImageCourse struct {
	Id        int    `gorm:"primaryKey"`
	CourseId  int    `gorm:"not null"`
	ImageType string `gorm:"not null"`
	ImageAlt  string
}
