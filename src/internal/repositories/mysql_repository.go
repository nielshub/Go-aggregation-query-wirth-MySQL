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

func (mS *MySqlRepository) CreateDataBaseAndTable() error {
	_, err := mS.client.Exec("CREATE DATABASE IF NOT EXISTS dataset")
	if err != nil {
		return err
	}
	_, err = mS.client.Exec(`
	CREATE TABLE IF NOT EXISTS dataset (
		pk_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		action_timestamp DATE NOT NULL,
		client_event varchar(45) NOT NULL
	)  ENGINE=INNODB;`)
	if err != nil {
		return err
	}
	return nil
}

func (mS *MySqlRepository) IngestFileData(ctx context.Context, path string) error {
	mysql.RegisterLocalFile(path)
	_, err := mS.client.Exec("LOAD DATA LOCAL INFILE '" + path + "' INTO TABLE dataset")
	if err != nil {
		return err
	}
	return nil
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
	// _, err := mS.client.Exec("ALTER TABLE dataset ADD UNIQUE INDEX pk_id (user_id, actiontimestamp, clientevent);")
	// if err != nil {
	// 	return err
	// }
	return nil
}

/*DELETE c1 FROM tablename c1
INNER JOIN tablename c2
WHERE
    c1.id > c2.id AND
    c1.unique_field = c2.unique_field;
*/
