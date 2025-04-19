package models_symbol_list

import (
	"backend/internal/enums"
	enums_symbol_list "backend/internal/enums/symbol_list"
)

type UpdateRequestModel struct {
	Group         string                       `json:"group" validate:"omitempty,alpha,uppercase"`
	Volume        int                          `json:"volume" validate:"number,gte=0"`
	SortColumn    enums_symbol_list.SortColumn `json:"sortColumn" validate:"required,symbolListSortColumn"`
	SortDirection enums.SortDirection          `json:"sortDirection" validate:"required,sortDirection"`
}
