package models

import "time"

type Customer struct {
	ID               uint       `json:"id" db:"id"`
	FirstName        string     `json:"firstName" db:"first_name"`
	LastName         string     `json:"lastName" db:"last_name"`
	MiddleName       string     `json:"middleName" db:"middle_name"`
	FeCustomerUserID *string    `json:"feCustomerUserId" db:"fe_customer_user_id"`
	TimeZone         *string    `json:"timeZone" db:"time_zone"`
	Email            *string    `json:"email" db:"email"`
	CheckWord        *string    `json:"checkWord" db:"check_word"`
	GenderID         *uint      `json:"genderId" db:"gender_id"`
	MaritalStatus    *bool      `json:"maritalStatus" db:"marital_status"`
	ChildrenAmount   *uint      `json:"childrenAmount" db:"children_amount"`
	CreateDate       *time.Time `json:"createDate" db:"create_date"`
	CustomerTypeID   *uint      `json:"customerTypeId" db:"customer_type_id"`
	OrganizationID   *uint      `json:"organizationId" db:"org_id"`
	MaidenName       *string    `json:"maidenName" db:"maiden_name"`
	MaritalStatusID  *uint      `json:"maritalStatusId" db:"marital_status_id"`
	ChildrenAmountID *uint      `json:"childrenAmountId" db:"children_amount_id"`
	MaidenFirstname  *string    `json:"maidenFirstname" db:"maiden_firstname"`
}
