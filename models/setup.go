package models

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("banking_simulation.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&User{}, &Account{}, &Investment{}, &Transaction{}, &Instrument{}, &TransactionCategory{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	if isDatabaseEmpty() {
		createInitialData()
	}
}

func isDatabaseEmpty() bool {
	var userCount int64
	DB.Model(&User{}).Count(&userCount)
	return userCount == 0
}

func generateFakeIBAN(bankCode string) string {
	countryCode := "NL"
	checkDigits := strconv.Itoa(rand.Intn(99) + 10)
	accountNumber := strconv.Itoa(rand.Intn(9999999999) + 1000000000)
	return strings.ToUpper(countryCode + checkDigits + bankCode + accountNumber)
}

func createInitialData() {
	var users []User
	for i := 0; i < 100; i++ {
		user := User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Password:  "password",
			Token:     faker.UUIDDigit(),
		}
		users = append(users, user)
	}
	DB.Create(&users)

	var accounts []Account
	var bankCode string
	for {
		bankCode = faker.Word()
		if len(bankCode) >= 4 {
			break
		}
	}
	bankCode = strings.ToUpper(bankCode[:4])
	for _, user := range users {
		for j := 0; j < rand.Intn(3)+1; j++ {
			account := Account{
				UserID:        user.ID,
				Balance:       rand.Float64() * 10000,
				AccountNumber: generateFakeIBAN(bankCode),
			}
			accounts = append(accounts, account)
		}
	}
	DB.Create(&accounts)

	var instrumentCategories []InstrumentCategory
	for i := 0; i < 10; i++ {
		category := InstrumentCategory{
			Name: faker.Word(),
		}
		instrumentCategories = append(instrumentCategories, category)
	}

	var instruments []Instrument
	for i := 0; i < 100; i++ {
		var name string
		for {
			name = faker.Word()
			if len(name) >= 3 {
				break
			}
		}
		symbol := strings.ToUpper(name[:3])
		instrument := Instrument{
			Name:               name,
			Symbol:             symbol,
			Price:              rand.Float64() * 200,
			InstrumentCategory: instrumentCategories[rand.Intn(len(instrumentCategories))],
		}
		instruments = append(instruments, instrument)
	}
	DB.Create(&instruments)

	var investments []Investment
	for _, user := range users {
		for j := 0; j < rand.Intn(20); j++ {
			investment := Investment{
				UserID:       user.ID,
				InstrumentID: instruments[rand.Intn(len(instruments))].ID,
				Amount:       rand.Float64() * 1000,
			}
			investments = append(investments, investment)
		}
	}
	DB.Create(&investments)

	categories := []TransactionCategory{
		{Name: "Groceries"},
		{Name: "Utilities"},
		{Name: "Entertainment"},
		{Name: "Rent"},
		{Name: "Salary"},
		{Name: "Investment"},
		{Name: "Shopping"},
		{Name: "Miscellaneous"},
	}
	DB.Create(&categories)

	var transactions []Transaction
	for _, fromAccount := range accounts {
		for j := 0; j < rand.Intn(10); j++ {
			toAccount := accounts[rand.Intn(len(accounts))]
			if toAccount.ID != fromAccount.ID {
				transaction := Transaction{
					AccountID:   fromAccount.ID,
					ToAccountID: toAccount.ID,
					Amount:      rand.Float64() * 500,
					CategoryID:  categories[rand.Intn(len(categories))].ID,
				}
				transactions = append(transactions, transaction)
			}
		}
	}
	DB.Create(&transactions)
}
