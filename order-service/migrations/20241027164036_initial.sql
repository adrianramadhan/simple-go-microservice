-- Create "orders" table
CREATE TABLE "public"."orders" ("id" serial NOT NULL, "product" character varying(255) NOT NULL, "user_id" bigint NOT NULL, "quantity" integer NOT NULL, "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
