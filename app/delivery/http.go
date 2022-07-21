package delivery

import (
	"core-banking/app"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Config struct {
	DB     *gorm.DB
	Router *echo.Echo
}

// Handler represent the httphandler
type Handler struct {
	uCoreBanking app.Usecase
}

// NewHTTPHandler ...
func NewHTTPHandler(
	r *echo.Echo,
	uCoreBanking app.Usecase,
) *echo.Echo {

	handler := &Handler{
		uCoreBanking: uCoreBanking,
	}

	r.GET("/healthcheck", handler.healthCheck)
	r.POST("/transaction", handler.Transaction)
	r.GET("/transactionhistory", handler.TransactionHistory)

	return r

}

func (handler Handler) healthCheck(c echo.Context) error {

	result := map[string]interface{}{
		"status": "working",
	}

	return c.JSON(http.StatusOK, result)
}
