package factories

import (
	"github.com/gabrielmrts/mybooks-golang-api/internal/services"
)

func GetMailService() *services.MailService {
	return services.NewMailService()
}
