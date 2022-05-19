FROM golang:1.17

WORKDIR /app

COPY . .

RUN go build -o GithubStatsGenerator

CMD ./GithubStatsGenerator