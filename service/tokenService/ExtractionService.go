package tokenService

import (
	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
)

func ExtractFieldFromToken(FieldName string, decodedtoken jwt.MapClaims) interface{} {

	if FieldName == USER_ID {
		if decodedtoken[FieldName] == nil {
			logger.Warnf("Error decodedtoken['%s']: %s", FieldName, decodedtoken[FieldName])
			return nil
		}

		// extract userid
		userId := int(decodedtoken[FieldName].(float64))
		return userId
	}

	return nil
}
