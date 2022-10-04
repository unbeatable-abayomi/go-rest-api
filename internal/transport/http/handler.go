package http

import (
	"fmt"
	"net/http"
	"strconv"
	"unbeatable-abayomi/go-rest-api/internal/comment"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service

type Handler struct{
	Router *mux.Router
	Service *comment.Service
}
//NewHandler -- returns a pointer to  a Handler
func NewHandler(service *comment.Service) *Handler{
	return &Handler{
		Service: service,
	}
}
// SetupRoutes -- sets up all the routes for our application
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router= mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "I am alive")
	})

}
//GetComment --retrieve a comment by ID
func (h *Handler)GetComment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10,64)

	if err != nil{
		fmt.Fprint(w, "Unable to parse UNIT From ID")
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil{
		fmt.Fprint(w, "Error Retrieving Comment By ID")
	}
  
	fmt.Fprintf(w, "%+v", comment)
	
}

//GetAllComments - retrives all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request){
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrive all comments")
	}
	fmt.Fprintf(w, "%+v", comments)
}
//PostComment -- adds a new comment
func (h *Handler) PostComment (w http.ResponseWriter, r *http.Request){
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
		
	})
	if err != nil{
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comment)

}



//UpdateComment --update comment by ID
func (h *Handler)UpdateComment(w http.ResponseWriter, r *http.Request){
  comment, err := h.Service.UpdateComment(1, comment.Comment{
	Slug: "/new",
  })
  if err != nil{
	fmt.Fprintf(w, "Failed to Update comment")
  }
  fmt.Fprintf(w, "%+v", comment)
}
//DeleteComment -- deletes a comment by ID
func (h *Handler)DeleteComment(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10,64)
	if err != nil{
		fmt.Fprintf(w, "Failed to Parse unit from ID")
	}
	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by ID")
	}
	fmt.Fprintf(w, "Successsfully deleted comment")
}

//just a test

func (d *Handler) GetAbayomi(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Successsfully deleted comment")
}