#!/usr/bin/env bash

action(){

    showUsage(){
        local scriptFileBaseName
        scriptFileBaseName=$(basename "$0")
        echo "usage: $scriptFileBaseName "
        echo
        echo "       /?        Show usage."
        

        return 0;
    }

    main(){
        [[ $# -eq 0 ]] && showUsage && return 1;
        [[ "$*" =~ "/?" ]] && showUsage && return 0;

        return 0;
    }

    main "$@"
}

action "$@"
