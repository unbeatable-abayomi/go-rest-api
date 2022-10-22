package http

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "html"
	"net/http"
	_ "strconv"
	"unbeatable-abayomi/go-rest-api/internal/comment"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service

type Handler struct{
	Router *mux.Router
	Service *comment.Service
}

type Response struct{
	Message string
	Error   string
}
//NewHandler -- returns a pointer to  a Handler
func NewHandler(service *comment.Service) *Handler{
	return &Handler{
		Service: service,
	}
}


//BaiscAuth a handy middleware function that will provide basic auth around specfic endpoints

func BasicAuth (original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
     return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic Auth Endpoint Hit")
		user,pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok{
			original(w,r)
		}else{
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(w,"not authorized", errors.New("Not Authorized")) 
		}
		
	 }
}

//LoggingMiddleware adds Middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler{
     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path": r.URL.Path,
			}).Info("Handled Request")
		next.ServeHTTP(w,r)
	 })
}
// SetupRoutes -- sets up all the routes for our application
func (h *Handler) SetupRoutes(){
	//fmt.Println("Setting Up Routes")
	log.Info("Setting Up Routes")
	h.Router= mux.NewRouter()
    h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	//h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment", BasicAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	//h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")
	//h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")
	
	
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		//fmt.Fprint(w, "I am alive")
		// w.Header().Set("Content-Type","application/json; charset=UTF-8")
		// w.WriteHeader(http.StatusOK)
		// if err := json.NewEncoder(w).Encode(Response{Message:"I am Alive"}); err != nil{
		// 	panic(err)
		// }
		if err := sendOkResponse(w,Response{Message:"I am Alive"}); err != nil{
			panic(err)
		}
	})

}

// //GetComment --retrieve a comment by ID
// func (h *Handler)GetComment(w http.ResponseWriter, r *http.Request){
// 	// w.Header().Set("Content-Type","application/json; charset=UTF-8")
// 	// w.WriteHeader(http.StatusOK)
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	i, err := strconv.ParseUint(id, 10,64)

// 	if err != nil{
// 		//fmt.Fprint(w, "Unable to parse UNIT From ID")
// 		sendErrorResponse(w, "Unable to parse UNIT From ID",err)
// 		return
// 	}
// 	comment, err := h.Service.GetComment(uint(i))
// 	if err != nil{
// 		//fmt.Fprint(w, "Error Retrieving Comment By ID")
// 		sendErrorResponse(w, "Error Retrieving Comment By ID", err)
// 		return
// 	}
// //    if err := json.NewEncoder(w).Encode(comment); err != nil{
// // 	 panic(err)
// //    }
//    if err := sendOkResponse(w,comment); err != nil{
// 	 panic(err)
//    }
// 	//fmt.Fprintf(w, "%+v", comment)
	
// }

// //GetAllComments - retrives all comments from the comment service
// func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request){
// 	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
// 	//w.WriteHeader(http.StatusOK)
// 	comments, err := h.Service.GetAllComments()
// 	if err != nil {
// 		//fmt.Fprintf(w, "Failed to retrive all comments")
// 		sendErrorResponse(w, "Failed to retrive all comments", err)
// 		return
// 	}
// 	//fmt.Fprintf(w, "%+v", comments)
// 	// if err := json.NewEncoder(w).Encode(comments); err != nil{
// 	// 	panic(err)
// 	//   }
// 	if err := sendOkResponse(w,comments); err != nil{
// 		panic(err)
// 	  }
// }
// //PostComment -- adds a new comment
// func (h *Handler) PostComment (w http.ResponseWriter, r *http.Request){
// 	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
// 	//w.WriteHeader(http.StatusOK)
	
// 	var comment comment.Comment

// 	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
// 		 //fmt.Fprintf(w, "Failed to decode JSON Body")
// 		 sendErrorResponse(w, "Failed to decode JSON Body", err)
// 		 return
// 	}

// 	// comment, err := h.Service.PostComment(comment.Comment{
// 	// 	Slug: "/",
		
// 	// })
// 	comment, err := h.Service.PostComment(comment)
// 	if err != nil{
// 		//fmt.Fprintf(w, "Failed to post new comment")
// 		sendErrorResponse(w, "Failed to post new comment", err)
// 		return
// 	}

// 	//fmt.Fprintf(w, "%+v", comment)

// 	// if err := json.NewEncoder(w).Encode(comment); err != nil{
// 	// 	panic(err)
// 	//   }
// 	if err := sendOkResponse(w,comment); err != nil{
// 		panic(err)
// 	  }

// 	  //fmt.Fprintf(w, "%+v", comment)

// }



// //UpdateComment --update comment by ID
// func (h *Handler)UpdateComment(w http.ResponseWriter, r *http.Request){
// 	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
// 	//w.WriteHeader(http.StatusOK)
// 	var comment comment.Comment

// 	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
// 		// fmt.Fprintf(w, "Failed to decode JSON Body")
// 		 sendErrorResponse(w, "Failed to decode JSON Body",err)
// 		 return
// 	}
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	commentID, err := strconv.ParseUint(id, 10,64)
// 	if err != nil{
// 		//fmt.Fprint(w, "Failed to parse unit from ID")
// 		sendErrorResponse(w,"Failed to parse unit from ID",err)
// 		return
// 	}

// 	comment, err = h.Service.UpdateComment(uint(commentID), comment)
//             if err != nil{
// 				//fmt.Fprint(w,"Failed to update comment")
// 				sendErrorResponse(w,"Failed to update comment",err)
// 				return
// 			}
// 		//	fmt.Fprintf(w,"%+v", comment)
// 		// if err := json.NewEncoder(w).Encode(comment); err != nil{
// 		// 	panic(err)
// 		//   }
// 		if err := sendOkResponse(w, comment); err != nil{
// 			panic(err)
// 		  }
	
//   }
	
// // 	comment, err := h.Service.UpdateComment(1, comment.Comment{
// // 	Slug: "/new",
// //   })
// //   if err != nil{
// // 	fmt.Fprintf(w, "Failed to Update comment")
// //   }
// //   fmt.Fprintf(w, "%+v", comment)
// // }
// //DeleteComment -- deletes a comment by ID
// func (h *Handler)DeleteComment(w http.ResponseWriter, r *http.Request){
// 	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
// 	//w.WriteHeader(http.StatusOK)
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	commentID, err := strconv.ParseUint(id, 10,64)
// 	if err != nil{
// 		//fmt.Fprintf(w, "Failed to Parse unit from ID")
// 		sendErrorResponse(w, "Failed to Parse unit from ID", err)
// 		return
// 	}
// 	err = h.Service.DeleteComment(uint(commentID))
// 	if err != nil {
// 		//fmt.Fprintf(w, "Failed to delete comment by ID")
// 		sendErrorResponse(w, "Failed to delete comment by ID",err)
// 		return
// 	}
// 	//fmt.Fprintf(w, "Successsfully deleted comment")

//  if err = sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil{
// 	panic(err)
//  }

// 	// if err := json.NewEncoder(w).Encode(Response{Message:"Comment Succesffully deleted"}); err != nil{
// 	// 	panic(err)
// 	// }
// }

//just a test

func (d *Handler) GetAbayomi(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Successsfully deleted comment")
}



func sendOkResponse (w http.ResponseWriter, resp interface{}) error{
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)

}


func sendErrorResponse(w http.ResponseWriter, message string, err error){
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil{
		panic(err)
	}
}
