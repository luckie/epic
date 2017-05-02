package epic

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

/*
func InstallPostgreSQLDatabase(adminUser string, adminPassword string, server string, epicPassword string) (string, error) {

	//err := dropPostgreSQLSchemaTablesAndConstraints(adminUser, adminPassword, server)
	//if err != nil {
	//	return err
	//}
	//err = dropPostgreSQLUserAndDatabase(adminUser, adminPassword, server)
	//if err != nil {
	//	return err
	//}
	err := createPostgreSQLUserAndDatabase(adminUser, adminPassword, server, epicPassword)
	if err != nil {
		return "", err
	}
	//err = createPostgreSQLSchemaTablesAndConstraints(adminUser, adminPassword, server)
	err = createPostgreSQLSchemaTablesAndConstraints(epicPassword, server)
	if err != nil {
		return "", err
	}
	appID, err := createPostgreSQLEpicAppAdminUser(epicPassword, server)
	if err != nil {
		return err
	}
	return appID, nil
}


func UninstallPostgreSQLDatabase(adminUser string, adminPassword string, server string) error {

	err := dropPostgreSQLSchemaTablesAndConstraints(adminUser, adminPassword, server)
	if err != nil {
		return err
	}
	err = dropPostgreSQLUserAndDatabase(adminUser, adminPassword, server)
	if err != nil {
		return err
	}
	return nil

}
*/

/*
==================================================
CREATE EPIC ROLE AND EPIC DATABASE WITH ADMIN USER
==================================================
*/

func createPostgreSQLUserAndDatabase(adminUser string, adminPassword string, server string, epicPassword string) error {

	var err error
	conn := "postgres://" + adminUser + ":" + adminPassword + "@" + server + "/postgres?sslmode=disable"

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return errors.Wrap(err, "Error in sql.Open()")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "Error in db.Ping()")
	}

	stmt, err := db.Prepare(createEpicRoleQuery(epicPassword))
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createUserQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createRoleQuery()")
	}

	stmt, err = db.Prepare(createDatabaseQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createDatabaseQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createDatabaseQuery()")
	}

	stmt, err = db.Prepare(grantDbPrivQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for grantDbPrivQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for grantDbPrivQuery()")
	}

	return nil

}

func createEpicRoleQuery(password string) string {
	return "CREATE ROLE epic WITH LOGIN SUPERUSER PASSWORD '" + password + "';"
}

func createDatabaseQuery() string {
	return "CREATE DATABASE epic OWNER epic;"
}

func grantDbPrivQuery() string {
	return "GRANT ALL PRIVILEGES ON DATABASE epic to epic;"
}

/*
=================================================
CREATE EPIC SCHEMA AND EPIC TABLES WITH EPIC USER
=================================================
*/

//func createPostgreSQLSchemaTablesAndConstraints(adminUser string, adminPassword string, server string) error {
func createPostgreSQLSchemaTablesAndConstraints(epicPassword string, server string) error {

	var err error
	//conn := "postgres://" + adminUser + ":" + adminPassword + "@" + server + "/epic?sslmode=disable"
	conn := "postgres://epic:" + epicPassword + "@" + server + "/epic?sslmode=disable"

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return errors.Wrap(err, "Error in sql.Open()")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "Error in db.Ping()")
	}

	stmt, err := db.Prepare(createSchemaQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createSchemaQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createSchemaQuery()")
	}

	stmt, err = db.Prepare(grantTablePrivQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for grantTablePrivQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for grantTablePrivQuery()")
	}

	stmt, err = db.Prepare(createTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createTagTableQuery()")
	}

	stmt, err = db.Prepare(alterTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterTagTableQuery()")
	}

	stmt, err = db.Prepare(createContentTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createContentTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createContentTagTableQuery()")
	}

	stmt, err = db.Prepare(alterContentTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterContentTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterContentTagTableQuery()")
	}

	stmt, err = db.Prepare(createConfigTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createConfigTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createConfigTableQuery()")
	}

	stmt, err = db.Prepare(alterConfigTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterConfigTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterConfigTableQuery()")
	}

	stmt, err = db.Prepare(createApplicationUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createApplicationUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createApplicationUserTableQuery()")
	}

	stmt, err = db.Prepare(alterApplicationUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationUserTableQuery()")
	}

	stmt, err = db.Prepare(createUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createUserTableQuery()")
	}

	stmt, err = db.Prepare(alterUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterUserTableQuery()")
	}

	stmt, err = db.Prepare(createEntryTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createEntryTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createEntryTableQuery()")
	}

	stmt, err = db.Prepare(alterEntryTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterEntryTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterEntryTableQuery()")
	}

	stmt, err = db.Prepare(createApplicationTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createApplicationTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createApplicationTableQuery()")
	}

	stmt, err = db.Prepare(alterApplicationTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationTableQuery()")
	}

	stmt, err = db.Prepare(createContentTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createContentTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createContentTableQuery()")
	}

	stmt, err = db.Prepare(alterContentTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterContentTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterContentTableQuery()")
	}

	stmt, err = db.Prepare(createLocaleTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createLocaleTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createLocaleTableQuery()")
	}

	stmt, err = db.Prepare(alterLocaleTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterLocaleTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterLocaleTableQuery()")
	}

	stmt, err = db.Prepare(alterTagTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterTagTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterTagTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterContentTagTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterContentTagTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterContentTagTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterConfigTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterConfigTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterConfigTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterApplicationUserTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationUserTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationUserTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterUserTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterUserTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterUserTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(createUserIndexQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createUserIndexQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createUserIndexQuery()")
	}

	stmt, err = db.Prepare(alterEntryTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterEntryTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterEntryTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterApplicationTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(createApplicationIndexQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createApplicationIndexQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createApplicationIndexQuery()")
	}

	stmt, err = db.Prepare(alterContentTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterContentTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterContentTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterLocaleTablePrimaryKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterLocaleTablePrimaryKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterLocaleTablePrimaryKeyQuery()")
	}

	stmt, err = db.Prepare(alterConfigTableForeignKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterConfigTableForeignKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterConfigTableForeignKeyQuery()")
	}

	stmt, err = db.Prepare(alterApplicationUserTableForeignKey1Query())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationUserTableForeignKey1Query()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationUserTableForeignKey1Query()")
	}

	stmt, err = db.Prepare(alterApplicationUserTableForeignKey2Query())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterApplicationUserTableForeignKey2Query()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterApplicationUserTableForeignKey2Query()")
	}

	stmt, err = db.Prepare(alterEntryTableForeignKey1Query())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterEntryTableForeignKey1Query()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterEntryTableForeignKey1Query()")
	}

	stmt, err = db.Prepare(alterEntryTableForeignKey2Query())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterEntryTableForeignKey2Query()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterEntryTableForeignKey2Query()")
	}

	stmt, err = db.Prepare(alterContentTableForeignKeyQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for alterContentTableForeignKeyQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for alterContentTableForeignKeyQuery()")
	}

	return nil
}

func createSchemaQuery() string {
	return "CREATE SCHEMA epic;"
}

func grantTablePrivQuery() string {
	return "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA epic TO epic;"
}

func createTagTableQuery() string {
	return "CREATE TABLE epic.tag (id uuid NOT NULL, application_id uuid NOT NULL, value text NOT NULL) WITH (OIDS=FALSE);"
}

func alterTagTableQuery() string {
	return "ALTER TABLE epic.tag OWNER TO epic;"
}

func createContentTagTableQuery() string {
	return "CREATE TABLE epic.content_tag (content_id uuid NOT NULL, tag_id uuid NOT NULL) WITH (OIDS=FALSE);"
}

func alterContentTagTableQuery() string {
	return "ALTER TABLE epic.content_tag OWNER TO epic;"
}

func createConfigTableQuery() string {
	return "CREATE TABLE epic.config (id uuid NOT NULL, application_id uuid NOT NULL, name text NOT NULL, value text, updated_at timestamp(6) NOT NULL) WITH (OIDS=FALSE);"
}

func alterConfigTableQuery() string {
	return "ALTER TABLE epic.config OWNER TO epic;"
}

func createApplicationUserTableQuery() string {
	return "CREATE TABLE epic.application_user (application_id uuid NOT NULL, user_id uuid NOT NULL) WITH (OIDS=FALSE);"
}

func alterApplicationUserTableQuery() string {
	return "ALTER TABLE epic.application_user OWNER TO epic;"
}

func createUserTableQuery() string {
	return "CREATE TABLE epic.user (id uuid NOT NULL, first_name text, last_name text, username text NOT NULL, password text, salt text, token text, private_key text, public_key text, email text, token_expires timestamp(6) WITH TIME ZONE) WITH (OIDS=FALSE);"
}

func alterUserTableQuery() string {
	return "ALTER TABLE epic.user OWNER TO epic;"
}

func createEntryTableQuery() string {
	return "CREATE TABLE epic.entry (id uuid NOT NULL, content_id uuid NOT NULL, locale_id uuid NOT NULL, data text, timestamp timestamp(6) NOT NULL) WITH (OIDS=FALSE);"
}

func alterEntryTableQuery() string {
	return "ALTER TABLE epic.entry OWNER TO epic;"
}

func createApplicationTableQuery() string {
	return "CREATE TABLE epic.application (id uuid NOT NULL, name text NOT NULL, code text) WITH (OIDS=FALSE);"
}

func alterApplicationTableQuery() string {
	return "ALTER TABLE epic.application OWNER TO epic;"
}

func createContentTableQuery() string {
	return "CREATE TABLE epic.content (id uuid NOT NULL, application_id uuid NOT NULL, name text NOT NULL, description text, timestamp timestamp(6) NOT NULL) WITH (OIDS=FALSE);"
}

func alterContentTableQuery() string {
	return "ALTER TABLE epic.content OWNER TO epic;"
}

func createLocaleTableQuery() string {
	return "CREATE TABLE epic.locale (id uuid NOT NULL, name text NOT NULL, code text NOT NULL) WITH (OIDS=FALSE);"
}

func alterLocaleTableQuery() string {
	return "ALTER TABLE epic.locale OWNER TO epic;"
}

func alterTagTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.tag ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterContentTagTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.content_tag ADD PRIMARY KEY (content_id, tag_id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterConfigTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.config ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterApplicationUserTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.application_user ADD PRIMARY KEY (application_id, user_id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterUserTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.user ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func createUserIndexQuery() string {
	return "CREATE UNIQUE INDEX  user_id_key ON epic.user USING btree(id pg_catalog.uuid_ops ASC NULLS LAST);"
}

func alterEntryTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.entry ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterApplicationTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.application ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func createApplicationIndexQuery() string {
	return "CREATE UNIQUE INDEX  application_id_key ON epic.application USING btree(id pg_catalog.uuid_ops ASC NULLS LAST);"
}

func alterContentTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.content ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterLocaleTablePrimaryKeyQuery() string {
	return "ALTER TABLE epic.locale ADD PRIMARY KEY (id) NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterConfigTableForeignKeyQuery() string {
	return "ALTER TABLE epic.config ADD CONSTRAINT config_application_id_fkey FOREIGN KEY (application_id) REFERENCES epic.application (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterApplicationUserTableForeignKey1Query() string {
	return "ALTER TABLE epic.application_user ADD CONSTRAINT application_user_application_id_fkey FOREIGN KEY (application_id) REFERENCES epic.application (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterApplicationUserTableForeignKey2Query() string {
	return "ALTER TABLE epic.application_user ADD CONSTRAINT application_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES epic.user (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterEntryTableForeignKey1Query() string {
	return "ALTER TABLE epic.entry ADD CONSTRAINT fk_entry_content FOREIGN KEY (content_id) REFERENCES epic.content (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterEntryTableForeignKey2Query() string {
	return "ALTER TABLE epic.entry ADD CONSTRAINT fk_entry_locale FOREIGN KEY (locale_id) REFERENCES epic.locale (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

func alterContentTableForeignKeyQuery() string {
	return "ALTER TABLE epic.content ADD CONSTRAINT fk_content_application FOREIGN KEY (application_id) REFERENCES epic.application (id) ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;"
}

/*
==================================================
CREATE EPIC APPLICATION AND EPIC ADMIN USER
==================================================
*/

//func createPostgreSQLEpicAppAdminUser(adminUser string, adminPassword string, server string, epicPassword string) error {
func createPostgreSQLEpicAppAdminUser(epicPassword string, server string) (string, error) {

	var err error
	conn := "postgres://epic:" + epicPassword + "@" + server + "/epic?sslmode=disable"

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return "", errors.Wrap(err, "Error in sql.Open()")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return "", errors.Wrap(err, "Error in db.Ping()")
	}

	/*
	stmt, err := db.Prepare(createEpicUserQuery(epicPassword))
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for createSchemaQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for createSchemaQuery()")
	}
	*/

	appID, err := CreateApplication("Epic", "epic")
	if err != nil {
		return "", errors.Wrap(err, "Error in createPostgreSQLEpicAppAdminUser() | CreateApplication()")
	}

	u := User{
		ID: uuid.NewV4(),
		Username: "epic",
		Password: epicPassword,
		AppID: appID,
	}
	uPtr := &u
	_, err = CreateUser(uPtr)
	if err != nil {
		return "", errors.Wrap(err, "Error in createPostgreSQLEpicAppAdminUser() | Epic application was created, but CreateUser() failed.")
	}

	return appID.String(), nil
}

/*
func createEpicUserQuery(password string) string {
	return ""
}

func createPostgreSQLEpicApplication(adminUser string, adminPassword string, server string, appName string) (string, error) {
	return "", nil
}
*/


/*
=================================================
DROP EPIC SCHEMA AND EPIC TABLES WITH EPIC USER
=================================================
*/

func dropPostgreSQLSchemaTablesAndConstraints(adminUser string, adminPassword string, server string) error {

	var err error
	conn := "postgres://" + adminUser + ":" + adminPassword + "@" + server + "/epic?sslmode=disable"

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return errors.Wrap(err, "Error in sql.Open()")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "There is not an Epic database to connect to.  Error in db.Ping().")
	}

	stmt, err := db.Prepare(dropTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropTagTableQuery()")
	}

	stmt, err = db.Prepare(dropContentTagTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropContentTagTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropContentTagTableQuery()")
	}

	stmt, err = db.Prepare(dropConfigTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropConfigTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropConfigTableQuery()")
	}

	stmt, err = db.Prepare(dropApplicationUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropApplicationUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropApplicationUserTableQuery()")
	}

	stmt, err = db.Prepare(dropUserTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropUserTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropUserTableQuery()")
	}

	stmt, err = db.Prepare(dropEntryTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropEntryTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropEntryTableQuery()")
	}

	stmt, err = db.Prepare(dropContentTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropContentTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropContentTableQuery()")
	}

	stmt, err = db.Prepare(dropApplicationTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropApplicationTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropApplicationTableQuery()")
	}

	stmt, err = db.Prepare(dropLocaleTableQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropLocaleTableQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropLocaleTableQuery()")
	}

	stmt, err = db.Prepare(dropSchemaQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropSchemaQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropSchemaQuery()")
	}

	return nil
}

//func createUserIndexQuery() string {
//	return "CREATE UNIQUE INDEX  user_id_key ON epic.user USING btree(id pg_catalog.uuid_ops ASC NULLS LAST);"
//}

//func createApplicationIndexQuery() string {
//	return "CREATE UNIQUE INDEX  application_id_key ON epic.application USING btree(id pg_catalog.uuid_ops ASC NULLS LAST);"
//}

func dropTagTableQuery() string {
	return "DROP TABLE IF EXISTS epic.tag;"
}

func dropContentTagTableQuery() string {
	return "DROP TABLE IF EXISTS epic.content_tag;"
}

func dropConfigTableQuery() string {
	return "DROP TABLE IF EXISTS epic.config;"
}

func dropApplicationUserTableQuery() string {
	return "DROP TABLE IF EXISTS epic.application_user;"
}

func dropUserTableQuery() string {
	return "DROP TABLE IF EXISTS epic.user;"
}

func dropEntryTableQuery() string {
	return "DROP TABLE IF EXISTS epic.entry;"
}

func dropContentTableQuery() string {
	return "DROP TABLE IF EXISTS epic.content;"
}

func dropApplicationTableQuery() string {
	return "DROP TABLE IF EXISTS epic.application;"
}

func dropLocaleTableQuery() string {
	return "DROP TABLE IF EXISTS epic.locale;"
}

func dropSchemaQuery() string {
	return "DROP SCHEMA IF EXISTS epic CASCADE;"
}

/*
==================================================
DROP EPIC ROLE AND EPIC DATABASE WITH ADMIN USER
==================================================
*/

func dropPostgreSQLUserAndDatabase(adminUser string, adminPassword string, server string) error {

	var err error
	conn := "postgres://" + adminUser + ":" + adminPassword + "@" + server + "/postgres?sslmode=disable"

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return errors.Wrap(err, "Error in sql.Open()")
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "Error in db.Ping().")
	}

	stmt, err := db.Prepare(dropDatabaseQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropDatabaseQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropDatabaseQuery()")
	}

	stmt, err = db.Prepare(dropRoleQuery())
	if err != nil {
		return errors.Wrap(err, "Error in db.Prepare() for dropUserQuery()")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return errors.Wrap(err, "Error in statement.Exec() for dropRoleQuery()")
	}

	return nil

}

func dropDatabaseQuery() string {
	return "DROP DATABASE IF EXISTS epic;"
}

func dropRoleQuery() string {
	return "DROP ROLE IF EXISTS epic;"
}
