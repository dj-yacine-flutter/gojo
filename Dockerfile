#########
# Build #
#########

FROM golang:1.21.5-alpine3.19 AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# Create an unprivileged user
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /gojo
COPY . .

# Use go mod
RUN go mod download -x
RUN go mod verify

# The following line compiles your Go application with specific build flags.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -installsuffix nocgo \
	-o main .
    
RUN chmod +x main


#########
#  Run  #
#########

FROM alpine:3.19
WORKDIR /gojo


# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the compiled binary from the builder stage to the final image.
COPY --from=builder /gojo/main .

# Copy the environment file.
COPY gojo.env .

# Use an unprivileged user.
USER appuser:appuser

# Specify the command to run your application when the container starts.
CMD [ "./main" ]