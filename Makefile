all: build

build:
	echo "Building gofred for all platforms"
	GOOS=darwin GOARCH=amd64 go build -o ./dist/gofred-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/gofred-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -o ./dist/gofred-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o ./dist/gofred-linux-arm64
	GOOS=windows GOARCH=amd64 go build -o ./dist/gofred-windows-amd64.exe
	GOOS=windows GOARCH=arm64 go build -o ./dist/gofred-windows-arm64.exe

	cd ./dist && tar -czvf gofred-darwin-amd64.tar.gz gofred-darwin-amd64
	cd ./dist && tar -czvf gofred-darwin-arm64.tar.gz gofred-darwin-arm64
	cd ./dist && tar -czvf gofred-linux-amd64.tar.gz gofred-linux-amd64
	cd ./dist && tar -czvf gofred-linux-arm64.tar.gz gofred-linux-arm64
	cd ./dist && tar -czvf gofred-windows-amd64.tar.gz gofred-windows-amd64.exe
	cd ./dist && tar -czvf gofred-windows-arm64.tar.gz gofred-windows-arm64.exe

	rm ./dist/gofred-darwin-amd64 ./dist/gofred-darwin-arm64 ./dist/gofred-linux-amd64 ./dist/gofred-linux-arm64 ./dist/gofred-windows-amd64.exe ./dist/gofred-windows-arm64.exe

clean:
	rm -rf ./dist