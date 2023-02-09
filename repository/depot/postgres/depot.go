package postgres

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/common/postgres"
	"github.com/mansoorceksport/warehouse-stocking/entity"
	"github.com/mansoorceksport/warehouse-stocking/repository/depot"
)

type DepotPostgres struct {
	postgres *postgres.Postgres
}

func NewPostgresWarehouse(postgres *postgres.Postgres) depot.Repository {
	p := &DepotPostgres{}
	p.postgres = postgres
	return p
}

func (dp *DepotPostgres) Add(ctx context.Context, aw aggregate.Warehouse) error {
	sqlQuery := "insert into depot(name) values($1)"
	_, err := dp.postgres.DB.QueryContext(ctx, sqlQuery, aw.GetName())
	if err != nil {
		return err
	}
	return nil
}

func (dp *DepotPostgres) GetById(ctx context.Context, uuid uuid.UUID) (aggregate.Warehouse, error) {
	sqlQuery := "select uuid, name from depot where uuid = $1"
	row := dp.postgres.DB.QueryRowContext(ctx, sqlQuery, uuid)
	d := &entity.Depot{}
	err := row.Scan(&d.ID, &d.Name)
	if err != nil {
		return aggregate.Warehouse{}, err
	}
	return aggregate.Warehouse{}, nil
}
