package posting

func CreatePosting(postingRequest PostingRequest) PostingResponse {
	posting := NewPosting(postingRequest)
	// TODO : save posting
	return NewPostingResponse(posting)
}