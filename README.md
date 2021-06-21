Software Delivery Engine Explorations
---
### Onboarding
We are always onboarding ourselves, team members, customers, partners, clients, vendors  

### MacOS Setup
For these exporations, I use [go-task](https://taskfile.dev/#/) to gather exploration steps and reproduce them,
and virtualbox, vmware and vagrant to manage virtual machines

**Install brew**
```
% /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

**Install golang**
```
% brew install go
```

**Install go-task**
```
% brew install go-task/tap/go-task
```

[Getting familiar with go-task](https://taskfile.dev/#/)  
```
% task -h  
% task -l
```

### Configuration Formats

**CUE**  
The need for schemas and expressive configuration language are needed. In the past I have used [Hocon](https://github.com/lightbend/config) for the configuration language. I will explore using [Cue](https://cuelang.org) to define schemas and configurations for [infrastructure layers](doc/infrastructure-layers.png)
