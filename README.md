Software Delivery Engine Explorations
---
### MacOS Setup
For these exporations, I use [go-task](https://taskfile.dev/#/) to gather exploration steps and reproduce them,
and [multipass](https://multipass.run) to manage virtual machines

**install brew**
```
% /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

**install golang**
```
% brew install go
```

**install go-task**
```
% brew install go-task/tap/go-task
```

**install multipass**
```
https://multipass.run provides multiple installation methods
```

### Onboarding
We are always onboarding ourselves, team members, customers, partners, clients, vendors  

**FAASD**  
An early work in progress to make FaaS the center of the software ecosystem design. the first step will be to bootstrap a standalone ecosystem that includes a virtual machine orchestrator (multipass), a virtual machine, a docker registry, and finally a faasd service

Some of the first steps have been implemented, but not validated. 

To get a feel of how the multipass utility is being used 
```
% multipass --help
% multipass list
% multipass launch --name foo
% multipass list
% multipass exec foo -- lsb_release -a
% multipass stop foo
% multipass list
% multipass delete foo
% multipass list
% multipass purge
```

Getting familiar with go-task
```
cd internal/onboard/faasd
% task -h
% task -l
% task dev:vm:launch --summary
% task dev:vm:launch 
% task vm:list
% task dev:vm:version
% task dev:vm:ipaddr
% task dev:vm:packages
% task dev:docker:install --summary
% task dev:docker:install 
% task dev:vm:stop
% task dev:vm:delete
% task vm:purge --summary
% task vm:purge 
```

**CUE**  
The need for schemas and expressive configuration language are needed. In the past I have used [Hocon](https://github.com/lightbend/config) for the configuration language. I will expore using [Cue](https://cuelang.org) to define schemas and configurations for [infrastructure layers](doc/infrastructure-layers.png)
