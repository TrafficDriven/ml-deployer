.PHONY: ci \
	lint \
	helm_lint \
	setup \
	bump \
	docker-image \

# Install libvips
setup:
	@sudo apt-get install -y libvips libvips-tools libvips-dev

bump:
	@CONFIG_FILE=config.yml go run tools/bump/main.go

bump1:
	@CONFIG_FILE=config1.yml go run tools/bump/main.go

bump2:
	@CONFIG_FILE=config2.yml go run tools/bump/main.go

bump3:
	@CONFIG_FILE=config3.yml go run tools/bump/main.go

bump4:
	@CONFIG_FILE=config4.yml go run tools/bump/main.go

version:
	@go run tools/version/main.go

version1:
	@go run tools/version/main.go

docker-image:
	$(eval VERSION=$(shell go run tools/version/main.go))
	docker build . -t us-central1-docker.pkg.dev/tdt-platform/mintlist/ml-deployer:$(VERSION)
