package models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model

	Message string
}

type AnnouncementRequest struct {
	Message string `json:"message"`
}

type AnnouncementReturn struct {
	Message string `json:"message"`
}

func (value AnnouncementRequest) ToAnnouncement() *Announcement {
	var announcement = new(Announcement)
	announcement.Message = value.Message
	return announcement
}

func (value Announcement) ToAnnouncementReturn() *AnnouncementReturn {
	var announcementReturn = new(AnnouncementReturn)
	announcementReturn.Message = value.Message
	return announcementReturn
}
