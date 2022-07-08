if [ -f api_linux_amd64.tar.gz ]
then
      tar -xzvf api_linux_amd64.tar.gz
fi

git pull
git reset --hard origin/master
go build cmd/api/main.go

chmod +x api_linux
pkill api_linux
nohup ./api_linux >> log.output &