package config

const (
	// Parser names
	ParseCity        = "ParseCity"
	ParseCityList    = "ParseCityList"
	ParseProfile     = "ParseProfile"
	NilParser        = "NilParser"
	ParseMatchList   = "ParseMatchList"
	ParseContestList = "ParseContestList"
	// ElasticSearch
	ElasticIndex = "dating_show"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate limiting
	Qps = 20
)
