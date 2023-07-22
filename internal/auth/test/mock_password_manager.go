package auth
import(
	"fmt"
	"github.com/ssonumkar/weather-report-api/internal/log"
	)	

type MockSuccessPasswordManager struct{
	
}

func (m *MockSuccessPasswordManager) ComparePasswords(hashedPassword string, plainPassword string, logger log.CustomLogger) error{
	return nil
}
func (m *MockSuccessPasswordManager) EncryptPassword(plainPassword string, logger log.CustomLogger) (string, error){
	return "dummy_pass", nil
}

type MockFailPasswordManager struct{
	
}

func (m *MockFailPasswordManager) ComparePasswords(hashedPassword string, plainPassword string, logger log.CustomLogger) error{
	return fmt.Errorf("passwords did not match")
}
func (m *MockFailPasswordManager) EncryptPassword(plainPassword string, logger log.CustomLogger) (string, error){
	return "", fmt.Errorf("could not create password")

}