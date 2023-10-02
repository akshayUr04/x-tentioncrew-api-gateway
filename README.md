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


## Follow these steps to Run and Test API using docker

create .env file on the project root dir and add the below envs on it
```.env
#example
PORT="8080"
M2PORT="50051"
USERPORT="50052"

USER_SERVICE_URL="userService:50052"
SERVICE2_URL="service2:50052"

DBHOST="db"
DBNAME="x-tentioncrew_users"
DBUSER="postgres"
DBPORT="5432"
DBPASSWORD="mypassword"

REDISPORT="6379"
REDISHOST="redis
```