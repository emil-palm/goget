package config;

type RssParser struct {
	Sections map[string]Section
}

type Section struct {
	URLs []string
	Rules []Rule
	DownloadPath string
	Matching struct {
		List []string
		ExternalRss string
	}
}


type Rule struct {
	Pattern string
	Matching string
	First interface{}
	Other interface{}
	All interface{}
}

func ConvertRuleValueToString(val interface{}) (string,bool) {
	if str, ok := val.(string); ok {
		return str,ok
	} else {
		return "", ok
	}
}

func ConvertRuleValueToFloat(val interface{}) (float64,bool) {
	if flt, ok := val.(float64); ok {
		return flt, ok
	} else {
		return 0.0, ok
	}
}