#!/usr/bin/env just --justfile

alias h := usage
alias help := usage
alias i := readme
alias info := readme
alias l := license
alias b := build

name := 'kami'
version := '1.0.0'

# shows the command list
default:
	@just --list --unsorted

# shows how to use phantom
usage:
	@go run main.go -h

# shows the readme file
readme:
	@cat README.md

# shows the license file
license:
	@cat LICENSE

# count non-empty lines of code
sloc:
    @find . -type f -name "*.go" -exec cat {} + | sed '/^\s*$/d' | wc -l

# needed to compile to windows from a unix system
_install_tools_debian:
    #!/usr/bin/env bash
    set -euo pipefail

    sudo apt-get install gcc-multilib
    sudo apt-get install gcc-mingw-w64

# runs the golang build
build label=(name + '-' + version) goos=os() arcitecture=arch():
    #!/usr/bin/env bash
    set -euo pipefail

    archh={{ arcitecture }}
    if [[ {{ arcitecture }} == x86* ]]; then
        archh="amd${archh:4}"
    fi

    bits="32"
    if [[ {{ arcitecture }} =~ "64" ]]; then
        bits="64"
    fi

    native="i686-w64-mingw${bits}-g"
    
    echo "Starting build for {{ goos }} ($archh)"

    env \
        `# set enviroment variables`
        GOOS={{ goos }} \
        GOARCH=$archh \
        CGO_ENABLED=1 \
        CXX="${native}++" \
        CC="${native}cc" \
    `# build the program`
    go build \
        -o {{ label }}-{{ goos }}-$archh
    
    echo 'Finished building'