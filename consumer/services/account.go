package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo}
}

func (obj accountEventHandler) Handle(topic string, eventBytes []byte) {
	//Recieve topic and eventBytes for doing each topic process
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():

		event := &events.OpenAccountEvent{}
		// Unmarshal eventBytes to get OpenAccountEvent data then set to event
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}
		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("topic: ", topic)
		log.Println("event: ", event)
	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}

		// Unmarshal eventBytes to get DepositFundEvent data then set to event
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// Call repo find by id for getting account
		bankAccount, err := obj.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		// Sum amount with balance
		bankAccount.Balance += event.Amount

		// Save new data account
		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("topic: ", topic)
		log.Println("event: ", event)
	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := &events.WithdrawFundEvent{}
		// Unmarshal eventBytes to get WithdrawFundEvent data then set to event
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// Call repo find by id for getting account
		bankAccount, err := obj.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		// Minus amount with balance
		bankAccount.Balance -= event.Amount

		// Save new data account
		err = obj.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("topic: ", topic)
		log.Println("event: ", event)
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := &events.CloseAccountEvent{}
		// Unmarshal eventBytes to get CloseAccountEvent data then set to event
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// Call repo delete by
		err = obj.accountRepo.Delete(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("topic: ", topic)
		log.Println("event: ", event)
	default:
		log.Println("no event")
	}
}
