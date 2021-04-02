
#!/bin/bash

version=`git log --date=iso --pretty=format:"%cd @%h" -1`
if [ $? -ne 0 ]; then
    version="not a git repo"
fi

compile=`date +"%F %T %z"`" by "`go version`
author= "liangdu"
email="liangdu1992@gmail.com"
cat << EOF | gofmt > version.go
package version

const (
    Version = "$version"
    Compile = "$compile"
    Author= "$author"
    Email="$email"
)
EOF
