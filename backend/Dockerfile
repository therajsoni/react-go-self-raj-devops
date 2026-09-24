from golang:latest as builder 
label name="user management backend app"
workdir /app 
copy go.mod go.sum ./ 
run go mod download 
copy . .
run go build -o backend . 

from alpine:latest
workdir /app 
copy --from=builder --chown=appUser:appUser /app/backend ./
run adduser -D appUser
user appUser
expose 8080 
cmd ["./backend"]
