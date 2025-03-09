docker build -t pardha1024/dta-producer:latest producer/

docker build -t pardha1024/dta-consumer1:latest consumer1/

docker build -t pardha1024/dta-consumer1db:latest consumer1db/

docker login

docker push pardha1024/dta-producer:latest

docker push pardha1024/dta-consumer1:latest

docker push pardha1024/dta-consumer1db:latest