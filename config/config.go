package config

type Cloudinary struct {
	CloudName, ApiKey, ApiSecret string
}

var CloudConfig = &Cloudinary{"sandydev99", "917665392796572", "J6m152XVs7TyyfxJYn9oIHjPiGc"}

var TimeZone = "Asia/Kolkata"
var TimeSlots = []string{"10:00:00", "10:30:00", "11:00:00", "11:30:00", "12:00:00", "12:30:00", "13:00:00", "13:30:00", "14:00:00", "14:30:00", "15:00:00", "15:30:00", "16:00:00", "16:30:00", "17:00:00"}
