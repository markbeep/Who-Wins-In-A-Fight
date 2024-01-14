# Who Would Win in a Fight?

It's a simple question. You are given two characters. Who would win in a fight?

![Index Page](media/index_page.png)

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

## Generating Models

To work with the database, an SQL builder called [SQLBoiler](https://github.com/volatiletech/sqlboiler) is used. It generates model files from a running database instance.

First start up the postgres instance:

```sh
docker compose up
```

Then run the migrations as explained above. And lastly run SQLBoiler:

```sh
sqlboiler psql
```

If you don't want to install SQLBoiler, you can also execute it using the Dockerfile provided in `tools/`:

```sh
docker build -f tools/Dockerfile.sqlboiler . -t build-models
docker run -v $(PWD)/models:/app/models --network host build-models psql

```

_Note, on big changes clear the directory beforehand, because SQLBoiler does not delete any old files._

## Design

[Drawio design plans](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1#R7Zrfc5s4EMf%2FGj%2FmBpDB%2BLF20jRzbe6mvo6fZZBBV0CcLNs4f31XIPwDkda5GsRDZpzYrASIj7672pU9QvO0eOQ4j7%2BwkCQjxwqLEbofOY7t2gjepOWgLJalLBGnobKdDAv6QuqOyrqlIdlcdBSMJYLml8aAZRkJxIUNc872l93WLLm8a44johkWAU5065KGIq6svmud7J8IjeL6zvCAVUuK687KsIlxyPZnJvQwQnPOmKg%2BpcWcJJJezaU67%2BMrrceBcZKJa0749%2FvfT%2F8sU%2BJPn6z48NfLn%2FfPyzuvusoOJ1v1wGqw4lAT4GybhURexBqh2T6mgixyHMjWPUw62GKRJnBkw0d9UGqcO8IFKc5MapCPhKVE8AN0Ua1eTbCWzEQd70%2F83Voe8Rl75CkjVnMeHa99wgIfFJk3UJq0UPISuO1szeA5z3F5%2F21Z3XC3KeX8ATo447wo8dTt8CmS74ttFJGNqK8Hw6suWbVqkwEMpX0jOPtO5ixhHOwZy4i8JU2ShgknNMrgMCFreZqcBAra%2FqDMKQ1DeeXZBiaUZtHnstv9%2BGT5qvBKE4PT10kp4RhOJHCFWc5oJkrc7gxeMAFz6w935MJY53Bsn47hJbtzMWcZDB%2FTUiAEb8RePj%2BacSawwKvySQ0pDelKs50WpY27Epo%2FPHdEl4w827g3TrvzxidAW7z7ogFfbMjMbQn6ra6IulJZrekB%2B%2BIAVkbbbqF0I2f8THBI%2BIphHr67pHmXbF0d3V5dcqyJjaaRbdwtj4GpXiOn7nXB64jv9qhcDdWMrYyTQg1N%2BeMrSXUXwPTcHjQ1gFDfzLsGoKm27LSKyyHdyTuquHkK5QEAILyl6oEhwQkZLt9wFpYF82Elj8gOQiewr7pA8R9wmgvKsvpmK36K%2F8dloRzBK6vCm%2Baufgg19E4ChHnZO3oS%2FakE3kDXvdDdxirj66uM1%2Bcig%2FS8b84JFibYNHVjHo6e7hmSjTc4NEhD84yNoJk00Bw5nKFxJr2y0bORb3nCsAz7NJU7sOYpIcs4JT0VWcZMjohtk7AcmVwYy38Y%2FtbVmD5q8FQV9LNF7dVShxOozM7KjMsKRhYpeCtYVb2Vl7621rpFOJw2cqKxPmV2v06vp0SL7Sql4q2CPgc2clCIib8ONLrQ4gU%2BWa07cgG3Rxd4nBZhwXZflo9kt8TPL%2F7TIr9r2034%2FURcI9LC7epN4n43QFspdfh1xENI37%2BL%2BJ%2BbLTeVWdsGqN%2BnK7ZENoJ5EBt3yGbU8sxHLb1kuNFmy%2B9JqlFa%2Be61VafbFSg9vt9qs%2BWmpCZWn5st7aicQWpqMh2cpvS6axCaapIagqb0%2FfMhaGra3BQ2rym9Xh2EppqkhqAp%2FRdEC7zTC3rTIX3q6FlCZ18ot5PS0%2FZBkGoGqgGQ0jenB0Gq6X7mSTl65lkYxzQZ%2FzpKoV4p6WmneUq%2BPzRKesZpnpJtOb1hgsPTr3HLtrMfNaOHHw%3D%3D)
