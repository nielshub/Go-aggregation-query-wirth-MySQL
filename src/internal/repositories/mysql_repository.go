package repositories

import (
	"contentSquare/src/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (mS *MySqlRepository) CreateDataBaseAndTable() error {
	_, err := mS.client.Exec("CREATE DATABASE IF NOT EXISTS dataset")
	if err != nil {
		return err
	}
	_, err = mS.client.Exec(`
	CREATE TABLE IF NOT EXISTS dataset (
		pk_id BIGINT AUTO_INCREMENT PRIMARY KEY,
		user_id BIGINT NOT NULL,
		action_timestamp DATETIME NOT NULL,
		client_event varchar(45) NOT NULL,
		UNIQUE KEY unique_action (user_id,action_timestamp,client_event)
	)  ENGINE=InnoDB;`)
	if err != nil {
		return err
	}
	return nil
}

func (mS *MySqlRepository) IngestFileData(ctx context.Context, path string) error {
	mysql.RegisterLocalFile(path)
	_, err := mS.client.Exec(`
		LOAD DATA LOCAL INFILE '` + path + `' INTO TABLE dataset
		FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n'
		(user_id, action_timestamp, client_event)
		;`)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MySqlRepository) CountEvents(ctx context.Context, filters models.Filters) (int64, error) {
	var countValue int64
	var query string

	if filters.Event != "" && filters.UserId != "" {
		query = `
		SELECT COUNT(*) FROM dataset
		WHERE client_event = '` + filters.Event + `'
		AND user_id = '` + filters.UserId + `'
		AND action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	} else if filters.Event != "" && filters.UserId == "" {
		query = `
		SELECT COUNT(*) FROM dataset
		WHERE client_event = '` + filters.Event + `'
		AND action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	} else if filters.UserId != "" && filters.Event == "" {
		query = `
		SELECT COUNT(*) FROM dataset
		WHERE client_event = '` + filters.UserId + `'
		AND action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	} else if filters.UserId == "" && filters.Event == "" {
		query = `
		SELECT COUNT(*) FROM dataset
		AND action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	}
	queryOutput, err := ms.client.Query(query)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	defer queryOutput.Close()
	for queryOutput.Next() {
		if err := queryOutput.Scan(&countValue); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}

	return countValue, nil
}

func (ms *MySqlRepository) CountDistinctUsers(ctx context.Context, filters models.Filters) (int64, error) {
	var countValue int64
	var query string

	if filters.Event != "" {
		query = `
		SELECT COUNT(DISTINCT user_id) FROM dataset
		WHERE client_event = '` + filters.Event + `'
		AND action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	} else if filters.Event == "" {
		query = `
		SELECT COUNT(DISTINCT user_id) FROM dataset
		WHERE action_timestamp BETWEEN CAST('` + filters.DateFrom + `' AS DATETIME) AND CAST('` + filters.DateTo + `' AS DATETIME);
		`
	}
	queryOutput, err := ms.client.Query(query)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	defer queryOutput.Close()
	for queryOutput.Next() {
		if err := queryOutput.Scan(&countValue); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}

	return countValue, nil
}

func (ms *MySqlRepository) Exists(ctx context.Context, filters models.Filters) (bool, error) {
	var exists bool
	var query string

	if filters.Event != "" && filters.UserId != "" {
		query = `
		SELECT EXISTS(SELECT * FROM dataset
		WHERE client_event = '` + filters.Event + `'
		AND user_id = '` + filters.UserId + `');
		`
	} else {
		return false, errors.New("Missing filter for exisits query")
	}
	queryOutput, err := ms.client.Query(query)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	defer queryOutput.Close()
	for queryOutput.Next() {
		if err := queryOutput.Scan(&exists); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}

	return exists, nil
}

func (mS *MySqlRepository) RemoveDuplicates(ctx context.Context) error {
	_, err := mS.client.Exec(`
		create temporary 
			table dataset_copy (pk_id int);
		`)
	if err != nil {
		return err
	}
	_, err = mS.client.Exec(`
		insert into
			dataset_copy (pk_id)
		select
			pk_id
		from
			dataset ds
		where
			exists (
			select
				*
			from
				dataset ds2
			where
				ds2.client_event = ds.client_event
				and ds2.user_id = ds.user_id
				and ds2.action_timestamp = ds.action_timestamp
				and ds2.pk_id > ds.pk_id);
		`)
	if err != nil {
		return err
	}
	_, err = mS.client.Exec(`
		delete from 
			dataset 
		using 
			dataset, dataset_copy where dataset.pk_id=dataset_copy.pk_id;`)
	if err != nil {
		return err
	}
	return nil
}
