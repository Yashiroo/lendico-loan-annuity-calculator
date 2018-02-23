package intrstcalc

import "time"

type Loan struct {
	InitialAmount      float64        `json:"loanAmount,string"`
	InterestRate       float64        `json:"nominalRate,string"`
	EffectiveInterestRate	float64        `json:"effectiveInterestRate"`
	// Instalments are expressed in # of months
	Instalments        int        `json:"duration"`
	DateOfDisbursement time.Time        `json:"startDate"`
}

// NewLoan returns a new Loan struct with the given parameters
func NewLoan(amount, rate float64, duration int) *Loan{
	return &Loan{
		InitialAmount:amount,
		InterestRate: rate,
		Instalments: duration,
		EffectiveInterestRate:rate/12.00,
	}
}

func (l *Loan) SetEffectiveRate(){
	l.EffectiveInterestRate = format(l.InterestRate/12.0)
}
type Plan []Instalment

type Instalment struct {
	BorrowerPaymentAmount	float64        `json:"borrowerPaymentAmount"`
	Date	time.Time        `json:"date"`
	Principal	float64        `json:"principal"`
	Interest	float64        `json:"interest"`
	InitialOutstandingPrincipal	float64        `json:"initialOutstandingPrincipal"`
	RemainingOutstandingPrincipal	float64        `json:"remainingOutstandingPrincipal"`
}