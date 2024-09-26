# build a Docker container for a single architecture
singlearch:
	/usr/local/bin/docker build . -t arm64-issue:latest

# build a Docker container for linux/amd64 and linux/arm64
multiarch:
	/usr/local/bin/docker buildx build --platform linux/amd64,linux/arm64 .

# build the application. Don't invoke this directly
buildapplication:
	CGO_ENABLED=1 GOOS=linux go build -o /ims-run /app/cmd/ims/main.go

#clean:
#	rm -f arm64-issue
#	docker rmi -f arm64-issue:latest