package grpc

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/antosdaniel/go-presentation-generate-code/gen/api/grpc/payroll/v1"
	"github.com/antosdaniel/go-presentation-generate-code/internal/db/models"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

//go:generate mockgen -source payrollService.go -destination mocks/mocks.go -package mocks

//go:generate gowrap gen -g -p ../../gen/api/grpc/payroll/v1/payrollv1connect -i PayrollServiceHandler -t ./../../templates/log -o payrollServiceWithLogs.go

type payrollServiceServer struct {
	payrollRepo payrollRepo
}

type payrollRepo interface {
	Create(ctx context.Context, payrollID, tenantID string, payday time.Time) error
	AddPayslip(ctx context.Context, payslipID, payrollID string, grossPay, tax, netPay int) error
	Find(ctx context.Context, payrollID string) (*models.Payroll, models.PayslipSlice, error)
}

func (s *payrollServiceServer) AddPayroll(ctx context.Context, request *connect_go.Request[v1.AddPayrollRequest]) (*connect_go.Response[v1.AddPayrollResponse], error) {
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

	err = s.payrollRepo.Create(ctx, id, request.Msg.TenantId, payday)
	if err != nil {
		return nil, err
	}

	return &connect_go.Response[v1.AddPayrollResponse]{
		Msg: &v1.AddPayrollResponse{
			PayrollId: id,
		},
	}, nil
}

func (s *payrollServiceServer) AddPayslip(ctx context.Context, request *connect_go.Request[v1.AddPayslipRequest]) (*connect_go.Response[v1.AddPayslipResponse], error) {
	payslipID := uuid.NewString()
	err := s.payrollRepo.AddPayslip(
		ctx,
		payslipID,
		request.Msg.PayrollId,
		int(request.Msg.GrossPay),
		int(request.Msg.Tax),
		int(request.Msg.GrossPay-request.Msg.Tax),
	)
	if err != nil {
		return nil, err
	}

	return &connect_go.Response[v1.AddPayslipResponse]{
		Msg: &v1.AddPayslipResponse{
			PayslipId: payslipID,
		},
	}, nil
}

func (s *payrollServiceServer) GetPayroll(ctx context.Context, request *connect_go.Request[v1.GetPayrollRequest]) (*connect_go.Response[v1.GetPayrollResponse], error) {
	payroll, payslips, err := s.payrollRepo.Find(ctx, request.Msg.PayrollId)
	if err != nil {
		return nil, err
	}

	responsePayslips := make([]*v1.Payslip, len(payslips))
	for i, payslip := range payslips {
		responsePayslips[i] = &v1.Payslip{
			Id:        payslip.ID,
			TenantId:  payslip.TenantID,
			PayrollId: payslip.PayrollID,
			GrossPay:  int32(payslip.GrossPay),
			Tax:       int32(payslip.Tax),
			NetPay:    int32(payslip.NetPay),
		}
	}

	return &connect_go.Response[v1.GetPayrollResponse]{
		Msg: &v1.GetPayrollResponse{
			Payroll: &v1.Payroll{
				Id:       payroll.ID,
				TenantId: payroll.TenantID,
				Payday:   payroll.Payday.Format(time.DateOnly),
			},
			Payslips: responsePayslips,
		},
	}, nil
}
