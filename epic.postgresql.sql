
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
--  Records of tag
-- ----------------------------
BEGIN;
INSERT INTO "epic"."tag" VALUES ('88880fe5-c4f3-4333-bb84-2d5120132c3a', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'my-tag');
INSERT INTO "epic"."tag" VALUES ('a8c99e02-016b-4952-a04a-d4ffcf5b3b60', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'your-tag');
INSERT INTO "epic"."tag" VALUES ('d78edb7c-89e7-4f1d-a1ad-39ef1505c57d', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'her-tag');
INSERT INTO "epic"."tag" VALUES ('e6c240dc-d034-4913-9870-1a114fe0c17b', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'his-tag');
COMMIT;

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
--  Records of content_tag
-- ----------------------------
BEGIN;
INSERT INTO "epic"."content_tag" VALUES ('3fe5c0ff-680b-435d-91d7-75c4288f4e1d', 'e6c240dc-d034-4913-9870-1a114fe0c17b');
INSERT INTO "epic"."content_tag" VALUES ('3fe5c0ff-680b-435d-91d7-75c4288f4e1d', 'a8c99e02-016b-4952-a04a-d4ffcf5b3b60');
COMMIT;

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
--  Records of config
-- ----------------------------
BEGIN;
INSERT INTO "epic"."config" VALUES ('e63845d0-cac6-431f-b006-097c07838e99', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'letsencrypt', '{
	"Email": "",
	"Reg": {
		"body": {
			"resource": "reg",
			"id": 1460199,
			"key": {
				"kty": "EC",
				"crv": "P-384",
				"x": "OR7z-OW2YzeJsZbEKo57K6xorTYJJ26Igj2pvHq-Is56Z1erCETQjZjfYtcq4-Do",
				"y": "8jsq8lBONJpTz82GdF6BEufVJse-j3JJ9rF0fO2-vlESKWFSxyKwq1sa9iQwsxtt"
			},
			"contact": null,
			"agreement": "https://letsencrypt.org/documents/LE-SA-v1.0.1-July-27-2015.pdf"
		},
		"uri": "https://acme-v01.api.letsencrypt.org/acme/reg/1460199",
		"new_authzr_uri": "https://acme-v01.api.letsencrypt.org/acme/new-authz",
		"terms_of_service": "https://letsencrypt.org/documents/LE-SA-v1.0.1-July-27-2015.pdf"
	},
	"Key": "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDCyyGBwdqi9g9/5L36lVxdZJxemg7NE1LGTd46xtMU8ewWeVfk02m4g\nIX3Na8/zi/6gBwYFK4EEACKhZANiAAQ5HvP45bZjN4mxlsQqjnsrrGitNgknboiC\nPam8er4iznpnV6sIRNCNmN9i1yrj4OjyOyryUE40mlPPzYZ0XoES59Umx76Pckn2\nsXR87b6+URIpYVLHIrCrWxr2JDCzG20=\n-----END EC PRIVATE KEY-----\n",
	"Hosts": null,
	"Certs": {
		"test001.schuttsports.com": {
			"Cert": "-----BEGIN CERTIFICATE-----\nMIIESTCCAzGgAwIBAgISA/i1ndLf1T9jRFaumdoHTIV9MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0xNjA0MjYxNjA1MDBaFw0x\nNjA3MjUxNjA1MDBaMCMxITAfBgNVBAMTGHRlc3QwMDEuc2NodXR0c3BvcnRzLmNv\nbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA9spy5gnWa5pO6lZu2lAINTZi4R\ndDy2PF7rBmLTP5EQkgaOWbfk5Qv6b3WKuVLHSwMaxnuBQD2PIKgFjrudcVejggIZ\nMIICFTAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUF\nBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFKPgAGkRNlM8XX3TthjV2lQKY3Lr\nMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMHAGCCsGAQUFBwEBBGQw\nYjAvBggrBgEFBQcwAYYjaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0Lm9y\nZy8wLwYIKwYBBQUHMAKGI2h0dHA6Ly9jZXJ0LmludC14My5sZXRzZW5jcnlwdC5v\ncmcvMCMGA1UdEQQcMBqCGHRlc3QwMDEuc2NodXR0c3BvcnRzLmNvbTCB/gYDVR0g\nBIH2MIHzMAgGBmeBDAECATCB5gYLKwYBBAGC3xMBAQEwgdYwJgYIKwYBBQUHAgEW\nGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIGrBggrBgEFBQcCAjCBngyBm1Ro\naXMgQ2VydGlmaWNhdGUgbWF5IG9ubHkgYmUgcmVsaWVkIHVwb24gYnkgUmVseWlu\nZyBQYXJ0aWVzIGFuZCBvbmx5IGluIGFjY29yZGFuY2Ugd2l0aCB0aGUgQ2VydGlm\naWNhdGUgUG9saWN5IGZvdW5kIGF0IGh0dHBzOi8vbGV0c2VuY3J5cHQub3JnL3Jl\ncG9zaXRvcnkvMA0GCSqGSIb3DQEBCwUAA4IBAQAPxJ/HOfhoNk+KPe0+6McPMZut\nAte+OFYo9TWXzGYAnVGJZxGRyZBonOrRFHxcKBBFRu1HHpLmXw6/Ma3w1qGSJukR\nO0v8iUXdeTdZ9N2zTlZWlmTHJk1t1rImXm3CljRyPWOGs8SALgyKSjum1KC9a0Il\nYxjtsNQUFn8vFIxjMdTwwJcsjHWAoM7Fs5kft6FYtZ3fG8LnVdaPBLZe8bIM0nIY\nA/KBFMjzY6D9wxBiiIEPbVuxJvz0vRDD3Cqpu5jn0a537bxIVgKtCRshzia/foNK\n0rVosn95eu4gQWX17TnjEm53g3UF9vczpGOWf/a3HJLLu2bTMu1/+08/wKou\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----\n",
			"Key": "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIHAIdFQknK9fNhaZBVrPRcXtJI6CfTjpe6rk2KfatXS1oAoGCCqGSM49\nAwEHoUQDQgAED2ynLmCdZrmk7qVm7aUAg1NmLhF0PLY8XusGYtM/kRCSBo5Zt+Tl\nC/pvdYq5UsdLAxrGe4FAPY8gqAWOu51xVw==\n-----END EC PRIVATE KEY-----\n"
		}
	}
}', '2016-01-01 00:00:00');
INSERT INTO "epic"."config" VALUES ('aee4b48f-71e7-4930-8286-fdc766987c77', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'private-key', '-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----', '2016-05-01 00:00:00');
INSERT INTO "epic"."config" VALUES ('f008d0e2-5240-4097-a589-f874c8cacefa', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'public-key', '-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA6l3BHG1njK5r/avIUiid
56tvOiD7QAFfKq0WcMdZIvE+/aNJ9YevLEgoP51Ii0L41oaeeCoGqgccY/0owyya
B7GA3lJzyALhmbpEg6WMYJ33lMFrauRNHQshbxru2NWXjoFItee3mfXczUuGeRON
LshNZ8xZZlViPOlhUNbMfxMtsD55qJsV1zjnU8syWxF+Qs8GMMWpeO0+HHSWlP+w
c3slkbjCZspkM3LmcLGddV5lszK4DMAkqKuubLIYT/ENvgYXBE1x4nA8o8PMMsMq
PdFJpmR3CHbLpqCzWQ+J0GRDzjjPFqI5VcmkIJ0e0kqXWmANv4bswLAYSb+89G90
YwWA5F03d24XCbIPpITvVwDfHyipm3lawrTA3KuGyygfj+F79nkIsJY1pAxIsOcd
m4KYJNgw8gUVmF/C3EZdNeFNmP0G2GK4EffXcA03ziZ5jvTZidayicicvFd9mD15
e1xyogDwOdEYzdrHqSOg/J2rs85G58fRBsEzxqkwZ5eY9g77XF6TWwzKFBaZFT+F
HIAmAqrBBV1vmNUN7Wq3QWjrRpJE6v4ZwpJJaeme1UwTYbzhFdMJSaa+5Tu6o2jJ
GnyKY9uCeWATHUP1tXtysJz23W+aBdBJtjn4wjNMlcD/oGLG2SqAHwmvJUsK0YFA
J0AoXt9nTMekU/S4AbrJevUCAwEAAQ==
-----END PUBLIC KEY-----', '2016-05-01 00:00:00');
COMMIT;

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
	"token_expires" time(6),
	"private_key" text COLLATE "default",
	"public_key" text COLLATE "default",
	"email" text COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "epic"."user" OWNER TO "epic";

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
--  Records of entry
-- ----------------------------
BEGIN;
INSERT INTO "epic"."entry" VALUES ('957b0e47-f56a-4d5a-b16a-b8313e618f86', '3fe5c0ff-680b-435d-91d7-75c4288f4e1d', '57818858-1f30-4fdc-b280-3264a3ad6c4f', '{"id":"123e4567-e89b-12d3-a456-426655440000","locale":"us-EN","tags":{"my-tag","your-tag"},"data":{"id":"123","your-data":"whatever-you-want"}}', '2016-04-15 18:20:00');
INSERT INTO "epic"."entry" VALUES ('3e7b6d5f-2eec-4ebe-a5ff-d01aa2c6503c', '3fe5c0ff-680b-435d-91d7-75c4288f4e1d', '57818858-1f30-4fdc-b280-3264a3ad6c4f', '{"id":"123e4567-e89b-12d3-a456-426655440000","locale":"us-EN","tags":{"my-tag","their-tag"},"data":{"id":"456","your-data":"whatever-you-want"}}', '2016-04-19 13:14:37.436577');
COMMIT;

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
--  Records of locale
-- ----------------------------
BEGIN;
INSERT INTO "epic"."locale" VALUES ('57818858-1f30-4fdc-b280-3264a3ad6c4f', 'U.S. English', 'us-EN');
COMMIT;

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
--  Records of application
-- ----------------------------
BEGIN;
INSERT INTO "epic"."application" VALUES ('cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'schuttsports.com', 'schutt');
COMMIT;

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
--  Records of content
-- ----------------------------
BEGIN;
INSERT INTO "epic"."content" VALUES ('3fe5c0ff-680b-435d-91d7-75c4288f4e1d', 'cfff0d83-1041-490f-9d2b-6df0a6f6df03', 'test', 'test record', '2016-04-15 17:45:00');
COMMIT;

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
--  Primary key structure for table user
-- ----------------------------
ALTER TABLE "epic"."user" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX  "user_id_key" ON "epic"."user" USING btree("id" "pg_catalog"."uuid_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table application_user
-- ----------------------------
ALTER TABLE "epic"."application_user" ADD PRIMARY KEY ("application_id", "user_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table entry
-- ----------------------------
ALTER TABLE "epic"."entry" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Primary key structure for table locale
-- ----------------------------
ALTER TABLE "epic"."locale" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

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
