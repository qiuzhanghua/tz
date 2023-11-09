default: build

all: build-linux build-darwin build-windows build-loong64

loong64: build-loong64

pre:
	autotag write

build:
	mkdir -p bin
	go build -o bin/tz

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o tz_`autotag current`_linux_arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tz_`autotag current`_linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o tz_`autotag current`_linux_386

build-loong64:
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o tz_`autotag current`_linux.mipsle
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o tz_`autotag current`_linux.mips64le
 	CGO_ENABLED=0 GOOS=linux GOARCH=loong64  go build -o tz_`autotag current`_linux_loong64.loong64

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o tz_`autotag current`_darwin_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o tz_`autotag current`_darwin_arm64

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o tz_`autotag current`_windows_amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o tz_`autotag current`_windows_386.exe

 clean:
	rm -rf tz_* bin/