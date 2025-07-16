package postgres

import (
	"fmt"
	"strings"
)

// PaginationParams berisi parameter untuk pagination, filter dan sorting.
type PaginationParams struct {
	Limit   int
	Offset  int
	SortBy  string                 // Kolom untuk sorting
	SortAsc bool                   // true = ascending, false = descending
	Filter  map[string]interface{} // key = kolom, value = nilai filter
}

// BuildSQL mengembalikan potongan SQL WHERE, ORDER BY, LIMIT OFFSET untuk pagination.
func (p *PaginationParams) BuildSQL() (whereClause string, orderBy string, limitOffset string, args []interface{}) {
	args = []interface{}{}
	whereParts := []string{}

	i := 1
	for col, val := range p.Filter {
		whereParts = append(whereParts, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	if len(whereParts) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	if p.SortBy != "" {
		dir := "ASC"
		if !p.SortAsc {
			dir = "DESC"
		}
		orderBy = fmt.Sprintf("ORDER BY %s %s", p.SortBy, dir)
	}

	if p.Limit > 0 {
		limitOffset = fmt.Sprintf("LIMIT %d OFFSET %d", p.Limit, p.Offset)
	}

	return
}
