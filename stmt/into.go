package stmt

import (
	"github.com/ulule/loukoum/token"
	"github.com/ulule/loukoum/types"
)

// Into is the INTO clause.
type Into struct {
	Statement
	Table Table
}

// NewInto returns a new Into instance.
func NewInto(table Table) Into {
	return Into{
		Table: table,
	}
}

// Write expose statement as a SQL query.
func (into Into) Write(ctx *types.Context) {
	ctx.Write(token.Into.String())
	ctx.Write(" ")
	into.Table.Write(ctx)
}

// IsEmpty return true if statement is undefined.
func (into Into) IsEmpty() bool {
	return into.Table.IsEmpty()
}
