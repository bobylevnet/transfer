package Authok

import st "transfer/st"

/*access_token - созданный токен.
user_id - пользователь для которого нужно дать авторизацию.
client_id - приложение. Обратите внимание, что можно указать любое время жизни токена, а не только час времени, используемый по умолчанию в обычной авторизации.
expires - дата истечения токена.
scope - требуемые скоупы.*/

//Appbitrix - токен под которым приложение зарегистрированной в bitrix
type Appbitrix struct {
	apptoken string
}

func request() {

}

func Auth() st.Auhthb {

	request()

	auth := st.Auhthb{
		AccessToken: "88888888",
		UserID:      1,
		ClientID:    1,
		Expires:     60,
		Scope:       "",
		UserName:    "bobylev.ss",
	}

	return auth

}
