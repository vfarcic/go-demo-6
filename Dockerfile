FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-demo-6_linux"]
COPY ./bin/ /
