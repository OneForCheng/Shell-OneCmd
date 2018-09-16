#!/usr/bin/env bash

go(){
    local dataDictFile
    
    showUsage(){
        local scriptFileBaseName
        scriptFileBaseName=$(basename "$0")
        echo "usage: $scriptFileBaseName key"
        echo "       $scriptFileBaseName key value"
        echo "       $scriptFileBaseName -r key..."
        echo
        echo "       /?        Show usage."
        echo "       -r        Remove key-value."
        echo "       key       It's the 'key' of key-value."
        echo "       value     It's the 'value' of key-value."
        echo 
        echo "       It will try to use 'open' command to execute the 'value' of key-value."
        echo 

        return 0;
    }

    getScriptDirectory(){
        cd "$(dirname "$0")" && pwd
        return 0;
    }

    showAllKeyValues(){
        cat "$dataDictFile"
        return 0;        
    }

    openValue(){
        local value
        value=$(grep -E "^$1" "$dataDictFile" | head -1)
        if [[ "$value" == "" ]]; then
            open ./*"$1"*
        else
            open "${value#*=}"
        fi

        return 0;
    }

    setKeyValue(){
        local line
        local tempFile
        tempFile=$(mktemp)
        
        echo "$1"="$2">"$tempFile"
        while read -r line
        do
            if [[ "${line%%=*}" != "$1" ]]; then
                echo "$line">>"$tempFile"                
            fi
        done < "$dataDictFile"

        sort < "$tempFile" >"$dataDictFile"

        return 0;
    }

    removeKeyValue(){
        local line
        local tempFile
        local isFound
        
        isFound="false"
        tempFile=$(mktemp)

        while read -r line; do
            if [[ "${line%%=*}" == "$1" ]]; then
                isFound="true"
            else
                echo "$line">>"$tempFile"
            fi
        done < "$dataDictFile"

        if [[ "$isFound" == "true" ]]; then
            cp -af "$tempFile" "$dataDictFile"
        else
            echo No such key: "$1"
        fi

        return 0;
    }

    removeKeyValues(){
        local key
        for key in "$@"; do
            removeKeyValue "$key"
        done
        return 0;
    }

    init(){
        dataDictFile=$(getScriptDirectory)/$(basename "$0").dd
        [[ ! -f "$dataDictFile" ]] && touch "$dataDictFile"
        return 0;
    }

    main(){
        init

        [[ $# -eq 0 ]] && showAllKeyValues && return 0;
        [[ "$*" == *"/?"* ]] && showUsage && return 0;
        
        if [[ $# -eq 1 ]]; then
            openValue "$1"
        elif [[ "$1" == "-r" ]]; then
            removeKeyValues "${@:2}"
        elif [[ $# -eq 2 ]]; then
            setKeyValue "$1" "$2"
        else
            showUsage
            return 1; 
        fi

        return 0;
    }

    main "$@"
}

go "$@"
