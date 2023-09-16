CREATE TABLE "products" (
    "id" serial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "price" decimal(12, 2) NOT NULL ,
    "description" varchar NOT NULL
);

CREATE INDEX ON "products" ("name");