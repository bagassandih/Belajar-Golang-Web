package entity

type Product struct{
	ID int
	Name string
	Price int
	Stock int
}

func (x Product) StockStatus() string{
	var status string
	if x.Stock < 1 {
		status = "Stock Habis"
	} else if x.Stock < 2 {
		status = "Stock Tersisa 1"
	} else if x.Stock < 6 {
		status = "Stock Terbatas"
	} else {
		status = "Stock Aman"
	}

	return status
}