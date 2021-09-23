
# testgolang

application deployed in heroku

##Test it online in heroku
https://youtochigolanglogin.herokuapp.com/v1/multiplica/2/15


## Running Locally

git clone https://github.com/czendee/testgolang.git

cd testgolang

go build

go run main.go
     
      test in port 8090
      http://localhost:8090/v1/multiplica/4/1
      http://localhost:8090/v1/multiplica/0/0
      http://localhost:8090/v1/multiplica/1/1
      http://localhost:8090/v1/multiplica/4/tt
      http://localhost:8090/v1/multiplica/tt/1
      http://localhost:8090/v1/multiplica/tt/tt

go test

## Deploying to Heroku

automatic deploy with commit in the github repository

Try it:

   https://youtochigolanglogin.herokuapp.com/v1/multiplica/2/15

https://youtochigolanglogin.herokuapp.com/v1/multiplica/1/1


https://youtochigolanglogin.herokuapp.com/v1/multiplica/0/0

https://youtochigolanglogin.herokuapp.com/v1/multiplica/NAT/15

https://youtochigolanglogin.herokuapp.com/v1/multiplica/2/NAT

https://youtochigolanglogin.herokuapp.com/v1/multiplica/NAT/NAT

https://youtochigolanglogin.herokuapp.com/v1/multiplica/2/0

https://youtochigolanglogin.herokuapp.com/v1/multiplica/0/2
## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
