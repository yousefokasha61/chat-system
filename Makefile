run:
	docker network create shared_network || true
	docker-compose -f chat/docker-compose.yml up -d
	docker-compose -f chat_api/docker-compose.yml up -d