dockerup:
	. ./scripts/docker_compose_up.sh

dockerdown:
	. ./scripts/docker_compose_down.sh

migrateup:
	. ./.env; . ./scripts/migrateup.sh

migratedown:
	. ./.env; . ./scripts/migratedown.sh

.PHONY: migrateup migratedown dockerup dockerdown
