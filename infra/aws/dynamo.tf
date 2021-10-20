resource "aws_dynamodb_table" "dynamodb_table_users" {
  name           = "Users"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "user_id"

  attribute {
    name = "user_id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "dynamodb_table_orders" {
  name           = "Orders"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "order_id"

  attribute {
    name = "order_id"
    type = "S"
  }
}
