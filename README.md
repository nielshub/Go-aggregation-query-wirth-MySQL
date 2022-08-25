# ContentSquare code test

The purpose of this service is to create a datastore and allow aggregations with a simple API.

## Definition

Instructions from the readme has been used. Go lenguage has been choosed and MySQL database because of the performance and the capabilities of the querys.

All endpoints have been created and tested with the data.txt file given and properly tested with the exact results.

## Instructions to start the application on localhost

- Install docker and docker-compose
- Add the file with the data to the folder
- Create variables.env file inside env folder (right now there are no critical values but here we could save apikey and sensitive stuff needed for real environments) with following values:

```[env]
MYSQL_DSN = "user:password@tcp(mysqlDB:3306)/dataset"
VERSION = "1.0"
ENV = "LOCAL"
FILEPATH = "./data.txt"
```

- Ensure that the filepath is the correct one with the data file
- Run following command in main folder of the repo: docker-compose up --build
- In order to run the test manually in console run: go test -v ./...
- Postman collection has been done, please see attached in the repository. In order to make it work you can:
  - Check health
  - Count events
  - Exists
  - Count distinct users

## An explanation of the choices taken and assumptions made during development

A repository pattern an a hexagonal arquitecture has been applied. It is a very simple microservice but with this structure it should be easy to scalate with other services and handlers.
Moreover, with the docker image and docker compose is easy to implement in different environments using K8 and AWS stuff.

Unit test coverage is quite basic, mainly for the handlers which has some logic and the Db repo has been mocked with the proper interface defines in the ports.

## Additional questions

- If you had > 100GB of data, and you didn't code a solution that would handle so much data:
    If we continue using a RDBMS such as MySQL we should split all this data into smaller tables with smaller index that are more easy to manage, and maybe start using MySQL Cluster to manage different MySQL server and allow a better performance. also it should be possibe to migrate to HDFS (Hadoop Distributed File System).
- Are there any optimisations that you could have implemented but didn't have time to?
    Optimisations with maybe other techstack in DB more big data oriented. Optimisation when loading data into DB and checking repited rows.
- If you had a continuous stream in input, how would you manage to respond to aggregation queries?
    Use of automated time-based partitioning and local-only indexing to achieve high insert rates. With a TimescaleDB architecture. With some predefined aggregates it should just query for the new ones and combine it with the old results that we have with an older timestamp stored in a materialization table.
    Moreover, if we want a real-time aggregation we can combine the results from two tables, the old data (materialized one with predefined aggregations), and the introduced in real time (between thresholds), and from there generate the real time aggregations.
