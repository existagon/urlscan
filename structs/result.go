package structs

import "time"

type ResultData struct {
	Data struct {
		Requests []struct {
			Request struct {
				RequestID   string `json:"requestId"`
				LoaderID    string `json:"loaderId"`
				DocumentURL string `json:"documentURL"`
				Request     struct {
					URL     string `json:"url"`
					Method  string `json:"method"`
					Headers struct {
						UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
						UserAgent               string `json:"User-Agent"`
						AcceptLanguage          string `json:"accept-language"`
					} `json:"headers"`
					MixedContentType string `json:"mixedContentType"`
					InitialPriority  string `json:"initialPriority"`
					ReferrerPolicy   string `json:"referrerPolicy"`
					IsSameSite       bool   `json:"isSameSite"`
				} `json:"request"`
				Timestamp float64 `json:"timestamp"`
				WallTime  float64 `json:"wallTime"`
				Initiator struct {
					Type string `json:"type"`
				} `json:"initiator"`
				RedirectHasExtraInfo bool `json:"redirectHasExtraInfo"`
			} `json:"request"`
			Response struct {
				EncodedDataLength int    `json:"encodedDataLength"`
				DataLength        int    `json:"dataLength"`
				RequestID         string `json:"requestId"`
				Type              string `json:"type"`
				Response          struct {
					URL               string            `json:"url"`
					Status            int               `json:"status"`
					StatusText        string            `json:"statusText"`
					Headers           map[string]string `json:"headers"`
					MimeType          string            `json:"mimeType"`
					RemoteIPAddress   string            `json:"remoteIPAddress"`
					RemotePort        int               `json:"remotePort"`
					EncodedDataLength int               `json:"encodedDataLength"`
					Timing            struct {
						RequestTime              float64 `json:"requestTime"`
						ProxyStart               int     `json:"proxyStart"`
						ProxyEnd                 int     `json:"proxyEnd"`
						DNSStart                 float64 `json:"dnsStart"`
						DNSEnd                   float64 `json:"dnsEnd"`
						ConnectStart             float64 `json:"connectStart"`
						ConnectEnd               float64 `json:"connectEnd"`
						SslStart                 float64 `json:"sslStart"`
						SslEnd                   float64 `json:"sslEnd"`
						WorkerStart              int     `json:"workerStart"`
						WorkerReady              int     `json:"workerReady"`
						WorkerFetchStart         int     `json:"workerFetchStart"`
						WorkerRespondWithSettled int     `json:"workerRespondWithSettled"`
						SendStart                float64 `json:"sendStart"`
						SendEnd                  float64 `json:"sendEnd"`
						PushStart                int     `json:"pushStart"`
						PushEnd                  int     `json:"pushEnd"`
						ReceiveHeadersEnd        float64 `json:"receiveHeadersEnd"`
					} `json:"timing"`
					ResponseTime    float64 `json:"responseTime"`
					Protocol        string  `json:"protocol"`
					SecurityState   string  `json:"securityState"`
					SecurityDetails struct {
						Protocol                          string        `json:"protocol"`
						KeyExchange                       string        `json:"keyExchange"`
						KeyExchangeGroup                  string        `json:"keyExchangeGroup"`
						Cipher                            string        `json:"cipher"`
						CertificateID                     int           `json:"certificateId"`
						SubjectName                       string        `json:"subjectName"`
						SanList                           []string      `json:"sanList"`
						Issuer                            string        `json:"issuer"`
						ValidFrom                         int           `json:"validFrom"`
						ValidTo                           int           `json:"validTo"`
						SignedCertificateTimestampList    []interface{} `json:"signedCertificateTimestampList"`
						CertificateTransparencyCompliance string        `json:"certificateTransparencyCompliance"`
						ServerSignatureAlgorithm          int           `json:"serverSignatureAlgorithm"`
						EncryptedClientHello              bool          `json:"encryptedClientHello"`
					} `json:"securityDetails"`
				} `json:"response"`
				HasExtraInfo bool   `json:"hasExtraInfo"`
				Hash         string `json:"hash"`
				Size         int    `json:"size"`
				Asn          struct {
					IP          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"asn"`
				Geoip struct {
					Country     string    `json:"country"`
					Region      string    `json:"region"`
					Timezone    string    `json:"timezone"`
					City        string    `json:"city"`
					Ll          []float64 `json:"ll"`
					CountryName string    `json:"country_name"`
					Metro       int       `json:"metro"`
				} `json:"geoip"`
			} `json:"response,omitempty"`
		} `json:"requests"`
		Cookies []struct {
			Name         string `json:"name"`
			Value        string `json:"value"`
			Domain       string `json:"domain"`
			Path         string `json:"path"`
			Expires      int    `json:"expires"`
			Size         int    `json:"size"`
			HTTPOnly     bool   `json:"httpOnly"`
			Secure       bool   `json:"secure"`
			Session      bool   `json:"session"`
			Priority     string `json:"priority"`
			SameParty    bool   `json:"sameParty"`
			SourceScheme string `json:"sourceScheme"`
			SourcePort   int    `json:"sourcePort"`
		} `json:"cookies"`
		Console []struct {
			Message struct {
				Source    string      `json:"source"`
				Level     string      `json:"level"`
				Text      string      `json:"text"`
				Timestamp float64     `json:"timestamp"`
				URL       string      `json:"url"`
				Line      interface{} `json:"line"`
			} `json:"message"`
		} `json:"console"`
		Links []struct {
			Href string `json:"href"`
			Text string `json:"text"`
		} `json:"links"`
		Timing struct {
			BeginNavigation      time.Time `json:"beginNavigation"`
			FrameStartedLoading  time.Time `json:"frameStartedLoading"`
			FrameNavigated       time.Time `json:"frameNavigated"`
			FrameStoppedLoading  time.Time `json:"frameStoppedLoading"`
			DomContentEventFired time.Time `json:"domContentEventFired"`
		} `json:"timing"`
		Globals []struct {
			Prop string `json:"prop"`
			Type string `json:"type"`
		} `json:"globals"`
	} `json:"data"`
	Stats struct {
		ResourceStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Latency     int      `json:"latency"`
			Countries   []string `json:"countries"`
			Ips         []string `json:"ips"`
			Type        string   `json:"type"`
			Compression string   `json:"compression"`
			Percentage  int      `json:"percentage"`
		} `json:"resourceStats"`
		ProtocolStats []struct {
			Count         int      `json:"count"`
			Size          int      `json:"size"`
			EncodedSize   int      `json:"encodedSize"`
			Ips           []string `json:"ips"`
			Countries     []string `json:"countries"`
			SecurityState struct {
			} `json:"securityState"`
			Protocol string `json:"protocol"`
		} `json:"protocolStats"`
		TLSStats []struct {
			Count         int            `json:"count"`
			Size          int            `json:"size"`
			EncodedSize   int            `json:"encodedSize"`
			Ips           []string       `json:"ips"`
			Countries     []string       `json:"countries"`
			Protocols     map[string]int `json:"protocols"`
			SecurityState string         `json:"securityState"`
		} `json:"tlsStats"`
		ServerStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Ips         []string `json:"ips"`
			Countries   []string `json:"countries"`
			Server      string   `json:"server"`
		} `json:"serverStats"`
		DomainStats []struct {
			Count       int      `json:"count"`
			Ips         []string `json:"ips"`
			Domain      string   `json:"domain"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Countries   []string `json:"countries"`
			Index       int      `json:"index"`
			Initiators  []string `json:"initiators"`
			Redirects   int      `json:"redirects"`
		} `json:"domainStats"`
		RegDomainStats []struct {
			Count       int           `json:"count"`
			Ips         []string      `json:"ips"`
			RegDomain   string        `json:"regDomain"`
			Size        int           `json:"size"`
			EncodedSize int           `json:"encodedSize"`
			Countries   []interface{} `json:"countries"`
			Index       int           `json:"index"`
			SubDomains  []struct {
				Domain  string `json:"domain"`
				Country string `json:"country"`
			} `json:"subDomains"`
			Redirects int `json:"redirects"`
		} `json:"regDomainStats"`
		SecureRequests   int `json:"secureRequests"`
		SecurePercentage int `json:"securePercentage"`
		IPv6Percentage   int `json:"IPv6Percentage"`
		UniqCountries    int `json:"uniqCountries"`
		TotalLinks       int `json:"totalLinks"`
		Malicious        int `json:"malicious"`
		AdBlocked        int `json:"adBlocked"`
		IPStats          []struct {
			Requests int      `json:"requests"`
			Domains  []string `json:"domains"`
			IP       string   `json:"ip"`
			Asn      struct {
				IP          string `json:"ip"`
				Asn         string `json:"asn"`
				Country     string `json:"country"`
				Registrar   string `json:"registrar"`
				Date        string `json:"date"`
				Description string `json:"description"`
				Route       string `json:"route"`
				Name        string `json:"name"`
			} `json:"asn"`
			DNS struct {
			} `json:"dns"`
			Geoip struct {
				Country     string    `json:"country"`
				Region      string    `json:"region"`
				Timezone    string    `json:"timezone"`
				City        string    `json:"city"`
				Ll          []float64 `json:"ll"`
				CountryName string    `json:"country_name"`
				Metro       int       `json:"metro"`
			} `json:"geoip"`
			Size        int         `json:"size"`
			EncodedSize int         `json:"encodedSize"`
			Countries   []string    `json:"countries"`
			Index       int         `json:"index"`
			Ipv6        bool        `json:"ipv6"`
			Redirects   int         `json:"redirects"`
			Count       interface{} `json:"count"`
			Rdns        struct {
				IP  string `json:"ip"`
				Ptr string `json:"ptr"`
			} `json:"rdns,omitempty"`
		} `json:"ipStats"`
	} `json:"stats"`
	Meta struct {
		Processors struct {
			Umbrella struct {
				Data []struct {
					Hostname string `json:"hostname"`
					Rank     int    `json:"rank"`
				} `json:"data"`
			} `json:"umbrella"`
			Geoip struct {
				Data []struct {
					IP    string `json:"ip"`
					Geoip struct {
						Country     string    `json:"country"`
						Region      string    `json:"region"`
						Timezone    string    `json:"timezone"`
						City        string    `json:"city"`
						Ll          []float64 `json:"ll"`
						CountryName string    `json:"country_name"`
						Metro       int       `json:"metro"`
					} `json:"geoip"`
				} `json:"data"`
			} `json:"geoip"`
			Rdns struct {
				Data []struct {
					IP  string `json:"ip"`
					Ptr string `json:"ptr"`
				} `json:"data"`
			} `json:"rdns"`
			Asn struct {
				Data []struct {
					IP          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"data"`
			} `json:"asn"`
			Wappa struct {
				Data []struct {
					Confidence []struct {
						Confidence int    `json:"confidence"`
						Pattern    string `json:"pattern"`
					} `json:"confidence"`
					ConfidenceTotal int    `json:"confidenceTotal"`
					App             string `json:"app"`
					Icon            string `json:"icon"`
					Website         string `json:"website"`
					Categories      []struct {
						Name     string `json:"name"`
						Priority int    `json:"priority"`
					} `json:"categories"`
				} `json:"data"`
			} `json:"wappa"`
		} `json:"processors"`
	} `json:"meta"`
	Task struct {
		UUID          string        `json:"uuid"`
		Time          time.Time     `json:"time"`
		URL           string        `json:"url"`
		Visibility    string        `json:"visibility"`
		Method        string        `json:"method"`
		Source        string        `json:"source"`
		Tags          []interface{} `json:"tags"`
		ReportURL     string        `json:"reportURL"`
		ScreenshotURL string        `json:"screenshotURL"`
		DomURL        string        `json:"domURL"`
	} `json:"task"`
	Page struct {
		URL     string `json:"url"`
		Domain  string `json:"domain"`
		Country string `json:"country"`
		City    string `json:"city"`
		Server  string `json:"server"`
		IP      string `json:"ip"`
		Asn     string `json:"asn"`
		Asnname string `json:"asnname"`
	} `json:"page"`
	Lists struct {
		Ips          []string `json:"ips"`
		Countries    []string `json:"countries"`
		Asns         []string `json:"asns"`
		Domains      []string `json:"domains"`
		Servers      []string `json:"servers"`
		Urls         []string `json:"urls"`
		LinkDomains  []string `json:"linkDomains"`
		Certificates []struct {
			SubjectName string `json:"subjectName"`
			Issuer      string `json:"issuer"`
			ValidFrom   int    `json:"validFrom"`
			ValidTo     int    `json:"validTo"`
		} `json:"certificates"`
		Hashes []string `json:"hashes"`
	} `json:"lists"`
	Verdicts struct {
		Overall struct {
			Score       int      `json:"score"`
			Categories  []string `json:"categories"`
			Brands      []string `json:"brands"`
			Tags        []string `json:"tags"`
			Malicious   bool     `json:"malicious"`
			HasVerdicts bool     `json:"hasVerdicts"`
		} `json:"overall"`
		Urlscan struct {
			Score      int      `json:"score"`
			Categories []string `json:"categories"`
			Brands     []struct {
				Key      string   `json:"key"`
				Name     string   `json:"name"`
				Country  []string `json:"country"`
				Vertical []string `json:"vertical"`
			} `json:"brands"`
			Tags        []string `json:"tags"`
			Malicious   bool     `json:"malicious"`
			HasVerdicts bool     `json:"hasVerdicts"`
		} `json:"urlscan"`
		Community struct {
			Score      int      `json:"score"`
			Categories []string `json:"categories"`
			Brands     []struct {
				Key      string   `json:"key"`
				Name     string   `json:"name"`
				Country  []string `json:"country"`
				Vertical []string `json:"vertical"`
			} `json:"brands"`
			VotesTotal     int  `json:"votesTotal"`
			VotesMalicious int  `json:"votesMalicious"`
			VotesBenign    int  `json:"votesBenign"`
			Malicious      bool `json:"malicious"`
			HasVerdicts    bool `json:"hasVerdicts"`
		} `json:"community"`
	} `json:"verdicts"`
	Submitter struct {
	} `json:"submitter"`
}
