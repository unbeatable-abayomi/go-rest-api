package http

import (
	"encoding/json"

	_"html"
	"net/http"
	"strconv"
	"unbeatable-abayomi/go-rest-api/internal/comment"

	"github.com/gorilla/mux"
	
)


//GetComment --retrieve a comment by ID
func (h *Handler)GetComment(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type","application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10,64)

	if err != nil{
		//fmt.Fprint(w, "Unable to parse UNIT From ID")
		sendErrorResponse(w, "Unable to parse UNIT From ID",err)
		return
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil{
		//fmt.Fprint(w, "Error Retrieving Comment By ID")
		sendErrorResponse(w, "Error Retrieving Comment By ID", err)
		return
	}
//    if err := json.NewEncoder(w).Encode(comment); err != nil{
// 	 panic(err)
//    }
   if err := sendOkResponse(w,comment); err != nil{
	 panic(err)
   }
	//fmt.Fprintf(w, "%+v", comment)
	
}

//GetAllComments - retrives all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		//fmt.Fprintf(w, "Failed to retrive all comments")
		sendErrorResponse(w, "Failed to retrive all comments", err)
		return
	}
	//fmt.Fprintf(w, "%+v", comments)
	// if err := json.NewEncoder(w).Encode(comments); err != nil{
	// 	panic(err)
	//   }
	if err := sendOkResponse(w,comments); err != nil{
		panic(err)
	  }
}
//PostComment -- adds a new comment
func (h *Handler) PostComment (w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	
	var comment comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		 //fmt.Fprintf(w, "Failed to decode JSON Body")
		 sendErrorResponse(w, "Failed to decode JSON Body", err)
		 return
	}

	// comment, err := h.Service.PostComment(comment.Comment{
	// 	Slug: "/",
		
	// })
	comment, err := h.Service.PostComment(comment)
	if err != nil{
		//fmt.Fprintf(w, "Failed to post new comment")
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}

	//fmt.Fprintf(w, "%+v", comment)

	// if err := json.NewEncoder(w).Encode(comment); err != nil{
	// 	panic(err)
	//   }
	if err := sendOkResponse(w,comment); err != nil{
		panic(err)
	  }

	  //fmt.Fprintf(w, "%+v", comment)

}



//UpdateComment --update comment by ID
func (h *Handler)UpdateComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	var comment comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		// fmt.Fprintf(w, "Failed to decode JSON Body")
		 sendErrorResponse(w, "Failed to decode JSON Body",err)
		 return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10,64)
	if err != nil{
		//fmt.Fprint(w, "Failed to parse unit from ID")
		sendErrorResponse(w,"Failed to parse unit from ID",err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
            if err != nil{
				//fmt.Fprint(w,"Failed to update comment")
				sendErrorResponse(w,"Failed to update comment",err)
				return
			}
		//	fmt.Fprintf(w,"%+v", comment)
		// if err := json.NewEncoder(w).Encode(comment); err != nil{
		// 	panic(err)
		//   }
		if err := sendOkResponse(w, comment); err != nil{
			panic(err)
		  }
	
  }
	
// 	comment, err := h.Service.UpdateComment(1, comment.Comment{
// 	Slug: "/new",
//   })
//   if err != nil{
// 	fmt.Fprintf(w, "Failed to Update comment")
//   }
//   fmt.Fprintf(w, "%+v", comment)
// }
//DeleteComment -- deletes a comment by ID
func (h *Handler)DeleteComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type","application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10,64)
	if err != nil{
		//fmt.Fprintf(w, "Failed to Parse unit from ID")
		sendErrorResponse(w, "Failed to Parse unit from ID", err)
		return
	}
	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		//fmt.Fprintf(w, "Failed to delete comment by ID")
		sendErrorResponse(w, "Failed to delete comment by ID",err)
		return
	}
	//fmt.Fprintf(w, "Successsfully deleted comment")

 if err = sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil{
	panic(err)
 }

	// if err := json.NewEncoder(w).Encode(Response{Message:"Comment Succesffully deleted"}); err != nil{
	// 	panic(err)
	// }
}