package models

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
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

func generateTwoDecimalFloat(value float64) float64 {
	return float64(int(value*100)) / 100
}

func generateRandomDateBetween(startDate, endDate time.Time) time.Time {
	min := startDate.Unix()
	max := endDate.Unix()
	delta := max - min

	randomTime := min + rand.Int63n(delta)

	return time.Unix(randomTime, 0).UTC()
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
				Balance:       generateTwoDecimalFloat(rand.Float64() * 10000),
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
	DB.Create(&instrumentCategories)

	var instruments []Instrument
	for i := 0; i < 1000; i++ {
		var name string
		for {
			name = faker.UUIDDigit()
			if len(name) >= 3 {
				break
			}
		}
		symbol := strings.ToUpper(name[:4])
		instrument := Instrument{
			Name:                 name,
			Symbol:               symbol,
			Price:                generateTwoDecimalFloat(rand.Float64() * 200),
			InstrumentCategoryID: instrumentCategories[rand.Intn(len(instrumentCategories))].ID,
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
				Amount:       generateTwoDecimalFloat(rand.Float64() * 1000),
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
		for j := 0; j < rand.Intn(20000)+400; j++ {
			toAccount := accounts[rand.Intn(len(accounts))]
			startDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
			endDate := time.Now()
			if toAccount.ID != fromAccount.ID {
				transaction := Transaction{
					AccountID:   fromAccount.ID,
					ToAccountID: toAccount.ID,
					Amount:      generateTwoDecimalFloat(rand.Float64() * 500),
					CategoryID:  categories[rand.Intn(len(categories))].ID,
					Datetime:    generateRandomDateBetween(startDate, endDate),
				}
				transactions = append(transactions, transaction)
			}
		}
		sort.Slice(transactions, func(i, j int) bool {
			return transactions[i].Datetime.Before(transactions[j].Datetime)
		})
		DB.Create(&transactions)
		transactions = []Transaction{}
	}

}
