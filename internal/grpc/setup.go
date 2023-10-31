package grpc

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/antosdaniel/go-presentation-generate-code/gen/grpc/payroll/payrollv1/payrollv1connect"
	"github.com/antosdaniel/go-presentation-generate-code/internal/db/repos"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	address = "0.0.0.0"
)

var port = getOptionalEnv("PORT", "8000")

func New(db *sql.DB) *http.Server {
	payrollService := newPayrollServiceHandlerWithLog(
		newWithAuth(
			&payrollServiceServer{repos.NewPayrollRepo(db)},
		))
	path, handler := payrollv1connect.NewPayrollServiceHandler(payrollService)

	mux := http.NewServeMux()
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
