# install する docker image を指定
FROM golang:1.17.1

# current directory を指定
WORKDIR /go/src/app

ADD . /go/src/app

# OS(ubuntu のライブラリの更新)
# ホットリロード用のライブラリ
# 実行権限変更
RUN apt update \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

EXPOSE 8000