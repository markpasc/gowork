package main

import (
    "fmt"
    "os"
    "path"
    "strings"
)

func main() {
    // Where is the workspace?
    cwd, error := os.Getwd()
    if error != nil {
        fmt.Println(error)
        os.Exit(40)
    }

    localPath := "."
    if len(os.Args) > 1 {
        localPath = os.Args[1]
    }

    workspace := path.Join(cwd, localPath)

    for _, dirname := range []string{"bin", "pkg", "src"} {
        error = os.MkdirAll(path.Join(workspace, dirname), os.FileMode(0755))
        if error != nil {
            fmt.Println(error)
            os.Exit(40)
        }
    }

    activate := `
# This file must be used with "source bin/activate" *from bash*
# you cannot run it directly

deactivate () {
    if [ -n "$_OLD_PS1" ]; then
        PS1="$_OLD_PS1"
        export PS1
        unset _OLD_PS1
    fi

    if [ -n "$_OLD_PATH" ]; then
        PATH="$_OLD_PATH"
        export PATH
        unset _OLD_PATH
    fi

    if [ ! "$1" = "nondestructive" ]; then
        GOPATH="$_OLD_GOPATH"
        export GOPATH
        unset _OLD_GOPATH
    fi

    if [ -n "$BASH" -o -n "$ZSH_VERSION" ]; then
        hash -r
    fi

    unset GOWORKSPACE
    if [ ! "$1" = "nondestructive" ]; then
        unset -f deactivate
    fi
}

# unset irrelevant variables
deactivate nondestructive

GOWORKSPACE="%s"
export GOWORKSPACE

_OLD_PS1="$PS1"
PS1="(` + "`" + `basename "$GOWORKSPACE"` + "`" + `)$PS1"
export PS1

_OLD_PATH="$PATH"
PATH="$PATH:$GOWORKSPACE/bin"
export PATH

_OLD_GOPATH="$GOPATH"
if [ -n "$GOPATH" ]; then
    GOPATH="$GOWORKSPACE:$GOPATH"
else
    GOPATH="$GOWORKSPACE"
fi
export GOPATH

if [ -n "$BASH" -o -n "$ZSH_VERSION" ]; then
    hash -r
fi

    `
    activate = strings.TrimSpace(activate)
    activate = fmt.Sprintf(activate, workspace)

    // WRITE IT OUT
    f, error := os.Create(path.Join(workspace, "bin", "activate"))
    if error != nil {
        fmt.Println(error)
        os.Exit(40)
    }

    _, error = f.WriteString(activate)
    if error != nil {
        fmt.Println(error)
        os.Exit(40)
    }

    f.Close()
}
