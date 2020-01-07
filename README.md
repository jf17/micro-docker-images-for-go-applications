## BUILD :
### CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
### sudo docker build -t my-go-app .

## RUN :
### sudo docker run -p 3000:3000 my-go-app

## OPEN browser :
### localhost:3000
