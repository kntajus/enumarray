# Retrieving an array of enums with sqlc/pgx

1. Create a PostgreSQL database, then run the two `CREATE` statements in `query.sql` against it.

2. Run this code from the root directory with:
```
DATABASE_URL=<your_db_conn> go run .
```

3. Results in:
```
panic: reflect.Value.Convert: value of type *[]db.Fruit cannot be converted to type *[]string
```
