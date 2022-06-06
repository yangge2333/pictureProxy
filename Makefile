all:
	mkdir -p ./bin
	go build -o bin/pic-proxy main.go

install:
	/bin/cp -f bin/pic-proxy /usr/bin/pic-proxy
	/bin/cp -f ./pic-proxy.service /lib/systemd/system/
	mkdir  -p  /usr/bin/template
	mkdir  -p  /usr/bin/config
	mkdir  -p  /usr/bin/public
	/bin/cp -r ./template/. /usr/bin/template
	/bin/cp -f ./config/config.toml /usr/bin/config
	systemctl daemon-reload
	systemctl enable pic-proxy.service; systemctl restart pic-proxy.service ; systemctl status pic-proxy.service

update:
	/bin/cp -f bin/pic-proxy /usr/bin/pic-proxy
	/bin/cp -f ./pic-proxy.service /lib/systemd/system/
	/bin/cp -f ./config/config.toml /usr/bin/config
	mkdir  -p  /usr/bin/template
	mkdir  -p  /usr/bin/config
	mkdir  -p  /usr/bin/public
	rm -rf /usr/bin/template/*
	/bin/cp -r ./template/. /usr/bin/template
	systemctl daemon-reload
	systemctl enable pic-proxy.service; systemctl restart pic-proxy.service ; systemctl status pic-proxy.service

clean:
	systemctl disable pic-proxy.service; systemctl stop pic-proxy.service
	rm -f /lib/systemd/system/pic-proxy.service
	rm -rf ./bin
	rm -f /usr/bin/pic-proxy
	systemctl daemon-reload
