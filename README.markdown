# Go Remote Control (grc)

Control systems remotely via HTTP.

## Running

Set up .env first:

```sh
HTTP_BIND_ADDRESS=127.0.0.1
HTTP_BIND_PORT=6000
HTTP_KEY_FILE=key.pem
HTTP_CERT_FILE=cert.pem
HTTP_AUTH_TOKEN=token
SCRIPT_DIR="./scripts"
SCRIPT_EXT="sh"
```

You will need an SSL certificate and key that correspond to the environment variables in the .env file.

Run foreman:

```sh
foreman run
```

## Remote Calls

To get the service status:

```
curl https://host:6000/status -H "Authorize: token"
```

To start the service:

```
curl https://host:6000/start -H "Authorize: token"
```

To stop the service:

```
curl https://host:6000/stop -H "Authorize: token"
```

To get the grc version:

```
curl https://host:6000/ -H "Authorize: token"
```
