package modules

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// constants
const OWNER_ID = int64(1833850637)
const YOUTUBE_API_KEY = "AIzaSyAEz0eRkbsEE7TrHGKmd_iXh4AmYJlMKDs"
const ResolveURL = "https://707c80624779.up.railway.app/username?username="
const TelegraphToken = "11b5d2394a13e4dc286ab2c29a7df8a2b02844c76e7cf7a7300d6e5420fd"

// vars
var BOT_USERNAME = b.Me.Username
var BOT_NAME = b.Me.FirstName
var BOT_ID = b.Me.ID
var StartTime = time.Now()
var Client = http.Client{Timeout: time.Second * 10}

//types
type FakeID struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City     string `json:"city"`
			State    string `json:"state"`
			Country  string `json:"country"`
			Postcode int    `json:"postcode"`
		} `json:"location"`
		Email string `json:"email"`
		Dob   struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"dob"`
		Phone string `json:"phone"`
		Cell  string `json:"cell"`
		Nat   string `json:"nat"`
	} `json:"results"`
}

type InstSearch struct {
	Users []struct {
		Position int `json:"position"`
		User     struct {
			Pk                         string      `json:"pk"`
			Username                   string      `json:"username"`
			FullName                   string      `json:"full_name"`
			IsPrivate                  bool        `json:"is_private"`
			ProfilePicURL              string      `json:"profile_pic_url"`
			ProfilePicID               string      `json:"profile_pic_id"`
			IsVerified                 bool        `json:"is_verified"`
			FollowFrictionType         int         `json:"follow_friction_type"`
			HasAnonymousProfilePicture bool        `json:"has_anonymous_profile_picture"`
			HasHighlightReels          bool        `json:"has_highlight_reels"`
			LatestReelMedia            int         `json:"latest_reel_media"`
			LiveBroadcastID            interface{} `json:"live_broadcast_id"`
			ShouldShowCategory         bool        `json:"should_show_category"`
			Seen                       int         `json:"seen"`
		} `json:"user,omitempty"`
	} `json:"users"`
}

type PintrestResp struct {
	ResourceResponse struct {
		Status string `json:"status"`
		Data   struct {
			Results []struct {
				Objects []struct {
					RecentPinImages struct {
						One92X []struct {
							URL           string `json:"url"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							DominantColor string `json:"dominant_color"`
						} `json:"192x"`
					} `json:"recent_pin_images"`
				} `json:"objects"`
			} `json:"results"`
		} `json:"data"`
	} `json:"resource_response"`
}

type UrbanDict struct {
	List []struct {
		Definition string `json:"definition"`
		ThumbsUp   int    `json:"thumbs_up"`
		Author     string `json:"author"`
		Word       string `json:"word"`
		Example    string `json:"example"`
		ThumbsDown int    `json:"thumbs_down"`
	} `json:"list"`
}

type Bin struct {
	Number struct {
		Length int `json:"length"`
	} `json:"number"`
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Prepaid bool   `json:"prepaid"`
	Country struct {
		Numeric  string `json:"numeric"`
		Alpha2   string `json:"alpha2"`
		Name     string `json:"name"`
		Emoji    string `json:"emoji"`
		Currency string `json:"currency"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"bank"`
}

type TGraph struct {
	Ok     bool `json:"ok"`
	Result struct {
		URL string `json:"url"`
	} `json:"result"`
}

//cookies
var InstagramCookies = `mid=YheZOwALAAENlSyZzIkARG87nhp0; ig_did=BAC63AEF-60D7-43E7-8C79-55913556F7AF; ig_nrcb=1; fbm_124024574287414=base_domain=.instagram.com; fbsr_124024574287414=Nr4qLfABjfAAcCABhrzg1iSfA-ZQTmss9Mbb5epyY6o.eyJ1c2VyX2lkIjoiMTAwMDQwNzY1MTU3ODcxIiwiY29kZSI6IkFRRG9rOGlqZk5SZkJZamF0bzM0anYyUjhTeWVDQlR1dEMzb0Y3M3R3NU9EM2ZpcG8tcy1QeG1LZDVyeVRiOGYxMmtZOHc1dHoyLXQ2NHhpOUsyaWh3WC1tRnFOQWRwU0Y4UTBpX1Z1S1NKZ2tYeEF1ZUwxUkhwSWJTQVNmSVhFb25lRGJaVE56VW44VXYyV3N0dmVlemJLUzQ1bzg0N2lRclRhQS0tZTJSRS1obDBHcEZyclh6RWVrVm11UmZDWXU3NmF0U2hnbC1rOEVQclhLNkZJN1hXc0hGRHBqeF9tNEh5UXZBbmlaNGlWUVFnQlJra3VhZDJoNHJzY2tRSVlQSlRBLWV5Q19EQWFxMnItTm1zWk9nMjNWRnFLQlhvVjBSUXhmQi1GVTYzaElsSzZzVzZRdFlqVktydnYwTFY5angxUDQ5MHZKY1JKTVJHX0lQWFVwVU9sIiwib2F1dGhfdG9rZW4iOiJFQUFCd3pMaXhuallCQUxsYUNQMGE0Y0g2RjlMRk1jOEVaQWJqNjV4Ymo5dFNsVElyU1JSSTJtWkJlajE2QzhRSWd5N0ZuV2w0bWxjeWE3V1FHdXJrbGVxcXdkajF0NnpKWkN3Nkg1cDZJUTFORVpBWUFRZkxXZFVaQ1pDY2o4WkI1VmRPRUJzNzdVUkZnMHpTQ0lFM2h2MVZLQzByR0Nydjk5MHlaQXNkbjlRUnZxbHlwZTZvbDRsV3dhU0RLd1RKajFVWkQiLCJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImlzc3VlZF9hdCI6MTY0NTg4NzUwM30; csrftoken=5ZeeJnVKxCVy7e6usH0Ixmch65CBEaqR; ds_user_id=52090581237; sessionid=52090581237%3ASd6yOszWExqzwt%3A12; shbid="1118\05452090581237\0541677423521:01f7d101bfd6414ceac2be62ecaf63a9f24173cc51f4b24ffeffaba6b6eb458435159c3a"; shbts="1645887521\05452090581237\0541677423521:01f7db2111c0f4b7c91e5315961e2cd9cce6b0bdcfc17dc57ae36a24572dfec402feb7d8"; fbsr_124024574287414=Nr4qLfABjfAAcCABhrzg1iSfA-ZQTmss9Mbb5epyY6o.eyJ1c2VyX2lkIjoiMTAwMDQwNzY1MTU3ODcxIiwiY29kZSI6IkFRRG9rOGlqZk5SZkJZamF0bzM0anYyUjhTeWVDQlR1dEMzb0Y3M3R3NU9EM2ZpcG8tcy1QeG1LZDVyeVRiOGYxMmtZOHc1dHoyLXQ2NHhpOUsyaWh3WC1tRnFOQWRwU0Y4UTBpX1Z1S1NKZ2tYeEF1ZUwxUkhwSWJTQVNmSVhFb25lRGJaVE56VW44VXYyV3N0dmVlemJLUzQ1bzg0N2lRclRhQS0tZTJSRS1obDBHcEZyclh6RWVrVm11UmZDWXU3NmF0U2hnbC1rOEVQclhLNkZJN1hXc0hGRHBqeF9tNEh5UXZBbmlaNGlWUVFnQlJra3VhZDJoNHJzY2tRSVlQSlRBLWV5Q19EQWFxMnItTm1zWk9nMjNWRnFLQlhvVjBSUXhmQi1GVTYzaElsSzZzVzZRdFlqVktydnYwTFY5angxUDQ5MHZKY1JKTVJHX0lQWFVwVU9sIiwib2F1dGhfdG9rZW4iOiJFQUFCd3pMaXhuallCQUxsYUNQMGE0Y0g2RjlMRk1jOEVaQWJqNjV4Ymo5dFNsVElyU1JSSTJtWkJlajE2QzhRSWd5N0ZuV2w0bWxjeWE3V1FHdXJrbGVxcXdkajF0NnpKWkN3Nkg1cDZJUTFORVpBWUFRZkxXZFVaQ1pDY2o4WkI1VmRPRUJzNzdVUkZnMHpTQ0lFM2h2MVZLQzByR0Nydjk5MHlaQXNkbjlRUnZxbHlwZTZvbDRsV3dhU0RLd1RKajFVWkQiLCJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImlzc3VlZF9hdCI6MTY0NTg4NzUwM30; rur="ATN\05452090581237\0541677423705:01f708822e68dd8c2fb6172f0468b08692c147823db1322c292b479ee26370ecd8a64267"`
var PinterestCookies = `_pinterest_referrer=https://www.google.com/; csrftoken=252cb8f31a7abce1ad3bd54bdb09e212; _routing_id="8eb4877c-bd2c-42ec-86ee-cbf941dc3e7a"; sessionFunnelEventLogged=1; g_state={"i_l":0}; _auth=1; _pinterest_sess=TWc9PSZaY1dZak8xbThQbHVoK2d3REJibDcvS0x6STY4YWQ5eXUwQjRRQnRHU25MdXUveExwMFVhb05kTFV2YWJEd3U3b1YvVGRKUXIyOWhNTXBNTVEwT0RnODB1SzBBY1RBV09MNzNHUXExTUZua2s4aWhvZnY1d0VKME9GTWt3eEozK3dEYWJ1WUF3T2JpVEd2dGFSMGwwVkVuRDA0RDVrZVJBaERrS0czTmI4d01wSWZmQktIWVMzODFhTm5EV0l6QmszNjhtT0Zqdlg3b0ZQUGcxdlNxQXJGeG5qYUNxK2dqdk1sTVo0alQ3L3V3OElpV0NKVklNZit2S1J6VHpIWW1MYkNXazR6QUQyc3QvM0FnMldPSGVIZXdNd2RhVUlhNGJLMC9KSHZ2empzVGQ3eGdadnZ3MS9YWjljc3FnL3VjSms5L3BEd0VvWHZJem1YWWp3M3Z2UjAyODVuMEd4alQ3aEFMQlVTSnpvNDMweXdtVlgrZTlKZVhCQTBSbTl4VVhPMXdhU3p5bkdub25OMXVqRk84RTd4TmZ0UmdKYmpWbXlVdDVqMzFiOWo2Y0xmd3A3SzdnSG1SWnlFa3hRbzRvQnc4QVdJK0ptc08zRmxSbXM1ckVBSFFtK1dHckhvMnRQMW9aUXhRaDYvUFpCUDJhZkcvUU1MZzNWQk05b1loa29zNUY3MnJwZHVybTV5R294NmljY1RwYkdLTDlONkQweDRzM09ZM1NMNmdHMnoxRVhFakJMNEdtVHVDY3Q4M3VaZ2I0Q201U0lZcmZGUGdQcFJ3c3lMZ0I5SS9wclFnWm5SWTdoUDJxVms2Tm5ReEdnMDBXMDI4Qkd0VlVnQ3JXWmtTemVXaHZncndRTFZVUUZrUHVCdmJjV05PUEdpdU5YV2FVVE1sajUzbVZXdzJPeWZ6NVNjak1meVhXNG15K1hUTjFnQm90YkpBTEtFRHZ6YndsQXdXYkFWcG9ZTjczUHlIdXkvbUxOblFsMTlKU2J3MVhnL1lkdDJBVzdBdW0rVjB3cEluZDAwN212UUR1SUdXdTBnPT0mdjV6bHFEZ2hrd2cvTUU0VGJSYllSZ2NNcFZJPQ==; _b="AWE8wCyJ3SZDna0mrM3dRqAlBU5cT0RDGYpjd4nzFS++g5yM3ahjTivMA7KKR5J7XiE="; cm_sub=denied`

var notes_help = "✨ Here is the help for **Notes:**\n**Command for Members**\n**->** `/get notename`: get the note with this notename\n**-** #notename: same as /get\n**->** `/notes`: list all saved notes in this chat\n**Command for Admins**\n**->** `/save notename notedata`: saves notedata as a note with name notename, reply to a message or document to save it\n**->** `/clear notename`: clear note with this name\n**->** `/privatenote on/yes/off/no`: whether or not to send the note in PM. Write del besides on/off to delete hashtag message on group.\n**Note**\n **-** Only admins can use This module\n **-** To save a document (like photo, audio, etc.), reply to a document or media then type /save\n **-** Need help for parsing text? Check /markdownhelp\nSave data for future users with notes!\nNotes are great to save random tidbits of information; a phone number, a nice gif, a funny picture - anything!\nAlso you can save a text/document with buttons, you can even save it in here."

var help = bson.M{"notes": notes_help}

var PLUGIN_LIST = []string{"admin", "bans", "chatbot", "feds", "greetings", "inline", "lock", "misc", "notes", "pin", "stickers", "warns"}

var help_caption = `
Hey!, My name is Mika.
I am a group management bot, here to help you get around and keep the order in your groups!
I have lots of handy features.
So what are you waiting for?
Add me in your groups and give me full rights to make me function well.`

var COUNTRY_CODES = bson.M{"Australia": "AU", "Brazil": "BR", "Canada": "CA", "Switzerland": "CH", "Germany": "DE", "France": "FR", "Netherlands": "NL", "Russia": "RU", "Spain": "ES", "Turkey": "TR", "United Kingdom": "GB", "United States": "US", "SK": "Sweden"}
var CODE_C = []string{"AU", "BR", "CA", "CH", "DE", "DK", "ES", "FI", "FR", "GB", "IE", "IR", "NO", "NL", "NZ", "TR", "US"}
var AFK_STR = []string{
	"<b>%s</b> is here!",
	"<b>%s</b> is back!",
	"<b>%s</b> is now in the chat!",
	"<b>%s</b> is awake!",
	"<b>%s</b> is back online!",
	"<b>%s</b> is finally here!",
	"Welcome back! <b>%s</b>",
	"Where is <b>%s</b>?\nIn the chat!",
	"Pro <b>%s</b>, is back alive!",
}

var stripe_1 = `
<b>⌥ Gateway ✑ %s</b>
<b>CC ✑</b> <code>%s|%s|%s|%s</code>
<b>⌥ Status ✑ %s</b> %s %s
<b>⌥ Response ✑</b> %s

<b>⎋ Card Details: %s</b>
<b>⎋ Time: %ds</b>
<b>✁Checked by</b> <b>%s</b> [%s]
`
var CNT = []string{"locks", "flood", "filters", "get", "notes", "saved", "adminlist", "info", "warns", "rules", "approval"}
