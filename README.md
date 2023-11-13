<p align="center">
  <img src="https://socialify.git.ci/mortum5/callback/image?description=1&descriptionEditable=&font=Inter&issues=1&language=1&name=1&owner=1&pattern=Signal&pulls=1&stargazers=1&theme=Light"     alt="callback" width="640" height="320" />
</p>

![Repository Top Language](https://img.shields.io/github/languages/top/mortum5/callback)
![Github Open Issues](https://img.shields.io/github/issues/mortum5/callback)
![GitHub contributors](https://img.shields.io/github/contributors/mortum5/callback)

# Callback service

## Task Description

Write a rest-service that listens on `localhost:9090` for POST requests on /callback.

Run the go service attached to this task. It will send requests to your service at a fixed interval of 5 seconds.

The request body will look like this:

```js
{
    "object_ids": [1,2,3,4,5,6]
}
```

The amount of IDs varies with each request. Expect up to 200 IDs.

Every ID is linked to an object whose details can be fetched from the provided service. Our service listens on localhost:9010/objects/:id and returns the following response:

```js
{
    "id": <id>,
    "online": true|false
}
```

Note that this endpoint has an unpredictable response time between 300ms and 4s!

Your task is to request the object information for every incoming object_id and filter the objects by their "online" status. Store all objects in a PostgreSQL database along with a timestamp when the object was last seen.

Let your service delete objects in the database when they have not been received for more than 30 seconds.

Important: due to business constraints, we are not allowed to miss any callback to our service.

Write code so that all errors are properly recovered and that the endpoint is always available.

Optimize for very high throughput so that this service could work in production.

Bonus: some comments in the code to explain the more complicated parts are appreciated it is a nice bonus if you provide some way to set up the things needed for us to

Test your code.

## Using

```sh
$> docker compose up -d
$> go run cmd/callback/main.go
$> go run cmd/tester/tester.go
```