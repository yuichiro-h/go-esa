go-esa
------

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yuichiro-h/go-esa"
)

func main() {
	client := esa.New(&esa.Config{AccessToken: "accessToken"})
	res, err := client.GetTeams(&esa.GetTeamsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
```

### Supported resources

-	[x] GET /v1/teams
-	[x] GET /v1/teams/:team_name
-	[x] GET /v1/teams/:team_name/stats
-	[x] GET /v1/teams/:team_name/members
-	[x] GET /v1/teams/:team_name/posts
-	[x] GET /v1/teams/:team_name/posts/:post_number
-	[x] POST /v1/teams/:team_name/posts
-	[x] PATCH /v1/teams/:team_name/posts/:post_number
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number
-	[x] GET /v1/teams/:team_name/posts/:post_number/comments
-	[x] GET /v1/teams/:team_name/comments/:comment_id
-	[x] POST /v1/teams/:team_name/posts/:post_number/comments
-	[x] PATCH /v1/teams/:team_name/comments/:comment_id
-	[ ] DELETE /v1/teams/:team_name/comments/:comment_id
-	[x] GET /v1/teams/:team_name/posts/:post_number/stargazers
-	[x] POST /v1/teams/:team_name/posts/:post_number/star
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number/star
-	[x] GET /v1/teams/:team_name/comments/:comment_id/stargazers
-	[x] POST /v1/teams/:team_name/comments/:comment_id/star
-	[ ] DELETE /v1/teams/:team_name/comments/:comment_id/star
-	[x] GET /v1/teams/:team_name/posts/:post_number/watchers
-	[x] POST /v1/teams/:team_name/posts/:post_number/watch
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number/watch
-	[x] GET /v1/user
