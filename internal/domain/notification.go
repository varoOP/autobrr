package domain

import (
	"context"
	"time"
)

type NotificationRepo interface {
	List(ctx context.Context) ([]Notification, error)
	Find(ctx context.Context, params NotificationQueryParams) ([]Notification, int, error)
	FindByID(ctx context.Context, id int) (*Notification, error)
	Store(ctx context.Context, notification Notification) (*Notification, error)
	Update(ctx context.Context, notification Notification) (*Notification, error)
	Delete(ctx context.Context, notificationID int) error
}

type NotificationSender interface {
	Send(event NotificationEvent, payload NotificationPayload) error
	CanSend(event NotificationEvent) bool
}

type Notification struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Type      NotificationType `json:"type"`
	Enabled   bool             `json:"enabled"`
	Events    []string         `json:"events"`
	Token     string           `json:"token"`
	APIKey    string           `json:"api_key"`
	Webhook   string           `json:"webhook"`
	Title     string           `json:"title"`
	Icon      string           `json:"icon"`
	Username  string           `json:"username"`
	Host      string           `json:"host"`
	Password  string           `json:"password"`
	Channel   string           `json:"channel"`
	Rooms     string           `json:"rooms"`
	Targets   string           `json:"targets"`
	Devices   string           `json:"devices"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type NotificationPayload struct {
	Subject        string
	Message        string
	Event          NotificationEvent
	ReleaseName    string
	Filter         string
	Indexer        string
	InfoHash       string
	Size           uint64
	Status         ReleasePushStatus
	Action         string
	ActionType     ActionType
	ActionClient   string
	Rejections     []string
	Protocol       ReleaseProtocol       // torrent
	Implementation ReleaseImplementation // irc, rss, api
	Timestamp      time.Time
}

type NotificationType string

const (
	NotificationTypeDiscord    NotificationType = "DISCORD"
	NotificationTypeIFTTT      NotificationType = "IFTTT"
	NotificationTypeJoin       NotificationType = "JOIN"
	NotificationTypeMattermost NotificationType = "MATTERMOST"
	NotificationTypeMatrix     NotificationType = "MATRIX"
	NotificationTypePushBullet NotificationType = "PUSH_BULLET"
	NotificationTypePushover   NotificationType = "PUSHOVER"
	NotificationTypeRocketChat NotificationType = "ROCKETCHAT"
	NotificationTypeSlack      NotificationType = "SLACK"
	NotificationTypeTelegram   NotificationType = "TELEGRAM"
)

type NotificationEvent string

const (
	NotificationEventPushApproved       NotificationEvent = "PUSH_APPROVED"
	NotificationEventPushRejected       NotificationEvent = "PUSH_REJECTED"
	NotificationEventPushError          NotificationEvent = "PUSH_ERROR"
	NotificationEventAppUpdateAvailable NotificationEvent = "APP_UPDATE_AVAILABLE"
	NotificationEventIRCHealth          NotificationEvent = "IRC_HEALTH"
	NotificationEventTest               NotificationEvent = "TEST"
)

type NotificationEventArr []NotificationEvent

type NotificationQueryParams struct {
	Limit   uint64
	Offset  uint64
	Sort    map[string]string
	Filters struct {
		Indexers   []string
		PushStatus string
	}
	Search string
}
