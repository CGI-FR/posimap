FROM gcr.io/distroless/base
ARG BIN
COPY /bin/posch /posch
CMD ["/posch"]
