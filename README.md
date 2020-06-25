# Patal Web Backend
Patal Web is repository for https://palembangdigital.org/

### Getting Started
To run this project localy, make sure minimum requirements are fulfilled.
- `go` version 1.10 or higher

### Running in Local Machine
1. **Make sure go is installed as global command** (first time only)

2. **Clone this project and go to the root project to install all dependencies** (first time only)
```
// clone the repository
> git clone git@github.com:palembang-digital/web-backend.git

// change directory to root project web-frontend folder
> cd web-backend
> change config-example.yaml > config.yaml

// install all the dependencies
> go mod tidy
```
3. **While still in root project build and run the app**
```
// build and run the app
> go build -v -o web-backend
> ./web-backend


// now go to http://localhost:8000/ in your browser to check the app.
```