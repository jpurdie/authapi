package database

import (
	"github.com/go-pg/pg"
)

// User represents the client for company_user table

type PingStore struct {
	db *pg.DB
}

// NewAdmAccountStore returns an AccountStore.
func NewPingStore(db *pg.DB) *PingStore {
	return &PingStore{
		db: db,
	}
}
func (s *PingStore) Ping() error {
	return nil
}

//
//// Custom errors
//var (
//	ErrCompAlreadyExists  = echo.NewHTTPError(http.StatusConflict, "Company name already exists.")
//	ErrEmailAlreadyExists = echo.NewHTTPError(http.StatusConflict, "Email already exists.")
//)
//
//// Create creates a new user on database
//func (env *Env) Create(cu authapi.CompanyUser) (authapi.CompanyUser, error) {
//
//	var n string
//	tempDb, _ := postgres.Init()
//	_, err := tempDb.QueryOne(pg.Scan(&n), "SELECT now() ")
//	tempDb.Close()
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Connected to database at: " + n)
//
//	var company = new(authapi.Company)
//	tempDb, err = postgres.Init()
//	if err != nil {
//		panic(err)
//	}
//	count, err := tempDb.Model(company).Where("lower(name) = ? and deleted_at is null", strings.ToLower(cu.Company.Name)).Count()
//	tempDb.Close()
//	if err != nil {
//		return authapi.CompanyUser{}, err
//	}
//	if count > 0 {
//		return authapi.CompanyUser{}, ErrCompAlreadyExists
//	}
//	var user = new(authapi.User)
//
//	tempDb, _ = postgres.Init()
//	count, err = tempDb.Model(user).Where("lower(email) = ? and deleted_at is null", strings.ToLower(cu.User.Email)).Count()
//	tempDb.Close()
//
//	if err != nil {
//		return authapi.CompanyUser{}, err
//	}
//	if count > 0 {
//		return authapi.CompanyUser{}, ErrEmailAlreadyExists
//	}
//
//	tempDb, _ = postgres.Init()
//	print(tempDb.PoolStats().TotalConns)
//	tempDb, _ = postgres.Init()
//	tx, err := tempDb.Begin()
//	tx.Model(cu.Company).Insert()
//	cu.User.CompanyID = cu.Company.ID
//	tx.Model(cu.User).Insert()
//	cu.UserID = cu.User.ID
//	cu.CompanyID = cu.Company.ID
//	tx.Model(&cu).Insert()
//	trErr := tx.Commit()
//	if trErr != nil {
//		tx.Rollback()
//	}
//	tempDb.Close()
//	return cu, err
//}

//func (env *Env) List() ([]authapi.Company, error) {
//	var companies []authapi.Company
//	tempDb, _ := postgres.Init()
//	defer tempDb.Close()
//
//	_, err := tempDb.QueryOne(pg.Scan(&n), "SELECT now() ")
//	return companies, err
//}

//// View returns single user by ID
//func (co Company) View(db orm.DB, id int) (authapi.Company, error) {
//	var company authapi.Company
//	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name"
//	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id"
//	WHERE ("user"."id" = ? and deleted_at is null)`
//	_, err := db.QueryOne(&company, sql, id)
//	return company, err
//}
//
//// Update updates user's contact info
//func (co Company) Update(db orm.DB, company authapi.Company) error {
//	_, err := db.Model(&company).WherePK().UpdateNotZero()
//	return err
//}
//
//// List returns list of all users retrievable for the current user, depending on role

//
//// Delete sets deleted_at for a user
//func (co Company) Delete(db orm.DB, company authapi.Company) error {
//	return db.Delete(&company)
//}