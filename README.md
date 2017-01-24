go-esa
------

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CotoCoto/go-esa"
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
-	[ ] POST /v1/teams/:team_name/posts
-	[ ] PATCH /v1/teams/:team_name/posts/:post_number
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number
-	[ ] GET /v1/teams/:team_name/posts/:post_number/comments
-	[ ] GET /v1/teams/:team_name/comments/:comment_id
-	[ ] POST /v1/teams/:team_name/posts/:post_number/comments
-	[ ] PATCH /v1/teams/:team_name/comments/:comment_id
-	[ ] DELETE /v1/teams/:team_name/comments/:comment_id Star
-	[ ] GET /v1/teams/:team_name/posts/:post_number/stargazers
-	[ ] POST /v1/teams/:team_name/posts/:post_number/star
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number/star
-	[ ] GET /v1/teams/:team_name/comments/:comment_id/stargazers
-	[ ] POST /v1/teams/:team_name/comments/:comment_id/star
-	[ ] DELETE /v1/teams/:team_name/comments/:comment_id/star Watch
-	[ ] GET /v1/teams/:team_name/posts/:post_number/watchers
-	[ ] POST /v1/teams/:team_name/posts/:post_number/watch
-	[ ] DELETE /v1/teams/:team_name/posts/:post_number/watch
-	[ ] GET /v1/user
