# Copyright 2024 Google LLC

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     https://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

img = "us-east4-docker.pkg.dev/nam-sdbx/ontrac/redis-hc-sidecar"
version = "1.0.4"
gcpProject = "nam-sdbx"

image: build-image push-image
remote-image: remote-build-image push-image

.PHONY: server
server:
	@go build -o server

.PHONY: build-image
build-image:
	@echo "Building image: $(img):$(version)"
	@docker build -t $(img):$(version) .

.PHONY: remote-build-image
remote-build-image:
	@echo "Building image: $(img):$(version)"
	@gcloud builds submit --tag $(img):$(version) --project $(gcpProject)

.PHONY: push-image
push-image:
	@echo "Pushing image: $(img):$(version)"
	@docker push $(img):$(version)

.PHONY: test-hc
test-hc: server
	@echo "Starting Redis Server"
	@redis-server ./examples/redis.conf & REDIS_PWD=testpwd ./server