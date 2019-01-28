package connection

type Track struct {
	ID         string   `json:"id,omitempty" bson:"id"`
	Name       string   `json:"name,omitempty" bson:"name"`
	Artist     []Artist `json:"artist,omitempty" bson:"artist"`
	Album      Album    `json:"album,omitempty" bson:"album"`
	Popularity int      `json:"popularity,omitempty" bson:"popularity"`
	Duration   int      `json:"duration,omitempty" bson:"duration"`
}

type Artist struct {
	ID    string `json:"id,omitempty" bson:"id"`
	Name  string `json:"name,omitempty" bson:"name"`
	Image Image  `json:"image,omitempty" bson:"image"`
}
type Album struct {
	ID     string   `json:"id,omitempty" bson:"id"`
	Name   string   `json:"name,omitempty" bson:"name"`
	Artist []Artist `json:"artist,omitempty" bson:"artist"`
	Cover  Image    `json:"cover,omitempty" bson:"cover"`
}

type Image struct {
	Small  string `json:"64,omitempty" bson:"small"`
	Medium string `json:"300,omitempty" bson:"medium"`
	Large  string `json:"640,omitempty" bson:"large"`
}

type SongMessage struct {
	Type    string
	RoomKey string
	Track   Track
}
