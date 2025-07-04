package user_project

import (
	"database/sql"
)

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (f optionFunc) apply(c *config) { f(c) }

type config struct {
	SQLTx *sql.Tx
}

func applyAllOptions(opts []Option) config {
	res := config{}
	for _, o := range opts {
		o.apply(&res)
	}
	return res
}

func WithSQLTransaction(tx *sql.Tx) optionFunc {
	return func(c *config) {
		c.SQLTx = tx
	}
}
