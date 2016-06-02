

-- ----------------------------
--  Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS "epic"."tag";
CREATE TABLE "epic"."tag" (
	"id" uuid NOT NULL,
	"application_id" uuid NOT NULL,
	"value" text NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."tag" OWNER TO "epic";

-- ----------------------------
--  Table structure for content_tag
-- ----------------------------
DROP TABLE IF EXISTS "epic"."content_tag";
CREATE TABLE "epic"."content_tag" (
	"content_id" uuid NOT NULL,
	"tag_id" uuid NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."content_tag" OWNER TO "epic";

-- ----------------------------
--  Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS "epic"."config";
CREATE TABLE "epic"."config" (
	"id" uuid NOT NULL,
	"application_id" uuid NOT NULL,
	"name" text NOT NULL COLLATE "default",
	"value" text COLLATE "default",
	"updated_at" timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."config" OWNER TO "epic";

-- ----------------------------
--  Table structure for application_user
-- ----------------------------
DROP TABLE IF EXISTS "epic"."application_user";
CREATE TABLE "epic"."application_user" (
	"application_id" uuid NOT NULL,
	"user_id" uuid NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."application_user" OWNER TO "epic";

-- ----------------------------
--  Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "epic"."user";
CREATE TABLE "epic"."user" (
	"id" uuid NOT NULL,
	"first_name" text COLLATE "default",
	"last_name" text COLLATE "default",
	"username" text NOT NULL COLLATE "default",
	"password" text COLLATE "default",
	"salt" text COLLATE "default",
	"token" text COLLATE "default",
	"private_key" text COLLATE "default",
	"public_key" text COLLATE "default",
	"email" text COLLATE "default",
	"token_expires" timestamp(6) WITH TIME ZONE
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."user" OWNER TO "epic";

-- ----------------------------
--  Table structure for entry
-- ----------------------------
DROP TABLE IF EXISTS "epic"."entry";
CREATE TABLE "epic"."entry" (
	"id" uuid NOT NULL,
	"content_id" uuid NOT NULL,
	"locale_id" uuid NOT NULL,
	"data" text COLLATE "default",
	"timestamp" timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."entry" OWNER TO "epic";

-- ----------------------------
--  Table structure for application
-- ----------------------------
DROP TABLE IF EXISTS "epic"."application";
CREATE TABLE "epic"."application" (
	"id" uuid NOT NULL,
	"name" text NOT NULL COLLATE "default",
	"code" text COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."application" OWNER TO "epic";

-- ----------------------------
--  Table structure for content
-- ----------------------------
DROP TABLE IF EXISTS "epic"."content";
CREATE TABLE "epic"."content" (
	"id" uuid NOT NULL,
	"application_id" uuid NOT NULL,
	"name" text NOT NULL COLLATE "default",
	"description" text COLLATE "default",
	"timestamp" timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."content" OWNER TO "epic";

-- ----------------------------
--  Table structure for locale
-- ----------------------------
DROP TABLE IF EXISTS "epic"."locale";
CREATE TABLE "epic"."locale" (
	"id" uuid NOT NULL,
	"name" text NOT NULL COLLATE "default",
	"code" text NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."locale" OWNER TO "epic";

-- ----------------------------
--  Primary key structure for table tag
-- ----------------------------
ALTER TABLE "epic"."tag" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table content_tag
-- ----------------------------
ALTER TABLE "epic"."content_tag" ADD PRIMARY KEY ("content_id", "tag_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table config
-- ----------------------------
ALTER TABLE "epic"."config" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table application_user
-- ----------------------------
ALTER TABLE "epic"."application_user" ADD PRIMARY KEY ("application_id", "user_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table user
-- ----------------------------
ALTER TABLE "epic"."user" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX  "user_id_key" ON "epic"."user" USING btree("id" "pg_catalog"."uuid_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table entry
-- ----------------------------
ALTER TABLE "epic"."entry" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table application
-- ----------------------------
ALTER TABLE "epic"."application" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table application
-- ----------------------------
CREATE UNIQUE INDEX  "application_id_key" ON "epic"."application" USING btree("id" "pg_catalog"."uuid_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table content
-- ----------------------------
ALTER TABLE "epic"."content" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table locale
-- ----------------------------
ALTER TABLE "epic"."locale" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Foreign keys structure for table config
-- ----------------------------
ALTER TABLE "epic"."config" ADD CONSTRAINT "config_application_id_fkey" FOREIGN KEY ("application_id") REFERENCES "epic"."application" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Foreign keys structure for table application_user
-- ----------------------------
ALTER TABLE "epic"."application_user" ADD CONSTRAINT "application_user_application_id_fkey" FOREIGN KEY ("application_id") REFERENCES "epic"."application" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "epic"."application_user" ADD CONSTRAINT "application_user_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "epic"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Foreign keys structure for table entry
-- ----------------------------
ALTER TABLE "epic"."entry" ADD CONSTRAINT "fk_entry_content" FOREIGN KEY ("content_id") REFERENCES "epic"."content" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "epic"."entry" ADD CONSTRAINT "fk_entry_locale" FOREIGN KEY ("locale_id") REFERENCES "epic"."locale" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Foreign keys structure for table content
-- ----------------------------
ALTER TABLE "epic"."content" ADD CONSTRAINT "fk_content_application" FOREIGN KEY ("application_id") REFERENCES "epic"."application" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;
