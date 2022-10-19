#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------
FROM --platform=$TARGETPLATFORM golang:1.18 as builder

WORKDIR /workdir

RUN apt-get update && apt-get -y install git procps lsb-release  && \
    apt-get autoremove -y \
        && apt-get clean -y \
        && rm -rf /var/lib/apt/lists/*

ADD . .
RUN go env &&  ls -al && pwd
RUN cd src && ./all.bash && cd - && \
      ls -al


FROM --platform=$TARGETPLATFORM golang:1.18
#ENV GOPROXY https://goproxy.cn,direct


# Install git, process tools, lsb-release (common in install instructions for CLIs)
RUN apt-get update && apt-get -y install git procps lsb-release  && \
    apt-get autoremove -y \
        && apt-get clean -y \
        && rm -rf /var/lib/apt/lists/*
#RUN mkdir base && sed -i 's/dl-cdn.alpinelinux.org/opentuna.cn/g' /etc/apk/repositories && apk update --no-cache && \
#    apk add --no-cache ca-certificates tzdata  git procps  lsb-release build-base openssh curl bash &&  rm -rf /var/cache/apk/*

# Clean up
#RUN apt-get autoremove -y \
#    && apt-get clean -y \
#    && rm -rf /var/lib/apt/lists/*


#ENV GOROOT /go
COPY --from=builder /workdir $GOROOT
COPY --from=builder /workdir/bin $GOPATH/bin
COPY --from=builder /workdir/pkg $GOPATH/pkg
# Set the default shell to bash instead of sh

