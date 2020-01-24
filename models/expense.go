package expense;

import (
	"time"
	"github.com/thomasraydeniscool/fundse-web-api/utils"
);

type Expense struct {
	name string
	amount int
}


func create(expense Expense) {
	collection, err := utils.GetCollection("expenses");

	ctx, cancel := utils.GetContext(5*time.Second);

	collection.InsertOne(ctx, expense);
}