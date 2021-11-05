package token_manager

import (
    "api/golang-song-api/confi"
	//"math/rand"
	//"encoding/hex"
	"time"
    "strings"
    "errors"
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
)

//const TokenLength = 32

type JWTData struct {
    // Standard claims are the standard jwt claims from the IETF standard
    // https://tools.ietf.org/html/rfc7519
    jwt.StandardClaims
    CustomClaims map[string]string `json:"custom,omitempty"`
}

/*func GenerateSecureToken() string {
    b := make([]byte, TokenLength)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}*/

func GeneratJwcToken(user string) string {
	// Creating a signed token
	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		CustomClaims: map[string]string{
			"userName": user,
		},
	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString([]byte(confi.SECRET))
    return tokenString
}

func ValidateJwcToken(r *http.Request) (*JWTData, error) {
    jwtToken := r.Header.Get("Authorization")
	corsHell := r.Header.Get("Sec-Fetch-Mode")

	if (jwtToken == "" && corsHell != "") {
		return nil, nil
	}

    authArr := strings.Split(jwtToken, ".")

    if len(authArr) != 3 {
    	err_msg := "Authentication header is invalid: " + jwtToken
        println(err_msg)
        return nil, errors.New(err_msg)
    }

    claims, err := jwt.ParseWithClaims(jwtToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
        if jwt.SigningMethodHS256 != token.Method {
            return nil, errors.New("Invalid signing algorithm")
        }
        return []byte(confi.SECRET), nil
    })

	data := claims.Claims.(*JWTData)

    return data, err
}

func GetNewTokenExpirationDate() time.Time {
	return time.Now().Add(time.Hour * 1)
}

func GetTokenRefreshDate() time.Time {
    return time.Now().Add(time.Minute * (-30))
}