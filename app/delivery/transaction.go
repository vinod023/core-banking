package delivery

import (
	"core-banking/models"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func (handler Handler) Transaction(c echo.Context) error {
	var transaction models.Transaction
	if err := json.NewDecoder(c.Request().Body).Decode(&transaction); err != nil {
		return err
	}

	userTransaction, err := handler.uCoreBanking.UserTransaction(transaction)
	if err != nil {
		errorData := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorData)
	}

	resp := models.Response{
		Status:        "success",
		TransactionID: userTransaction.ID,
		Balance:       userTransaction.Balance,
	}

	return c.JSON(http.StatusOK, resp)
}
