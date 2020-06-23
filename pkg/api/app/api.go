package app

import (
	"github.com/go-pg/pg"
	"github.com/jpurdie/authapi/pkg/api/database"
	authMw "github.com/jpurdie/authapi/pkg/utl/middleware/auth"
	"github.com/labstack/echo"
	"net/http"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
)

// API provides admin application resources and handlers.
type API struct {
	Companies *CompanyResource
}

// NewAPI configures and returns admin application API.
func NewAPI(db *pg.DB) (*API, error) {

	companyStore := database.NewCompanyStore(db)
	company := NewCompanyResource(companyStore)

	api := &API{
		Companies: company,
	}
	return api, nil
}
func (a *API) Router(r *echo.Group) {

	companies := r.Group("/companies")
	a.Companies.router(companies)
	authMiddleware := authMw.Middleware()
	r.Use(authMiddleware)

	r.GET("/ping", func(c echo.Context) error {
		cc := c.(*authMw.CustomContext)
		x := cc.Sub()
		if 1 == 2 {
			print(x)
		}
		return c.String(http.StatusOK, "pong")
	})

}
