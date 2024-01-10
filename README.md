Requires tailwindcss, templ and go 1.21

To run locally, execute:
```sh
tailwindcss -i static/tw.css -o static/main.css --minify && templ generate && go run main.go
```
### Using Docker

```
docker build -t compare .
docker run -p 3000:3000 compare
```
