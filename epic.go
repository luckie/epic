package main
/*
import (
	"fmt"
	"io"
	//"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	//"github.com/urfave/cli"
)

var (

	database string
	port int
	err error

	//Trace   *log.Logger
	//Info    *log.Logger
	//Warning *log.Logger
	//Error   *log.Logger
)

func main() {
  var db string
	var port int
	var appCode string

	var dbArg string
	var portArg string
	//var appArg string

	dbEnv 	:= os.Getenv("EPIC_DATABASE")
	portEnv	:= os.Getenv("EPIC_PORT")

	var err error

  app := cli.NewApp()

  app.Flags = []cli.Flag {
	cli.StringFlag{
		Name:        "app, a",
		Usage:       "Epic application.",
		Destination: &portArg,
		EnvVar: 		 "EPIC_APP",
	},
    cli.StringFlag{
      Name:        "database, d",
      Usage:       "Connection string for Epic database.",
      Destination: &dbArg,
      EnvVar: 		 "EPIC_DATABASE",
    },
    cli.StringFlag{
      Name:        "port, p",
			Usage:       "Epic service TCP port number.",
      Destination: &portArg,
      EnvVar: 		 "EPIC_PORT",
    },
  }

  app.Action = func(c *cli.Context) error {
		Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		if dbArg != "" {
			db = dbArg
			os.Setenv("EPIC_DATABASE", dbArg)
		} else if dbEnv != "" {
			db = dbEnv
		} else {
			log.Fatal("The database connection string is not specified by the EPIC_DATABASE environmental variable, nor by the CLI --database (-d) flag.")
		}
		if portArg != "" {
			if port, err = strconv.Atoi(portArg); err != nil {
				log.Fatal("The TCP port number in the CLI --database (-d) flag should be an integer, instead of " + portArg + ".")
			}
			os.Setenv("EPIC_PORT", portArg)
		} else if portEnv != "" {
			if port, err = strconv.Atoi(portEnv); err != nil {
				log.Fatal("The TCP port number in the EPIC_PORT environmental variable should be an integer, instead of " + portEnv + ".")
			}
		} else {
			port = 443
		}
		initSQLDatabase(db)
		Serve("", port, appCode)
		return nil
	}

  app.Run(os.Args)
}



func Init(

	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func old_main() {
	//to make Compile
	var appCode string

	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	switch len(os.Args) {
    case 1:
			fmt.Println("Epic is launching using environmental variables for configuration.")
			//
    case 2:
			if strings.ToLower(os.Args[1]) == "help" {
				help()
			} else {
				//
			}
    case 3:
        fmt.Println("three")
    }


	if len(os.Args) == 1 {
		fmt.Println("Epic is launching using environmental variables for configuration.")
		//
	}

	if len(os.Args) < 3 || len(os.Args) > 4 {
		//
		return
	} else {
		dbPattern := "^postgres:\\/\\/[a-z0-9-_]*:[a-z0-9-_]*@[a-z0-9-_]*\\/[a-z0-9-_]*"
		dbRegex, _ := regexp.Compile(dbPattern)
		if dbRegex.MatchString(os.Args[1]) != true {
			fmt.Println("There is a problem with the format of your database connection string.")
			fmt.Println("Your database connection string: '" + os.Args[1] + "'")
			fmt.Println("Example database connection string: 'postgres://username:password@localhost/epic?sslmode=disable'")
			os.Exit(1)
		} else {
			urlPattern := "[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)"
			urlRegex, _ := regexp.Compile(urlPattern)
			if urlRegex.MatchString(os.Args[2]) != true {
				fmt.Println("There is a problem with the format of your host URL.")
				fmt.Println("Your host URL: '" + os.Args[2] + "'")
				fmt.Println("Example host URL: 'example.com' or 'epic.example.com'")
				fmt.Println("It should not include 'http://' or 'https://' or a trailing '/'.")
				os.Exit(1)
			} else {
				if len(os.Args) > 3 {
					port, err := strconv.Atoi(os.Args[3])
					if err != nil {
						fmt.Println("A port number may optionally be included as the third parameter for HTTP (not HTTPS).")
						fmt.Println("If the port number is not specified, the server will default to HTTPS on port 443.")
						fmt.Println("If a port number other than 443 is specified, the server will use HTTP on the specified port.")
						fmt.Println("Any port number below 1025 requires administrative privileges (sudo).")
						os.Exit(1)
					} else {
						initSQLDatabase(os.Args[1])
						//initServer(os.Args[2], port)
						Serve(os.Args[2], port, appCode)
					}
				} else {
					initSQLDatabase(os.Args[1])
					//initServer(os.Args[2], 443)
					Serve(os.Args[2], 443, appCode)
				}

			}
		}
	}
}

func help() {
	fmt.Println("The Epic database connection string must be included as the first parameter.")
	fmt.Println("e.g. 'postgres://username:password@localhost/epic?sslmode=disable'")
	fmt.Println("The Epic host URL string must be included as the second parameter.")
	fmt.Println("e.g. 'example.com' or 'epic.example.com'")
	fmt.Println("A port number may optionally be included as the third parameter for HTTP (not HTTPS).")
	fmt.Println("If the port number is not specified, the server will default to HTTPS on port 443.")
	fmt.Println("If a port number other than 443 is specified, the server will use HTTP on the specified port.")
	fmt.Println("Any port number below 1025 requires administrative privileges (sudo).")
	fmt.Println("Here is an example of launching Epic using TLS (HTTPS):")
	fmt.Println("'sudo epic postgres://username:password@localhost/epic?sslmode=disable example.com'")
	fmt.Println("Here is an example of launching Epic without TLS (HTTP) on port 8080:")
	fmt.Println("'epic postgres://username:password@localhost/epic?sslmode=disable example.com 8080'")
	fmt.Println("Available API Calls:")
	fmt.Println("GET /content/{uuid}       |  e.g. PUT /content/123e4567-e89b-12d3-a456-426655440000")
	fmt.Println("PUT /content/{uuid}       |  e.g. PUT /content/123e4567-e89b-12d3-a456-426655440000")
	fmt.Println("GET /app/{app}/tag/{tag}  |  e.g. GET /app/my-app/tag/my-tag")
}
*/