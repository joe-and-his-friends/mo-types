package util

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func StructToProtobuf(input interface{}, output interface{}) {
	iVal := reflect.ValueOf(input)
	oVal := reflect.ValueOf(output)

	for iVal.Kind() != reflect.Struct {
		if iVal.IsZero() {
			return
		}

		iVal = iVal.Elem()
	}

	msg := oVal.Interface().(proto.Message)
	fields := msg.ProtoReflect().Descriptor().Fields()

	for i := 0; i < iVal.NumField(); i++ {
		ifName := iVal.Type().Field(i).Name
		ifVal := iVal.Field(i)
		ofName := LowerFirstLetter(ifName)

		ofD := fields.ByJSONName(ofName)

		// fmt.Printf("%v: %v\n", ifName, ofD)

		if ifVal.Kind() == reflect.Slice {
			if ofD == nil {
				continue
			}
			ofVal := msg.ProtoReflect().Mutable(ofD)

			for j := 0; j < ifVal.Len(); j++ {
				ifsVal := ifVal.Index(j)

				if ifsVal.Kind() == reflect.Ptr {
					ele := ofVal.List().NewElement()
					StructToProtobuf(ifsVal.Interface(), ele.Message().Interface())
					ofVal.List().Append(ele)
				} else {
					if ifsVal.Type().String() == "primitive.ObjectID" && ofD.Kind() == protoreflect.StringKind {
						ofVal.List().Append(protoreflect.ValueOf(ifsVal.Interface().(primitive.ObjectID).Hex()))
					} else if ofD.Kind() == protoreflect.Int32Kind {
						ofVal.List().Append(protoreflect.ValueOf(int32(ifsVal.Int())))
					} else if ofD.Kind() == protoreflect.Int64Kind {
						ofVal.List().Append(protoreflect.ValueOf(ifsVal.Int()))
					} else if ofD.Kind() == protoreflect.StringKind {
						ofVal.List().Append(protoreflect.ValueOf(ifsVal.String()))
					} else {
						ofVal.List().Append(protoreflect.ValueOf(ifsVal.Interface()))
					}
				}
			}
		} else if ifVal.Kind() == reflect.Ptr {
			if !ifVal.IsZero() && ofD != nil {
				ofVal := msg.ProtoReflect().Mutable(ofD)
				StructToProtobuf(ifVal.Interface(), ofVal.Message().Interface())
			}
		} else {
			if ifVal.IsValid() && ofD != nil {
				if ifVal.Type().String() == "primitive.ObjectID" && ofD.Kind() == protoreflect.StringKind {
					msg.ProtoReflect().Set(ofD, protoreflect.ValueOf(ifVal.Interface().(primitive.ObjectID).Hex()))
				} else if ofD.Kind() == protoreflect.Int32Kind {
					msg.ProtoReflect().Set(ofD, protoreflect.ValueOf(int32(ifVal.Int())))
				} else if ofD.Kind() == protoreflect.Int64Kind {
					msg.ProtoReflect().Set(ofD, protoreflect.ValueOf(ifVal.Int()))
				} else if ofD.Kind() == protoreflect.StringKind {
					msg.ProtoReflect().Set(ofD, protoreflect.ValueOf(ifVal.String()))
				} else {
					msg.ProtoReflect().Set(ofD, protoreflect.ValueOf(ifVal.Interface()))
				}
			}
		}
	}
}

func ProtobufToStruct(input interface{}, output interface{}) {
	iVal := reflect.ValueOf(input)
	msg := iVal.Interface().(proto.Message)
	fields := msg.ProtoReflect().Descriptor().Fields()
	oVal := reflect.ValueOf(output)

	if iVal.IsZero() {
		return
	}

	for iVal.Kind() != reflect.Struct {
		iVal = iVal.Elem()
	}

	for oVal.Kind() != reflect.Struct {
		oVal = oVal.Elem()
	}

	for i := 0; i < fields.Len(); i++ {
		ifD := fields.Get(i)
		ifName := string(ifD.Name())
		ofName := SnakeCaseToCamelCase(ifName)
		ifVal := iVal.FieldByName(ofName)

		// fmt.Printf("%v: %v\n", ifName, ifVal)

		if ifVal.IsValid() {
			ofVal := oVal.FieldByName(ofName)

			if ifD.IsList() {
				if ofVal.IsValid() {
					iList := msg.ProtoReflect().Get(ifD).List()

					oList := reflect.MakeSlice(ofVal.Type(), 0, iList.Len())

					for j := 0; j < iList.Len(); j++ {
						ifsVal := iList.Get(j)
						iV := reflect.ValueOf(ifsVal.Interface())

						if iV.Kind() == reflect.Ptr {
							oV := reflect.New(ofVal.Type().Elem().Elem())
							ProtobufToStruct(ifsVal.Message().Interface(), oV.Interface())
							// fmt.Println(oV)
							oList = reflect.Append(oList, oV)
						} else {
							if ofVal.Type().Elem().Kind() == reflect.Int {
								oList = reflect.Append(oList, reflect.ValueOf((int(ifsVal.Int()))))
							} else if ofVal.Type().Elem().String() == "primitive.DateTime" {
								oList = reflect.Append(oList, reflect.ValueOf(primitive.DateTime(ifsVal.Int())))
							} else if ofVal.Type().Elem().String() == "primitive.ObjectID" {
								id, _ := primitive.ObjectIDFromHex(ifsVal.String())
								oList = reflect.Append(oList, reflect.ValueOf(id))
							} else if ofVal.Type().Elem().Kind() == reflect.String {
								oList = reflect.Append(oList, reflect.ValueOf(ifsVal.String()))
							} else if ofVal.Type().Elem().Kind() == reflect.Float64 {
								oList = reflect.Append(oList, reflect.ValueOf(ifsVal.Float()))
							} else {
								oList = reflect.Append(oList, reflect.ValueOf(ifsVal.Interface()))
							}

						}

					}

					ofVal.Set(oList)
				}
			} else if ifD.Kind() == protoreflect.MessageKind {
				if !ifVal.IsNil() && ofVal.IsValid() {
					ofVal.Set(reflect.New(ofVal.Type().Elem()))
					ProtobufToStruct(msg.ProtoReflect().Get(ifD).Message().Interface(), ofVal.Interface())
				}
			} else {
				if ofVal.IsValid() {
					if ofVal.Kind() == reflect.Int {
						ofVal.Set(reflect.ValueOf(int(ifVal.Int())))
					} else if ofVal.Type().String() == "primitive.DateTime" {
						ofVal.Set(reflect.ValueOf(primitive.DateTime(ifVal.Int())))
					} else if ofVal.Type().String() == "primitive.ObjectID" {
						id, _ := primitive.ObjectIDFromHex(ifVal.String())
						ofVal.Set(reflect.ValueOf(id))
					} else if ofVal.Kind() == reflect.String {
						ofVal.SetString(ifVal.String())
					} else if ofVal.Kind() == reflect.Float64 {
						ofVal.Set(reflect.ValueOf(ifVal.Float()))
					} else {
						ofVal.Set(reflect.ValueOf(ifVal.Interface()))
					}
				}
			}
		}

	}
}
