# Build
##ARM64
```
env GOOS=linux GOARCH=arm64 go build -o iaa-tlc-arm64 cmd/traffic-light-control-unit/main.go
```
##Linux
```
env GOOS=linux go build -o iaa-tlc-linux cmd/traffic-light-control-unit/main.go
```
##Windows
```
env GOOS=windows go build -o iaa-tlc-win cmd/traffic-light-control-unit/main.go
```