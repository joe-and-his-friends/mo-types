package decoder

import (
	"mo-service/mo-types/model"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type PointTransactionDecoder struct {
}

func (r *PointTransactionDecoder) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	dReader, err := vr.ReadDocument()
	if err != nil {
		return err
	}

	for {
		elem, elemReader, elemErr := dReader.ReadElement()
		if elemErr != nil {
			// if elemErr == io.EOF{
			// 	break
			// }
			return nil
		}

		// fmt.Println(elem)

		switch elem {
		case "systemtransactiondetails":
			details := model.SystemTransactionDetails{}
			r.DecodeValue(dc, elemReader, reflect.ValueOf(&details).Elem())
			val.FieldByName("Details").Set(reflect.ValueOf(details))

		case "thirdpartytransactiondetails":
			details := model.ThirdPartyTransactionDetails{}
			r.DecodeValue(dc, elemReader, reflect.ValueOf(&details).Elem())
			val.FieldByName("Details").Set(reflect.ValueOf(details))

		case "_id":
			fVal, err := elemReader.ReadObjectID()
			if err != nil {
				return err
			}

			val.FieldByName("ID").Set(reflect.ValueOf(fVal))
		case "accumulated":
			fVal, err := elemReader.ReadBoolean()
			if err != nil {
				return err
			}

			val.FieldByName("Accumulated").SetBool(fVal)
		case "createdat":
			fVal, err := elemReader.ReadDateTime()
			if err != nil {
				return err
			}

			val.FieldByName("CreatedAt").SetInt(fVal)
		case "updatedat":
			fVal, err := elemReader.ReadDateTime()
			if err != nil {
				return err
			}

			val.FieldByName("UpdatedAt").SetInt(fVal)
		case "sourceentity":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("SourceEntity").SetString(fVal)
		case "type":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("Type").SetString(fVal)
		case "amount":
			fVal, err := elemReader.ReadInt64()
			if err != nil {
				return err
			}

			val.FieldByName("Amount").SetInt(fVal)
		case "userid":
			fVal, err := elemReader.ReadObjectID()
			if err != nil {
				return err
			}

			val.FieldByName("UserID").Set(reflect.ValueOf(&fVal))
		case "remarks":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("Remarks").SetString(fVal)
		case "customeremail":
			fVal, err := elemReader.ReadString()

			if err != nil {
				return err
			}
			val.FieldByName("CustomerEmail").SetString(fVal)
		case "customerphone":
			fVal, err := elemReader.ReadString()

			if err != nil {
				return err
			}
			val.FieldByName("CustomerPhone").SetString(fVal)
		case "customername":
			fVal, err := elemReader.ReadString()

			if err != nil {
				return err
			}
			val.FieldByName("CustomerName").SetString(fVal)
		case "referenceid":
			fVal, err := elemReader.ReadObjectID()
			if err != nil {
				return err
			}

			val.FieldByName("ReferenceID").Set(reflect.ValueOf(fVal))
		case "merchantid":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("MerchantID").SetString(fVal)
		case "merchantname":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("MerchantName").SetString(fVal)
		case "orderid":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("OrderID").SetString(fVal)
		case "ordernumber":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("OrderNumber").SetString(fVal)
		case "orderpayment":
			payment := model.OrderPayment{}
			r.DecodeValue(dc, elemReader, reflect.ValueOf(&payment).Elem())
			val.FieldByName("OrderPayment").Set(reflect.ValueOf(&payment))
		case "paymentmethodid":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("PaymentMethodID").Set(reflect.ValueOf(fVal))
		case "paymentstatus":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("PaymentStatus").Set(reflect.ValueOf(fVal))
		case "paymenttype":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("PaymentType").Set(reflect.ValueOf(fVal))
		case "total":
			fVal, err := elemReader.ReadDouble()
			if err != nil {
				return err
			}

			val.FieldByName("Total").Set(reflect.ValueOf(fVal))
		case "totallabel":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("TotalLabel").Set(reflect.ValueOf(fVal))
		case "nametranslations":
			translations := model.NameTranslations{}
			r.DecodeValue(dc, elemReader, reflect.ValueOf(&translations).Elem())
			val.FieldByName("NameTranslations").Set(reflect.ValueOf(&translations))
		case "en":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("En").Set(reflect.ValueOf(fVal))
		case "zhhant":
			fVal, err := elemReader.ReadString()
			if err != nil {
				return err
			}

			val.FieldByName("ZhHant").Set(reflect.ValueOf(fVal))
		default:
			err := elemReader.Skip()
			if err != nil {
				return err
			}
		}
	}
}
