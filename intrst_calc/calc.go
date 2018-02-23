package intrstcalc

import (
	"math"
	"time"
	"fmt"
)
const(
	DaysInMonth = 30
	DaysInYear = 360
)
func (l *Loan) CalculateInterest(iop float64) float64 {
	interest := (l.InterestRate * DaysInMonth * iop) / DaysInYear
	//fmt.Printf("interest: %f		--		FormattedInterest: %f\n",interest, format(interest))
	return format(interest)
	//return interest
}

func (l *Loan) Annuity(lastRemainingOP float64) (float64){

	nominator := l.InitialAmount * l.EffectiveInterestRate
	denominator := 1.00 - math.Pow((1.00 + l.EffectiveInterestRate), float64(-l.Instalments))
	//fmt.Printf("Nominator: %f	Denominator:%f\n",nominator, denominator)
	annuity := nominator / denominator
	if lastRemainingOP < annuity{
		return format(lastRemainingOP)
		//return lastRemainingOP
	}

	return format(annuity)
	//return annuity
}


// Principal calculates the principal amount for this given initial outstanding principal
func (l *Loan) Principal(annuity, iop float64) float64{
	var principal float64
	interest := l.CalculateInterest(iop)
	if interest > iop{
		principal = iop
	}else{
		// l.Annuity()
		principal = annuity - interest
	}
	//fmt.Printf("principal: %f		--		Formattedprincipal: %f\n",principal, format(principal))
	return format(principal)
	//return principal
}
// NextInstalment calculates the next instalment based on the given remainingOP
func (l *Loan) NextInstalment(lastInstalmentDate time.Time, lastRemainingOP float64) Instalment{
	var i Instalment
	i.Date = lastInstalmentDate
	i.InitialOutstandingPrincipal = lastRemainingOP
	annuity := l.Annuity(lastRemainingOP)

	i.BorrowerPaymentAmount = annuity
	// if lastRemainingOP is less than the annuity, it means this is the last instalment
	if lastRemainingOP < annuity{
		i.BorrowerPaymentAmount = lastRemainingOP
	}
	i.Interest = l.CalculateInterest(i.InitialOutstandingPrincipal)
	i.Principal = l.Principal(i.BorrowerPaymentAmount, i.InitialOutstandingPrincipal)

	i.RemainingOutstandingPrincipal = i.InitialOutstandingPrincipal - i.Principal

	fmt.Println(i)
	return i
}


// formatCents formats the given float to have 2 decimals at most and add 1 to compensate for the rest of the decimals
func format(f float64) float64{

	//return round(float64(int(f * 100))/100.00)
	return round(f)/100
}

func round (f float64) float64{
	remainder, fraction :=  math.Modf(f*100)
	if 0.5 - fraction > 0{
		return remainder
	}
	return remainder + 1

}