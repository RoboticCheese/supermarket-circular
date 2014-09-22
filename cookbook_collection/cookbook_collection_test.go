package cookbook_collection

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/RoboticCheese/supermarket-circular/cookbook"
)

var http_response = "{\"mailcatcher\":{\"0.1.0\":{\"location_type\":\"opscode\",\"location_path\":\"https://supermarket.getchef.com/api/v1\",\"download_url\":\"https://supermarket.getchef.com/api/v1/cookbooks/mailcatcher/versions/0.1.0/download\",\"dependencies\":{}}}}"

func start_httptest() (*httptest.Server) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, http_response)
  }))
  return ts
}

func Test_Merge_1(t *testing.T) {
  oldcc := CookbookCollection{[]cookbook.Cookbook{
    {"apache", []string{"0.1.0"}},
    {"bpache", []string{"1.0.0"}},
  }}
  newcc := CookbookCollection{[]cookbook.Cookbook{
    {"cpache", []string{"2.0.0"}},
  }}
  res, err := oldcc.Merge(newcc)
  if err != nil {
    t.Fatalf("Expected no error, got: %s", err)
  }
  if len(res.Cookbooks) != 3 {
    t.Fatalf("Expected res length of 3, got: %d", len(res.Cookbooks))
  }
  if len(oldcc.Cookbooks) != 3 {
    t.Fatalf("Expected merged length of 3, got: %d", len(oldcc.Cookbooks))
  }
}

func Test_contains_1(t *testing.T) {
  cb := cookbook.Cookbook{"apache", []string{}}
  cc := CookbookCollection{}
  res := cc.contains(cb)
  if res != false {
    t.Fatalf("Expected false, got: %s", res)
  }
}

func Test_contains_2(t *testing.T) {
  cb := cookbook.Cookbook{"apache", []string{}}
  cc := CookbookCollection{[]cookbook.Cookbook{cb}}
  res := cc.contains(cb)
  if res != true {
    t.Fatalf("Expected true, got: %s", res)
  }
}

func Test_cc_from_universe_1(t *testing.T) {
  ts := start_httptest()
  defer ts.Close()
  cc := CookbookCollection{}
  input := []byte(http_response)
  res, err := cc.cc_from_universe(input)
  if err != nil {
    t.Fatalf("Expected no error, got: %s", err)
  }
  if len(res.Cookbooks) != 1 {
    t.Fatalf("Expected return length 1, got: %d", len(res.Cookbooks))
  }
  if len(cc.Cookbooks) != 1 {
    t.Fatalf("Expected collection length 1, got: %d", len(cc.Cookbooks))
  }
  if res.Cookbooks[0].Name != "mailcatcher" {
    t.Fatalf("Expected mailcatcher cookbook, got: %s", res.Cookbooks[0].Name)
  }
  if cc.Cookbooks[0].Name != "mailcatcher" {
    t.Fatalf("Expected mailcatcher cookbook, got: %s", cc.Cookbooks[0].Name)
  }
  if len(res.Cookbooks[0].Versions) != 1 {
    t.Fatalf("Expected Versions length 1, got: %d", len(res.Cookbooks[0].Versions))
  }
  if len(cc.Cookbooks[0].Versions) != 1 {
    t.Fatalf("Expected Versions length 1, got: %d", len(cc.Cookbooks[0].Versions))
  }
  if res.Cookbooks[0].Versions[0] != "0.1.0" {
    t.Fatalf("Expected version 0.1.0, got: %s", res.Cookbooks[0].Versions[0])
  }
  if cc.Cookbooks[0].Versions[0] != "0.1.0" {
    t.Fatalf("Expected version 0.1.0, got: %s", cc.Cookbooks[0].Versions[0])
  }
}

func Test_universe_1(t *testing.T) {
  ts := start_httptest()
  defer ts.Close()
  cc := CookbookCollection{}
  univ, err := cc.universe(ts.URL)
  if err != nil {
    t.Fatalf("Expected no error from universe, got: %s", err)
  }
  res := string(univ)
  if res != http_response {
    t.Fatalf("Expected HTTP response '%s', got '%s'", http_response, res)
  }
}
