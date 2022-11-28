.PHONY: webapp
webapp: webapp-build webapp-container
	@echo Done build webapp

webapp-build:
	@echo start build webapp ...
	cd webapp && pnpm build

webapp-container: tag
	@echo build webapp docker image
	docker compose up --build -d webapp

# get webapp git tag version and modify docker-compose file config
webversion := $(shell git describe) 
tag:
	@echo verison is $(webversion)
	sed -r -i 's/(webapp:)(v.+)/\1$(webversion)/' ./docker-compose.yml

