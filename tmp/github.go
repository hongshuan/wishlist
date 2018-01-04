// file: github.go
package github

// A Client manages communication with the GitHub API.
type Client struct {
	clientMu sync.Mutex   // protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Base URL for uploading files.
	UploadURL *url.URL

	// User agent used when communicating with the GitHub API.
	UserAgent string

	rateMu     sync.Mutex
	rateLimits [categories]Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the GitHub API.
	Activity       *ActivityService
	Admin          *AdminService
	Apps           *AppsService
	Authorizations *AuthorizationsService
	Gists          *GistsService
	Git            *GitService
	Gitignores     *GitignoresService
	Issues         *IssuesService
	Licenses       *LicensesService
	Marketplace    *MarketplaceService
	Migrations     *MigrationService
	Organizations  *OrganizationsService
	Projects       *ProjectsService
	PullRequests   *PullRequestsService
	Reactions      *ReactionsService
	Repositories   *RepositoriesService
	Search         *SearchService
	Users          *UsersService
}

type service struct {
	client *Client
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	uploadURL, _ := url.Parse(uploadBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, UploadURL: uploadURL}
	c.common.client = c
	c.Activity = (*ActivityService)(&c.common)
	c.Admin = (*AdminService)(&c.common)
	c.Apps = (*AppsService)(&c.common)
	c.Authorizations = (*AuthorizationsService)(&c.common)
	c.Gists = (*GistsService)(&c.common)
	c.Git = (*GitService)(&c.common)
	c.Gitignores = (*GitignoresService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
	c.Licenses = (*LicensesService)(&c.common)
	c.Marketplace = &MarketplaceService{client: c}
	c.Migrations = (*MigrationService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.PullRequests = (*PullRequestsService)(&c.common)
	c.Reactions = (*ReactionsService)(&c.common)
	c.Repositories = (*RepositoriesService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", mediaTypeV3)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// file: activity.go ==========
package github

// GitHub API docs: https://developer.github.com/v3/activity/
type ActivityService service

type Feeds struct { ... }

func (s *ActivityService) ListFeeds(ctx context.Context) (*Feeds, *Response, error) {
	req, err := s.client.NewRequest("GET", "feeds", nil)
	if err != nil {
		return nil, nil, err
	}

	f := &Feeds{}
	resp, err := s.client.Do(ctx, req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
// ActivityService has a lot of methods

// file: admin.go ==========
package github

// GitHub API docs: https://developer.github.com/v3/enterprise/
type AdminService service

// GitHub API docs: https://developer.github.com/v3/enterprise-admin/admin_stats/
func (s *AdminService) GetAdminStats(ctx context.Context) (*AdminStats, *Response, error) {
	u := fmt.Sprintf("enterprise/stats/all")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	m := new(AdminStats)
	resp, err := s.client.Do(ctx, req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}
// AdminService has a lot of methods

// file: apps.go ==========
package github

// GitHub API docs: https://developer.github.com/v3/apps/
type AppsService service

// GitHub API docs: https://developer.github.com/v3/apps/#get-a-single-github-app
func (s *AppsService) Get(ctx context.Context, appSlug string) (*App, *Response, error) {
	var u string
	if appSlug != "" {
		u = fmt.Sprintf("apps/%v", appSlug)
	} else {
		u = "app"
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launches.
	req.Header.Set("Accept", mediaTypeIntegrationPreview)

	app := new(App)
	resp, err := s.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}
// AppsService has a lot of methods

type IssuesService service
type MigrationService service
type AuthorizationsService service
type GistsService service
type GitignoresService service
type OrganizationsService service
type PullRequestsService service
type ReactionsService service
type RepositoriesService service
type SearchService service
type UsersService service
