package categorizer

type Item struct {
	ItemId string `json:"itemId"`
	FeedId string `json:"feedId"`
	Title  string `json:"title"`
}

type AllItems struct {
	Items []Item
}
