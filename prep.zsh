#! /bin/zsh

[[ -f base.go ]] && rm base.go
while read line; do
    case $line in
        (__BULMA__)
            cat assets/bulma.min.css >> base.go
            print >> base.go
            ;;

        (__FA__)
            cat assets/fa.min.js >> base.go
            print >> base.go
            ;;

        (*)
            print $line >> base.go
    esac
done < base.go.in
