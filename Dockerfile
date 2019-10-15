FROM ccr.ccs.tencentyun.com/66super/go_base:1.12.7-v20191011

RUN mkdir -p /go/src/stroage_api
WORKDIR /go/src/stroage_api
COPY . .
CMD bee run
