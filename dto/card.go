package dto

// PCard is a data transfer object for pcard data.
type PCard struct {
	Brand       string
	Category    string
	Model       string
	Name        string
	SKU         string
	Price       string
	Currency    string
	Source      string
	Img         string
	Properties  string
	Description string
	File        string
}

// NewPCard creates a new PCard.
func NewPCard(brand, category, model, name, sku, price, currency, source, img, properties, description, file string) *PCard {
	return &PCard{
		Brand:       brand,
		Category:    category,
		Model:       model,
		Name:        name,
		SKU:         sku,
		Price:       price,
		Currency:    currency,
		Source:      source,
		Properties:  properties,
		Description: description,
		Img:         img,
		File:        file,
	}
}

// String returns a string representation of the PCard.
func (p *PCard) String() []string {

	str := []string{p.Brand, p.Category, p.Model, p.Name, p.SKU, p.Price, p.Currency, p.Source, p.Img, p.Properties, p.Description, p.File}

	// brand, category, model, name, sku, price, currency, source
	return str
}