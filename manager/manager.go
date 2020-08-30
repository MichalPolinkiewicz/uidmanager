package manager

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"uidmanager/config"
)

const (
	UIDParamName        = "uid"
	BidderParamName     = "bidder"
	SameSiteCookieName  = "SSCookie"
	SameSiteCookieValue = "1"

	defaultTTL          = 14 * 24 * time.Hour
	chromeStr       = "Chrome/"
	chromeiOSStr    = "CriOS/"
	chromeMinVer    = 67
	chromeStrLen    = len(chromeStr)
	chromeiOSStrLen = len(chromeiOSStr)
)

func CreateCookie(name, value string, secure bool, sameSite http.SameSite, expires time.Time) *http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   "",
		Expires:  expires,
		Secure:   secure,
		SameSite: sameSite,
	}
	return &cookie
}

func GetBidderDetails(r *http.Request, supportedBidders map[string]*config.BidderDetails) (*config.BidderDetails, error) {
	bidderName := r.URL.Query().Get(BidderParamName)
	if bidderName == "" {
		return nil, BidderParamRequiredError
	}
	if bidderDetails, ok := supportedBidders[bidderName]; !ok {
		return nil, UnsupportedBidderError
	} else {
		return bidderDetails, nil
	}
}

func GetUID(r *http.Request) (string, error) {
	uid := r.URL.Query().Get(UIDParamName)
	if uid == "" {
		return "", EmptyUIDError
	}
	return uid, nil
}

func SiteCookieCheck(ua string) bool {
	result := false
	index := strings.Index(ua, chromeStr)
	criOSIndex := strings.Index(ua, chromeiOSStr)
	if index != -1 {
		result = checkChromeBrowserVersion(ua, index, chromeStrLen)
	} else if criOSIndex != -1 {
		result = checkChromeBrowserVersion(ua, criOSIndex, chromeiOSStrLen)
	}
	return result
}

func GetExpiry(bidderDetails *config.BidderDetails) time.Time {
	ttl := defaultTTL
	if bidderDetails.TTL != 0 {
		ttl = bidderDetails.TTL
	}
	return time.Now().Add(ttl)
}

func SyncAllowed() bool {
	return true
}

func checkChromeBrowserVersion(ua string, index int, chromeStrLength int) bool {
	result := false
	vIndex := index + chromeStrLength
	dotIndex := strings.Index(ua[vIndex:], ".")
	if dotIndex == -1 {
		dotIndex = len(ua[vIndex:])
	}
	version, _ := strconv.Atoi(ua[vIndex : vIndex+dotIndex])
	if version >= chromeMinVer {
		result = true
	}
	return result
}
