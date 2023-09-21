include .env.dev
export

db-create:
	docker run --name ${DB_NAME} -d -p ${DB_PORT}:27017 -e MONGO_INITDB_ROOT_USERNAME=${DB_USERNAME} -e MONGO_INITDB_ROOT_PASSWORD=${DB_PASSWORD} -e MONGO_INITDB_DATABASE=admin mongo:6-jammy
db-rm:
	docker stop ${DB_NAME} && docker rm ${DB_NAME}
