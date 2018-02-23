# Loan Plan Generator

This software provides packages to help calculate a plan for a client Loan.

You can use the `loanplangenerator` package to create a Loan and get a Plan for your loan.

You can also use the HTTP service through the following endpoint:

`POST` **http://localhost/generate-plan**

with the given payload:
```
{
 "loanAmount": "5000",
 "nominalRate": "5.0",
 "duration": 24,
 "startDate": "2018-01-01T00:00:01Z"
 }
 ```
 
You should expect an array of instalments:
```
{[
{
"borrowerPaymentAmount": "219.36",
"date": "2018-01-01T00:00:00Z",
"initialOutstandingPrincipal": "5000.00",
"interest": "20.83",
"principal": "198.53",
"remainingOutstandingPrincipal": "4801.47",
},
{
"borrowerPaymentAmount": "219.36",
"date": "2018-02-01T00:00:00Z",
"initialOutstandingPrincipal": "4801.47",
"interest": "20.00",
"principal": "199.36",
"remainingOutstandingPrincipal": "4638",
},
...
{
"borrowerPaymentAmount": "219.30",
"date": "2020-01-01T00:00:00Z",
"initialOutstandingPrincipal": "218.39",
"interest": "0.91",
"principal": "218.39",
"remainingOutstandingPrincipal": "0",
}
]}
```