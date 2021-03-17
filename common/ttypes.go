// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package common

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type TDeviceType int64

const (
	TDeviceType_CPU TDeviceType = 0
	TDeviceType_GPU TDeviceType = 1
)

func (p TDeviceType) String() string {
	switch p {
	case TDeviceType_CPU:
		return "CPU"
	case TDeviceType_GPU:
		return "GPU"
	}
	return "<UNSET>"
}

func TDeviceTypeFromString(s string) (TDeviceType, error) {
	switch s {
	case "CPU":
		return TDeviceType_CPU, nil
	case "GPU":
		return TDeviceType_GPU, nil
	}
	return TDeviceType(0), fmt.Errorf("not a valid TDeviceType string")
}

func TDeviceTypePtr(v TDeviceType) *TDeviceType { return &v }

func (p TDeviceType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TDeviceType) UnmarshalText(text []byte) error {
	q, err := TDeviceTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type TDatumType int64

const (
	TDatumType_SMALLINT            TDatumType = 0
	TDatumType_INT                 TDatumType = 1
	TDatumType_BIGINT              TDatumType = 2
	TDatumType_FLOAT               TDatumType = 3
	TDatumType_DECIMAL             TDatumType = 4
	TDatumType_DOUBLE              TDatumType = 5
	TDatumType_STR                 TDatumType = 6
	TDatumType_TIME                TDatumType = 7
	TDatumType_TIMESTAMP           TDatumType = 8
	TDatumType_DATE                TDatumType = 9
	TDatumType_BOOL                TDatumType = 10
	TDatumType_INTERVAL_DAY_TIME   TDatumType = 11
	TDatumType_INTERVAL_YEAR_MONTH TDatumType = 12
	TDatumType_POINT               TDatumType = 13
	TDatumType_LINESTRING          TDatumType = 14
	TDatumType_POLYGON             TDatumType = 15
	TDatumType_MULTIPOLYGON        TDatumType = 16
	TDatumType_TINYINT             TDatumType = 17
	TDatumType_GEOMETRY            TDatumType = 18
	TDatumType_GEOGRAPHY           TDatumType = 19
)

func (p TDatumType) String() string {
	switch p {
	case TDatumType_SMALLINT:
		return "SMALLINT"
	case TDatumType_INT:
		return "INT"
	case TDatumType_BIGINT:
		return "BIGINT"
	case TDatumType_FLOAT:
		return "FLOAT"
	case TDatumType_DECIMAL:
		return "DECIMAL"
	case TDatumType_DOUBLE:
		return "DOUBLE"
	case TDatumType_STR:
		return "STR"
	case TDatumType_TIME:
		return "TIME"
	case TDatumType_TIMESTAMP:
		return "TIMESTAMP"
	case TDatumType_DATE:
		return "DATE"
	case TDatumType_BOOL:
		return "BOOL"
	case TDatumType_INTERVAL_DAY_TIME:
		return "INTERVAL_DAY_TIME"
	case TDatumType_INTERVAL_YEAR_MONTH:
		return "INTERVAL_YEAR_MONTH"
	case TDatumType_POINT:
		return "POINT"
	case TDatumType_LINESTRING:
		return "LINESTRING"
	case TDatumType_POLYGON:
		return "POLYGON"
	case TDatumType_MULTIPOLYGON:
		return "MULTIPOLYGON"
	case TDatumType_TINYINT:
		return "TINYINT"
	case TDatumType_GEOMETRY:
		return "GEOMETRY"
	case TDatumType_GEOGRAPHY:
		return "GEOGRAPHY"
	}
	return "<UNSET>"
}

func TDatumTypeFromString(s string) (TDatumType, error) {
	switch s {
	case "SMALLINT":
		return TDatumType_SMALLINT, nil
	case "INT":
		return TDatumType_INT, nil
	case "BIGINT":
		return TDatumType_BIGINT, nil
	case "FLOAT":
		return TDatumType_FLOAT, nil
	case "DECIMAL":
		return TDatumType_DECIMAL, nil
	case "DOUBLE":
		return TDatumType_DOUBLE, nil
	case "STR":
		return TDatumType_STR, nil
	case "TIME":
		return TDatumType_TIME, nil
	case "TIMESTAMP":
		return TDatumType_TIMESTAMP, nil
	case "DATE":
		return TDatumType_DATE, nil
	case "BOOL":
		return TDatumType_BOOL, nil
	case "INTERVAL_DAY_TIME":
		return TDatumType_INTERVAL_DAY_TIME, nil
	case "INTERVAL_YEAR_MONTH":
		return TDatumType_INTERVAL_YEAR_MONTH, nil
	case "POINT":
		return TDatumType_POINT, nil
	case "LINESTRING":
		return TDatumType_LINESTRING, nil
	case "POLYGON":
		return TDatumType_POLYGON, nil
	case "MULTIPOLYGON":
		return TDatumType_MULTIPOLYGON, nil
	case "TINYINT":
		return TDatumType_TINYINT, nil
	case "GEOMETRY":
		return TDatumType_GEOMETRY, nil
	case "GEOGRAPHY":
		return TDatumType_GEOGRAPHY, nil
	}
	return TDatumType(0), fmt.Errorf("not a valid TDatumType string")
}

func TDatumTypePtr(v TDatumType) *TDatumType { return &v }

func (p TDatumType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TDatumType) UnmarshalText(text []byte) error {
	q, err := TDatumTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type TEncodingType int64

const (
	TEncodingType_NONE         TEncodingType = 0
	TEncodingType_FIXED        TEncodingType = 1
	TEncodingType_RL           TEncodingType = 2
	TEncodingType_DIFF         TEncodingType = 3
	TEncodingType_DICT         TEncodingType = 4
	TEncodingType_SPARSE       TEncodingType = 5
	TEncodingType_GEOINT       TEncodingType = 6
	TEncodingType_DATE_IN_DAYS TEncodingType = 7
)

func (p TEncodingType) String() string {
	switch p {
	case TEncodingType_NONE:
		return "NONE"
	case TEncodingType_FIXED:
		return "FIXED"
	case TEncodingType_RL:
		return "RL"
	case TEncodingType_DIFF:
		return "DIFF"
	case TEncodingType_DICT:
		return "DICT"
	case TEncodingType_SPARSE:
		return "SPARSE"
	case TEncodingType_GEOINT:
		return "GEOINT"
	case TEncodingType_DATE_IN_DAYS:
		return "DATE_IN_DAYS"
	}
	return "<UNSET>"
}

func TEncodingTypeFromString(s string) (TEncodingType, error) {
	switch s {
	case "NONE":
		return TEncodingType_NONE, nil
	case "FIXED":
		return TEncodingType_FIXED, nil
	case "RL":
		return TEncodingType_RL, nil
	case "DIFF":
		return TEncodingType_DIFF, nil
	case "DICT":
		return TEncodingType_DICT, nil
	case "SPARSE":
		return TEncodingType_SPARSE, nil
	case "GEOINT":
		return TEncodingType_GEOINT, nil
	case "DATE_IN_DAYS":
		return TEncodingType_DATE_IN_DAYS, nil
	}
	return TEncodingType(0), fmt.Errorf("not a valid TEncodingType string")
}

func TEncodingTypePtr(v TEncodingType) *TEncodingType { return &v }

func (p TEncodingType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TEncodingType) UnmarshalText(text []byte) error {
	q, err := TEncodingTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - Type
//  - Encoding
//  - Nullable
//  - IsArray
//  - Precision
//  - Scale
//  - CompParam
//  - Size
type TTypeInfo struct {
	Type      TDatumType    `thrift:"type,1" json:"type"`
	Nullable  bool          `thrift:"nullable,2" json:"nullable"`
	IsArray   bool          `thrift:"is_array,3" json:"is_array"`
	Encoding  TEncodingType `thrift:"encoding,4" json:"encoding"`
	Precision int32         `thrift:"precision,5" json:"precision"`
	Scale     int32         `thrift:"scale,6" json:"scale"`
	CompParam int32         `thrift:"comp_param,7" json:"comp_param"`
	Size      int32         `thrift:"size,8" json:"size,omitempty"`
}

func NewTTypeInfo() *TTypeInfo {
	return &TTypeInfo{
		Size: -1,
	}
}

func (p *TTypeInfo) GetType() TDatumType {
	return p.Type
}

func (p *TTypeInfo) GetEncoding() TEncodingType {
	return p.Encoding
}

func (p *TTypeInfo) GetNullable() bool {
	return p.Nullable
}

func (p *TTypeInfo) GetIsArray() bool {
	return p.IsArray
}

func (p *TTypeInfo) GetPrecision() int32 {
	return p.Precision
}

func (p *TTypeInfo) GetScale() int32 {
	return p.Scale
}

func (p *TTypeInfo) GetCompParam() int32 {
	return p.CompParam
}

var TTypeInfo_Size_DEFAULT int32 = -1

func (p *TTypeInfo) GetSize() int32 {
	return p.Size
}
func (p *TTypeInfo) IsSetSize() bool {
	return p.Size != TTypeInfo_Size_DEFAULT
}

func (p *TTypeInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TTypeInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := TDatumType(v)
		p.Type = temp
	}
	return nil
}

func (p *TTypeInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		temp := TEncodingType(v)
		p.Encoding = temp
	}
	return nil
}

func (p *TTypeInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Nullable = v
	}
	return nil
}

func (p *TTypeInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.IsArray = v
	}
	return nil
}

func (p *TTypeInfo) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Precision = v
	}
	return nil
}

func (p *TTypeInfo) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Scale = v
	}
	return nil
}

func (p *TTypeInfo) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.CompParam = v
	}
	return nil
}

func (p *TTypeInfo) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.Size = v
	}
	return nil
}

func (p *TTypeInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TTypeInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TTypeInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("type", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("nullable", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nullable: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Nullable)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nullable (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nullable: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("is_array", thrift.BOOL, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:is_array: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.IsArray)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.is_array (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:is_array: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("encoding", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:encoding: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Encoding)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.encoding (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:encoding: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("precision", thrift.I32, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:precision: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Precision)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.precision (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:precision: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("scale", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:scale: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Scale)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.scale (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:scale: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("comp_param", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:comp_param: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.CompParam)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.comp_param (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:comp_param: ", p), err)
	}
	return err
}

func (p *TTypeInfo) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetSize() {
		if err := oprot.WriteFieldBegin("size", thrift.I32, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:size: ", p), err)
		}
		if err := oprot.WriteI32(int32(p.Size)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.size (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:size: ", p), err)
		}
	}
	return err
}

func (p *TTypeInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TTypeInfo(%+v)", *p)
}
