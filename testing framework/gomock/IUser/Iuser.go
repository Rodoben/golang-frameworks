package iuser

//go:generate mockgen -destination=../mocks/mockrunner.go -package=mocks gomock/learn/iuser IuserInterface
type IuserInterface interface {
	AddCell(int, string, bool, int) error
	UpdateCell(int, string) error
	DeleteCell(int) error
	GetCell(int) error
	GetAllCell() error
}
