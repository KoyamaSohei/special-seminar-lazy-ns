FROM golang
WORKDIR /go/src/github.com/KoyamaSohei/special-seminar-lazy-ns
ENV GO111MODULE=on
COPY . .
RUN go build
CMD ./special-seminar-lazy-ns
