CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY,
  "product" VARCHAR(255) NOT NULL,
  "user_id" BIGINT NOT NULL,
  "quantity" INT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
