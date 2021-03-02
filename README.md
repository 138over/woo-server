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

### Onboarding
We are always onboarding ourselves, team members, customers, partners, clients, vendors  

**FAASD**  
An early exploration in progress to make FaaS the center of the software ecosystem design. The first step will be to bootstrap a standalone ecosystem that includes a virtual machine orchestrator (vagrant), a hypervisor (vmware fusion), a virtual machine (ubuntu 20.010), a docker registry, and finally a faasd service


Getting familiar with go-task
```
cd internal/onboard/faasd
% task -h
% task -l
% task faasd:provision --summary
```

**CUE**  
The need for schemas and expressive configuration language are needed. In the past I have used [Hocon](https://github.com/lightbend/config) for the configuration language. I will expore using [Cue](https://cuelang.org) to define schemas and configurations for [infrastructure layers](doc/infrastructure-layers.png)
