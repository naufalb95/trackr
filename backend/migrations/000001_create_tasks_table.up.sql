CREATE TABLE IF NOT EXISTS tasks (
  "id" TEXT NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "status" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS tasks_title_idx ON tasks ("title");
CREATE INDEX IF NOT EXISTS tasks_status_idx ON tasks ("status");