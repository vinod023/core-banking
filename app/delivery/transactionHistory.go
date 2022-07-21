package delivery

import (
	"core-banking/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (handler Handler) TransactionHistory(c echo.Context) error {

	var userID int
	if c.FormValue("userId") != "" {
		userID, _ = strconv.Atoi(c.FormValue("userId"))
	}

	var amount float64
	if c.FormValue("amount") != "" {
		amount, _ = strconv.ParseFloat(c.FormValue("amount"), 64)
	}

	var balance float64
	if c.FormValue("balance") != "" {
		balance, _ = strconv.ParseFloat(c.FormValue("balance"), 64)
	}

	var startDate time.Time
	if c.FormValue("startDate") != "" {
		startDate, _ = time.Parse("2006-01-02", c.FormValue("startDate"))
	}

	var endDate time.Time
	if c.FormValue("endDate") != "" {
		endDate, _ = time.Parse("2006-01-02", c.FormValue("endDate"))
	}

	transactionType := c.FormValue("transactionType")

	transHisSearchFilter := models.TransHistorySearchFilter{
		UserID:          uint(userID),
		Amount:          amount,
		Balance:         balance,
		TransactionType: transactionType,
		StartDate:       &startDate,
		EndDate:         &endDate,
	}

	transactionHistory, err := handler.uCoreBanking.GetTransactionHistory(transHisSearchFilter)
	if err != nil {
		errorData := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorData)
	}

	return c.JSON(http.StatusOK, transactionHistory)
}
