package cdn

import (
	"time"

	"github.com/reedchan7/aliyungo/common"
)

const (
	Web                       = "web"
	Download                  = "download"
	Video                     = "video"
	LiveStream                = "liveStream"
	Ipaddr                    = "ipaddr"
	Domain                    = "domain"
	OSS                       = "oss"
	Domestic                  = "domestic"
	Overseas                  = "overseas"
	Global                    = "global"
	ContentType               = "Content-Type"
	CacheControl              = "Cache-Control"
	ContentDisposition        = "Content-Disposition"
	ContentLanguage           = "Content-Language"
	Expires                   = "Expires"
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	AccessControlMaxAge       = "Access-Control-Max-Age"
)

var CdnTypes = []string{Web, Download, Video, LiveStream}
var SourceTypes = []string{Ipaddr, Domain, OSS}
var Scopes = []string{Domestic, Overseas, Global}
var HeaderKeys = []string{ContentType, CacheControl, ContentDisposition, ContentLanguage, Expires, AccessControlAllowMethods, AccessControlAllowOrigin, AccessControlMaxAge}

type CdnCommonResponse struct {
	common.Response
}

type Domains struct {
	DomainName   string
	Cname        string
	CdnType      string
	DomainStatus string
	GmtCreated   string
	GmtModified  string
	Description  string
}

type MonitorDataItem struct {
	TimeStamp         string
	QueryPerSecond    string
	BytesHitRate      string
	BytesPerSecond    string
	RequestHitRate    string
	AverageObjectSize string
}

type TaskItem struct {
	TaskId       string
	ObjectPath   string
	Status       string
	CreationTime time.Time
}

type LogDetail struct {
	LogName   string
	LogPath   string
	LogSize   int32
	StartTime string
	EndTime   string
}
