package requests

type LogExternal struct {
	UUID             string `json:"uuid" bson:"uuid"`
	ServiceUUID      string `json:"service_uuid" bson:"service_uuid"`
	APIUUID          string `json:"api_uuid" bson:"api_uuid"`
	IPAddress        string `json:"ip_address" bson:"ip_address"`
	OfficeID         int    `json:"office_id" bson:"office_id"`
	ReqFromCitizenID string `json:"req_from_citizen_id" bson:"req_from_citizen_id"`
	ReqCitizenID     string `json:"req_citizen_id" bson:"req_citizen_id"`
	UserID           int    `json:"user_id" bson:"user_id"`
	DurationTime     string `json:"duration_time" bson:"duration_time"`
	StatusCode       int    `json:"status_code" bson:"status_code"`
}

type LogInternal struct {
	UUID             string `json:"uuid" bson:"uuid"`
	ServiceUUID      string `json:"service_uuid" bson:"service_uuid"`
	APIUUID          string `json:"api_uuid" bson:"api_uuid"`
	IPAddress        string `json:"ip_address" bson:"ip_address"`
	ApplicationID    int    `json:"application_id" bson:"application_id"`
	ReqFromCitizenID string `json:"req_from_citizen_id" bson:"req_from_citizen_id"`
	ReqCitizenID     string `json:"req_citizen_id" bson:"req_citizen_id"`
	UserID           int    `json:"user_id" bson:"user_id"`
	DurationTime     string `json:"duration_time" bson:"duration_time"`
	StatusCode       string `json:"status_code" bson:"status_code"`
}

type ExternalAccessLog struct {
	UUID         string `json:"uuid" bson:"uuid"`
	ServiceUUID  string `json:"service_uuid" bson:"service_uuid"`
	APIUUID      string `json:"api_uuid" bson:"api_uuid"`
	IPAddress    string `json:"ip_address" bson:"ip_address"`
	OfficeID     int    `json:"office_id" bson:"office_id"`
	CitizenID    string `json:"citizen_id" bson:"citizen_id"`
	UserID       int    `json:"user_id" bson:"user_id"`
	DurationTime string `json:"duration_time" bson:"duration_time"`
	StatusCode   string `json:"status_code" bson:"status_code"`
}

type InternalAccessLog struct {
	UUID          string `json:"uuid" bson:"uuid"`
	ServiceUUID   string `json:"service_uuid" bson:"service_uuid"`
	APIUUID       string `json:"api_uuid" bson:"api_uuid"`
	IPAddress     string `json:"ip_address" bson:"ip_address"`
	ApplicationID int    `json:"application_id" bson:"application_id"`
	CitizenID     string `json:"citizen_id" bson:"citizen_id"`
	UserID        int    `json:"user_id" bson:"user_id"`
	DurationTime  string `json:"duration_time" bson:"duration_time"`
	StatusCode    string `json:"status_code" bson:"status_code"`
}

type ProxyLog struct {
	UUID             string `json:"uuid" bson:"uuid"`
	ServiceUUID      string `json:"service_uuid" bson:"service_uuid"`
	ReqFromCitizenID string `json:"req_from_citizen_id" bson:"req_from_citizen_id"`
	ReqCitizenID     string `json:"req_citizen_id" bson:"req_citizen_id"`
	UserID           int    `json:"user_id" bson:"user_id"`
}

type TrafficLog struct {
	RequestID    string `json:"request_id" bson:"request_id"`
	UserAgent    string `json:"user_agent" bson:"user_agent"`
	IPAddress    string `json:"ip_address" bson:"ip_address"`
	ServiceUUID  string `json:"service_uuid" bson:"service_uuid"`
	APIUUID      string `json:"api_uuid" bson:"api_uuid"`
	StatusCode   string `json:"status_code" bson:"status_code"`
	DurationTime string `json:"duration_time" bson:"duration_time"`
	UserID       int    `json:"user_id" bson:"user_id"`
}
