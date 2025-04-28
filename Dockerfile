FROM gcr.io/distroless/base
ARG BIN
COPY /bin/posimap /posimap
CMD ["/posimap"]
