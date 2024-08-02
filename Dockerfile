# 1. 先指定当前镜像的基础镜像是什么
FROM golang:1.14.2

# 2. 描述这个镜像的作者以及联系方式（可选）
MAINTAINER echin<197795131@qq.com>

# 3. 设置镜像的标签信息（可选）
LABEL version="1.0"
LABEL description="This is the first dockerfile"

# 4. 环境变量,可以设置多个
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 5. 在构建镜像时，需要执行的命令和脚本
RUN mkdir -p /go/src/app

# 6. 将主机中的指定文件拷贝到镜像中目标路径  ADD <src>... <dest>
ADD . /go/src/app
ADD ["./config", "/go/src/app/config"]

# 7. 指定工作目录,如果不存在会自动创建
WORKDIR /go/src/app

# 8. 镜像数据卷绑定，可以将主机的指定目录挂载到容器中
VOLUME ["/data"]

# 9. 容器对外暴露的端口，仅仅是容器内暴露的端口，而没有和主机端口进行关联
EXPOSE 8080

# 10. 容器启动时执行的命令
# CMD 和 ENTRYPOINT 选择一个，描述构建镜像完成后，启动容器时默认执行的脚本
# 两者区别： CMD 可以被 docker run 命令行参数替换，ENTRYPOINT 不会被替换
CMD ["go", "run", "main.go"]
ENTRYPOINT ["go", "run", "main.go"]

# ------------
# 1. 设置变量，在镜像中定义一个变量，可以在构建镜像时通过 --build-arg 参数传递
ARG app_env=production
# ARG golang_env
# FROM golang:$golang_env
# docker build --build-arg app_env=development -t myapp:dev .

# 2. 设置容器的用户，可以是用户名或者 UID
USER root

# 3. ONBUILD 指令，当构建一个被继承的 Dockerfile 时运行命令
ONBUILD ADD . /go/src/app

# 4. STOPSIGNAL 指令，设置容器停止时的信号
STOPSIGNAL SIGTERM

# 5. HEALTHCHECK 指令，设置容器健康检查的命令
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost/ || exit 1
