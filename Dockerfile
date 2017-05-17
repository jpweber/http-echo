FROM golang
COPY http-echo /bin/http-echo
ENTRYPOINT ["/bin/http-echo"]
CMD ["-h"]
