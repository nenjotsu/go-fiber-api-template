package entity

import (
	"time"

	"gorm.io/gorm"
)

type GetForexFactoryRequestBody struct {
	Day   uint32
	Month uint32
	Year  uint32
}

type ForexFactory struct {
	gorm.Model
	Currency       string    `gorm:"not null;column:currency" json:"currency"`
	Event          string    `gorm:"null;column:event" json:"event"`
	Impact         string    `gorm:"not null;column:impact" json:"impact"`
	Time           string    `gorm:"unique;not null;column:time" json:"time"`
	DateTimeUtc    time.Time `gorm:"not null;column:date_time_utc" json:"dateTimeUtc"`
	DateTimeManila time.Time `gorm:"not null;column:date_time_manila" json:"dateTimeManila"`
}

type GetForexFactoryRequest struct {
	Day   uint32 `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month uint32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year  uint32 `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

type CreateForexFactoryRequest struct {
	Currency string `json:"currency,omitempty"`
	Event    string `json:"event,omitempty"`
	Impact   string `json:"impact,omitempty"`
	Time     string `json:"time,omitempty"`
}

type GetForexFactoryResponse struct {
	Success      bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message      string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Forexfactory []*ForexFactory `protobuf:"bytes,3,rep,name=forexfactory,proto3" json:"forexfactory,omitempty"`
}
