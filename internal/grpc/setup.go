package grpc

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/antosdaniel/go-presentation-generate-code/internal/db/repos"
	"github.com/antosdaniel/go-presentation-generate-code/internal/grpc/payroll/payrollv1/payrollv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	address = "0.0.0.0"
)

var port = getOptionalEnv("PORT", "8000")

func New(db *sql.DB) *http.Server {
	mux := http.NewServeMux()
	payrollService := NewPayrollServiceHandlerWithLog(&payrollServiceServer{repos.NewPayrollRepo(db)})
	path, handler := payrollv1connect.NewPayrollServiceHandler(payrollService)
	mux.Handle(path, handler)

	addr := fmt.Sprintf("%s:%s", address, port)
	return &http.Server{ //nolint:gosec
		Addr:    addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}

func getOptionalEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
