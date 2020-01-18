### Workspace Orchestration Services
```
Optional
    brew install go-task/tap/go-task
    task -l
    task: Available tasks for this project:
    * build:                        go build main.go
    * lint:                         golint $(go list ./...)
    * run:web:service:              go run main.go service --name web
    * run:workspace:service:        go run main.go service --name workspace
    * test:                         run unit, functional and system test
    * test:functional:              functional test
    * test:system:                  system test
    * test:unit:                    unit test
    * utils:                        download and install go utilities
```

