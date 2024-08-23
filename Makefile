migrateup:
	. ./.env; . ./scripts/migrateup.sh

migratedown:
	. ./.env; . ./scripts/migratedown.sh

.PHONY: migrateup migratedown
