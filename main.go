package main

import (
    "fmt"
    "os"
    // "encoding/json"

    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
)

var personalAccessToken = os.Getenv("GH_ACCESS_TOKEN")

type TokenSource struct {
    AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

func main() {
  args := os.Args

  if len(os.Args) == 1 {
    fmt.Println("No arguments")
    return
  }

  tokenSource := &TokenSource{
    AccessToken: personalAccessToken,
  }

  oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
  client := github.NewClient(oauthClient)

  ghSection := args[1]
  switch ghSection {
    case "issues":
      issues, _, err := client.Issues.ListByRepo("user_name", "repo_name", nil)
      for _, i := range issues {
        fmt.Printf("#%d: %s\n", *i.Number, *i.Title)
      }
      if err != nil {
        fmt.Println(err)
      }
      return
  }
  fmt.Printf("Argument %s is not valid\n", ghSection)
}
