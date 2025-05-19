.PHONY: update

# Docker services
DOCKER_COMPOSE_PROJECT_NAME=imagimation_analyzer_services_local
LOCAL_DOCKER_COMPOSE_FILE=./docker/local/docker-compose.yml

up_local_services:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) up -d

down_local_services:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) down

restart_local_services: down_local_services up_local_services

local_services_logs:
	docker compose -p $(DOCKER_COMPOSE_PROJECT_NAME) -f $(LOCAL_DOCKER_COMPOSE_FILE) logs

# Deployment
DEFAULT_SERVICE_NAME = imagimation_analyzer_service

get_updates:
	git pull
	rm exercise_manager
	go build -o exercise_manager cmd/api/main.go
	sudo systemctl restart $(or $(SERVICE_NAME), $(DEFAULT_SERVICE_NAME))
	sudo systemctl --no-pager -l status $(or $(SERVICE_NAME), $(DEFAULT_SERVICE_NAME))
