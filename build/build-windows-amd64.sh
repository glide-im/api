cd ../cmd/api || exit

export CGO_ENABLED=0
export GOOS=windows
export GOHOSTOS=windows
export GOARCH=amd64

echo 'build...'
go build -o api.exe
echo 'build complete'
cp ../../config/config.toml config.toml
tar -czvf ./api_windows_amd64.tar.gz api.exe config.toml
rm config.toml
rm api
read -p 'complete.'