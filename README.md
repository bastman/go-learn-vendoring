# go-learn-vendoring
simple guide to vendoring in golang. (uses rubigo as dependency manager - which is quite similar to php composer)

## install rubigo dependency manager
- download from: https://github.com/yaa110/rubigo/releases

## A. build this project
    $ rubigo install
    $ rubigo update


## B. create a project from scratch
    $ rubigo init

        --> creates rubigo.json

### add a dependency
    $ rubigo get github.com/ahmetb/go-linq

        --> installs a package into projects vendor folder
        --> adds the dependency into rubigo.json
        --> creates/updates rubigo.lock file

### install / update dependencies - according to rubigo.json / rubigo.lock
    $ rubigo install
    $ rubigo update




