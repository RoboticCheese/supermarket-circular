package cookbook

type Cookbook struct {
  Name string
  Versions []string
}

func (c *Cookbook) Merge(new_c Cookbook) ([]string, error) {
  for _, version := range new_c.Versions {
    if !c.contains(version) {
      c.Versions = append(c.Versions, version)
    }
  }
  return c.Versions, nil
}

func (c *Cookbook) contains(version string) (bool) {
  for _, v := range c.Versions {
    if v == version {
      return true
    }
  }
  return false
}
