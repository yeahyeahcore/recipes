package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm/logger"
)

type ConnectDeps struct {
	Configuration *Configuration
	Timeout       time.Duration
}

func Connect(ctx context.Context, deps *ConnectDeps) (*pgxpool.Pool, error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		deps.Configuration.User,
		deps.Configuration.Password,
		deps.Configuration.Host,
		deps.Configuration.Port,
		deps.Configuration.DatabaseName,
	)

	ticker := time.NewTicker(1 * time.Second)
	timeoutExceeded := time.After(deps.Timeout)

	defer ticker.Stop()

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", deps.Timeout)
		case <-ticker.C:
			connection, err := pgxpool.New(ctx, uri)
			if err == nil {
				return connection, nil
			}

			logger.Default.Error(context.Background(), fmt.Sprintf("failed to connect to db <%s>", uri))
		}
	}
}
