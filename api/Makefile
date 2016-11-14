#Backend 

dependencies:
	go get -v -t ./...

artifacts:
	mkdir artifacts \
	&& go build \
	&& mv api artifacts/

clean:
	rm -rf artifacts/

rebuild: clean artifacts
