package parser

import "time"

type Operation struct {
	Key                       string
	Datetime                  time.Time
	SrcOwner                  string
	SrcTicker                 string
	SrcSum                    float32
	SrcAccruedInterestTicker  string
	SrcAccruedInterestSum     float32
	DestOwner                 string
	DestTicker                string
	DestSum                   float32
	DestAccruedInterestTicker string
	DestAccruedInterestSum    float32
	CommissionTicker          string
	CommissionSum             float32
}

type OperationParser interface {
	GetOperations() []Operation
}

type ownerTypeStruct struct {
	User       string
	UserCache  string
	Tariff     string
	Commission string
	Tax        string
}

var OwnerType = ownerTypeStruct{
	User:       "user",
	UserCache:  "userCache",
	Tariff:     "tariff",
	Commission: "commission",
	Tax:        "tax",
}
