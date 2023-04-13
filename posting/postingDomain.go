package posting

import "time"

type Posting struct {
	Title string
	Content string
	CreatedAt time.Time
}

type PostingRequest struct {
	Title string
	Content string
}

type PostingResponse struct {
	Title string
	Content string
	CreatedAt time.Time
}

func NewPosting(postingRequest PostingRequest) Posting {
	return Posting {
		Title: postingRequest.Title,
		Content: postingRequest.Content,
		CreatedAt: time.Now(),
	}
}

func NewPostingResponse(posting Posting) PostingResponse {
	return PostingResponse {
		Title: posting.Title,
		Content: posting.Content,
		CreatedAt: time.Now(),
	}
}