package servers

type Server struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Addresses struct {
		IPv4 struct {
			Address map[string]interface{}
		}
		IPv6 struct {
			Address map[string]interface{}
		}
	} `json:"addresses"`
	OS     string `json:"operatingSystem"`
	Ram    int    `json:"ram"`
	Status string `json:"status"`
}
