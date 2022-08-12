package repositories

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	client *sql.DB
}

func NewMySqlRepository() *MySqlRepository {
	return &MySqlRepository{
		client: CreateDBClient(context.Background(), "mysql"),
	}
}

func CreateDBClient(ctx context.Context, driverName string) *sql.DB {
	client, err := sql.Open(driverName, os.Getenv("MYSQL_DSN"))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func (mS *MySqlRepository) IngestFileData(ctx context.Context, path string) error {
	mysql.RegisterLocalFile(path)
	_, err := mS.client.Exec("LOAD DATA LOCAL INFILE '" + path + "' INTO TABLE dataset")
	if err != nil {
		return err
	}
	return nil
}
