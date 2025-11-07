package classes

// import "fmt"

// // ✅ Singleton Logger
// type Logger struct{}

// func (l *Logger) Log(message string) {
// 	fmt.Println("Message:", message)
// }

// var loggerInstance *Logger

// func GetLoggerInstance() *Logger {
// 	if loggerInstance == nil {
// 		fmt.Println("Created Logger Instance.")
// 		loggerInstance = &Logger{}
// 	}

// 	return loggerInstance
// }

// // ✅ Service (Dependency Injection)
// type Service struct{ logger *Logger }

// // Constructor: Logger dependency inject হচ্ছে বাইরে থেকে
// func NewService(logger *Logger) *Service {
// 	return &Service{logger: logger}
// }

// func (s *Service) User(name string) {
// 	s.logger.Log("Creating user: " + name)
// 	s.logger.Log("Created user: " + name)
// }

// func main() {
// 	// Singleton instance একবারই তৈরি হবে
// 	logger := GetLoggerInstance()

// 	// Dependency Injection: একই logger inject করলাম service এ
// 	service1 := NewService(logger)
// 	service2 := NewService(logger)

// 	service1.User("Alice")
// 	service2.User("Bob")
// }
