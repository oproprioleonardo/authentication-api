package query

type UserProfileSearch struct {
	UUID       string `bson:"uuid,omitempty" query:"uuid,omitempty"`
	Name       string `bson:"name,omitempty" query:"name,omitempty"`
	OnlineMode bool   `bson:"onlineMode" query:"onlineMode"`
	Email      string `bson:"email,omitempty" query:"email,omitempty"`
	Phone      string `bson:"query,omitempty" query:"phone,omitempty"`
}
