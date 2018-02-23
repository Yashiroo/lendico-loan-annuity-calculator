package intrstcalc

import (
	"testing"
	"time"
)

func TestAnnuityAmount(t *testing.T){
	var months int = 24
	var rate float64 = 0.05
	var amount float64= 5000.00
	loan := NewLoan(amount, rate, months)

	expectedResult := 219.36
	actualResult := loan.Annuity(loan.InitialAmount)
	if expectedResult != actualResult{
		t.Errorf("Got %f, expected %f",actualResult, expectedResult)
	}
}

func TestNewLoan(t *testing.T) {
	l := NewLoan(5000.00, 0.05, 24)
	expectedAmount := 5000.00
	expectedRate := 0.05
	expectedInstalments := 24
	expectedEffectiveInterestRate := expectedRate/12.00

	if l.InitialAmount != expectedAmount{
		t.Errorf("Incorrect amount: expected %f, got %f",expectedAmount, l.InitialAmount)
	}
	if l.InterestRate != expectedRate{
		t.Errorf("Incorrect rate: expected %f, got %f",expectedRate, l.InterestRate)
	}
	if l.Instalments != expectedInstalments{
		t.Errorf("Incorrect instalments: expected %f, got %f", expectedInstalments, l.Instalments)
	}
	if l.EffectiveInterestRate != expectedEffectiveInterestRate{
		t.Errorf("Incorrect effective interest rate: expected %f, got %f", expectedEffectiveInterestRate, l.EffectiveInterestRate)
	}
}

func TestLoan_Principal(t *testing.T) {
	loan := NewLoan(5000.00, 0.05, 24)

	expectedPrincipal := 198.53
	actualPrincipal := loan.Principal(219.36, 5000.00)
	if actualPrincipal != expectedPrincipal{
		t.Errorf("Incorrect calculated principal: expected %f, got %f", expectedPrincipal, actualPrincipal)
	}

	expectedPrincipal = 199.36
	actualPrincipal = loan.Principal(219.36, 4801.47)
	if actualPrincipal != expectedPrincipal{
		t.Errorf("Incorrect calculated principal: expected %f, got %f", expectedPrincipal, actualPrincipal)
	}

}


func TestLoan_NextInstalment(t *testing.T) {
	loan := NewLoan(5000.00, 0.05, 24)

	nextDate, err := time.Parse("01.01.2018", "01.02.2018")
	if err != nil{
		t.Fail()
	}
	loan.DateOfDisbursement, err = time.Parse("01.01.2018", "01.01.2018")
	if err != nil{
		t.Fail()
	}

	nextBorrowerPayment := 219.36
	nextIOP := 5000.00
	nextROP := 4801.47
	nextPrincipal := 198.53
	nextInterest := 20.83

	actualInstalment := loan.NextInstalment(loan.DateOfDisbursement, 5000.00)
	if actualInstalment.Date != nextDate{
		t.Errorf("Date: expected %s, got %s", actualInstalment.Date, nextDate)
	}
	if actualInstalment.BorrowerPaymentAmount != nextBorrowerPayment{
		t.Errorf("BorrowerPayment: expected %f, got %f", nextBorrowerPayment, actualInstalment.BorrowerPaymentAmount)
	}
	if actualInstalment.InitialOutstandingPrincipal != nextIOP{
		t.Errorf("IOP: expected %f, got %f", nextIOP, actualInstalment.InitialOutstandingPrincipal)
	}
	if actualInstalment.RemainingOutstandingPrincipal != nextROP{
		t.Errorf("ROP: expected %f, got %f", nextROP, actualInstalment.RemainingOutstandingPrincipal)
	}
	if actualInstalment.Principal != nextPrincipal{
		t.Errorf("Principal: expected %f, got %f", nextPrincipal, actualInstalment.Principal)
	}
	if actualInstalment.Interest != nextInterest{
		t.Errorf("interest: expected %f, got %f", nextInterest, actualInstalment.Interest)
	}
}