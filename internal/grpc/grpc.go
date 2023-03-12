package grpc

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	payrollv1 "github.com/antosdaniel/go-presentation-generate-code/internal/grpc/payroll/v1"
	"github.com/antosdaniel/go-presentation-generate-code/internal/grpc/payroll/v1/payrollv1connect"
	connect_go "github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var port = getOptionalEnv("PORT", "8000")

const address = "0.0.0.0"

func StartServer() error {
	mux := http.NewServeMux()
	path, handler := payrollv1connect.NewPayrollServiceHandler(&payrollServiceServer{})
	mux.Handle(path, handler)
	return http.ListenAndServe( //nolint:gosec
		fmt.Sprintf("%s:%s", address, port),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type payrollServiceServer struct{}

func (s *payrollServiceServer) AddPayroll(ctx context.Context, request *connect_go.Request[payrollv1.AddPayrollRequest]) (*connect_go.Response[payrollv1.AddPayrollResponse], error) {
	log.Printf("add payroll on %q", request.Msg.Payday)

	return &connect_go.Response[payrollv1.AddPayrollResponse]{
		Msg: &payrollv1.AddPayrollResponse{
			PayrollId: "payroll-1",
		},
	}, nil
}

func (s *payrollServiceServer) AddPayslip(ctx context.Context, request *connect_go.Request[payrollv1.AddPayslipRequest]) (*connect_go.Response[payrollv1.AddPayslipResponse], error) {
	log.Printf("add payslip to payroll %q", request.Msg.PayrollId)

	return &connect_go.Response[payrollv1.AddPayslipResponse]{
		Msg: &payrollv1.AddPayslipResponse{
			PayslipId: "payslip-1",
		},
	}, nil
}

func (s *payrollServiceServer) GetPayroll(ctx context.Context, request *connect_go.Request[payrollv1.GetPayrollRequest]) (*connect_go.Response[payrollv1.GetPayrollResponse], error) {
	return &connect_go.Response[payrollv1.GetPayrollResponse]{
		Msg: &payrollv1.GetPayrollResponse{
			Payroll:  nil,
			Payslips: nil,
		},
	}, nil
}

func getOptionalEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
