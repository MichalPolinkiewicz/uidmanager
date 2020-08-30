package handlers

import (
	"fmt"
	"net/http"
	"uidmanager/config"
	"uidmanager/manager"
)

func SetBidderUIDEndpoint(cfg *config.Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bidderDetails, err := manager.GetBidderDetails(r, cfg.Bidders)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if !manager.SyncAllowed() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		uid, err := manager.GetUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		bidderUIDCookieName := fmt.Sprintf("%s_%s", bidderDetails, manager.UIDParamName)
		bidderUIDCookieExpiry := manager.GetExpiry(bidderDetails)
		if setSiteCookie := manager.SiteCookieCheck(r.UserAgent()); setSiteCookie {
			http.SetCookie(w, manager.CreateCookie(manager.SameSiteCookieName, manager.SameSiteCookieValue, true, http.SameSiteNoneMode, bidderUIDCookieExpiry))
			http.SetCookie(w, manager.CreateCookie(bidderUIDCookieName, uid, false, http.SameSiteNoneMode, bidderUIDCookieExpiry))
		} else {
			http.SetCookie(w, manager.CreateCookie(bidderUIDCookieName, uid, false, 0, bidderUIDCookieExpiry))
		}
	}
}
