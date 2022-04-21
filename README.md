### QueryPerformance API for PostgreSQL
#### Version 1.0

Api that returns performance  data for a postgreSQL database using the extension pg_stat_statements.
Supports sorting by all columns as well as filtering and pagnation.
Built using Fiber framework


### Run project:
docker-compose up -d

### Stop project:

docker-compose down

### Test:

```make test```

## Integration testing:

```make test-integration```

# API Documentation


## /performance

Returns a list of queries and there mean,total,calls and rows retrived or affected

### Endpoint definition

`performance/`

### HTTP method

<span class="label label-primary">GET</span>

### Parameters

| Parameter | Description | Data Type |

|-----------|------|-----|-----------|

| sort_by | *Optional*. Column to order result by. Default is unordered. | string |

| asc | *Optional*. Sort by ASC order. Default is false and will sort by DESC order. | bool |

| offset | *Optional*. Pagnation offset, default 0 | uint

| limit | *Optional*. Pagnation limit, default is show all | uint

| statement | *Optional*. Only queries containing Statement,Multiple values will show queries with either, default show all | string

| Mincalls | *Optional*. Only queries with calls equal or greater than mincalls, default 0 | int




### Sample request

```curl --get --include 'http://127.0.0.1:7070/performance?sort_by=mean_exec_time'```

### Sample response

```json
[
  {
    "query": "SELECT extname FROM pg_extension WHERE extname = $1",
    "calls": 1,
    "mean_exec_time": 0.123666,
    "total_exec_time": 0.123666,
    "stddev_exec_time": 0,
    "rows": 1
  },
  {
    "query": "INSERT INTO users(username, password)\nVALUES ($1, $2)",
    "calls": 2,
    "mean_exec_time": 0.0444785,
    "total_exec_time": 0.088957,
    "stddev_exec_time": 0.0336045,
    "rows": 2
  },
  {
    "query": "SELECT $1 FROM pg_database WHERE datname = $2",
    "calls": 1,
    "mean_exec_time": 0.007,
    "total_exec_time": 0.007,
    "stddev_exec_time": 0,
    "rows": 0
  }
]
```

