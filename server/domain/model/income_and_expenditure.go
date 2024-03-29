package model

type IncomeAndExpenditureListByQuery struct {
	UserID             string
	OccurrenceDateFrom string
	OccurrenceDateTo   string
}

type IncomeAndExpenditure struct {
	ID             int64
	UserID         string
	Name           string
	Amount         int32
	OccurrenceDate string
	UserTagID      int64
	Classification string
}
