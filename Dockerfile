FROM golang AS build-env
ADD . /src
RUN cd /src && make

FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-demo-6"]
COPY --from=build-env /src/bin/go-demo-6 /

