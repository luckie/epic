package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
  "log"
	"net/http"
  //"strings"
	"time"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/gorilla/mux"
  "github.com/gorilla/context"
  "github.com/satori/go.uuid"
)

type key int
const TokenKey key = 0
const AdminKey key = 1

type S3PutRequest struct {
  Bucket  string `json:"bucket"`
  Key     string `json:"key"`
  URL     string `json:"url"`
  Error     string `json:"error"`
}

type NewUUID struct {
	UUID 		string `json:"uuid"`
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func ListContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func ReadContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  contentID := vars["id"]
  e, err := NewestEntryForContentID(contentID)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	//w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func UpdateContentHandler(w http.ResponseWriter, r *http.Request) {
  verifyAdmin(w,r)
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
    log.Println(err)
  }
	w.WriteHeader(http.StatusCreated)
}

func ListTagsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To be implemented.\n")
}

func CreateTagHandler(w http.ResponseWriter, r *http.Request) {
  verifyAdmin(w,r)
  vars := mux.Vars(r)
  tag := vars["tag"]
  appID := vars["app-uuid"]
  err := CreateTag(tag, appID)
  if err != nil {
    log.Println(err)
  }
}

func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
  verifyAdmin(w,r)
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

func AssignTagToContentHandler(w http.ResponseWriter, r *http.Request) {
  verifyAdmin(w,r)
  vars := mux.Vars(r)
  contentID := vars["content-uuid"]
  tag := vars["tag"]

  err := TagContent(contentID, tag)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	w.WriteHeader(http.StatusOK)
}

func AssetUploadURLHandler(w http.ResponseWriter, r *http.Request) {
  verifyAdmin(w,r)
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
  svc := s3.New(session.New(),&aws.Config{
    Region: aws.String("us-east-1"),
    Credentials: credentials.NewEnvCredentials(),
    },)
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
  token := r.Header.Get("Authorization")
  err := Logout(token)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	w.WriteHeader(http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
  //verifyAdmin(w,r)
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

func AuthenticateTokenHandler(w http.ResponseWriter, r *http.Request) {
	//var token string
	token := r.Header.Get("Authorization")
	//if token, ok := context.Get(r, TokenKey).(string); ok {
	//
		if len(token) > 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	//} else {
	//	fmt.Println("no token")
	//	w.WriteHeader(http.StatusUnauthorized)
	//}
}

func NewUUIDHandler(w http.ResponseWriter, r *http.Request) {
	uuid := NewUUID{}
  uuid.UUID = UUID()
  if err := json.NewEncoder(w).Encode(uuid); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func AuthHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
	  //w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Accept", "application/json")
    authHeader := r.Header.Get("Authorization")
    if len(authHeader) > 0 {
      err := Authenticate(authHeader)
      if err != nil {
				//fmt.Println("AuthHandler FINAL ERR: " + err.Error())
				//w.WriteHeader(http.StatusUnauthorized)
      } else {
				w.Header().Set("Authorization", authHeader)
				context.Set(r, TokenKey, authHeader)
				context.Set(r, AdminKey, true)

			}

    }
		h.ServeHTTP(w, r)
  })
}

func verifyAdmin(w http.ResponseWriter, r *http.Request) {
  if context.Get(r, AdminKey).(bool) != true {
    w.WriteHeader(http.StatusUnauthorized)
  }
  return
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
