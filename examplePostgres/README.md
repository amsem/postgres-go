# this go project show how we could connect a postgresql DB with pq driver

## before running the program make sure to spin up a docker container with postgresql
## ,create a user to perform the queries and create a db
## make sure that the user have all granted permision over the DB
## run this commands 
### GRANT ALL ON DATABASE <<DB>> TO <<user>>;
### ALTER DATABASE <<DB>> OWNER TO <<user >>;
### GRANT USAGE, CREATE ON SCHEMA PUBLIC TO <<user>>;
