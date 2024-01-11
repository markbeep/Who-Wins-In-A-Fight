# Compare

Ever wanted to compare some list of images to find out which one is truly the best? Fret no more, with this tool you can upload images with some description and have them get battled out.

## Development

Requires tailwindcss, templ and go 1.21

To run locally, execute:

```sh
tools/run.sh
```

This will also attempt to start auto reloading on changes.

If that fails, you can also execute the steps manually.

You'll first have to first export the environment variables from the .env file:

```sh
# Neat way to load env variables
set -o allexport
source ./.env
set +o allexport
```

Then you have to build the tailwindcss file, generate the templ components, and then run the webserver:

```sh
tailwindcss -i static/tw.css -o static/main.css \
&& templ generate \
&& go run main.go
```

This has to be restarted manually after every change.

### Using Docker

```
docker build -t compare .
docker run -p 3000:3000 compare
```
