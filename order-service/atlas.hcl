env "local" {
  url = "postgres://postgres:password@localhost:5432/order_db?sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}
