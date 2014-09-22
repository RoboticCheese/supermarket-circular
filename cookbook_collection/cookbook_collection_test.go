package cookbook_collection

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/RoboticCheese/supermarket-circular/cookbook"
)

var http_response = "{\"mailcatcher\":{\"0.1.0\":{\"location_type\":\"opscode\",\"location_path\":\"https://supermarket.getchef.com/api/v1\",\"download_url\":\"https://supermarket.getchef.com/api/v1/cookbooks/mailcatcher/versions/0.1.0/download\",\"dependencies\":{}}}}"

func start_httptest(response string) (*httptest.Server) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, response)
  }))
  return ts
}

func Test_Update_1(t *testing.T) {
  ts := start_httptest(http_response)
  cc, err := new(CookbookCollection).NewFromUniverse(ts.URL)
  if err != nil {
    t.Fatalf("Expected no error, got: %s", err)
  }
  ts.Close()
  new_response := "{\"mailcatcher\":{\"0.1.0\":{\"location_type\":\"opscode\",\"location_path\":\"https://supermarket.getchef.com/api/v1\",\"download_url\":\"https://supermarket.getchef.com/api/v1/cookbooks/mailcatcher/versions/0.1.0/download\",\"dependencies\":{}}, \"0.2.0\":{\"location_type\":\"opscode\",\"location_path\":\"https://supermarket.getchef.com/api/v1\",\"download_url\":\"https://supermarket.getchef.com/api/v1/cookbooks/mailcatcher/versions/0.2.0/download\",\"dependencies\":{}}},\"nailcatcher\":{\"0.1.0\":{\"location_type\":\"opscode\",\"location_path\":\"https://supermarket.getchef.com/api/v1\",\"download_url\":\"https://supermarket.getchef.com/api/v1/cookbooks/nailcatcher/versions/0.1.0/download\",\"dependencies\":{}}}}"
  ts = start_httptest(new_response)
  defer ts.Close()
  cc.URL = ts.URL
  diff, err := cc.Update()
  if err != nil {
    t.Fatalf("Expected no error, got: %s", err)
  }
  if len(diff.Cookbooks) != 2 {
    t.Fatalf("Expected diff of 2 cookbooks, got: %d", len(diff.Cookbooks))
  }
  if len(diff.Cookbooks[0].Versions) != 1 {
    t.Fatalf("Expected 1 version, got: %d", len(diff.Cookbooks[0].Versions))
  }
  if len(diff.Cookbooks[1].Versions) != 1 {
    t.Fatalf("Expected 1 version, got: %d", len(diff.Cookbooks[1].Versions))
  }
  if len(cc.Cookbooks) != 2 {
    t.Fatalf("Expected cc of 2 cookbooks, got: %d", len(cc.Cookbooks))
  }
  if len(cc.Cookbooks[0].Versions) != 2 {
    t.Fatalf("Expected 2 versions, got: %d", len(cc.Cookbooks[0].Versions))
  }
  if len(cc.Cookbooks[1].Versions) != 1 {
    t.Fatalf("Expected 1 version, got: %d", len(cc.Cookbooks[1].Versions))
  }
}

func Test_NewFromUniverse_1(t *testing.T) {
  ts := start_httptest(http_response)
  defer ts.Close()
  cc, err := new(CookbookCollection).NewFromUniverse(ts.URL)
  if err != nil {
    t.Fatalf("Expected no error, got: %s", err)
  }
  if cc.URL != ts.URL {
    t.Fatalf("Expected set collection URL, got: %s", cc.URL)
  }
  if len(cc.Cookbooks) != 1 {
    t.Fatalf("Expected collection length 1, got: %d", len(cc.Cookbooks))
  }
  if cc.Cookbooks[0].Name != "mailcatcher" {
    t.Fatalf("Expected mailcatcher cookbook, got: %s", cc.Cookbooks[0].Name)
  }
  if len(cc.Cookbooks[0].Versions) != 1 {
    t.Fatalf("Expected Versions length 1, got: %d", len(cc.Cookbooks[0].Versions))
  }
  if cc.Cookbooks[0].Versions[0] != "0.1.0" {
    t.Fatalf("Expected version 0.1.0, got: %s", cc.Cookbooks[0].Versions[0])
  }
}

func Test_Merge_1(t *testing.T) {
  oldcc := CookbookCollection{"", []cookbook.Cookbook{
    {"apache", []string{"0.1.0"}},
    {"bpache", []string{"1.0.0"}},
  }}
  newcc := CookbookCollection{"", []cookbook.Cookbook{
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

func Test_Contains_1(t *testing.T) {
  cb := cookbook.Cookbook{"apache", []string{}}
  cc := CookbookCollection{}
  res := cc.Contains(cb)
  if res != false {
    t.Fatalf("Expected false, got: %s", res)
  }
}

func Test_Contains_2(t *testing.T) {
  cb := cookbook.Cookbook{"apache", []string{}}
  cc := CookbookCollection{"", []cookbook.Cookbook{cb}}
  res := cc.Contains(cb)
  if res != true {
    t.Fatalf("Expected true, got: %s", res)
  }
}

func Test_cc_from_universe_1(t *testing.T) {
  ts := start_httptest(http_response)
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
  ts := start_httptest(http_response)
  defer ts.Close()
  cc := CookbookCollection{}
  cc.URL = ts.URL
  univ, err := cc.universe()
  if err != nil {
    t.Fatalf("Expected no error from universe, got: %s", err)
  }
  res := string(univ)
  if res != http_response {
    t.Fatalf("Expected HTTP response '%s', got '%s'", http_response, res)
  }
}
