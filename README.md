# Yourbeat

Beat with metrics you really want to know!

My aim is to implement flexible and configurable metrics calculation
Config will look somewhat like this:
```
connectors:
  - connector_name:
    type: mysql
    user: username
    password: securepassword
    database: dbname
  - connector_name2:
    type: redis
    host: host.redis
    port: 6379
metrics:
  - queue_size:
    connector_name2: llen queue
  - last_task:
    connector_name2: get last_task
  - tasks_in_progress:
    connector_name: SELECT COUNT(*) FROM tasks where state='RUNNING'

```