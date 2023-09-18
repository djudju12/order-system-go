ALTER TABLE products ADD COLUMN "status" varchar NOT NULL DEFAULT 'ACTIVE';
ALTER TABLE products ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT (now());
ALTER TABLE products ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT (now());