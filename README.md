Software Delivery Engine Explorations
---
### MacOS Setup
For these exporations, I use [go-task](https://taskfile.dev/#/) to gather exploration steps and run them repeatabily
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

**faasd**  
is an early work in progress. the first step will be to bootstrap a standalone ecosystem that includes a virtual machine orchestrator (multipass), a virtual machine, a docker registry, and finally a faasd service

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
% task vm:launch --summary
% task vm:launch
% task vm:list
% task docker:install --summary
% task docker:install 
% task docker-compose:install
% task docker:hello-world
```
