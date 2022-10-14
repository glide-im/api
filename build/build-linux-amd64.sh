cd ../cmd/api || exit

export CGO_ENABLED=0
export GOOS=linux
export GOHOSTOS=linux
export GOARCH=amd64

echo 'build...'
go build -o api_linux
echo 'build complete'
cp ../../config/config.toml config.toml
tar -czvf ./api_linux_amd64.tar.gz api_linux config.toml
rm config.toml
rm api_linux
read -p 'complete.'