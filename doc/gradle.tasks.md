Model Driven Gradle Example
---
Configuration is defined externally in JSON or YAML
```

task('download-tomcat', dependsOn: model.downloadTomcat.dependsOn, type:Download) {
    ext.config  = model.downloadTomcat
    description   config.description
    group         config.group
    src           config.src
    dest          config.dest
    overwrite     config.overwrite
}

task('install-tomcat', dependsOn: model.installTomcat.dependsOn, type:Copy) {
    ext.config  = model.installTomcat
    description   config.description
    group         config.group
    from          tarTree(config.from)
    into          config.into
}
    
task('download-war', dependsOn: model.downloadWar.dependsOn, type: Download) {
    ext.config  = model.downloadWar
    description   config.description
    group         config.group
    src           config.src
    dest          config.dest
    overwrite     config.overwrite
}

task('install-war', dependsOn: model.installWar.dependsOn, type: Copy) {
    ext.config  = model.installWar
    description   config.description
    group         config.group
    from          config.from
    into          config.into
}

task('download-config', dependsOn: model.downloadConfig.dependsOn, type: Download) {
    ext.config  = model.downloadConfig
    description   config.description
    group         config.group
    src           config.src
    dest          config.dest
    overwrite     config.overwrite
}

task('unpack-config', dependsOn: model.unpackConfig.dependsOn, type: Copy) {
    ext.config  = model.unpackConfig
    description   config.description
    group         config.group
    from          zipTree(config.from)
    into          config.into
}

task('install-config', dependsOn: model.installConfig.dependsOn) {
    ext.config   = model.installConfig
    description   config.description
    group         config.group
    from          config.from
    into          config.into
    filter(ReplaceTokens, tokens: config.tokens)
}

```


