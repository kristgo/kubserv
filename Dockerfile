FROM scratch

ENV PORT 8080

EXPOSE $PORT

COPY kubserv /

CMD ["/kubserv"]
