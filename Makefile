docker:
	docker build -t mayapada-test .

run:
	docker-compose up --build -d

stop:
	docker-compose down