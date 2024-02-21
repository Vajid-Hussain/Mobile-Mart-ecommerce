FROM golang:1.22-alpine AS firstStage
WORKDIR /MobileMart
COPY . /MobileMart/
RUN go mod download
RUN go build -o ./cmd/api/binary ./cmd/api/main.go

FROM scratch
COPY --from=firstStage /MobileMart/cmd/api/binary /MobileMart/
COPY --from=firstStage /MobileMart/.env /MobileMart/
COPY --from=firstStage /MobileMart/template /MobileMart/template/
WORKDIR /MobileMart
CMD [ "/MobileMart/binary" ]