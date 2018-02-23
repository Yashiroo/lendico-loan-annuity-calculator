package loanplangenerator

import (
	"github.com/yashiroo/lendico-loan-annuity-calculator/intrst_calc"
	"time"
	"errors"
	"fmt"
)

func GeneratePlan(loan *intrstcalc.Loan) intrstcalc.Plan{
	//check we have all the required data (4 inputs)
	if err := validateLoan(loan); err != nil{
		// we should return an error here
		fmt.Printf("Error validating the loan: %s", err)
		return intrstcalc.Plan{}
	}
	fmt.Printf("Generating plan for following loan:\n%v",*loan)
	var plan intrstcalc.Plan
	var lastRemainingOP float64 = loan.InitialAmount
	var lastInstalmentDate = loan.DateOfDisbursement
	for i:=0; i<loan.Instalments; i++{
		var inst intrstcalc.Instalment
		inst = loan.NextInstalment(lastInstalmentDate, lastRemainingOP)
		plan = append(plan, inst)

		lastInstalmentDate = lastInstalmentDate.Add(720 * time.Hour)
		lastRemainingOP = inst.RemainingOutstandingPrincipal
	}

	return plan
}

// validateLoan makes sure the minimum parameters are available in the passed struct
func validateLoan(l *intrstcalc.Loan) error{
	if  l == nil{
		return errors.New("load is nil!")
	}
	if l.InitialAmount == 0{
		return errors.New("initial amount is 0 !")
	}
	if l.InterestRate == 0 {
		return errors.New("interest rate is 0 !")
	}
	if l.Instalments == 0{
		return errors.New("number of instalments is 0 !")
	}

	return nil
}