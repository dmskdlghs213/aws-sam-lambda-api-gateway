# aws-sam-lambda-api-gateway

## build

```
make build
```

## local test

```
make start
```

## deploy

```
make deploy
```


## stg/prd test

```
curl -X GET "https://yyu9zwuf9b.execute-api.ap-northeast-1.amazonaws.com/Stage/hello?name=masa"
```

```
Hello, take
```