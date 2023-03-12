package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/antosdaniel/go-presentation-generate-code/internal/db/models"
	payrollv1 "github.com/antosdaniel/go-presentation-generate-code/internal/grpc/payroll/v1"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type payrollServiceServer struct {
	db *sql.DB
}

func (s *payrollServiceServer) AddPayroll(ctx context.Context, request *connect_go.Request[payrollv1.AddPayrollRequest]) (*connect_go.Response[payrollv1.AddPayrollResponse], error) {
	log.Printf("add payslip to payroll %q", request.Msg.Payday)

	payday, err := time.Parse(time.DateOnly, request.Msg.Payday)
	if err != nil {
		return nil, fmt.Errorf("invalid payday: %w", err)
	}

	if request.Msg.TenantId == "" {
		return nil, fmt.Errorf("tenant ID is missing")
	}

	id := request.Msg.PayrollId
	if id == "" {
		id = uuid.NewString()
	}

	payroll := models.Payroll{
		ID:       id,
		TenantID: request.Msg.TenantId,
		Payday:   payday,
	}
	err = payroll.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("could not insert payroll: %w", err)
	}

	return &connect_go.Response[payrollv1.AddPayrollResponse]{
		Msg: &payrollv1.AddPayrollResponse{
			PayrollId: payroll.ID,
		},
	}, nil
}

func (s *payrollServiceServer) AddPayslip(ctx context.Context, request *connect_go.Request[payrollv1.AddPayslipRequest]) (*connect_go.Response[payrollv1.AddPayslipResponse], error) {
	log.Printf("add payslip to payroll %q", request.Msg.PayrollId)

	payrollID := request.Msg.PayrollId
	payroll, err := models.FindPayroll(ctx, s.db, payrollID)
	if err != nil {
		return nil, fmt.Errorf("could not find payroll %q: %w", payrollID, err)
	}

	payslip := models.Payslip{
		ID:        uuid.NewString(),
		TenantID:  payroll.TenantID,
		PayrollID: payroll.ID,
		GrossPay:  int(request.Msg.GrossPay),
		Tax:       int(request.Msg.Tax),
		NetPay:    int(request.Msg.GrossPay - request.Msg.Tax),
	}
	err = payslip.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("could not insert payslip: %w", err)
	}

	return &connect_go.Response[payrollv1.AddPayslipResponse]{
		Msg: &payrollv1.AddPayslipResponse{
			PayslipId: payslip.ID,
		},
	}, nil
}

func (s *payrollServiceServer) GetPayroll(ctx context.Context, request *connect_go.Request[payrollv1.GetPayrollRequest]) (*connect_go.Response[payrollv1.GetPayrollResponse], error) {
	payrollID := request.Msg.PayrollId
	payroll, err := models.FindPayroll(ctx, s.db, payrollID)
	if err != nil {
		return nil, fmt.Errorf("could not find payroll %q: %w", payrollID, err)
	}

	payslips, err := payroll.Payslips().All(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("could not find payslips for payroll %q: %w", payrollID, err)
	}

	responsePayslips := make([]*payrollv1.Payslip, len(payslips))
	for i, payslip := range payslips {
		responsePayslips[i] = &payrollv1.Payslip{
			Id:        payslip.ID,
			TenantId:  payslip.TenantID,
			PayrollId: payslip.PayrollID,
			GrossPay:  int32(payslip.GrossPay),
			Tax:       int32(payslip.Tax),
			NetPay:    int32(payslip.NetPay),
		}
	}

	return &connect_go.Response[payrollv1.GetPayrollResponse]{
		Msg: &payrollv1.GetPayrollResponse{
			Payroll: &payrollv1.Payroll{
				Id:       payroll.ID,
				TenantId: payroll.TenantID,
				Payday:   payroll.Payday.Format(time.DateOnly),
			},
			Payslips: responsePayslips,
		},
	}, nil
}
