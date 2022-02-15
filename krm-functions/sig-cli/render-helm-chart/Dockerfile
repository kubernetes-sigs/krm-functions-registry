# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.17-alpine3.13
ENV CGO_ENABLED=0
WORKDIR /go/src/

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .
RUN go build -o /usr/local/bin/function ./

RUN apk update && apk add --no-cache curl
ARG HELM_VERSION="v3.8.0"
RUN curl -fsSL -o /helm-${HELM_VERSION}-linux-amd64.tar.gz https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
     tar -zxvf /helm-${HELM_VERSION}-linux-amd64.tar.gz && \
     mv ./linux-amd64/helm /usr/local/bin/helm

#############################################

FROM alpine:3.13
COPY --from=0 /usr/local/bin/function /usr/local/bin/function
COPY --from=0 /usr/local/bin/helm /usr/local/bin/helm

ENV PATH /usr/local/bin:$PATH
ENV HELM_PATH_CACHE /var/cache
ENV HELM_CONFIG_HOME /tmp/helm/config
ENV HELM_CACHE_HOME /tmp/helm/cache

ENTRYPOINT ["function"]