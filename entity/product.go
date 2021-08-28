package entity

type Product struct {
	ID int
	Name string
	Price int
	Stock int
}

// membuat method untuk dipanggil di template
func (product Product) StockStatus() string {
	status := ""
	if product.Stock < 3 {
		status = "Stok hampir habis"	
	} else if product.Stock < 10 {
		status = "Stok terbatas"
	}

	return status
}