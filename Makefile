# Previous
start: build
	@./bin/main

build:
	@go build -o ./bin ./cmd/api/main.go

# start: 
# 	@docker compose up --build
	
# stop:
# 	@docker-compose rm -v --force --stop
# 	@docker rmi ticket-booking