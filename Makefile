#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

SOURCEDIR=.

SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
BINARY=myprogram

# LDFLAG values
VERSION=latest
BUILD=`git rev-parse HEAD`
BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`

# TODO: automate other flags
# LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.GitCommit=${BUILD} -X main.BuildDate=${BUILD_DATE} -X main.Build=`git rev-parse HEAD` "
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Binary=${BINARY}"

# Build the project
build: clean
	go build ${LDFLAGS} -o ${BINARY}

# Run all tests
test:
	@echo "Testing"
	go test ./...

# Run the unit tests
unit_tests:
	@echo "Testing"
	go test ./... -tags=unit

# Run the integration tests
integration_test:
	@echo "Launch the integration tests."
	go test -v ./... -tags=integration

format:
	@echo "Formatting"
	go fmt ./...

lint: format
	@echo "Linting"
	golint .

install:
	go install

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi

.PHONY: clean install build deps updatedeps format lint test integration_test
