package classes

import "fmt"

type Database struct{}

func (d *Database) GetData() string {
	return "User data from database"
}

type Service struct {
	db *Database
}

func NewService() *Service {
	// ❌ Service নিজেই dependency তৈরি করছে (tight coupling)
	return &Service{db: &Database{}}
}

func (s *Service) Serve() {
	fmt.Println(s.db.GetData())
}

func tight_coupling() {
	x := NewService()
	x.Serve()

	y := Service{}.db.GetData()
	fmt.Println(y)
}
