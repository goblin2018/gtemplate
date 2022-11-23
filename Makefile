.PHONY: webapp
webapp: webapp-build webapp-container
	@echo Done build webapp

webapp-build:
	@echo start build webapp ...
	cd webapp && pnpm build

webapp-container:
	@echo build webapp docker image
	docker compose up --build -d webapp

