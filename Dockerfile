FROM golang:1.12.5-stretch
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt update && apt upgrade -y && apt install -y vim git gcc build-essential libffi-dev
WORKDIR /app
COPY . .
