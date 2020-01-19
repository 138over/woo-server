### Workspace Orchestration Services
```
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

### Experimentation with FaaS
```
Drive service configuration from external source, where handlers can live anywhere. 
Within the code, as a web service, or as a function (FaaS)

Handlers are candidates to live anywhere, whether that be state, error handling, or 
any non-DOM tasks
```

### Experimentation with Inline event configuration
```
task build run:web:service
open browser http://127.0.0.1:3000
open developer tools to observe console log
click Start Lifecycle

class ServiceStructure {
    constructor(route) {
        this._route = route;
    }

    static configure(route) {
        Lifecycle.createProcess('run-arbitrary-flow', ServiceLifecycleHandler)
            .addTask({
                description     : 'User clicked start flow button',
                eventType       : 'dom',
                onEvent         : 'click',
                selector        : '#arbitrary-flow',
                triggerEvent    : 'start:flow' 
            })
            .addTask({ 
                description     : "Display flow configuration form",
                eventType       : 'app',
                onEvent         : 'start:flow',
                handler         : 'formVisible',
                params          : { selector: '.arbitrary-flow-form' },
                publishEvent    : 'form:visible',
                mode            : 'debug'
            });
    
        Lifecycle.createProcess('run-foo-flow', ServiceLifecycleHandler)
            .addTask({ 
                description     : "Display flow configuration form",
                eventType       : 'app',
                onEvent         : 'start:flow',
                handler         : 'formVisible',
                params          : { selector: '.foo-form' },
                publishEvent    : 'form:visible' 
            });

        return new ServiceStructure(route)
    }
}

```

### Experimentation with External event configuration
```
task build run:web:service
open browser http://127.0.0.1:3000/lifecycle

TODO: implement method to read configuration from http://127.0.0.1:3000/lifecycle
      and dynamically create the processes for each name lifecycle
```

