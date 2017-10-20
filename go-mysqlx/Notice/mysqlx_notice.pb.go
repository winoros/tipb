// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/pingcap/tipb/go-mysqlx/Notice/mysqlx_notice.proto

/*
	Package Mysqlx_Notice is a generated protocol buffer package.

	Notices

	A notice

	* is sent from the server to the client
	* may be global or relate to the current message sequence

	It is generated from these files:
		github.com/pingcap/tipb/go-mysqlx/Notice/mysqlx_notice.proto

	It has these top-level messages:
		Frame
		Warning
		SessionVariableChanged
		SessionStateChanged
*/
package Mysqlx_Notice

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	Mysqlx_Datatypes "github.com/pingcap/tipb/go-mysqlx/Datatypes"

	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Frame_Scope int32

const (
	Frame_GLOBAL Frame_Scope = 1
	Frame_LOCAL  Frame_Scope = 2
)

var Frame_Scope_name = map[int32]string{
	1: "GLOBAL",
	2: "LOCAL",
}
var Frame_Scope_value = map[string]int32{
	"GLOBAL": 1,
	"LOCAL":  2,
}

func (x Frame_Scope) Enum() *Frame_Scope {
	p := new(Frame_Scope)
	*p = x
	return p
}
func (x Frame_Scope) String() string {
	return proto.EnumName(Frame_Scope_name, int32(x))
}
func (x *Frame_Scope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Frame_Scope_value, data, "Frame_Scope")
	if err != nil {
		return err
	}
	*x = Frame_Scope(value)
	return nil
}
func (Frame_Scope) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlxNotice, []int{0, 0} }

type Warning_Level int32

const (
	Warning_NOTE    Warning_Level = 1
	Warning_WARNING Warning_Level = 2
	Warning_ERROR   Warning_Level = 3
)

var Warning_Level_name = map[int32]string{
	1: "NOTE",
	2: "WARNING",
	3: "ERROR",
}
var Warning_Level_value = map[string]int32{
	"NOTE":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x Warning_Level) Enum() *Warning_Level {
	p := new(Warning_Level)
	*p = x
	return p
}
func (x Warning_Level) String() string {
	return proto.EnumName(Warning_Level_name, int32(x))
}
func (x *Warning_Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Warning_Level_value, data, "Warning_Level")
	if err != nil {
		return err
	}
	*x = Warning_Level(value)
	return nil
}
func (Warning_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlxNotice, []int{1, 0} }

type SessionStateChanged_Parameter int32

const (
	SessionStateChanged_CURRENT_SCHEMA      SessionStateChanged_Parameter = 1
	SessionStateChanged_ACCOUNT_EXPIRED     SessionStateChanged_Parameter = 2
	SessionStateChanged_GENERATED_INSERT_ID SessionStateChanged_Parameter = 3
	SessionStateChanged_ROWS_AFFECTED       SessionStateChanged_Parameter = 4
	SessionStateChanged_ROWS_FOUND          SessionStateChanged_Parameter = 5
	SessionStateChanged_ROWS_MATCHED        SessionStateChanged_Parameter = 6
	SessionStateChanged_TRX_COMMITTED       SessionStateChanged_Parameter = 7
	SessionStateChanged_TRX_ROLLEDBACK      SessionStateChanged_Parameter = 9
	SessionStateChanged_PRODUCED_MESSAGE    SessionStateChanged_Parameter = 10
	SessionStateChanged_CLIENT_ID_ASSIGNED  SessionStateChanged_Parameter = 11
)

var SessionStateChanged_Parameter_name = map[int32]string{
	1:  "CURRENT_SCHEMA",
	2:  "ACCOUNT_EXPIRED",
	3:  "GENERATED_INSERT_ID",
	4:  "ROWS_AFFECTED",
	5:  "ROWS_FOUND",
	6:  "ROWS_MATCHED",
	7:  "TRX_COMMITTED",
	9:  "TRX_ROLLEDBACK",
	10: "PRODUCED_MESSAGE",
	11: "CLIENT_ID_ASSIGNED",
}
var SessionStateChanged_Parameter_value = map[string]int32{
	"CURRENT_SCHEMA":      1,
	"ACCOUNT_EXPIRED":     2,
	"GENERATED_INSERT_ID": 3,
	"ROWS_AFFECTED":       4,
	"ROWS_FOUND":          5,
	"ROWS_MATCHED":        6,
	"TRX_COMMITTED":       7,
	"TRX_ROLLEDBACK":      9,
	"PRODUCED_MESSAGE":    10,
	"CLIENT_ID_ASSIGNED":  11,
}

func (x SessionStateChanged_Parameter) Enum() *SessionStateChanged_Parameter {
	p := new(SessionStateChanged_Parameter)
	*p = x
	return p
}
func (x SessionStateChanged_Parameter) String() string {
	return proto.EnumName(SessionStateChanged_Parameter_name, int32(x))
}
func (x *SessionStateChanged_Parameter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionStateChanged_Parameter_value, data, "SessionStateChanged_Parameter")
	if err != nil {
		return err
	}
	*x = SessionStateChanged_Parameter(value)
	return nil
}
func (SessionStateChanged_Parameter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorMysqlxNotice, []int{3, 0}
}

// Common Frame for all Notices
//
// ===================================================== =====
// .type                                                 value
// ===================================================== =====
// :protobuf:msg:`Mysqlx.Notice::Warning`                1
// :protobuf:msg:`Mysqlx.Notice::SessionVariableChanged` 2
// :protobuf:msg:`Mysqlx.Notice::SessionStateChanged`    3
// ===================================================== =====
//
// :param type: the type of the payload
// :param payload: the payload of the notification
// :param scope: global or local notification
//
type Frame struct {
	Type             *uint32      `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Scope            *Frame_Scope `protobuf:"varint,2,opt,name=scope,enum=Mysqlx.Notice.Frame_Scope,def=1" json:"scope,omitempty"`
	Payload          []byte       `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxNotice, []int{0} }

const Default_Frame_Scope Frame_Scope = Frame_GLOBAL

func (m *Frame) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Frame) GetScope() Frame_Scope {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return Default_Frame_Scope
}

func (m *Frame) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-side warnings and notes
//
// ``.scope`` == ``local``
//   ``.level``, ``.code`` and ``.msg`` map the content of
//
//   .. code-block:: sql
//
//     SHOW WARNINGS
//
// ``.scope`` == ``global``
//   (undefined) will be used for global, unstructured messages like:
//
//   * server is shutting down
//   * a node disconnected from group
//   * schema or table dropped
//
// ========================================== =======================
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =======================
// ``.type``                                  1
// ``.scope``                                 ``local`` or ``global``
// ========================================== =======================
//
// :param level: warning level: Note or Warning
// :param code: warning code
// :param msg: warning message
type Warning struct {
	Level            *Warning_Level `protobuf:"varint,1,opt,name=level,enum=Mysqlx.Notice.Warning_Level,def=2" json:"level,omitempty"`
	Code             *uint32        `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	Msg              *string        `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Warning) Reset()                    { *m = Warning{} }
func (m *Warning) String() string            { return proto.CompactTextString(m) }
func (*Warning) ProtoMessage()               {}
func (*Warning) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxNotice, []int{1} }

const Default_Warning_Level Warning_Level = Warning_WARNING

func (m *Warning) GetLevel() Warning_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_Warning_Level
}

func (m *Warning) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Warning) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// Notify clients about changes to the current session variables
//
// Every change to a variable that is accessable through:
//
// .. code-block:: sql
//
//   SHOW SESSION VARIABLES
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  2
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param namespace: namespace that param belongs to
// :param param: name of the variable
// :param value: the changed value of param
type SessionVariableChanged struct {
	Param            *string                  `protobuf:"bytes,1,req,name=param" json:"param,omitempty"`
	Value            *Mysqlx_Datatypes.Scalar `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *SessionVariableChanged) Reset()         { *m = SessionVariableChanged{} }
func (m *SessionVariableChanged) String() string { return proto.CompactTextString(m) }
func (*SessionVariableChanged) ProtoMessage()    {}
func (*SessionVariableChanged) Descriptor() ([]byte, []int) {
	return fileDescriptorMysqlxNotice, []int{2}
}

func (m *SessionVariableChanged) GetParam() string {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return ""
}

func (m *SessionVariableChanged) GetValue() *Mysqlx_Datatypes.Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

// Notify clients about changes to the internal session state
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  3
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param param: parameter key
// :param value: updated value
type SessionStateChanged struct {
	Param            *SessionStateChanged_Parameter `protobuf:"varint,1,req,name=param,enum=Mysqlx.Notice.SessionStateChanged_Parameter" json:"param,omitempty"`
	Value            *Mysqlx_Datatypes.Scalar       `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *SessionStateChanged) Reset()                    { *m = SessionStateChanged{} }
func (m *SessionStateChanged) String() string            { return proto.CompactTextString(m) }
func (*SessionStateChanged) ProtoMessage()               {}
func (*SessionStateChanged) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxNotice, []int{3} }

func (m *SessionStateChanged) GetParam() SessionStateChanged_Parameter {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return SessionStateChanged_CURRENT_SCHEMA
}

func (m *SessionStateChanged) GetValue() *Mysqlx_Datatypes.Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Frame)(nil), "Mysqlx.Notice.Frame")
	proto.RegisterType((*Warning)(nil), "Mysqlx.Notice.Warning")
	proto.RegisterType((*SessionVariableChanged)(nil), "Mysqlx.Notice.SessionVariableChanged")
	proto.RegisterType((*SessionStateChanged)(nil), "Mysqlx.Notice.SessionStateChanged")
	proto.RegisterEnum("Mysqlx.Notice.Frame_Scope", Frame_Scope_name, Frame_Scope_value)
	proto.RegisterEnum("Mysqlx.Notice.Warning_Level", Warning_Level_name, Warning_Level_value)
	proto.RegisterEnum("Mysqlx.Notice.SessionStateChanged_Parameter", SessionStateChanged_Parameter_name, SessionStateChanged_Parameter_value)
}
func (m *Frame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Frame) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(*m.Type))
	}
	if m.Scope != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(*m.Scope))
	}
	if m.Payload != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Warning) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Warning) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Level != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(*m.Level))
	}
	if m.Code == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(*m.Code))
	}
	if m.Msg == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(len(*m.Msg)))
		i += copy(dAtA[i:], *m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SessionVariableChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionVariableChanged) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Param == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(len(*m.Param)))
		i += copy(dAtA[i:], *m.Param)
	}
	if m.Value != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(m.Value.Size()))
		n1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SessionStateChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionStateChanged) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Param == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(*m.Param))
	}
	if m.Value != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMysqlxNotice(dAtA, i, uint64(m.Value.Size()))
		n2, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64MysqlxNotice(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32MysqlxNotice(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMysqlxNotice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Frame) Size() (n int) {
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovMysqlxNotice(uint64(*m.Type))
	}
	if m.Scope != nil {
		n += 1 + sovMysqlxNotice(uint64(*m.Scope))
	}
	if m.Payload != nil {
		l = len(m.Payload)
		n += 1 + l + sovMysqlxNotice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Warning) Size() (n int) {
	var l int
	_ = l
	if m.Level != nil {
		n += 1 + sovMysqlxNotice(uint64(*m.Level))
	}
	if m.Code != nil {
		n += 1 + sovMysqlxNotice(uint64(*m.Code))
	}
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovMysqlxNotice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SessionVariableChanged) Size() (n int) {
	var l int
	_ = l
	if m.Param != nil {
		l = len(*m.Param)
		n += 1 + l + sovMysqlxNotice(uint64(l))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMysqlxNotice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SessionStateChanged) Size() (n int) {
	var l int
	_ = l
	if m.Param != nil {
		n += 1 + sovMysqlxNotice(uint64(*m.Param))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMysqlxNotice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMysqlxNotice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlxNotice(x uint64) (n int) {
	return sovMysqlxNotice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Frame) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Frame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Frame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var v Frame_Scope
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Frame_Scope(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Scope = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxNotice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Warning) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Warning: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Warning: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var v Warning_Level
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Warning_Level(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Level = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Code = &v
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxNotice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SessionVariableChanged) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionVariableChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionVariableChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Param = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &Mysqlx_Datatypes.Scalar{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxNotice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SessionStateChanged) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionStateChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionStateChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param", wireType)
			}
			var v SessionStateChanged_Parameter
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (SessionStateChanged_Parameter(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Param = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &Mysqlx_Datatypes.Scalar{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxNotice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxNotice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMysqlxNotice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlxNotice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlxNotice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMysqlxNotice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlxNotice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMysqlxNotice(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMysqlxNotice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlxNotice   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/pingcap/tipb/go-mysqlx/Notice/mysqlx_notice.proto", fileDescriptorMysqlxNotice)
}

var fileDescriptorMysqlxNotice = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0xba, 0xac, 0xf4, 0x6d, 0x2b, 0xc6, 0x9b, 0x46, 0x34, 0x4d, 0x55, 0xe9, 0x85,
	0x22, 0xa1, 0x4c, 0xda, 0x09, 0x8d, 0x53, 0x6a, 0xbb, 0x5d, 0x44, 0x9a, 0x4c, 0x4e, 0xca, 0x76,
	0x22, 0xf2, 0xda, 0x30, 0x8a, 0xd2, 0xa6, 0x24, 0xd9, 0xb4, 0x7d, 0x0a, 0x8e, 0xf0, 0x91, 0x38,
	0x72, 0xe3, 0x8a, 0xb6, 0x2f, 0x82, 0xec, 0xb4, 0x13, 0x4c, 0x9c, 0xb8, 0xbd, 0xbf, 0xf3, 0xff,
	0xbd, 0xf7, 0xf7, 0x8b, 0x0c, 0xcf, 0x87, 0x37, 0xc5, 0xe7, 0xf4, 0xfa, 0xc0, 0xcf, 0xca, 0xe9,
	0x38, 0x39, 0x98, 0x69, 0x15, 0xcf, 0xb5, 0xb2, 0x17, 0x79, 0x56, 0x66, 0x64, 0xab, 0xb2, 0xd8,
	0x95, 0x65, 0xef, 0xc5, 0x92, 0x60, 0xb2, 0x94, 0xe5, 0xcd, 0x22, 0x29, 0x56, 0xd0, 0x64, 0x75,
	0x50, 0x71, 0x9d, 0x2f, 0x08, 0xcc, 0x7e, 0x2e, 0x67, 0x09, 0x21, 0xb0, 0xa6, 0x3e, 0x58, 0xa8,
	0x6d, 0x74, 0xb7, 0x84, 0xae, 0xc9, 0x6b, 0x30, 0x8b, 0x71, 0xb6, 0x48, 0x2c, 0xa3, 0x8d, 0xba,
	0xcd, 0xc3, 0x3d, 0xfb, 0xaf, 0x29, 0xb6, 0x06, 0xed, 0x50, 0x39, 0x8e, 0xd6, 0x07, 0x5e, 0xd0,
	0x73, 0x3c, 0x51, 0x01, 0xc4, 0x82, 0xfa, 0x42, 0xde, 0xa4, 0x99, 0x9c, 0x58, 0xb5, 0x36, 0xea,
	0x6e, 0x8a, 0x95, 0xec, 0xb4, 0xc0, 0xd4, 0x04, 0x01, 0x58, 0x32, 0x18, 0x91, 0x06, 0x98, 0x5e,
	0x40, 0x1d, 0x0f, 0x1b, 0x9d, 0xaf, 0x08, 0xea, 0xa7, 0x32, 0x9f, 0x4f, 0xe7, 0x17, 0xe4, 0x0d,
	0x98, 0x69, 0x72, 0x95, 0xa4, 0x16, 0xd2, 0xf3, 0xf7, 0x1f, 0xcc, 0x5f, 0xda, 0x6c, 0x4f, 0x79,
	0x8e, 0xea, 0xa7, 0x8e, 0xf0, 0x5d, 0x7f, 0x20, 0x2a, 0x46, 0x5d, 0x68, 0x9c, 0x4d, 0x54, 0x76,
	0x7d, 0x21, 0x55, 0x13, 0x0c, 0xb5, 0x59, 0x71, 0x61, 0xd5, 0xda, 0x46, 0xb7, 0x21, 0x54, 0xd9,
	0x79, 0x09, 0xa6, 0xc6, 0xc9, 0x63, 0x58, 0xf3, 0x83, 0x88, 0x63, 0x44, 0x36, 0x60, 0xd5, 0x0a,
	0x1b, 0x2a, 0x19, 0x17, 0x22, 0x10, 0xb8, 0xd6, 0x79, 0x0f, 0xbb, 0x61, 0x52, 0x14, 0xd3, 0x6c,
	0xfe, 0x4e, 0xe6, 0x53, 0x79, 0x9e, 0x26, 0xf4, 0xa3, 0x9c, 0x5f, 0x24, 0x13, 0xb2, 0x03, 0xe6,
	0x42, 0xe6, 0x72, 0xa6, 0x97, 0xd7, 0x10, 0x95, 0x20, 0x36, 0x98, 0x57, 0x32, 0xbd, 0xac, 0xb6,
	0xb7, 0x71, 0x68, 0xad, 0xd2, 0xdf, 0xff, 0x14, 0x3b, 0x1c, 0xcb, 0x54, 0xe6, 0xa2, 0xb2, 0x75,
	0xee, 0x0c, 0xd8, 0x5e, 0x0e, 0x08, 0x4b, 0x59, 0xde, 0x77, 0xef, 0xfd, 0xd9, 0xbd, 0x79, 0xf8,
	0xea, 0xc1, 0x16, 0xfe, 0x81, 0xd8, 0x27, 0xca, 0x9f, 0x94, 0x49, 0xfe, 0xbf, 0x59, 0x7e, 0x22,
	0x68, 0xdc, 0x37, 0x21, 0x04, 0x9a, 0x74, 0x24, 0x04, 0xf7, 0xa3, 0x38, 0xa4, 0xc7, 0x7c, 0xe8,
	0x60, 0x44, 0xb6, 0xe1, 0x89, 0x43, 0x69, 0x30, 0xf2, 0xa3, 0x98, 0x9f, 0x9d, 0xb8, 0x82, 0x33,
	0x6c, 0x90, 0x67, 0xb0, 0x3d, 0xe0, 0x3e, 0x17, 0x4e, 0xc4, 0x59, 0xec, 0xfa, 0x21, 0x17, 0x51,
	0xec, 0x32, 0x5c, 0x23, 0x4f, 0x61, 0x4b, 0x04, 0xa7, 0x61, 0xec, 0xf4, 0xfb, 0x9c, 0x46, 0x9c,
	0xe1, 0x35, 0xd2, 0x04, 0xd0, 0x47, 0xfd, 0x60, 0xe4, 0x33, 0x6c, 0x12, 0x0c, 0x9b, 0x5a, 0x0f,
	0x9d, 0x88, 0x1e, 0x73, 0x86, 0xd7, 0x15, 0x14, 0x89, 0xb3, 0x98, 0x06, 0xc3, 0xa1, 0x1b, 0x29,
	0xa8, 0xae, 0x92, 0xa8, 0x23, 0x11, 0x78, 0x1e, 0x67, 0x3d, 0x87, 0xbe, 0xc5, 0x0d, 0xb2, 0x03,
	0xf8, 0x44, 0x04, 0x6c, 0x44, 0x39, 0x8b, 0x87, 0x3c, 0x0c, 0x9d, 0x01, 0xc7, 0x40, 0x76, 0x81,
	0x50, 0xcf, 0x55, 0x91, 0x5d, 0x16, 0x3b, 0x61, 0xe8, 0x0e, 0x7c, 0xce, 0xf0, 0x46, 0xcf, 0xfe,
	0x7e, 0xdb, 0x42, 0x3f, 0x6e, 0x5b, 0xe8, 0xd7, 0x6d, 0x0b, 0x7d, 0xbb, 0x6b, 0x3d, 0x82, 0xfd,
	0x71, 0x36, 0xb3, 0xf5, 0xfb, 0xb0, 0xc7, 0x9f, 0xaa, 0xe2, 0xba, 0x7a, 0x1e, 0xe7, 0x97, 0x1f,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x90, 0x2d, 0x0e, 0x7d, 0x03, 0x00, 0x00,
}
