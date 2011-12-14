// Code generated by protoc-gen-go from "instruction.proto"
// DO NOT EDIT!

package tritium

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error

type Instruction_InstructionType int32

const (
	Instruction_BLOCK         Instruction_InstructionType = 0
	Instruction_FUNCTION_CALL Instruction_InstructionType = 1
	Instruction_IMPORT        Instruction_InstructionType = 2
	Instruction_TEXT          Instruction_InstructionType = 3
	Instruction_LOCAL_VAR     Instruction_InstructionType = 4
)

var Instruction_InstructionType_name = map[int32]string{
	0: "BLOCK",
	1: "FUNCTION_CALL",
	2: "IMPORT",
	3: "TEXT",
	4: "LOCAL_VAR",
}
var Instruction_InstructionType_value = map[string]int32{
	"BLOCK":         0,
	"FUNCTION_CALL": 1,
	"IMPORT":        2,
	"TEXT":          3,
	"LOCAL_VAR":     4,
}

func NewInstruction_InstructionType(x Instruction_InstructionType) *Instruction_InstructionType {
	e := Instruction_InstructionType(x)
	return &e
}
func (x Instruction_InstructionType) String() string {
	return proto.EnumName(Instruction_InstructionType_name, int32(x))
}

type Instruction struct {
	Type             *Instruction_InstructionType `protobuf:"varint,1,req,name=type,enum=tritium.Instruction_InstructionType" json:"type"`
	Value            *string                      `protobuf:"bytes,2,opt,name=value" json:"value"`
	ObjectId         *int32                       `protobuf:"varint,3,opt,name=object_id" json:"object_id"`
	Children         []*Instruction               `protobuf:"bytes,4,rep,name=children" json:"children"`
	Arguments        []*Instruction               `protobuf:"bytes,5,rep,name=arguments" json:"arguments"`
	FunctionId       *int32                       `protobuf:"varint,6,opt,name=function_id" json:"function_id"`
	XXX_unrecognized []byte
}

func (this *Instruction) Reset()         { *this = Instruction{} }
func (this *Instruction) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("tritium.Instruction_InstructionType", Instruction_InstructionType_name, Instruction_InstructionType_value)
}
