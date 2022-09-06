test:
	go test -v ./...

mock:
	mockgen -source=database/repository.go -destination=database/mock/mock_repository.go -package=mock