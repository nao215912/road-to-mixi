up:
	cp ./app/.air.norml.conf ./app/.air.toml
	docker-compose up

debug:
	cp ./app/.air.debug.conf ./app/.air.toml
	docker-compose up

stop:
	docker-compose stop

db:
	docker-compose exec db mysql -u root

test_db:
	docker-compose exec test_db mysql -u root

app:
	docker-compose exec app bash

test:
	docker-compose exec app go test ./...


del:
	docker system prune -af


.PHONY: up debug stop del app test