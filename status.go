package customrequestresponsewriter

/*
1  .. 10 general
11 .. 20 middleware
21 .. 30 service
31 .. 40 client
41 .. 50 vacant
51 .. 60 server
*/
const (
	STATUSAUTHORIZED      = 11
	STATUSNOTAUTHORIZED   = 12
	STATUSSERVICEFOUND    = 21
	STATUSSERVICENOTFOUND = 22
	STATUSOK              = 1
	STATUSNOTOK           = 2
	STATUSSERVERERROR     = 51
)

var Status = map[int]string{
	STATUSAUTHORIZED:      "AUTHORIZED",
	STATUSNOTAUTHORIZED:   "NOT AUTHORIZED",
	STATUSSERVICEFOUND:    "FOUND",
	STATUSSERVICENOTFOUND: "NOT FOUND",
	STATUSOK:              "OK",
	STATUSNOTOK:           "NOT OK",
	STATUSSERVERERROR:     "SERVER ERROR",
}
