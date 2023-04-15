# football-data-analysis

This package acquires data from various datasources on football and writes them into a db.

## run code

### create environment variables

```bash
export XAPISPORTSKEY=[YOUR-API-KEY]
export SPORTSBETTINGDB='host=localhost user=postgres password=postgres dbname=postgres port=5434'
```

## host local database

```bash
docker compose up
```

### run service

```bash
cd football-api-sports
go install
go run .
```
