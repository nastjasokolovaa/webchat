.PHONY: dev_env_start
dev_env_start:
	docker-compose -f ./docker/docker-compose.yml up -d chatdb

.PHONY: dev_migrate
dev_migrate:
	docker-compose -f ./docker/docker-compose.yml up goose

.PHONY: dev_env_stop
dev_env_stop:
	docker-compose -f ./docker/docker-compose.yml stop
