auth-up:
	cd auth-app/ && docker-compose up -d 

auth-down:
	cd auth-app/ && docker-compose down 

fetch-up:
	cd fetch-app/ && docker-compose up -d 

fetch-down:
	cd fetch-app/ && docker-compose down 