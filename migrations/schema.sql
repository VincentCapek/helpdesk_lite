CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"email" TEXT NOT NULL,
"password_hash" TEXT NOT NULL,
"role" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "tickets" (
"id" TEXT PRIMARY KEY,
"title" TEXT NOT NULL,
"description" TEXT NOT NULL,
"status" TEXT NOT NULL,
"priority" TEXT NOT NULL,
"category" TEXT NOT NULL,
"user_id" char(36) NOT NULL,
"agent_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE INDEX "tickets_status_idx" ON "tickets" (status);
CREATE INDEX "tickets_priority_idx" ON "tickets" (priority);
CREATE INDEX "tickets_user_id_idx" ON "tickets" (user_id);
CREATE TABLE IF NOT EXISTS "comments" (
"id" TEXT PRIMARY KEY,
"body" TEXT NOT NULL,
"internal" bool NOT NULL,
"ticket_id" char(36) NOT NULL,
"user_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE INDEX "comments_ticket_id_idx" ON "comments" (ticket_id);
CREATE INDEX "comments_user_id_idx" ON "comments" (user_id);
