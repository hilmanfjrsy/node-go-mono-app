auth-up:
	cd auth-app/ && docker-compose up -d 

auth-down:
	cd auth-app/ && docker-compose down 

auth-start:
	cd auth-app/ && docker-compose start

auth-stop:
	cd auth-app/ && docker-compose stop

fetch-up:
	cd fetch-app/ && docker-compose up -d 

fetch-down:
	cd fetch-app/ && docker-compose down 

fetch-start:
	cd fetch-app/ && docker-compose start

fetch-stop:
	cd fetch-app/ && docker-compose stop