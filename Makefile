build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

MODULES=api-gateway auth-user-profile vocabulary learning-engine scheduler-notification
lint:
	for m in ${MODULES}; do golangci-lint run $$m; done
