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

## Migrations

Database migrations are handled using sql-migrate and can be executed with the following command (requires the database to be running):

```sh
docker run \
    -v $(pwd)/migrations:/migrations \
    --network host migrate/migrate \
    -path /migrations/ \
    -database 'postgres://docker:docker@localhost:5432/compare?sslmode=disable' \
    down -all
```

## Design

[Drawio design plans](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=Compare%20Design#R7Vxbb%2BI4FP4t%2B8BjV7kTHgttd0Y7e9Ewoz6bxBC3IWYdU%2Bj8%2Bj3OhZDYdEAisTuq1ApyktgnX871s9uRO1vv%2F2Bok%2FxFY5yOHCvej9y7keM4tj%2BBDyF5LSW259qlZMVIXMkawZz8wJXQqqRbEuO8dSGnNOVk0xZGNMtwxFsyxBjdtS9b0rQ96watsCSYRyiVpY8k5kkpDX2rkX%2FCZJXUM9tWdWaN6osrQZ6gmO6ORO79yJ0xSnn5bb2f4VSgV%2BNS3vdw4uxBMYYzfs4NT8%2F%2Ffv72uMbh5LOVvP7z48%2B7vx9vvHKUF5RuqweulOWvNQKMbrMYi0GskTvdJYTj%2BQZF4uwOXjrIEr5O4ciGr9VwmHG8P6mnfXh6sBtM15izV7ikNpoKr9pixtXxroHfr60jOYLeDSohql756jB0gwp8qYC5ACRfAVKQwrTTJYVnOkYr%2BG9L6xM3eWHNt3CB4232BTr1efi2Ep%2BfAdp9PRooVw5YnpPeBGAq5Dln9BnPaEoZyDOaYTEhSdOOCKVklcFhipfiNvFSCBj2bSVekzgWI09zeJskW30pLrvzGsnXClwhonD7Mi3sN4EbMYww3VCS8QJsfwo%2FAP%2FM%2Bt0f%2BaDrDI7t5hh%2BxOWMz2gG6iNSmAJGOd%2FhXCjHKEccLYon1WNmrmxmocrK%2BjKywDxPHFvGueK4P1ecMYw4%2FvBFDb4oGZpuZwzNc8ZOvAp87b446dEX6VogU4%2B3YI0ffrin7lQZ2LLp2c6Q7lkrYJJ%2FdmOYAQ5q2%2F156BeMYswWFLH4wycNSJlKp%2FQHdUq5lyTrla3dMW27jdRYUcUqw9cBvutDJXeUU7rQjpTbCfQT50yk%2Botgcr0PNqU%2F2HeRMsGmVDVrGZhj8iJmrAJnE8sjeFrM5Ng%2BB5XghgwVHyiLCzrrdSGO8AvETsC%2BvCTGecTIhhOavVmrFRqcSAsXvbv6ISrVewkQ%2Bs3ekRP3pwLwDnT9G7rXyTKBDE0wZJKpC70jZOq%2BfXBsunajHxzXULOxQ93IyLWJMWajHxy5GmlV98NbT63RG441KB3kyLn1wFDosJ%2Bfx51h4ZGZIE1xZ9ylKnSH5HpgEwLPeGIaOMaUOV270R6TXTmVfyM81YFN2MHGUZDzznhQcORsfnfciOiHSNEODgyRnNO%2Fb1KKRCtH1miFcwNgchWc1cAwnV5PPNU0F3Sk1DJfyqva4Qleteq9eensZy9%2FXDi7dWL226LXVzX3FlrQ9jTlHVt%2BsZZ9Ka9SL6G7UU1h7BIkJuAJAeO3oqKAI3nDXAysbqEFoNogeFqBt%2FiTilZ%2FiyQ5yZ0zDIoe8dZtSlyw3mjLafkwxdDnkvdXiBbdMvfAeB%2FTMhNFtOiNlnEVbUBRzP12aTA9RmzkuDHC4TKS4IUzQRTixbKn8KtYoRo2%2FPoGbj%2Bz3S55ce6uF78%2FnORsrqlQtg%2BrlaZUyr5q49SVFjnvY3I6OKvD8MfqZj9OqHtHkK9ciTJgddPrOKQ7UaRJ5VKU1xtWcp7kyuZVO1i%2BwqyG3criy9zZHU6xgiC6CKxOgbEMIxwpC4xF6Hu%2B1Re8CoZpWHhrBYxz28A2zm0DmY4zxG27YOl328D5hdxWgle%2F25pDf0pBTT%2B5F5jGf8pFiHZmLzCRAJVx0t6CB3L7NEcv1w1jw9AbEriqHbhKcPvLqHLT8BXnWDSl84gqVo%2FfQbKQarxzC5f%2BTFhuN949ylJK1o%2By3KjUKKPi7lxgLUfVdwD22G2DreK0Dmz0IGDXUaoF9gpnmBVbBywOaMiJ%2Fh1AHZyx0D4w1HIln3Au%2Frr8VgzsPIjVsJs8gWe82TIY7gHl8fMTyieO63jPT%2FAr79y%2BMD%2F6OIw9Ffahs3AhQ18He7%2BzoKJcyHcHxV4u8%2BeJalfVL7vG5fk%2FfydXotPhsPlnA8W5o%2F%2FZ4N7%2FDw%3D%3D)
