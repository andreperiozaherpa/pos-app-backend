package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// NotificationService menangani pengiriman dan manajemen notifikasi ke user (email/SMS/WhatsApp/dll).
type NotificationService interface {
	SendNotification(ctx context.Context, notif *models.Notification) error
	ScheduleNotification(ctx context.Context, notif *models.Notification, sendAt time.Time) error
	CancelScheduledNotification(ctx context.Context, notificationID uuid.UUID) error
	GetNotificationStatus(ctx context.Context, notificationID uuid.UUID) (string, error)
	ListNotificationsByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Notification, error)
	ExportNotifications(ctx context.Context, userID uuid.UUID, from, to time.Time) ([]byte, error)
}
