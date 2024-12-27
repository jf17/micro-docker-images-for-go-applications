## 1. Clone repo:
   ```sh
   git clone https://github.com/jf17/micro-docker-images-for-go-applications
   ```
  ```sh
   cd micro-docker-images-for-go-applications
   ```


## 2. BUILD :
   ```sh
   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
   ```
   ```sh
   sudo docker build -t my-go-app .
   ```

## 3. RUN :
   ```sh
sudo docker run -p 3000:3000 my-go-app
   ```

## 4. OPEN browser :
 http://localhost:3000
