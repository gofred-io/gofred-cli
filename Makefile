all: build

build:
	echo "Building gofred for all platforms"
	GOOS=darwin GOARCH=amd64 go build -o ./dist/gofred-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/gofred-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -o ./dist/gofred-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o ./dist/gofred-linux-arm64
	GOOS=windows GOARCH=amd64 go build -o ./dist/gofred-windows-amd64.exe
	GOOS=windows GOARCH=arm64 go build -o ./dist/gofred-windows-arm64.exe

clean:
	rm -rf ./dist