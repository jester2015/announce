package repository

import (
	"absent.com/absentapi/database"
	"absent.com/absentapi/models"
)

func GetAnnouncements(limit int) ([]models.Announcement, error) {
	var announcements []models.Announcement
	if err := database.SqlLiteClient.Limit(limit).Find(&announcements).Error; err != nil {
		return announcements, err
	}
	return announcements, nil
}

func GetAnnouncement(id uint) (*models.Announcement, error) {
	var announcement *models.Announcement
	if err := database.SqlLiteClient.First(&announcement, id).Error; err != nil {
		return announcement, err
	}
	return announcement, nil
}

func CreateAnnouncement(message string) (models.Announcement, error) {
	announcement := models.Announcement{Message: message}
	result := database.SqlLiteClient.Create(&announcement).Error

	if result != nil {
		return announcement, result
	}
	return announcement, nil
}

func UpdateAnnouncement(announcement models.Announcement) (*models.Announcement, error) {
	announce := &announcement
	result := database.SqlLiteClient.Save(&announce).Error
	if result != nil {
		return announce, result
	}
	return announce, nil
}

func DeleteAnnouncement(id uint) error {
	announce := new(models.Announcement)
	result := database.SqlLiteClient.Delete(&announce, id).Error
	if result != nil {
		return result
	}

	return nil
}
