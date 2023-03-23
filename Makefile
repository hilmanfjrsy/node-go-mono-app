auth-up:
	cd auth-app/ && docker-compose up -d 

auth-down:
	cd auth-app/ && docker-compose down --rmi local

auth-start:
	cd auth-app/ && docker-compose start

auth-stop:
	cd auth-app/ && docker-compose stop

fetch-up:
	cd fetch-app/ && docker-compose up -d 

fetch-down:
	cd fetch-app/ && docker-compose down --rmi local

fetch-start:
	cd fetch-app/ && docker-compose start

fetch-stop:
	cd fetch-app/ && docker-compose stop

test:
	cd fetch-app/ && go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && ENV_STAGE=test