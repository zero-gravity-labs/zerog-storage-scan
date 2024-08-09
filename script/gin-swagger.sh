# gin middleware to automatically generate RESTful API documentation with Swagger 2.0.
swag fmt
swag init -g api.go --parseDependency -d ./api/storage --instanceName storage --md ./docs/markdown/storage
swag init -g api.go --parseDependency -d ./api/da  --instanceName da --md ./docs/markdown/da