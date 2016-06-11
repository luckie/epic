package main

import(

  "crypto/ecdsa"
  "crypto/elliptic"
  "regexp"
  "errors"
  "time"
  "encoding/base64"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/scrypt"
	"crypto/rand"
  "fmt"
  "github.com/satori/go.uuid"
)

const (
    SALT_BYTES = 64
    HASH_BYTES = 64
)

type User struct {
  ID            uuid.UUID `json:"id"`
  FirstName     string    `json:"first-name"`
  LastName      string    `json:"last-name"`
  Email         string    `json:"email"`
  AppID         uuid.UUID `json:"app-id"`
  Username      string    `json:"username"`
  Password      string    `json:"password"`
  Salt          string    `json:"salt"`
  PrivateKey    string    `json:"private-key"`
  PublicKey     string    `json:"public-key"`
  Token         *string    `json:"token"`
  TokenExpires  *time.Time `json:"token-expires"`
}

type UserCryptoBootstrap struct {
  PlainText     string    `json:"plain-text"`
  Hash          string    `json:"hash"`
  Salt          string    `json:"salt"`
  PrivateKey    string    `json:"private-key"`
  PublicKey     string    `json:"public-key"`
}

func Authenticate(data string) error {
  //return nil
  pattern := `^([A-Za-z0-9-])*.([A-Za-z0-9-])*.([A-Za-z0-9-])*$`
  //loginPattern := `\b([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}):::([A-Za-z0-9!@#_-]+?):::([A-Za-z0-9]+?)\b`
  regex, _ := regexp.Compile(pattern)
  if regex.MatchString(data) {
    err := parseToken(data)
    if err != nil {
      return errors.New("parseToken | " + err.Error())
    }
    return nil
  } else {
    return errors.New("400 Bad Request: Invalid format for JWT authentication token.")
  }
}

func Login(user *User) (*User, error) {

  username  :=  user.Username
  password  :=  user.Password
  appID     :=  user.AppID.String()

  stmt, err := db.Prepare("select epic_user.id, epic_user.first_name, epic_user.last_name, epic_user.email, epic_user.username, epic_user.password, epic_user.salt, epic_user.token, epic_user.token_expires, epic_user.private_key, epic_user.public_key from epic.user as epic_user inner join epic.application_user on application_user.user_id = epic_user.id inner join epic.application on application_user.application_id = application.id where epic_user.username = $1 and application.id = $2")
  if err != nil {
    return nil, errors.New("Prepare select from user | " + err.Error())
  }
  defer stmt.Close()
  err = stmt.QueryRow(username, appID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.Salt, &user.Token, &user.TokenExpires, &user.PrivateKey, &user.PublicKey)
  if err != nil {
    return nil, errors.New("QueryRow / Scan select from user | " + err.Error())
  }
  hashedPassword, err := hashPassword(password, user.Salt)
  if err != nil {
    return nil, errors.New("hashPassword | " + err.Error())
  }
  if hashedPassword != user.Password {
    return nil, fmt.Errorf("'%v' is not the correct password.", password)
  }
  if err := createToken(user); err != nil {
    return nil, errors.New("createToken | " + err.Error())
  }
  return user, nil
}

func Logout(tokenID string) error {
  return nil
}

func CreateUser(user *User) (*User, error) {
  if err := validatePassword(user.Password); err != nil {
    return nil, errors.New("validatePassword | " + err.Error())
  }
  if user.ID.String() == "" {
    user.ID = uuid.NewV4()
  }
  privateKey, publicKey, err := createKeys()
  if err != nil {
    return nil, errors.New("createKeys | " + err.Error())
  }
  salt, err := createSalt(SALT_BYTES)
  if err != nil {
    return nil, errors.New("createSalt | " + err.Error())
  }
  hashedPassword, err := hashPassword(user.Password, salt)
  if err != nil {
    return nil, errors.New("hashPassword | " + err.Error())
  }
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.New("Transaction Begin | " + err.Error())
	}
	defer tx.Rollback()
  e_stmt1, err := tx.Prepare("insert into epic.user (id, first_name, last_name, email, username, password, salt, private_key, public_key) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
  if err != nil {
  	return nil, errors.New("Prepare insert into epic.user | " + err.Error())
  }
  defer e_stmt1.Close()
  _, err = e_stmt1.Exec(user.ID.String(), user.FirstName, user.LastName, user.Email, user.Username, hashedPassword, salt, privateKey, publicKey)
  if err != nil {
  	return nil, errors.New("Exec insert into epic.user | " + err.Error())
  }
	e_stmt2, err := tx.Prepare("insert into epic.application_user (user_id, application_id) values ($1, $2)")
	if err != nil {
		return nil, errors.New("Prepare insert into epic.application_user | " + err.Error())
	}
	defer e_stmt2.Close()
	_, err = e_stmt2.Exec(user.ID.String(), user.AppID.String())
	if err != nil {
		return nil, errors.New("Exec insert into epic.application_user | " + err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.New("Transaction Commit | " + err.Error())
	}
	return user, nil
}

func GetUser(id string) (User, error) {

  user := User{}
  return user, nil
}

func GetAllUsers() ([]User, error) {
  var users []User
  return users, nil
}



func UUID() string {
  return uuid.NewV4().String()
}

func hashPassword(password string, salt string) (string, error) {
  dk, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
  if err != nil {
    return "", errors.New("scrypt.Key | " + err.Error())
  }
  return base64.StdEncoding.EncodeToString(dk), nil
  //return string(dk), nil
}

func createSalt(size int) (string, error) {
    b := make([]byte, size)
    _, err := rand.Read(b)
    if err != nil {
      return "", errors.New("rand.Read | " + err.Error())
    }
//    return string(b), nil
      return base64.StdEncoding.EncodeToString(b), nil
}

func validatePassword(password string) error {
  l := len(password)
  if l < 6 {
    return fmt.Errorf("Password is only %v characters long.  It must be at least 6 characters long.", l)
  }
  return nil
}

func createToken(user *User) error {

  tokenExpires := time.Now().Add(time.Hour * 168).UTC()
  user.TokenExpires = &tokenExpires

  // Create the token
  //token := jwt.New(jwt.SigningMethodES384)
  token := jwt.New(jwt.SigningMethodHS256)

  // Set some claims
  token.Claims["id"] = user.ID.String()
  token.Claims["first_name"] = user.FirstName
  token.Claims["last_name"] = user.LastName
  token.Claims["email"] = user.Email
  token.Claims["username"] = user.Username
  token.Claims["admin"] = true
  token.Claims["exp"] = user.TokenExpires.Unix()
  token.Claims["iat"] = time.Now().UTC().Unix()
  tokenString, err := token.SignedString([]byte("secret"))
  if err != nil {
    return errors.New("SignedString | " + err.Error())
  }

  stmt, err := db.Prepare("update epic.user set token=$1, token_expires=$2 where id=$3")
  if err != nil {
  	return errors.New("Prepare | " + err.Error())
  }
  defer stmt.Close()
  _, err = stmt.Exec(tokenString, tokenExpires, user.ID.String())
  if err != nil {
  	return errors.New("Exec | " + err.Error())
  }
  user.Token = &tokenString
  return nil
}

func parseToken(requestToken string) error {

  stmt, err := db.Prepare("select epic_user.id, epic_user.first_name, epic_user.last_name, epic_user.email, epic_user.username, epic_user.password, epic_user.salt, epic_user.token, epic_user.token_expires, epic_user.private_key, epic_user.public_key from epic.user as epic_user where epic_user.token = $1")
  if err != nil {
    return errors.New("Prepare | " + err.Error())
  }
  defer stmt.Close()
  var user User
  err = stmt.QueryRow(requestToken).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.Salt, &user.Token, &user.TokenExpires, &user.PrivateKey, &user.PublicKey)
  if err != nil {
    return errors.New("QueryRow / Scan | " + err.Error())
  }
  token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    //if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
//    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//    }
    return []byte("secret"), nil
  })
  if err != nil {
    return errors.New("Parse | " + err.Error())
  }
  if token.Valid {
    err = validateToken(requestToken, user)
    if err != nil {
      return errors.New("validateToken | " + err.Error())
    }
    return nil
  } else {
    return errors.New("Valid? | Token not valid.")
  }
}

func validateToken(requestToken string, user User) error {
  if requestToken != *user.Token {
    return fmt.Errorf("Authentication token is invalid.")
  }
  if time.Now().After(*user.TokenExpires) {
    return fmt.Errorf("Authentication token has expired.")
  }
  return nil
}

func createKeys() (string, string, error) {
  publicKeyCurve := elliptic.P384()
  privateKey := new(ecdsa.PrivateKey)
  privateKey, err := ecdsa.GenerateKey(publicKeyCurve, rand.Reader)
  if err != nil {
    return "", "", errors.New("GenerateKey | " + err.Error())
  }
  var publicKey ecdsa.PublicKey
  publicKey = privateKey.PublicKey

  privateKeyString := fmt.Sprintf("%x", privateKey)
  publicKeyString := fmt.Sprintf("%x", publicKey)

  return privateKeyString, publicKeyString, nil
}
