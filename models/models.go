package models

type DBRecord interface {
	IsValid() bool
}
