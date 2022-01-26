
# BUILD STAGE
FROM golang:alpine AS builder

LABEL maintainer="kman@vmware.com"

WORKDIR  /main

COPY  . .

# RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main/main .




# RUN STAGE

FROM alpine:latest

WORKDIR  /root

COPY --from=builder /main .

EXPOSE 8080

CMD [ "./main" ]






# #  build stage
# FROM golang:latest AS builder

# LABEL maintainer="kman@vmware.com"

# WORKDIR  /app

# ADD . /app

# COPY . .

# # run stage
# FROM golang:latest

# WORKDIR /app
# COPY --from=builder /app .


# EXPOSE 8080

# CMD [ "/app/main" ]



