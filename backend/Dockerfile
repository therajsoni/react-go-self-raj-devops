from golang:latest as builder 
label name="user management backend app"
workdir /app 
copy go.mod go.sum ./ 
run go mod download 
copy . .
run  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -o backend . 

from alpine:latest
workdir /app 
copy --from=builder  /app/backend ./backend
expose 8080 
cmd ["./backend"]
