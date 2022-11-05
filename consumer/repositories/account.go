package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	// Auto Create table ktb_banks
	db.Table("ktb_banks").AutoMigrate(&BankAccount{})
	return accountRepository{db}
}

func (a accountRepository) Save(bankAccount BankAccount) error {
	// Save bank account
	return a.db.Table("ktb_banks").Save(bankAccount).Error
}

func (a accountRepository) Delete(id string) error {
	// Delete bank account by id
	return a.db.Table("ktb_banks").Where("id=?", id).Delete(&BankAccount{}).Error
}

func (a accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	// Find all bank account
	err = a.db.Table("ktb_banks").Find(&bankAccounts).Error
	return bankAccounts, err
}

func (a accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	// Find bank account by id
	err = a.db.Table("ktb_banks").Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}
