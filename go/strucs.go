package main

type usermovie struct {
	Movie string `json:"movie"`
}

type rcmd struct {
	Movie     string   `json:"movie"`
	Recommend []string `json:"recommend"`
}
