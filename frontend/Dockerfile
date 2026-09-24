from node:latest as builder 
label name="user management app"
workdir /app 
copy package*.json ./ 
run npm install 
copy . .
run npm run build 

from nginx:alpine 
workdir /app 
copy /dist /usr/share/nginx/html/ 
run adduser -D rajsoni
user rajsoni
cmd ["nginx" , "-g" , "deamon off;"]
