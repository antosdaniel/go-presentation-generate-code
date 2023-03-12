package payroll

import "time"

type Payroll struct {
	ID       PayrollID
	TenantID TenantID
	Payday   time.Time
}

type PayrollID string

type Payslips []Payslip

type Payslip struct {
	ID        PayslipID
	TenatnID  TenantID
	PayrollID PayrollID

	GrossPay int
	Tax      int
	NetPay   int
}

type (
	PayslipID string
	TenantID  string
)
