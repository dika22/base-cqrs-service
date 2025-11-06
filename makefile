APP=auth-service
APP_EXECUTABLE=${APP}

serve-http:
	go run main.go serve-http

migrate:
	go run main.go migrate