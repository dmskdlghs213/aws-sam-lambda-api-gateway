# aws-sam-lambda-api-gateway

## build

```
sam build
```

## local test

```
sam local start-api
```

## deploy

```
sam deploy
```


## stg/prd test

```
curl -X GET "https://yyu9zwuf9b.execute-api.ap-northeast-1.amazonaws.com/Stage/hello?name=masa"
```

```
Hello, take
```