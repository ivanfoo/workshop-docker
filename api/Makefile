
SHELL = /bin/bash
DEPENDENCIES = github.com/gin-gonic/gin gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson

dependencies:
	for package in $(DEPENDENCIES); do \
		go get -v $${package}; \
	done; 

artifacts:
	@mkdir artifacts \
	&& go build \
	&& mv api artifacts/

clean:
	@rm -rf artifacts/

rebuild: clean artifacts
