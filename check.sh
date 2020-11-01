
#set -e
for f in $(ls -l go/ | grep drwx | awk '{print $9}')
do
    pushd go/$f > /dev/null
    echo $f
    go fmt .
    golint
    go test
    mod=${f//[-]/_}

    exercism submit ${mod}.go
    popd > /dev/null
    echo
done

