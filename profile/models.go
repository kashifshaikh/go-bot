package profile

import (
	"bot"

	creditcard "github.com/durango/go-credit-card"
)

type Profile struct {
	ID              bot.PK      `storm:"id,increment"`
	Name            string      `validate:"required,min=3,max=256"`
	Email           string      `validate:"required,email,min=3,max=256" storm:"index"`
	BillingAddress  Address     `validate:"required" storm:"inline"`
	ShippingAddress *Address    `json:",omitempty" validate:"omitempty,required" storm:"inline"`
	CreditCard      *CreditCard `json:",omitempty" validate:"omitempty,required" storm:"inline"`
	bot.Timestamps  `storm:"inline"`
}

type CreditCard struct {
	FirstName string `validate:"required,min=3,max=256"`
	LastName  string `validate:"required,min=3,max=256"`
	Type      string
	Number    string `json:",omitempty" validate:"required,min=12,max=19"`
	CVV       string `json:"cvv,omitempty" validate:"required,min=3,max=4"`
	ExpMonth  string `validate:"required,len=2,number"`
	ExpYear   string `validate:"required,len=4,number"`
	Last4     string `storm:"index"`
}

func (cc *CreditCard) Validate() (err error) {
	card := creditcard.Card{Number: cc.Number, Cvv: cc.CVV, Month: cc.ExpMonth, Year: cc.ExpYear}
	if err = card.Validate(true); err != nil {
		return err
	}
	company, _ := card.MethodValidate()
	cc.Type = company.Short
	cc.Last4, _ = card.LastFour()
	return nil
}

type Address struct {
	FirstName  string `validate:"required,min=3,max=256"`
	LastName   string `validate:"required,min=3,max=256"`
	Address1   string `validate:"required,min=3,max=256"`
	Address2   string `validate:"max=256"`
	City       string `validate:"required,min=3,max=256"`
	Province   string `validate:"required,min=2,max=256"`
	PostalCode string `validate:"required,min=3,max=10"`
	Country    string `validate:"required,min=2,max=256"`
}
