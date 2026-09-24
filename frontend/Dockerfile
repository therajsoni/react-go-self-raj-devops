from node:latest as builder 
label name="user management app"
workdir /app 
copy package*.json ./ 
run npm install 
copy . .
run npm run build 

from nginx:alpine 
workdir /app 
copy --from=builder /app/dist /usr/share/nginx/html/ 
cmd ["nginx" , "-g" , "daemon off;"]
