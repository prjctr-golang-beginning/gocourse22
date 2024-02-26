package providers

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"
	"gocourse22/cmd/flag"
	"gocourse22/internal/db"
)

func ProvideConnection(i *do.Injector) {
	do.ProvideNamed(i, "postgres", ProvidePostgresConnection)
}

func ProvidePostgresConnection(i *do.Injector) (*pgxpool.Pool, error) {
	c := do.MustInvoke[*cli.Context](i)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.String(flag.PostgresHost),
		c.Int(flag.PostgresPort),
		c.String(flag.PostgresUser),
		c.String(flag.PostgresPass),
		c.String(flag.PostgresDBName),
	)

	return db.NewConnectionPool(c.Context, dsn)
}
