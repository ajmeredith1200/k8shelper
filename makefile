install:
	@echo "Installing..."
	go install
	@echo "Done!"

lint:
	@echo "Linting..."
	golangci-lint run
	if [ $$? -eq 1 ]; then \
		echo "Linting failed!"; \
		exit 1; \
	fi
	@echo "Done!"

test:
	@echo "Testing..."
	go test -v ./...
	@echo "Done!"

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
	@echo "Done!"

run:
	@echo "Running..."
	go run main.go
	@echo "Done!"