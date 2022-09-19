#! /bin/bash

#default ENV is dev
env=dev

while test $# -gt 0; do
    case "$1" in 
        -env)
            shift
            if test $# -gt 0; then
                env=$
            fi
            #shift
            ;;
        *)
        break
        ;;
    esac
done

source .env
go build -o cmd/blog/blog cmd/blog/main.go
cmd/blog/blog -env $env &