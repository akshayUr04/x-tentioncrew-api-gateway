## Follow these steps to Run and Test API Gateway Service locally

#### Clone The Repository of API Gateway

git clone https://github.com/akshayUr04/x-tentioncrew-api-gateway.git

#### Change Directory To API Gateway

cd ./x-tentioncrew-api-gateway


#### Install tThe dependencies

go mod tidy 


#### Setup The Env
.env
PORT="port that you wan't to run API Gateway"

USER_SERVICE_URL="PortURL that the User Service running"

SERVICE2_URL="PortURl that the Service2 running"
#### Run The Application

go run ./cmd/main.go


##   Run the command to run the API using docker

docker compose up