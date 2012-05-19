# gowork #

gowork is a tool for creating Go workspaces. It also provides a script for easily shelling into and out of them, much like environments created by Python `virtualenv`, Perl `local::lib`, Ruby `rip` or `bundler`, etc. (In fact, the `activate` script it generates is cribbed mostly from virtualenv for Python.)

## Usage ##

    $ gowork myenv
    $ source myenv/bin/activate
    (myenv)$ go get github.com/nickoneill/go-dropbox  # or whatever
    (myenv)$ go list github.com/nickoneill/go-dropbox
    github.com/nickoneill/go-dropbox
    (myenv)$ deactivate
    $ $ go list github.com/nickoneill/go-dropbox
    can't load package: package github.com/nickoneill/go-dropbox: import
    "github.com/nickoneill/go-dropbox": cannot find package

## Requirements ##

* Go 1

## Installation ##

1. Run `go get github.com/markpasc/gowork` to install the `gowork` tool into your Go tree or current workspace.
