FROM scratch
ADD main /main
ENTRYPOINT ["/main"]
CMD ["127.0.0.1","8080"]