from node:latest as builder 
label name="user management app"
workdir /app 
copy package*.json ./ 
run npm install 
copy . .
run npm run build 

run ls -la

from nginx:alpine 
workdir /app 
run ls -la
copy /app/dist /usr/share/nginx/html/ 
run adduser -D rajsoni
user rajsoni
cmd ["nginx" , "-g" , "deamon off;"]
