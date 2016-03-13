CREATE TABLE "application" (
"id" uuid NOT NULL,
"name" text COLLATE "default" NOT NULL,
CONSTRAINT "application_pkey" PRIMARY KEY ("id")
)
WITHOUT OIDS;

ALTER TABLE "application" OWNER TO "epic";

CREATE TABLE "entry" (
"id" uuid NOT NULL,
"content-id" uuid NOT NULL,
"locale-id" uuid NOT NULL,
"value" text COLLATE "default",
"timestamp" timestamp(6) NOT NULL,
CONSTRAINT "content_pkey" PRIMARY KEY ("id")
)
WITHOUT OIDS;

ALTER TABLE "entry" OWNER TO "epic";

CREATE TABLE "content" (
"id" uuid NOT NULL,
"application-id" uuid NOT NULL,
"name" text COLLATE "default" NOT NULL,
"description" text COLLATE "default",
"timestamp" timestamp(6) NOT NULL,
CONSTRAINT "content-position_pkey" PRIMARY KEY ("id")
)
WITHOUT OIDS;

ALTER TABLE "content" OWNER TO "epic";

CREATE TABLE "locale" (
"id" uuid NOT NULL,
"name" text COLLATE "default" NOT NULL,
"code" text COLLATE "default" NOT NULL,
CONSTRAINT "localization_pkey" PRIMARY KEY ("id")
)
WITHOUT OIDS;

ALTER TABLE "locale" OWNER TO "epic";


ALTER TABLE "content" ADD CONSTRAINT "fk_content_application" FOREIGN KEY ("application-id") REFERENCES "application" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "entry" ADD CONSTRAINT "fk_entry_content" FOREIGN KEY ("content-id") REFERENCES "content" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "entry" ADD CONSTRAINT "fk_entry_locale" FOREIGN KEY ("locale-id") REFERENCES "locale" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
