# To-do API
to-do API written using:
* Gin Framework
* GORM

## Installation & Run
```
# Download the project
$ go get github.com/mvrsss/go-todo-api

# Download Gin
$ go get -u github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```
Setting DB in config/database.go
```
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host: "0.0.0.0",
			Port: 3306,
			User: "root",
			Name: "todoapp",
		},
	}
}
```
## Structure
```
├── app
│   ├── Routes.go        // Router
│   ├── controllers                
│   │   └── Tasks.go      
│   └── models
│   |   └── Model.go     // APi Model
|   └── migrate
|       └── Todo.go 
├── config
│   └── database.go        // Configuration
└── main.go
```
## API
**/todo**
* ```GET```: Get all tasks
* ```POST```: Create a new task

**/todo/:id**
* ```GET```: Get a task
* ```PUT```: Update a task
* ```DELETE```: Delete a task
