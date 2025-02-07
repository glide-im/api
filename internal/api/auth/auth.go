package auth

import (
	"errors"
	"fmt"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/auth"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/dao/wrapper/app"
	"github.com/glide-im/api/internal/dao/wrapper/category"
	"github.com/glide-im/api/internal/dao/wrapper/collect"
	"github.com/glide-im/api/internal/dao/wrapper/tm"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/messages"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Interface interface {
	AuthToken(info *route.Context, req *AuthTokenRequest) error
	SignIn(info *route.Context, req *SignInRequest) error
	Logout(info *route.Context) error
	Register(info *route.Context, req *RegisterRequest) error
}

var (
	ErrInvalidToken      = comm2.NewApiBizError(1001, "token 已失效，请重新登录")
	ErrSignInAccountInfo = comm2.NewApiBizError(1002, "密码错误")
	ErrReplicatedLogin   = comm2.NewApiBizError(1003, "replicated login")
)

var (
	host = []string{
		fmt.Sprintf("ws://%s/ws", "127.0.0.1:8080"),
	}
)

type AuthApi struct {
}

func (*AuthApi) AuthToken(ctx *route.Context, req *AuthTokenRequest) error {

	authInfo, err := auth.ParseToken(req.Token)
	if err != nil {
		return comm2.NewApiBizError(http.StatusUnauthorized, "token 已失效，请重新登录")
	}

	user, err := userdao.Dao.GetUser(authInfo.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	secret := user.MessageDeliverSecret
	if user.MessageDeliverSecret == "" {
		secret = randomStr(32)
		err = userdao.Dao.UpdateSecret(user.Uid, secret)
		if err != nil {
			return err
		}
	}

	credentials := gate.ClientAuthCredentials{
		Type:       0,
		UserID:     strconv.FormatInt(authInfo.Uid, 10),
		DeviceID:   strconv.FormatInt(authInfo.Device, 10),
		DeviceName: "",
		Secrets: &gate.ClientSecrets{
			MessageDeliverSecret: secret,
		},
		ConnectionID: "",
		Timestamp:    time.Now().UnixMilli(),
	}

	credential, err := auth.GenerateCredentials(&credentials)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	token, err := auth.GenerateTokenExpire(user.Uid, authInfo.Device, 3*24)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	resp := AuthResponse{
		Credential: credential,
		Token:      token,
		Uid:        authInfo.Uid,
		Servers:    host,
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp))
	return nil
}

func (*AuthApi) SignIn(ctx *route.Context, request *SignInRequest) error {
	user, err := userdao.Dao.GetUidInfoByLogin(request.Email, request.Password)
	if err != nil || user.Uid == 0 {
		if err == common.ErrNoRecordFound || user.Uid == 0 {
			return ErrSignInAccountInfo
		}
		return comm2.NewDbErr(err)
	}
	expire := 24 * 3
	if user.Role == 99 {
		expire = 24 * 360
	}
	token, err := auth.GenerateTokenExpire(user.Uid, request.Device, int64(expire))
	if err != nil {
		return comm2.NewDbErr(err)
	}

	tk := AuthResponse{
		Uid:      user.Uid,
		Token:    token,
		Servers:  host,
		NickName: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Device:   request.Device,
	}
	appProfile := app.AppDao.GetAppProfile(user.Uid)
	tk.App = appProfile
	resp := messages.NewMessage(ctx.Seq, comm2.ActionSuccess, tk)

	ctx.Uid = user.Uid
	ctx.Device = request.Device
	ctx.Response(resp)
	return nil
}

func (*AuthApi) SignInV2(ctx *route.Context, request *SignInRequest) error {
	user, err := userdao.Dao.GetUidInfoByLogin(request.Email, request.Password)
	if err != nil || user.Uid == 0 {
		if err == common.ErrNoRecordFound || user.Uid == 0 {
			return ErrSignInAccountInfo
		}
		return comm2.NewDbErr(err)
	}

	secret := user.MessageDeliverSecret
	if secret == "" {
		secret = randomStr(32)
		err = userdao.Dao.UpdateSecret(user.Uid, secret)
		user.MessageDeliverSecret = secret
		if err != nil {
			return err
		}
	}

	credentials := gate.ClientAuthCredentials{
		Type:       0,
		UserID:     strconv.FormatInt(user.Uid, 10),
		DeviceID:   strconv.FormatInt(request.Device, 10),
		DeviceName: "",
		Secrets: &gate.ClientSecrets{
			MessageDeliverSecret: secret,
		},
		ConnectionID: "",
		Timestamp:    time.Now().UnixMilli(),
	}

	credential, err := auth.GenerateCredentials(&credentials)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	expire := 24 * 3
	if user.Role == 99 {
		expire = 24 * 360
	}
	token, err := auth.GenerateTokenExpire(user.Uid, request.Device, int64(expire))
	if err != nil {
		return comm2.NewDbErr(err)
	}

	tk := AuthResponse{
		Uid:        user.Uid,
		Token:      token,
		Credential: credential,
		Servers:    host,
		NickName:   user.Nickname,
		Email:      user.Email,
		Phone:      user.Phone,
		Device:     request.Device,
	}
	tk.App = app.AppDao.GetAppProfile(user.Uid)
	ctx.Uid = user.Uid
	ctx.Device = request.Device
	ctx.ReturnSuccess(tk)
	return nil
}

func (*AuthApi) GuestRegister(ctx *route.Context, req *GuestRegisterRequest) error {

	avatar := req.Avatar
	nickname := req.Nickname

	if len(avatar) == 0 {
		avatar = ""
	}
	account := "guest_" + randomStr(32)
	if len(nickname) == 0 {
		nickname = account
	}

	hash := userdao.PasswordHash("-")
	secret := randomStr(32)
	u := &userdao.User{
		Email:                account,
		Account:              account,
		Phone:                account,
		Password:             hash,
		Nickname:             nickname,
		Avatar:               avatar,
		MessageDeliverSecret: secret,
	}
	err := userdao.UserInfoDao.AddUser(u)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	user, err := userdao.Dao.GetUidInfoByLogin(account, "-")
	if err != nil || user.Uid == 0 {
		if err == common.ErrNoRecordFound || user.Uid == 0 {
			return comm2.NewApiBizError(1011, "登录失败")
		}
		return comm2.NewDbErr(err)
	}

	token, err := auth.GenerateTokenExpire(user.Uid, auth.GUEST_DEVICE, 24*7)

	credentials := gate.ClientAuthCredentials{
		Type:       0,
		UserID:     strconv.FormatInt(user.Uid, 10),
		DeviceID:   strconv.FormatInt(auth.GUEST_DEVICE, 10),
		DeviceName: "",
		Secrets: &gate.ClientSecrets{
			MessageDeliverSecret: secret,
		},
		ConnectionID: "",
		Timestamp:    time.Now().UnixMilli(),
	}

	credential, err := auth.GenerateCredentials(&credentials)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	tk := AuthResponse{
		Credential: credential,
		Uid:        user.Uid,
		Token:      token,
		Servers:    host,
		NickName:   user.Nickname,
	}
	ctx.ReturnSuccess(&tk)
	return nil
}

func (*AuthApi) GuestRegisterV2(ctx *route.Context, req *GuestRegisterV2Request) error {
	fingerprintId := req.FingerprintId
	var err error
	isAccount := true

	app_id := app.AppDao.GetAppID(ctx.Context.GetHeader("Host-A"))
	if app_id == 0 {
		return comm2.NewApiBizError(4001, "访问异常")
	}

	secret := randomStr(32)
	user := &userdao.User{
		AppID:                app_id,
		Account:              fingerprintId,
		Password:             "",
		Nickname:             fingerprintId,
		FingerprintId:        fingerprintId,
		Avatar:               "",
		Role:                 2,
		MessageDeliverSecret: secret,
	}

	db.DB.Model(&userdao.User{}).Where("account = ?", fingerprintId).Find(&user)
	if user.Uid == 0 {
		isAccount = false
	}
	collectData := collect.GetUserUa(ctx)
	region := collect.GetIpAddr(collectData.Ip)

	if !isAccount {
		user.Nickname = fmt.Sprintf("%s(%s)", region, fingerprintId)
		err = userdao.UserInfoDao.AddGuestUser(user)
		if err != nil {
			return comm2.NewDbErr(err)
		}
	}

	collectData.AppID = app_id
	collectData.Device = "phone"
	collectData.Origin = req.Origin
	collectData.Uid = user.Uid
	collectData.Region = region
	collect.CollectDataDao.UpdateOrCreate(collectData)

	token, err := auth.GenerateTokenExpire(user.Uid, 3, 24*7)

	credentials := gate.ClientAuthCredentials{
		Type:       0,
		UserID:     strconv.FormatInt(user.Uid, 10),
		DeviceID:   strconv.FormatInt(auth.GUEST_DEVICE, 10),
		DeviceName: "",
		Secrets: &gate.ClientSecrets{
			MessageDeliverSecret: secret,
		},
		ConnectionID: "",
		Timestamp:    time.Now().UnixMilli(),
	}

	credential, err := auth.GenerateCredentials(&credentials)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	tk := GuestAuthResponse{
		Credential: credential,
		Uid:        user.Uid,
		Token:      token,
		Servers:    host,
		AppID:      app_id,
		NickName:   user.Nickname,
	}
	ctx.ReturnSuccess(&tk)
	return nil
}

func (*AuthApi) Register(ctx *route.Context, req *RegisterRequest) error {

	exists, err := userdao.UserInfoDao.AccountExists(req.Email)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if exists {
		return comm2.NewApiBizError(1004, "帐户已存在")
	}
	err = tm.VerifyCodeU.ValidateVerifyCode(req.Email, req.Captcha)
	if err != nil {
		return err
	}

	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	nickname := req.Nickname
	if len(nickname) == 0 {
		nickname = req.Email
	}
	secret := randomStr(32)
	u := &userdao.User{
		Account:              req.Email,
		Password:             userdao.PasswordHash(req.Password),
		Email:                req.Email,
		Nickname:             nickname,
		MessageDeliverSecret: secret,
		//Avatar:   nil,
	}
	err = userdao.UserInfoDao.AddUser(u)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	appU := &app.App{
		Name:   req.Email,
		Uid:    u.Uid,
		Status: 0,
		Logo:   "",
		Email:  req.Email,
		Phone:  "",
		Host:   "",
	}
	err = app.AppDao.AddApp(appU)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	userdao.UserInfoDao.UpdateAppId(u.Uid, appU.Id)
	category.CategoryUserDao.InitCategory(appU.Id)
	tm.VerifyCodeU.ClearLimit(req.Email)
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return err
}

func (*AuthApi) Forget(ctx *route.Context, req *RegisterRequest) error {

	exists, err := userdao.UserInfoDao.AccountExists(req.Email)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if !exists {
		return comm2.NewApiBizError(1004, "账户不存在")
	}

	err = tm.VerifyCodeU.ValidateVerifyCode(req.Email, req.Captcha)
	if err != nil {
		return err
	}
	var user userdao.User
	db.DB.Model(&userdao.User{}).Where("email = ?", req.Email).Find(&user)
	err = userdao.UserInfoDao.UpdatePassword(user.Uid, userdao.PasswordHash(req.Password))
	if err != nil {
		return comm2.NewDbErr(err)
	}

	tm.VerifyCodeU.ClearLimit(req.Email)
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return err
}

func (a *AuthApi) Logout(ctx *route.Context) error {
	err := userdao.Dao.DelAuthToken(ctx.Uid, ctx.Device)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	_ = im.IM.Exit(strconv.FormatInt(ctx.Uid, 10), strconv.FormatInt(ctx.Device, 10))
	return nil
}

func (a *AuthApi) VerifyCode(ctx *route.Context, req *VerifyCodeRequest) error {
	if req.Mode == "register" {
		exists, _ := userdao.UserInfoDao.AccountExists(req.Email)
		if exists {
			return errors.New("用户已存在，快去登录吧")
		}
	}
	if req.Mode == "login" || req.Mode == "forget" {
		exists, _ := userdao.UserInfoDao.AccountExists(req.Email)
		if !exists {
			return errors.New("用户不存在, 请先注册吧")
		}
	}
	err := tm.VerifyCodeU.SendVerifyCode(req.Email, "resources/auth/login.html")
	if err != nil {
		return err
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return err
}

func randomStr(n int) string {
	var l = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(l)
	for i := range b {
		b[i] = l[rand.Intn(length)]
	}
	return string(b)
}
