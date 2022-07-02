package api

import (
	"github.com/glide-im/api/internal/api/auth"
)

func initGuestRoute() {
	appGuestApi := auth.AuthApi{}
	postNoGuestAuth("/api/auth/guestV2", appGuestApi.GuestRegisterV2)
}

func postNoGuestAuth(path string, fn interface{}) {
	rt.POST(path, getHandler(path, fn))
}
func getNoGuestAuth(path string, fn interface{}) {
	rt.GET(path, getHandler(path, fn))
}

func guestPost(path string, fn interface{}) {
	useGuestAuth().POST(path, getHandler(path, fn))
}

func guestDelete(path string, fn interface{}) {
	useGuestAuth().DELETE(path, getHandler(path, fn))
}

func guestGet(path string, fn interface{}) {
	useGuestAuth().GET(path, getHandler(path, fn))
}
