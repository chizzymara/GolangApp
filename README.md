
# Golang-Epic-Api

The application retrieves information about the current free games on the website https://www.epicgames.com/store/en-US/free-games and sends a message to slack periodically.

## _Implementation_
1. Creating go structures (struct) to match the structure of the free games promotions payload when a GET request is called.
2. Creating a function "Get free games" that makes a GET api call to the website to  retrieve the response payload, have a loop that goes through the elements and identifying items with the "discountPercentage" of zero, in the "discountSetting".
3. Identified games are added to the "FreeGames" array.
4.  Function  "Main" and "SendSlackNotification" are used to send the contents of the "FreeGames" array to "SLACK_URL"
5. A make file was created to make it easier to run the app by simply using the "make" command.
6. A docker file was created to make a docker image and container of the application.
7. The application was also deployed using Kubernetes.


## _Requirement_
Slack url stored as an environmental variable.
```sh
export SLACK_URL="https://hooks.slack.com/services/avcdefghijk/lmnbvcza/hhgwdsvfsffa"
```

## _How to use_
There are four ways to make the app run:

1  Using the go run comand.
```sh
go run epic.go
```

2  Using the make comand.
```sh
make
```
3  Using the docker file.

- First build the image:
```sh
docker build -t golangapp:v1 .   
```

- Then run the image with the environmental variable:
```sh
docker run -e SLACK_URL=$SLACK_URL golangapp:v1  
```

Alternatively, the image can be pulled from docker hub to to skip the first step above.
```sh
docker pull chizzymara/goapp  
```

4  Using Kubernetes.

- Add the Slack app URL to the kubernetes/secret.yaml file:
- create the namespace:
```sh
kubectl apply -f    kubernetes/namespace.yaml
```
- create the secret:
```sh
kubectl apply -f    kubernetes/secret.yaml
```
- create the cronjob:
```sh
kubectl apply -f    kubernetes/cronjob.yaml
```


## _Further Improvements_
- Add unit test
- Add terraform module to deploy on AWS Lambda (Event bridge trigger)
