package handler

import (
	"fmt"

	"github.com/blackbarth/goopportunities.git/config"
	"gorm.io/gorm"
)

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type UpdateOpeningRequest struct {
	Role	 string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}


var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

func GetLogger() *config.Logger {
	return logger
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	config.CloseSQLite()
}


func (r *CreateOpeningRequest) Validate() error {

	if r == nil {
		return fmt.Errorf("Request malformed: nil request body")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}	
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("Param %s (type: %s) is required", name, typ)
}


func (r *UpdateOpeningRequest) Validate() error {

	// If any field is provided, validation is truthy	
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("At least one valid field must be provided")
	 return nil
}
