package controllers

import (
	"log"
	"net/http"
	"producer/commands"
	"producer/services"

	"github.com/labstack/echo/v4"
)

type AccountController interface {
	OpenAccount(c echo.Context) error
	DepositFund(c echo.Context) error
	WithdrawFund(c echo.Context) error
	CloseAccount(c echo.Context) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService}
}

func (obj accountController) OpenAccount(c echo.Context) error {
	command := commands.OpenAccountCommand{}

	err := c.Bind(&command)

	if err != nil {
		return err
	}

	id, err := obj.accountService.OpenAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "open account success",
		"id":      id,
	})
}

func (obj accountController) DepositFund(c echo.Context) error {
	command := commands.DepositFundCommand{}
	err := c.Bind(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "deposit fund success",
	})
}

func (obj accountController) WithdrawFund(c echo.Context) error {
	command := commands.WithdrawFundCommand{}
	err := c.Bind(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "withdraw fund success",
	})

}

func (obj accountController) CloseAccount(c echo.Context) error {
	command := commands.CloseAccountCommand{}
	err := c.Bind(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "close account success",
	})
}
