package hnclient

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

type Client struct {
	AgentCreator func() *gorequest.SuperAgent
}

func (self *Client) CreateAgent() *gorequest.SuperAgent {
	if self.AgentCreator != nil {
		return self.AgentCreator()
	}
	return gorequest.New()
}
func (self *Client) GetItem(id int) (item Item, err error) {
	_, _, errs := self.CreateAgent().Get(fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", id)).EndStruct(&item)
	if len(errs) != 0 {
		err = errs[0]
	}
	return
}

func (self *Client) GetUser(name string) (user User, err error) {
	_, _, errs := self.CreateAgent().Get(fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%v.json", name)).EndStruct(&user)
	if len(errs) != 0 {
		err = errs[0]
	}
	return
}
func (self *Client) GetMaxItem() (id int, err error) {
	_, b, errs := self.CreateAgent().Get("https://hacker-news.firebaseio.com/v0/maxitem.json").End()
	if len(errs) != 0 {
		err = errs[0]
	}
	id, err = strconv.Atoi(b)
	return
}

// https://hacker-news.firebaseio.com/v0/item/126809.json
// https://hacker-news.firebaseio.com/v0/user/jl.json
// https://hacker-news.firebaseio.com/v0/maxitem.json
// https://hacker-news.firebaseio.com/v0/topstories.json

type Item struct {
	Id          uint   `json:"id,omitempty"`          //The item's unique id.
	Deleted     bool   `json:"deleted,omitempty"`     //true if the item is deleted.
	Type        string `json:"type,omitempty"`        //The type of item. One of "job", "story", "comment", "poll", or "pollopt".
	By          string `json:"by,omitempty"`          //The username of the item's author.
	Time        int64  `json:"time,omitempty"`        //Creation date of the item, in Unix Time.
	Text        string `json:"text,omitempty"`        //The comment, story or poll text. HTML.
	Dead        bool   `json:"dead,omitempty"`        //true if the item is dead.
	Parent      uint   `json:"parent,omitempty"`      //The comment's parent: either another comment or the relevant story.
	Poll        uint   `json:"poll,omitempty"`        //The pollopt's associated poll.
	Kids        []uint `json:"kids,omitempty"`        //The ids of the item's comments, in ranked display order.
	Url         string `json:"url,omitempty"`         //The URL of the story.
	Score       int    `json:"score,omitempty"`       //The story's score, or the votes for a pollopt.
	Title       string `json:"title,omitempty"`       //The title of the story, poll or job.
	Parts       []int  `json:"parts,omitempty"`       //A list of related pollopts, in display order.
	Descendants int    `json:"descendants,omitempty"` //In the case of stories or polls, the total comment count.
}

type User struct {
	Id        string `json:"id,omitempty"`        //The user's unique username. Case-sensitive. Required.
	Delay     int    `json:"delay,omitempty"`     //Delay in minutes between a comment's creation and its visibility to other users.
	Created   int64  `json:"created,omitempty"`   //Creation date of the user, in Unix Time.
	Karma     int    `json:"karma,omitempty"`     //The user's karma.
	About     string `json:"about,omitempty"`     //The user's optional self-description. HTML.
	Submitted []uint `json:"submitted,omitempty"` //List of the user's stories, polls and comments.
}
