package nba

type Internal struct {
	ConsolidatedDomKey      string `json:"consolidatedDomKey"`
	DateTime                string `json:"pubDateTime"`
	EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	IgorPath                string `json:"igorPath"`
	RouteName               string `json:"routeName"`
	RouteValue              string `json:"routeValue"`
	XSLT                    string `json:"xslt"`
	XSLTForceRecompile      bool   `json:"xsltForceRecompile"`
	XSLTInCache             bool   `json:"xsltInCache"`
	XSLTCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
	XSLTTransformTimeMillis string `json:"xsltTransformTimeMillis"`
}
