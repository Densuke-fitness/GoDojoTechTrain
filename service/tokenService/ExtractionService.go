package tokenService

import (
	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
)

func ExtractFieldFromToken(FieldName string, FieldType Type, decodedtoken jwt.MapClaims) interface{} {

	if decodedtoken[FieldName] == nil {
		logger.Warnf("Error decodedtoken['%s']: %s", FieldName, decodedtoken[FieldName])
		return nil
	}

	if FieldType == TYPE_INT {
		return int(decodedtoken[FieldName].(float64))
	}

	return nil
}
