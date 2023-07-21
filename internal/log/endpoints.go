package log

type Endpoint struct {
	Name   string
	Url    string
	Method string
}

func NewEndpoint(name string, url string, method string) *Endpoint {
	return &Endpoint{
		Name:   name,
		Url:    url,
		Method: method,
	}
}

var (
	Auth                     = Endpoint{"Auth Middleware", "", ""}
	Login                    = Endpoint{"Login", "/api/login", "POST"}
	Logout                   = Endpoint{"Logout", "/api/logout", "POST"}
	Register                 = Endpoint{"Logout", "/api/register", "POST"}
	Weather_Get              = Endpoint{"Weather", "/api/weather", "GET"}
	Weather_Hist_Get         = Endpoint{"WeatherHistory", "/api/weather/history", "GET"}
	Weather_Hist_Post        = Endpoint{"WeatherHistory", "/api/weather/history", "POST"}
	Weather_Hist_Delete      = Endpoint{"WeatherHistoryDelete", "/api/weather/history/{history_id}", "DELETE"}
	Weather_Hist_Bulk_Delete = Endpoint{"WeatherHistoryBULKDelete", "/api/weather/history/bulk", "DELETE"}
)