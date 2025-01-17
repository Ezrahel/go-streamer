#DEV
build-dev:
	docker build -t gostreamer -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn . 

clean-dev:
	docker-composer -f containers/composes/dc.dev.yml.down

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up