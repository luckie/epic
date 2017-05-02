package epic

import (
	"context"
	"encoding/json"
	//"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"time"
)

//type key int
//const TokenKey key = 0
//const AdminKey key = 1

type S3PutRequest struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	URL    string `json:"url"`
	Error  string `json:"error"`
}

type UUID struct {
	UUID string `json:"uuid"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func ListContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func CreateContentReservationHandler(w http.ResponseWriter, r *http.Request) {

	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c := Content{}
	if err := json.Unmarshal(body, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = CreateContentReservation(&c)
	if err != nil {
		c.Error = err
	}
	id := ID{}
	id.ID = c.ID.String()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contentID := vars["id"]
	e, err := NewestEntryForContentID(contentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e := Entry{}
	if err := json.Unmarshal(body, &e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.ID = uuid.NewV4()
	e.Timestamp = time.Now()
	err = CreateEntryForContentID(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListTagsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	tag := vars["tag"]
	appID := vars["app-uuid"]
	err = CreateTag(tag, appID)
	if err != nil {
		log.Println(err)
	}
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "To be implemented.\n")
}

func ReadAllContentForTagHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	appID := vars["app-uuid"]
	e, err := AllContentForTag(tag, appID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadNewestLocalizedContentEntriesForTagHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	locale := vars["locale"]
	appID := vars["app-uuid"]
	e, err := NewestLocalizedContentEntriesForTag(tag, locale, appID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AssignTagToContentHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	contentID := vars["content-uuid"]
	tag := vars["tag"]
	AppID := vars["AppID"]

	err = TagContent(contentID, tag, AppID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {

	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app := App{}
	if err := json.Unmarshal(body, &app); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	appID, err := CreateApplication(app.Name, app.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		uuid := UUID{
			UUID: appID.String(),
		}
		//
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(uuid); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AssetUploadURLHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	put := S3PutRequest{}
	if err := json.Unmarshal(body, &put); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	svc := s3.New(session.New(), &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewEnvCredentials(),
	})
	por, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(put.Bucket),
		Key:    aws.String(put.Key),
	})
	url, err := por.Presign(15 * time.Minute)
	if err != nil {
		put.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		put.URL = url
		w.WriteHeader(http.StatusOK)
	}
	if err := json.NewEncoder(w).Encode(&put); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := User{}
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userPtr, err := Login(&user)
	if err != nil {
		if err.Error() == "QueryRow / Scan select from user | sql: no rows in result set" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		user = *userPtr
		w.Header().Set("Authorization", *user.Token)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	userIDi, err := fromContext("UserID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID, ok := userIDi.(string)
	if !ok {
		http.Error(w, "userID type assertion from interface{} to string failed.", http.StatusInternalServerError)
		return
	}
	err = Logout(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := User{}
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userPtr, err := CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		user = *userPtr
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("id")
	appIDi, err := fromContext("AppID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	appID, ok := appIDi.(string)
	if !ok {
		http.Error(w, "userID type assertion from interface{} to string failed.", http.StatusInternalServerError)
		return
	}


	if id != "" {
		u, err := GetUser(id, appID)
		user := *u
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		users, err := GetAllUsers(appID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			if err := json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userID := r.URL.Query().Get("user-id")
	appID := r.URL.Query().Get("app-id")
	if userID == "" {
		userIDi, err := fromContext("UserID", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		uID, ok := userIDi.(string)
		if !ok {
			http.Error(w, "userID type assertion from interface{} to string failed.", http.StatusInternalServerError)
			return
		}
		userID = uID
	}
	if appID == "" {
		appIDi, err := fromContext("AppID", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		aID, ok := appIDi.(string)
		if !ok {
			http.Error(w, "appID type assertion from interface{} to string failed.", http.StatusInternalServerError)
			return
		}
		appID = aID
	}
	err = DeleteUser(userID, appID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AuthenticateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, err := Authenticate(token)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Accept", "application/json")
		token := r.Header.Get("Authorization")

		if token != "" {
			claims, err := Authenticate(token)
			if err == nil {
				var epicCtxKey epicContextKey
				epicCtx := epicContext{}
				epicCtx.m = make(map[string]interface{}, 9)
				epicCtx.Set("Token", token)
				epicCtx.Set("UserID", claims["UserID"])
				epicCtx.Set("FirstName", claims["FirstName"])
				epicCtx.Set("LastName", claims["LastName"])
				epicCtx.Set("Email", claims["Email"])
				epicCtx.Set("Username", claims["Username"])
				epicCtx.Set("AppID", claims["AppID"])
				epicCtx.Set("TokenExpires", claims["exp"])
				epicCtx.Set("TokenCreated", claims["iat"])
				ctx := context.WithValue(r.Context(), epicCtxKey, epicCtx)
				h.ServeHTTP(w, r.WithContext(ctx))
			} else {

				http.Error(w, "AuthHandler | Error returned from Authenticate(token).  " + err.Error(), http.StatusUnauthorized)
			}
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func verifyAdmin(r *http.Request) error {
	// Temp using AppID below from within the tokenID as a proxy for validating a user is an admin.
	appIDi, err := fromContext("AppID", r)
	if err != nil {
		return errors.Wrap(err, "verifyAdmin() | Error from fromContext()")
	}
	appID, ok := appIDi.(string)
	if !ok {
		return errors.New("verifyAdmin() | appID is used to verify admin, but type assertion from interface{} to string failed.")
	}
	// This test below is insuffient. Needs real test.
	if appID == "" {
		return errors.New("verifyAdmin() | appID is used to verify admin, but appID is blank.")
	}
	return nil
}

func NewUUIDHandler(w http.ResponseWriter, r *http.Request) {
	uuid := UUID{}
	uuid.UUID = NewUUID()
	if err := json.NewEncoder(w).Encode(uuid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	err := verifyAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userIDi, err := fromContext("UserID", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID, ok := userIDi.(string)
	if !ok {
		http.Error(w, "userID type assertion from interface{} to string failed.", http.StatusInternalServerError)
		return
	}
	if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err = UpdatePassword(user.Password, userID)
	} else {
		err = UpdatePassword(user.Password, user.ID.String())
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(struct{}{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	//appID := r.URL.Query().Get("app-id")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var reset Reset
	if err := json.Unmarshal(body, &reset); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ResetPassword(reset.Email, reset.AppID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)


	/*
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/
}

func UserCryptoBootstrapHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	crypto := UserCryptoBootstrap{}
	if err := json.Unmarshal(body, &crypto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = validatePassword(crypto.PlainText)
	if err != nil {
		//
	}
	salt, err := createSalt(SALT_BYTES)
	if err != nil {
		//
	}
	hashedPassword, err := hashPassword(crypto.PlainText, crypto.Salt)
	if err != nil {
		//
	}
	privateKey, publicKey, err := createKeys()
	if err != nil {
		//
	}
	crypto.Hash = hashedPassword
	crypto.Salt = salt
	crypto.PrivateKey = privateKey
	crypto.PublicKey = publicKey
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&crypto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
UTILITIES
*/

func fromContext(key string, r *http.Request) (interface{}, error) {
	var epicCtxKey epicContextKey
	epicCtx, ok := r.Context().Value(epicCtxKey).(epicContext)
	if !ok {
		return nil, errors.New("fromContext() | Unable to get value out of Request.Context() and get epicContext struct using type assertion.")
	}
	val := epicCtx.Get(key)
	if val != nil {
		return val, nil
	} else {
		return nil, errors.New("fromContext() | Unable to get non-nil interface{} value out of epicContext struct.")
	}
}

type epicContextKey string

type epicContext struct {
	m map[string]interface{}
}

func (ec epicContext) Get(key string) interface{} {
	return ec.m[key]
}

func (ec epicContext) Set(key string, value interface{}) {
	ec.m[key] = value
}
