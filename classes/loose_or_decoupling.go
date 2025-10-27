package classes

// import "fmt"

// // Step 1: Database interface
// type Database interface {
// 	GetData() string
// }

// // Step 2: বাস্তব database implement করা
// type PostgresSqlDatabase struct{}

// func (m *PostgresSqlDatabase) GetData() string {
// 	return "Data from Postgres Sql Database"
// }

// // Step 3: Mock database implement (testing purpose)
// type MockDatabase struct{}

// func (m *MockDatabase) GetData() string {
// 	return "Mock data for testing"
// }

// // Step 4: Service struct, dependency inject করা হবে বাইরে থেকে
// type Service struct{ db Database }

// // Constructor — dependency parameter হিসেবে নিচ্ছে
// func NewService(db Database) *Service {
// 	return &Service{db: db}
// }

// func (s *Service) Serve() {
// 	fmt.Println(s.db.GetData())
// }

// func main() {
// 	// এখানে dependency inject করছি
// 	service1 := NewService(&PostgresSqlDatabase{})
// 	service1.Serve()

// 	// অন্য জায়গায় অন্য dependency inject করা যায় (e.g. mock)
// 	service2 := NewService(&MockDatabase{})
// 	service2.Serve()
// }
