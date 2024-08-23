# Migration folder from the perspective of the root directory of the project, since that is where the script will be called from.
migrations_folder="./database/migrations/"
migrate -path ./database/migrations/ -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$DB_ADDR:$DB_PORT/$POSTGRES_DB?sslmode=disable" -verbose up
