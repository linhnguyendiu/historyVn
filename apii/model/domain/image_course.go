package domain

type ImageCourse struct {
	ImageId   int    `gorm:"primaryKey"`
	CourseId  int    `gorm:"not null"`
	ImageType string `gorm:"not null"`
	ImageAlt  string
}
