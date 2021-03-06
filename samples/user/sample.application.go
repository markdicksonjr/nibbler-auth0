package main

import (
	"github.com/markdicksonjr/nibbler"
	"github.com/markdicksonjr/nibbler-auth0"
	"github.com/markdicksonjr/nibbler-sql"
	"github.com/markdicksonjr/nibbler-sql/session"
	NibUserSql "github.com/markdicksonjr/nibbler-sql/user"
	"github.com/markdicksonjr/nibbler/session"
	"github.com/markdicksonjr/nibbler/user"
	"log"
)

func main() {

	// configuration
	config, err := nibbler.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// allocate the sql extension, with all models
	// if database.uri is not configured, an in-memory SQLite database will be used
	sqlExtension := nibbler_sql.Extension{
		Models: []interface{}{
			nibbler.User{},
		},
	}

	// allocate session extension
	sessionExtension := session.Extension{
		SessionName: "auth0",
		StoreConnector: connectors.SqlStoreConnector{
			SqlExtension: &sqlExtension,
			Secret:       "somesecret",
		},
	}

	// allocate user extension, providing sql extension to it
	userExtension := user.Extension{
		PersistenceExtension: &NibUserSql.Extension{
			SqlExtension: &sqlExtension,
		},
	}

	// allocate user auth0 extension
	auth0Extension := auth0.UserExtension{
		Extension: auth0.Extension{
			SessionExtension:    &sessionExtension,
			LoggedInRedirectUrl: "/",
		},
		UserExtension: &userExtension,
	}

	// initialize the application, provide config, logger, extensions
	appContext := nibbler.Application{}
	if err = appContext.Init(config, nibbler.DefaultLogger{}, []nibbler.Extension{
		&sqlExtension,
		&userExtension,
		&sessionExtension,
		&auth0Extension,
		&SampleExtension{
			Auth0Extension: &auth0Extension,
		},
	}); err != nil {
		log.Fatal(err.Error())
	}

	// create a test user, if it does not exist
	emailVal := "someone@example.com"
	password := ""
	_, errCreate := userExtension.Create(&nibbler.User{
		Email:    &emailVal,
		Password: &password,
	})

	// assert the test user got created
	if errCreate != nil {
		log.Fatal(errCreate.Error())
	}

	// start the app
	if err = appContext.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
