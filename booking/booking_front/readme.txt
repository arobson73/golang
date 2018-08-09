npm i --save-dev awesome-typescript-loader babel-core babel-loader babel-preset-env css-loader file-loader source-map-loader style-loader ts-loader typescript url-loader webpack webpack-cli webpack-dev-server @types/react @types/react-dom @types/react-router-dom

npm i --save bootstrap jquery popper.js promise-polyfill react react-dom react-router-dom whatwg-fetch

#to dockerize this (make sure the backend stuff is running, see golang repo)
docker container build -t myevents/frontend .
#then
docker container run --name frontend -p 80:80 myevents/frontend	 
