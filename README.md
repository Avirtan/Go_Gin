# Exemple Go Gin and Vue Admin
## For Backend used Redis, Postgres, Gorm, WS, Swagger, JWT
## For Frontend used Vue3(TS), Vite, Vuex, Quasar 
Made a simple migration tool for Gorm
How work:
```bash
## for build exec file
cd service.*/migrations/
go build ../../pkg/utils/migrations.go
## for create migrations file
./migration -d=name_migration
## update method up in file name_migration####-##-##.go and run
go run *.go -a=up
```