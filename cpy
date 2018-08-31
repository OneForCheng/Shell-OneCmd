#!/usr/bin/env bash

cpy(){

    showUsage(){
        local scriptFileBaseName
        scriptFileBaseName=$(basename "$0")
        echo "usage: $scriptFileBaseName source... [directory]"
        echo
        echo "       /?        Show usage."
        echo
        echo "       It will try use 'OLDPWD' variable if directory is not exist,"
        echo "       but it needs '.' or 'source' to run this script."
        echo
        return 0;
    }

    copyFileToTargetPath(){
        [[ ! -a "$1" ]] && echo No such file or directory: "$1" && return 1;
        cp -af "$1" "$2"
        return 0;
    }

    copyFilesToTargetPath(){
        local i
        local file
        local args

        args=( "$@" )
        if [[ -d "${args[$#]}" && $# -ne 1 ]]; then
            for ((i=1; i<$#; i++)); do
              copyFileToTargetPath "${args[$i]}" "${args[$#]}"
            done
        elif [[ -d "$OLDPWD" ]]; then
            for file in "$@"; do
              copyFileToTargetPath "$file" "$OLDPWD"
            done
        else
            echo No such directory: "${args[$#]}"
        fi

        return 0;
    }

    main(){
        [[ $# -eq 0 ]] && showUsage && return 1;
        [[ "$*" == *"/?"* ]] && showUsage && return 0;

        copyFilesToTargetPath "$@"

        return 0;
    }

    main "$@"
}

cpy "$@"
