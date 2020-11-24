FROM golang:1.15

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN make
EXPOSE 8080

CMD [ "/app/bin/connection_test" ]