package req

type EkuiperRuleSetReq struct {
	Streams map[string]string `json:"streams"`
	Tables  map[string]string `json:"tables"`
	Rules   map[string]string `json:"rules"`
}

type EkuiperRuleSetRuleItemReq struct {
	ID      string                   `json:"id"`
	Sql     string                   `json:"sql,omitempty"`
	Graph   string                   `json:"graph,omitempty"`
	Actions []map[string]interface{} `json:"actions,omitempty"`
}

type EkuiperRuleSetImportReq struct {
	Content string `json:"content"`
}
