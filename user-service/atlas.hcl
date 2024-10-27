env "local" {
  url = "postgres://postgres:password@localhost:5432/user_db?sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}
