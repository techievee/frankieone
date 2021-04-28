## FrankieOne Test SDK REST API Service written in Golang using GIN Framework
> Golang 1.16  | GoGin Framework  | ZAP Logger | Docker | CircleCI

[![<techievee>](https://circleci.com/gh/techievee/frankieone.svg?style=svg)](<https://circleci.com/gh/techievee/frankieone>)

[![codecov](https://codecov.io/gh/techievee/frankieone/branch/master/graph/badge.svg?token=3ASLAUCRFD)](https://codecov.io/gh/techievee/frankieone)

## Run prebuilt image
To Run the pre-built docker image, that automatically builds with every new push
```
docker run -p 8081:8081 --name frankieone techievee/frankieone:latest
```
The application exposes 
*  8081 - <default>For TLS port ( TLS Enabled by default from the config with self-signed certificate)

## Generating the documentation

```
make swagger
```


## Configuration
The application configuration can be specified as YAML and their config location can be specified using the -cnf environment variable, which defaults to current directory

- config folder
  - app.yaml
    - app_env - prod: All debug logs are supressed in stdout, any other values: all logs enabled
    - services - For specifying the port and TLS options


## Building the solution

For building the solution, please 
- Install the gcc and other developer tools
- Copy the cert and config folders
- create a folder where the db file resides 
- RUN 
    - CGO_ENABLED=1 go build -o /frankieoneSDK
    
> Pre-built solution for MAC are available as ZIP in release folder

## API Endpoints

| SNo |           ENDPOINTS                | REST API | METHOD |                       DESCRIPTION                             |
|-----|------------------------------------|----------|--------|---------------------------------------------------------------|
|  1  | /isgood                          | Yes      |  POST   | gets all products.                                            |
|        |


## API Return Code

| SNo |           ENDPOINTS                | REST API | METHOD |                       DESCRIPTION                             |
|-----|------------------------------------|----------|--------|---------------------------------------------------------------|
|  1  | /isgood                          | Yes      |  POST   | 200- Success, 500- Internal Server Error.                     |


## Data Models

POST Endpoints requires the body to have the following json data

**DeviceCollection:**
```
[
  {
    "checkType": "DEVICE",
    "activityType": "SIGNUP",
    "checkSessionKey": "string",
    "activityData": [
      {
        "kvpKey": "ip.address",
        "kvpValue1": "1.23.45.123",
        "kvpType": "general.string"
      }
    ]
  },
  {
    "checkType": "DEVICE",
    "activityType": "SIGNUP",
    "checkSessionKey": "string1",
    "activityData": [
      {
        "kvpKey": "ip.address",
        "kvpValue": "1.23.45.123",
        "kvpType": "general.string"
      },
       {
        "kvpKey": "ip.address",
        "kvpValue": "1.23.45.123",
        "kvpType": "general.string"
      }
    ]
  }
]
```
Returns
**200: Success**
```
{
    "puppy": true
}
```
Returns
**500: Error**
```
{
    "code": 500,
    "message": "Key: 'DeviceCheckDetailsObject.CheckType' Error:Field validation for 'CheckType' failed on the 'oneof' tag"
}
```

