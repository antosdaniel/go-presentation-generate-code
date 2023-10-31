package grpc

import (
	"context"

	"github.com/antosdaniel/go-presentation-generate-code/gen/grpc/payroll/payrollv1"
	"github.com/antosdaniel/go-presentation-generate-code/gen/grpc/payroll/payrollv1/payrollv1connect"
	connect_go "github.com/bufbuild/connect-go"
)

type payrollServiceServerWithAuth struct {
	base payrollv1connect.PayrollServiceHandler
}

func newWithAuth(base payrollv1connect.PayrollServiceHandler) payrollv1connect.PayrollServiceHandler {
	return payrollServiceServerWithAuth{base: base}
}

func (p payrollServiceServerWithAuth) AddPayroll(ctx context.Context, c *connect_go.Request[payrollv1.AddPayrollRequest]) (*connect_go.Response[payrollv1.AddPayrollResponse], error) {
	// TODO: check permission
	return p.base.AddPayroll(ctx, c)
}

func (p payrollServiceServerWithAuth) AddPayslip(ctx context.Context, c *connect_go.Request[payrollv1.AddPayslipRequest]) (*connect_go.Response[payrollv1.AddPayslipResponse], error) {
	// TODO: check permission
	return p.base.AddPayslip(ctx, c)
}

func (p payrollServiceServerWithAuth) GetPayroll(ctx context.Context, c *connect_go.Request[payrollv1.GetPayrollRequest]) (*connect_go.Response[payrollv1.GetPayrollResponse], error) {
	// TODO: check permission
	return p.base.GetPayroll(ctx, c)
}
