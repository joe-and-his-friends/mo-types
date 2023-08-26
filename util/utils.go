package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func LowerFirstLetter(inputStr string) (str string) {
	str += strings.ToLower(string(inputStr[0]))
	str += inputStr[1:]
	return
}

func SnakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(v))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return
}

func CamelCaseToSnakeCase(inputCamelCaseStr string) (snakeCase string) {

	for k, v := range inputCamelCaseStr {
		if k == 0 {
			snakeCase = strings.ToLower(string(v))
		} else {
			if unicode.IsUpper(v) && unicode.IsLetter(v) {
				snakeCase += "_"
				snakeCase += strings.ToLower(string((v)))
			} else {
				snakeCase += string(v)
			}
		}
	}

	return

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func IsEmail(phoneOrEmail string) bool {
	return strings.Contains(phoneOrEmail, "@")
}

func DomainFromAddress(address string) string {
	strGroups := strings.Split(address, "@")

	if len(strGroups) > 1 {
		return strGroups[1]
	}

	return ""
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func DivideHans(input string) (output string) {
	runeValues := []rune(input)
	for i := 0; i < len(runeValues); i++ {
		runeValue := runeValues[i]

		if unicode.Is(unicode.Han, runeValues[i]) {
			if i > 1 {
				previousRuneValue := runeValues[i-1]
				if !unicode.Is(unicode.Han, previousRuneValue) {
					output += " "
				}
			}

			if i+1 < len(runeValues) {
				nextRuneValue := runeValues[i+1]
				if unicode.Is(unicode.Han, nextRuneValue) {
					output += string(runeValue) + string(nextRuneValue) + " "
				}
			}
		} else {
			output += string(runeValue)
		}
	}
	return output
}

func DivideWords(input string, depth int) (output string) {
	runeValues := []rune(input)

	for i := 0; i <= len(runeValues)-depth; i++ {
		str := string(runeValues[i])

		for j := 1; j < depth; j++ {
			str += string(runeValues[i+j])
		}

		output += str + " "
	}

	// fmt.Println(output)

	if depth < len(runeValues)-1 {
		output += DivideWords(input, depth+1)
	}

	return output
}

func DividePreprocessedWords(input string) (resultingStr string) {
	for _, str := range strings.Split(input, " ") {
		resultingStr += DivideWords(str, 1)
	}

	return resultingStr + input
}

func GenerateRedemptionCode(codeLength int, sed string) string {
	var code = ""
	var possible = "0123456789abcdefghijklmnopqrstuvwxyz" + sed
	max := len(possible)
	for i := 0; i < codeLength; i++ {
		code += string(possible[rand.Intn(max)])
	}
	return code
}

func InjectField(source interface{}, fieldName string, fieldValue interface{}) interface{} {
	bytes, err := bson.Marshal(source)

	if err != nil {
		return nil
	}

	injected := bson.M{}
	err = bson.Unmarshal(bytes, &injected)
	injected[fieldName] = fieldValue

	if err != nil {
		return nil
	}

	return injected
}

func StructToBsonDoc(source interface{}) bson.M {
	bytes, err := bson.Marshal(source)

	if err != nil {
		return nil
	}

	doc := bson.M{}
	err = bson.Unmarshal(bytes, &doc)

	if err != nil {
		return nil
	}

	return doc
}

func StructToBsonD(source interface{}) bson.D {
	bytes, err := bson.Marshal(source)

	if err != nil {
		return nil
	}

	doc := bson.D{}
	err = bson.Unmarshal(bytes, &doc)

	if err != nil {
		return nil
	}

	return doc
}

func StructToFlattenedBsonDoc(source interface{}) bson.M {
	return FlattenBsonDoc(StructToBsonDoc(source))
}

func FlattenBsonDoc(source bson.M) bson.M {
	prefixed := bson.M{}

	for k, v := range source {
		if subDoc, ok := v.(bson.M); ok {
			subSubDoc := FlattenBsonDoc(subDoc)
			for k2, v2 := range subSubDoc {
				prefixed[fmt.Sprintf("%s.%s", k, k2)] = v2
			}
		} else {
			prefixed[k] = v
		}
	}

	return prefixed
}

func StructToFlattenedBsonD(source interface{}) bson.D {
	return FlattenBsonD(StructToBsonD(source))
}

func FlattenBsonD(source bson.D) bson.D {
	prefixed := bson.D{}

	for _, v := range source {
		if subDoc, ok := v.Value.(bson.D); ok {
			subSubDoc := FlattenBsonD(subDoc)
			for _, v2 := range subSubDoc {
				// prefixed[fmt.Sprintf("%s.%s", k, k2)] = v2
				prefixed = append(prefixed, bson.E{Key: fmt.Sprintf("%s.%s", v.Key, v2.Key), Value: v2.Value})

			}
		} else {
			prefixed = append(prefixed, v)
		}
	}

	return prefixed
}

func PrefixToBsonDoc(source bson.M, prefix string) bson.M {
	prefixed := bson.M{}
	for k, v := range source {
		prefixed[prefix+k] = v
	}

	return prefixed
}

func BsonDocToJson(docs []bson.D) string {
	buf := bytes.NewBuffer(nil)
	vw, err := bsonrw.NewExtJSONValueWriter(buf, true, false)

	if err != nil {
		fmt.Println(err)
	}

	enc, err := bson.NewEncoder(vw)

	for _, doc := range docs {
		if err != nil {
			break
		}

		err = enc.Encode(doc)
		if err != nil {
			break
		}
	}

	return buf.String()
}

func ConvertStringToDatetime(dateStr string) (primitive.DateTime, error) {
	ymd := strings.Split(dateStr, "-")

	mi, _ := strconv.Atoi(ymd[1])
	di, _ := strconv.Atoi(ymd[2])

	normalized := fmt.Sprintf("%s-%02d-%02dT00:00:00Z", ymd[0], mi, di)
	time, err := time.Parse(time.RFC3339, normalized)

	if err != nil {
		return 0, err
	}

	return primitive.NewDateTimeFromTime(time), nil
}

func TruncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
}
