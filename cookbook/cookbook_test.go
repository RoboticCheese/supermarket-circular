package cookbook

import (
  "testing"
)

func Test_Merge_1(t *testing.T) {
  c := Cookbook{"test", []string{"0.1.0", "0.2.0"}}
  c2 := Cookbook{"test2", []string{"1.0.0"}}
  expected := []string{"0.1.0", "0.2.0", "1.0.0"}
  res, err := c.Merge(c2)
  if err != nil {
    t.Fatalf("Merge should not have returned an error")
  }
  if len(res) != 3 {
    t.Fatalf("Merge should have returned a three-element slice")
  }
  if len(c.Versions) != 3 {
    t.Fatalf("Merge should have updated Versions to a three-element slice")
  }
  for i, v := range res {
    if v != expected[i] {
      t.Fatalf("Merge should have returned updated Versions slice")
    }
  }
  for i, v := range c.Versions {
    if v != expected[i] {
      t.Fatalf("Merge should have updated cookbook Versions slice")
    }
  }
}

func Test_Merge_2(t *testing.T) {
  c := Cookbook{"test", []string{"0.1.0", "0.2.0"}}
  c2 := Cookbook{"test2", []string{"0.1.0", "0.2.0"}}
  expected := []string{"0.1.0", "0.2.0"}
  res, err := c.Merge(c2)
  if err != nil {
    t.Fatalf("Merge should not have returned an error")
  }
  if len(res) != 2 {
    t.Fatalf("Merge should have returned a three-element slice")
  }
  if len(c.Versions) != 2 {
    t.Fatalf("Merge should have updated Versions to a three-element slice")
  }
  for i, v := range res {
    if v != expected[i] {
      t.Fatalf("Merge should have returned updated Versions slice")
    }
  }
  for i, v := range c.Versions {
    if v != expected[i] {
      t.Fatalf("Merge should have updated cookbook Versions slice")
    }
  }
}

func Test_contains_1(t *testing.T) {
  c := Cookbook{"test", []string{"0.1.0", "0.2.0"}}
  if c.contains("0.2.0") != true {
    t.Fatalf("Cookbook should have contained version and returned true")
  }
}

func Test_contains_2(t *testing.T) {
  c := Cookbook{"test", []string{"0.1.0", "0.2.0"}}
  if c.contains("1.0.0") != false {
    t.Fatalf("Cookbook should have contained version and returned true")
  }
}
