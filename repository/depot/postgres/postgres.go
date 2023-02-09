package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/entity"
	"github.com/mansoorceksport/warehouse-stocking/repository/depot"
	"log"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgresWarehouse(connectionString string) depot.Repository {
	postgres := &Postgres{}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed to open database connection, %v", err)
	}
	postgres.db = db
	err = postgres.db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database, %v", err)
	}
	return &Postgres{}
}

func (p *Postgres) Add(ctx context.Context, aw aggregate.Warehouse) error {
	sqlQuery := "insert into depot(name) values($1)"
	_, err := p.db.QueryContext(ctx, sqlQuery, aw.GetName())
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetById(ctx context.Context, uuid uuid.UUID) (aggregate.Warehouse, error) {
	sqlQuery := "select uuid, name from depot where uuid = $1"
	row := p.db.QueryRowContext(ctx, sqlQuery, uuid)
	d := &entity.Depot{}
	err := row.Scan(&d.ID, &d.Name)
	if err != nil {
		return aggregate.Warehouse{}, err
	}
	return aggregate.Warehouse{}, nil
}
