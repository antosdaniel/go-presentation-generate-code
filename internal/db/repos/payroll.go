package repos

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/antosdaniel/go-presentation-generate-code/internal/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PayrollRepo struct {
	db *sql.DB
}

func NewPayrollRepo(db *sql.DB) *PayrollRepo {
	return &PayrollRepo{db}
}

func (r *PayrollRepo) Create(ctx context.Context, payrollID, tenantID string, payday time.Time) error {
	payroll := models.Payroll{
		ID:       payrollID,
		TenantID: tenantID,
		Payday:   payday,
	}
	err := payroll.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("could not insert payroll: %w", err)
	}
	return nil
}

func (r *PayrollRepo) AddPayslip(ctx context.Context, payslipID, payrollID string, grossPay, tax, netPay int) error {
	payroll, err := models.FindPayroll(ctx, r.db, payrollID)
	if err != nil {
		return fmt.Errorf("could not find payroll %q: %w", payrollID, err)
	}

	payslip := models.Payslip{
		ID:        payslipID,
		TenantID:  payroll.TenantID,
		PayrollID: payroll.ID,
		GrossPay:  grossPay,
		Tax:       tax,
		NetPay:    netPay,
	}
	err = payslip.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("could not insert payslip: %w", err)
	}
	return nil
}

func (r *PayrollRepo) Find(ctx context.Context, payrollID string) (*models.Payroll, models.PayslipSlice, error) {
	payroll, err := models.FindPayroll(ctx, r.db, payrollID)
	if err != nil {
		return nil, nil, fmt.Errorf("could not find payroll %q: %w", payrollID, err)
	}

	payslips, err := payroll.Payslips().All(ctx, r.db)
	if err != nil {
		return nil, nil, fmt.Errorf("could not find payslips for payroll %q: %w", payrollID, err)
	}

	return payroll, payslips, nil
}
