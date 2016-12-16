package goT411

import simplejson "github.com/bitly/go-simplejson"

//SubCategory represents a subsection for a Category
type SubCategory struct {
	ID   string `json:"id"`
	Pid  string `json:"pid"`
	Name string `json:"name"`
}

//SubCategories represent a list of SubCategory
type SubCategories []SubCategory

//Category represents a Category
type Category struct {
	ID      string        `json:"id"`
	Pid     string        `json:"pid"`
	Name    string        `json:"name"`
	SubCats SubCategories `json:"cats"`
}

//Categories represent a list of Category
type Categories []Category

//GetCat returns all category possible for a search
func (c Client) GetCat() (map[string]interface{}, error) {
	var out map[string]interface{}

	resp, err := c.requestURI("http://" + apiRoot + "/categories/tree/")
	if err != nil {
		return nil, err
	}
	// err = json.Unmarshal([]byte(resp), &out)
	json, err := simplejson.NewJson([]byte(resp))
	if err != nil {
		return nil, err
	}
	out, err = json.Map()
	if err != nil {
		return nil, err
	}
	return out, nil
}
