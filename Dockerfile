FROM scratch

ENV PORT 8000

EXPOSE $PORT

COPY kubserv /

CMD ["/kubserv"]
