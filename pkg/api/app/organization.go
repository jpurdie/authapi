package app

import (
	"github.com/google/uuid"
	"github.com/jpurdie/authapi"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Organization defines database operations for Organization.
type OrganizationStore interface {
	ListAccessible(u *authapi.User, includeInactive bool) ([]authapi.OrganizationUser, error)
	//List() error
}

// Organization Resource implements account management handler.
type OrganizationResource struct {
	Store OrganizationStore
}

func NewOrganizationResource(store OrganizationStore) *OrganizationResource {
	return &OrganizationResource{
		Store: store,
	}
}
func (rs *OrganizationResource) router(r *echo.Group) {
	log.Println("Inside Organization Router")
	r.GET("/ping", rs.ping)
	r.GET("", rs.listAuthorized)
}

var (
	ErrEmailAlreadyExists   = authapi.ErrorResp{Error: authapi.Error{CodeInt: http.StatusConflict, Message: "The user already exists"}}
	ErrPasswordsNotMatching = authapi.ErrorResp{Error: authapi.Error{CodeInt: http.StatusConflict, Message: "Passwords do not match"}}
	ErrPasswordNotValid     = authapi.ErrorResp{Error: authapi.Error{CodeInt: http.StatusConflict, Message: "Password is not in the required format"}}
	UnknownError            = authapi.ErrorResp{Error: authapi.Error{CodeInt: http.StatusConflict, Message: "There was a problem registering."}}
	ErrAuth0Unknown         = authapi.ErrorResp{Error: authapi.Error{CodeInt: http.StatusConflict, Message: "There was a problem registering with provider."}}
)

type listAuthorizedRespInner struct {
	OrgName string       `json:"name"`
	UUID    uuid.UUID    `json:"uuid"`
	Role    authapi.Role `json:"role"`
}

type listAuthorizedResp struct {
	Orgs []listAuthorizedRespInner `json:"orgs"`
}

func (rs *OrganizationResource) listAuthorized(c echo.Context) error {
	log.Println("Inside listAuthorized(first)")

	u := authapi.User{
		ExternalID: c.Get("sub").(string),
	}

	organizationUser, err := rs.Store.ListAccessible(&u, false)

	if err != nil {
		log.Println(err)
		if errCode := authapi.ErrorCode(err); errCode != "" {
			return c.JSON(http.StatusInternalServerError, ErrAuth0Unknown)
		}
		return c.JSON(http.StatusInternalServerError, ErrAuth0Unknown)

	}
	x := listAuthorizedResp{}
	for _, tempOrgUser := range organizationUser {
		temp := listAuthorizedRespInner{
			OrgName: tempOrgUser.Organization.Name,
			UUID:    tempOrgUser.Organization.UUID,
			Role:    *tempOrgUser.Role,
		}
		x.Orgs = append(x.Orgs, temp)
	}

	return c.JSON(http.StatusOK, x)
}

func (rs *OrganizationResource) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
