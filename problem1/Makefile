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


del:
	docker system prune -af


.PHONY: up debug stop del