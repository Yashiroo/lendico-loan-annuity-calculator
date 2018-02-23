package loanplangenerator

import (
	"testing"
	"github.com/yashiroo/lendico-loan-annuity-calculator/intrst_calc"
	"time"
	"fmt"
	"encoding/json"
)

func TestGeneratePlan(t *testing.T) {
	// amount should always be in cents
	loan := intrstcalc.NewLoan(500000.00, 0.05, 24)
	loan.DateOfDisbursement = time.Now()

	plan := GeneratePlan(loan)

	for _, v := range plan{
		js, _ := json.Marshal(v)
		fmt.Printf("%s\n\n",js)
	}
}
