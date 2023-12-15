Generator go struct from MySQL database

config.yaml
```
pad:
  host: "127.0.0.1:3306"
  user: "abc"
  password: "def"
  database: "test"
  tablePrefix: "tbl"
  modelSuffix: "Model"

pen:
  host: "127.0.0.1:3306"
  user: "abc"
  password: "def"
  database: "test2"
  tablePrefix: "tbl"
  modelSuffix: "Model"
```
```
tablePrefix set the prefix will trim
modelSuffix set the suffix will add after struct name
```

Usage

```
Usage: sql2struct [options]
Options:
  --config, -c Your config file path, default ./ (default )
  --database, -d Your database id in config.yaml (default )
  --help, -h Show help message (default false)
  --table, -t Your table name (default )
```
```
./sql2struct -d pad -t user_info

```


