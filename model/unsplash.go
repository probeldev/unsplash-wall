package model

type AppConfig struct {
	Locale  string `json:"locale"`
	AdEnv   string `json:"adEnv"`
	Captcha struct {
		Sitekey string `json:"sitekey"`
	} `json:"captcha"`
	ClientGeoRegion          string `json:"clientGeoRegion"`
	DirectAdsRegion          string `json:"directAdsRegion"`
	LogLevel                 string `json:"logLevel"`
	PseudoSameOriginDownload bool   `json:"pseudoSameOriginDownload"`
	RailsOrigin              string `json:"railsOrigin"`
	ShouldShowVisualSearch   bool   `json:"shouldShowVisualSearch"`
	ShouldUseSecureCookies   bool   `json:"shouldUseSecureCookies"`
	SnowplowTracker          string `json:"snowplowTracker"`
	UIDelay                  bool   `json:"uiDelay"`
	UnsplashEnv              string `json:"unsplashEnv"`
	UnsplashOrigin           string `json:"unsplashOrigin"`
	XPS                      struct {
		FreeSemanticPerf string `json:"free-semantic-perf"`
	} `json:"xps"`
}

type QueryState struct {
	Data struct {
		ID            string `json:"id"`
		PreviewPhotos []struct {
			ID        string `json:"id"`
			Slug      string `json:"slug"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
			BlurHash  string `json:"blur_hash"`
			URLs      struct {
				Raw     string `json:"raw"`
				Full    string `json:"full"`
				Regular string `json:"regular"`
				Small   string `json:"small"`
				Thumb   string `json:"thumb"`
				SmallS3 string `json:"small_s3"`
			} `json:"urls"`
			AssetType string `json:"asset_type"`
		} `json:"preview_photos"`
		Title       string `json:"title"`
		TotalPhotos int    `json:"total_photos"`
	} `json:"data"`
	DataUpdateCount    int         `json:"dataUpdateCount"`
	DataUpdatedAt      int64       `json:"dataUpdatedAt"`
	Error              interface{} `json:"error"`
	ErrorUpdateCount   int         `json:"errorUpdateCount"`
	ErrorUpdatedAt     int         `json:"errorUpdatedAt"`
	FetchFailureCount  int         `json:"fetchFailureCount"`
	FetchFailureReason interface{} `json:"fetchFailureReason"`
	FetchMeta          interface{} `json:"fetchMeta"`
	IsInvalidated      bool        `json:"isInvalidated"`
	Status             string      `json:"status"`
	FetchStatus        string      `json:"fetchStatus"`
}

type Query struct {
	// tate     QueryState `json:"state"`
	QueryKey  []string `json:"queryKey"`
	QueryHash string   `json:"queryHash"`
}

type Photo struct {
	UpdatedAt string `json:"updated_at"`
	AssetType string `json:"asset_type"`
	BlurHash  string `json:"blur_hash"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	URLs      struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
		SmallS3 string `json:"small_s3"`
	} `json:"urls"`
	AltDescription string      `json:"alt_description"`
	Color          string      `json:"color"`
	Description    interface{} `json:"description"`
	Height         int         `json:"height"`
	LikedByUser    bool        `json:"liked_by_user"`
	Likes          int         `json:"likes"`
	Links          struct {
		HTML     string `json:"html"`
		Download string `json:"download"`
	} `json:"links"`
	PromotedAt       interface{} `json:"promoted_at"`
	TopicSubmissions map[string]struct {
		Status     string `json:"status"`
		ApprovedOn string `json:"approved_on,omitempty"`
	} `json:"topic_submissions"`
	Width                    int           `json:"width"`
	CurrentUserCollectionIDs []interface{} `json:"current_user_collection_ids"`
	UserID                   string        `json:"userId"`
	Premium                  bool          `json:"premium"`
}

type User struct {
	UpdatedAt   string `json:"updated_at"`
	AcceptedTos bool   `json:"accepted_tos"`
	Bio         string `json:"bio"`
	FirstName   string `json:"first_name"`
	ForHire     bool   `json:"for_hire"`
	ID          string `json:"id"`
	LastName    string `json:"last_name"`
	Links       struct {
		Self      string `json:"self"`
		HTML      string `json:"html"`
		Photos    string `json:"photos"`
		Likes     string `json:"likes"`
		Portfolio string `json:"portfolio"`
	} `json:"links"`
	Location     string `json:"location"`
	Name         string `json:"name"`
	ProfileImage struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"profile_image"`
	Social struct {
		InstagramUsername string      `json:"instagram_username"`
		PortfolioURL      string      `json:"portfolio_url"`
		TwitterUsername   interface{} `json:"twitter_username"`
	} `json:"social"`
	TotalCollections   int    `json:"total_collections"`
	TotalLikes         int    `json:"total_likes"`
	TotalPhotos        int    `json:"total_photos"`
	TotalIllustrations int    `json:"total_illustrations"`
	Username           string `json:"username"`
}

type Root struct {
	AppConfig struct {
		Locale  string `json:"locale"`
		AdEnv   string `json:"adEnv"`
		Captcha struct {
			Sitekey string `json:"sitekey"`
		} `json:"captcha"`
		ClientGeoRegion          string `json:"clientGeoRegion"`
		DirectAdsRegion          string `json:"directAdsRegion"`
		LogLevel                 string `json:"logLevel"`
		PseudoSameOriginDownload bool   `json:"pseudoSameOriginDownload"`
		RailsOrigin              string `json:"railsOrigin"`
		ShouldShowVisualSearch   bool   `json:"shouldShowVisualSearch"`
		ShouldUseSecureCookies   bool   `json:"shouldUseSecureCookies"`
		SnowplowTracker          string `json:"snowplowTracker"`
		UIDelay                  bool   `json:"uiDelay"`
		UnsplashEnv              string `json:"unsplashEnv"`
		UnsplashOrigin           string `json:"unsplashOrigin"`
		XPS                      struct {
			FreeSemanticPerf string `json:"free-semantic-perf"`
		} `json:"xps"`
	} `json:"appConfig"`
	//QueryClientCache struct {
	//	Mutations []interface{} `json:"mutations"`
	//	Queries   []Query       `json:"queries"`
	// } `json:"queryClientCache"`
	ReduxInitialState struct {
		Entities struct {
			BriefSubmissions []interface{}    `json:"briefSubmissions"`
			Photos           map[string]Photo `json:"photos"`
			PhotosRemoteData struct {
				Full struct{} `json:"full"`
			} `json:"photosRemoteData"`
			Users         map[string]User `json:"users"`
			Collections   struct{}        `json:"collections"`
			Topics        struct{}        `json:"topics"`
			Notifications struct{}        `json:"notifications"`
			JobPosts      struct{}        `json:"jobPosts"`
			LandingPages  struct{}        `json:"landingPages"`
			KeywordTrend  struct{}        `json:"keywordTrend"`
		} `json:"entities"`
		Feeds struct {
			BriefSubmissionsFeeds []interface{} `json:"briefSubmissionsFeeds"`
			NotificationFeed      struct {
				Highlights struct {
					Tag string `json:"_tag"`
				} `json:"highlights"`
				Activity struct {
					Tag string `json:"_tag"`
				} `json:"activity"`
			} `json:"notificationFeed"`
			PhotoFeeds      [][]interface{} `json:"photoFeeds"`
			CollectionFeeds []interface{}   `json:"collectionFeeds"`
			UserFeeds       []interface{}   `json:"userFeeds"`
			TopicFeeds      struct{}        `json:"topicFeeds"`
		} `json:"feeds"`
		Uploader interface{} `json:"uploader"`
		Auth     interface{} `json:"auth"`
		UI       struct {
			ProgressBar struct {
				DataFetchingCounter int `json:"dataFetchingCounter"`
			} `json:"progressBar"`
			SayThanksCard struct {
				Tag string `json:"_tag"`
			} `json:"sayThanksCard"`
			Flash                      interface{} `json:"flash"`
			HomepageModulesSeed        int64       `json:"homepageModulesSeed"`
			SearchFormTrendingSearches []string    `json:"searchFormTrendingSearches"`
			ModuleTrends               [][]string  `json:"moduleTrends"`
			PlusFeedRouteCollections   interface{} `json:"plusFeedRouteCollections"`
			JoinPage                   interface{} `json:"joinPage"`
		} `json:"ui"`
		Searches                        struct{} `json:"searches"`
		PhotoSearchRelatedIllustrations struct{} `json:"photoSearchRelatedIllustrations"`
		SearchQueryLandingPages         struct{} `json:"searchQueryLandingPages"`
		VisualSearches                  struct{} `json:"visualSearches"`
		StaticData                      struct {
			JobPostIds interface{}   `json:"jobPostIds"`
			Timeline   []interface{} `json:"timeline"`
		} `json:"staticData"`
		Trends struct {
			MostInDemands struct {
				Pages  []interface{} `json:"pages"`
				Latest struct {
					Tag string `json:"_tag"`
				} `json:"latest"`
			} `json:"mostInDemands"`
			TrendingCategories struct {
				Pages  []interface{} `json:"pages"`
				Latest struct {
					Tag string `json:"_tag"`
				} `json:"latest"`
			} `json:"trendingCategories"`
		} `json:"trends"`
		PublicStats   struct{} `json:"publicStats"`
		Subscriptions struct {
			Prices struct {
				Tag string `json:"_tag"`
			} `json:"prices"`
			Trials []interface{} `json:"trials"`
		} `json:"subscriptions"`
		DirectAds struct {
			RenderedAdsByAdvertiser []interface{} `json:"renderedAdsByAdvertiser"`
		} `json:"directAds"`
		FeaturedLandingPages   interface{} `json:"featuredLandingPages"`
		BriefSubmissionsTotals interface{} `json:"briefSubmissionsTotals"`
	} `json:"reduxInitialState"`
}
