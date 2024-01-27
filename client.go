package pinboard

type Client struct {
	token string
	Notes NotesResource
	Posts PostsResource
	Tags  TagsResource
	User  UserResource
}

func New(token string) *Client {
	return &Client{
		token: token,
		Notes: NewNotesResource(token),
		Posts: NewPostsResource(token),
		Tags:  NewTagsResource(token),
		User:  NewUserResource(token),
	}
}
