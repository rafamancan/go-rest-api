package main

import (
  "log"
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "not found"}`))
}

func params(w http.ResponseWriter, r *http.Request){
  pathParams := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")

  userID := -1
  var err error
  if val, ok := pathParams["useID"]; ok {
    userID, _ := strconv.Atoi(val)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(fmt.Sprintf(`{"message": "need an user, receive": %d}`, userID)))
      return
    }
  }

  commentID := -1
  if val, ok := pathParams["commentID"]; ok {
    commentID, _ := strconv.Atoi(val)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(fmt.Sprintf(`{"message": "need a comment, receive": %d}`, commentID)))
      return
    }
  }


  query := r.URL.Query()
  location := query.Get("location")

  w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "localtion": %s}`, userID, commentID, location)))
}

func main() {
  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()
  r.HandleFunc("", get).Methods(http.MethodGet)
  r.HandleFunc("", post).Methods(http.MethodPost)
  r.HandleFunc("", put).Methods(http.MethodPut)
  r.HandleFunc("", delete).Methods(http.MethodDelete)
  r.HandleFunc("", notFound)

  api.HandleFunc("/user/{userId}/comment/{commentID}", params).Methods(http.MethodGet)

  log.Fatal(http.ListenAndServe(":8080", r))
}
