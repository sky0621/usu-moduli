# usu-moduli
List the modules used by the project

## env

go version 1.9.4

## dep

https://github.com/golang/dep

## go run example

go run main.go template.go -d {target fullpath directory} > {output markdown filename}

## go-assets-builder

go get -v github.com/jessevdk/go-assets-builder

go-assets-builder template/ > template.go

## gox

https://github.com/mitchellh/gox

gox -os="linux darwin windows" -arch="amd64"

## ghr

https://github.com/tcnksm/ghr

git config --global github.token "....."

export GITHUB_API=http://github.company.com/api/v3/

ghr v0.1.0 pkg/
