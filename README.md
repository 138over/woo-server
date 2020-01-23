### Workspace Orchestration Services
```
install brew https://brew.sh

% brew install go
% brew install go-task/tap/go-task
% task -l
task: Available tasks for this project:
* build:                go build main.go
* lint:                 golint $(go list ./...)
* run:svc:web:          go run main.go service --name web
* run:svc:workspace:    go run main.go service --name workspace
* test:                 run unit, functional and system test
* test:functional:      functional test 
* test:system:          system test
* test:unit:            unit test
* utils:                download and install go utilities
```

### Experimentation with FaaS
```
Drive service configuration from external source, where handlers can live anywhere. 
Within the code, as a web service, or as a function (FaaS)

Handlers are candidates to live anywhere, whether that be state, error handling, or 
any non-DOM tasks
```

### Experimentation with inline DAG/Event configuration
```
% task build run:svc:web

open browser http://127.0.0.1:3000
open developer tools to observe DAG generation in the console log

# an aribtary DAG

[
    {
        id: "start:flow",
        description: "Start flow button was clicked by user",
        eventType: "dom",
        onEvent: "click",
        selector: "#start-flow",
    },
    {
        id: "configure:flow",
        description: "Configure flow",
        eventType: "app",
        parentIds: ["start:flow"],
        handler: "configureFlow",
        params: { variant: "release" }
    },
    {
        id: "create:flow",
        description: "Create flow",
        eventType: "app",
        parentIds: ["configure:flow"],
        handler: "createFlow",
        params: { variant: "release" }
    },
    {
        id: "configure:report",
        description: "Report flow",
        eventType: "app",
        parentIds: ["configure:flow"],
        handler: "configureReport",
        params: { variant: "release" }
    },
    {
        id: "publish:configure",
        description: "Publish Flow Configuration",
        eventType: "app",
        parentIds: ["configure:report"],
        handler: "publishFlowConfiguration",
        params: { variant: "release" }
    },
    {
        id: "run:flow",
        description: "Run flow",
        eventType: "app",
        parentIds: ["create:flow"],
        handler: "runFlow",
        params: { variant: "release" }
    },
    {
        id: "completed:flow",
        description: "Completed flow",
        eventType: "app",
        parentIds: ["run:flow"],
        handler: "completeFlow",
        params: { variant: "release" }
    },
    {
        id: "publish:flow",
        description: "Publish flow",
        eventType: "app",
        parentIds: ["completed:flow"],
        handler: "publishFlow",
        params: { variant: "release" } 
    }   
]
```

### Experimentation with external DAG/Event configuration
```
task build run:svc:web
open browser http://127.0.0.1:3000/lifecycle

TODO: implement external integration
```

### Experimentation with DAGs and Svelte
[D3-DAG](https://github.com/erikbrinkman/d3-dag)  
[Svelte D3 Example](https://svelte.dev/repl/01a5774b53e9416584428c025668407b?version=3.15.0)  
```
TODO: implement or d3.js DAG visualiation
```

