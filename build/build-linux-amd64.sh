cd ../cmd/api || exit

export CGO_ENABLED=0
export GOOS=linux
export GOHOSTOS=linux
export GOARCH=amd64

echo 'build...'
go build
echo 'build complete'
cp ../../config/config.toml config.toml
tar -czvf ./api_linux_amd64.tar.gz api config.toml
rm config.toml
rm api
read -p 'complete.'