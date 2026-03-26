package repository

import (
	"time"
)

// ==========================
// Pagination
// ==========================

// Pagination defines standard pagination options for mobile apps
type Pagination struct {
	Limit  int `json:"limit"`  // Maximum number of records per request
	Offset int `json:"offset"` // Start index
}

// ==========================
// Filter
// ==========================

// FilterOperator defines the type of comparison
type FilterOperator string

const (
	OpEquals        FilterOperator = "="
	OpNotEquals     FilterOperator = "!="
	OpGreaterThan   FilterOperator = ">"
	OpLessThan      FilterOperator = "<"
	OpGreaterEquals FilterOperator = ">="
	OpLessEquals    FilterOperator = "<="
	OpLike          FilterOperator = "LIKE" // for text search
	OpIn            FilterOperator = "IN"
)

// Filter represents a single filter condition
type Filter struct {
	Field    string         `json:"field"`    // name of the field/column
	Operator FilterOperator `json:"operator"` // comparison operator
	Value    any            `json:"value"`    // value to compare
}

// ==========================
// Sorting
// ==========================

// SortDirection defines ascending or descending
type SortDirection string

const (
	SortAsc  SortDirection = "ASC"
	SortDesc SortDirection = "DESC"
)

// Sort defines a single sort condition
type Sort struct {
	Field     string        `json:"field"`     // field/column name
	Direction SortDirection `json:"direction"` // ASC or DESC
}

// ==========================
// Query Options (Filter + Sort + Pagination)
// ==========================

// QueryOptions combines filters, sorting, and pagination
type QueryOptions struct {
	Filters    []Filter   `json:"filters,omitempty"`    // optional filters
	Sorts      []Sort     `json:"sorts,omitempty"`      // optional sorting
	Pagination Pagination `json:"pagination,omitempty"` // optional pagination
}

// ==========================
// Helper: Date Range Filter
// ==========================

// DateRangeFilter helps filter by time intervals
type DateRangeFilter struct {
	Field string    `json:"field"` // field/column name
	From  time.Time `json:"from"`  // start datetime
	To    time.Time `json:"to"`    // end datetime
}
