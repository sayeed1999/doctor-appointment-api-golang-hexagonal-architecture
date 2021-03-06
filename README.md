# Golang Enterprise Project Structure | Hexagonal Architecture - Onion Architecture - Generic Repository Pattern

## Doctor Appointment Api

To run the project from your pc, 

- first, you have to get golang installed on your system
- second, git clone it, then run **go mod tidy** which installs all dependencies
- third, run **go run main.go** in terminal from root dir
- since, i have removed config.yaml file in git repo, you have to create ./config/config.yaml file and write it matching the Config struct in ./config/config.go file

---

In this golang project, I tried to apply all the enterprise level software design practices that I know e.g hexagonal architecture, onion architecture, clean architecture, CQRS, DDD, generic repository pattern..

---

## Golang Packages used:-

- **gorm** (ORM to communicate with the database)
- **gin/gonic** (golang framework to build apis)
- **gomail** (golang package for sending emails from your gmail)
- **go-redis** (golang package which implements redis)

---

## Architectures targeted:-

- **Generic repository pattern** (centralizes a common repository for database operations)
- **Onion architecture** (segregates a monolith project into controller layer, buisness layer, data access layer, domain/entity layer)
- **Hexagonal architecture**

---

## Project architecture explained:-

- **go.mod** & **go.sum** file are responsible for all the dependencies of the project
- Program execution starts from main.go
- /internal/core package is the core of the project which contains all buisness logic and is loosely coupled to the rest of the application via ports-adapters pattern
- First, the configuration files are initialized from ./config dir
- Second, a channel for sending mails gets created and a new goroutine is fired that keeps listening for mail sending tasks
- Third, the central generic repository gets initialized which is responsible for database transactions
- Fourth, redis server gets initialized which is responsible for the caching of data
- Fifth, the service layer gets initialized which calls on the repository layer for the database transactions and does caching, mail sending
- Sixth, the handlers are initialized which is the controller layer of this project. the client requests data from the handlers via http requests & the handlers communicate back and forth with the service layer.
- Seventh, gin router is created with default middlewares
- Eighth, middlewares are initialized
- Ninth, routes are initialized
- Tenth, the server is run on localhost:8080 until an error occurs..

---

Any more suggestions please let me know.. :)
