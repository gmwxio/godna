// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dna/config/config.v1.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/wxio/godna/pb/extensions/store"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
type Config_PluginOutDir_OutType int32

const (
	Config_PluginOutDir_NOTHING_SPECIAL Config_PluginOutDir_OutType = 0
	Config_PluginOutDir_GO_MODS         Config_PluginOutDir_OutType = 1
)

var Config_PluginOutDir_OutType_name = map[int32]string{
	0: "NOTHING_SPECIAL",
	1: "GO_MODS",
}

var Config_PluginOutDir_OutType_value = map[string]int32{
	"NOTHING_SPECIAL": 0,
	"GO_MODS":         1,
}

func (x Config_PluginOutDir_OutType) String() string {
	return proto.EnumName(Config_PluginOutDir_OutType_name, int32(x))
}

func (Config_PluginOutDir_OutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 0, 0}
}

//
type Config_Pass_Command int32

const (
	Config_Pass_NONE            Config_Pass_Command = 0
	Config_Pass_PROTOC          Config_Pass_Command = 1
	Config_Pass_MOD_INIT        Config_Pass_Command = 2
	Config_Pass_MOD_REQUIRE     Config_Pass_Command = 3
	Config_Pass_MOD_REPLACE     Config_Pass_Command = 4
	Config_Pass_MOD_TIDY        Config_Pass_Command = 5
	Config_Pass_GIT_ADD_NEXTTAG Config_Pass_Command = 6
	Config_Pass_GIT_TAG         Config_Pass_Command = 7
)

var Config_Pass_Command_name = map[int32]string{
	0: "NONE",
	1: "PROTOC",
	2: "MOD_INIT",
	3: "MOD_REQUIRE",
	4: "MOD_REPLACE",
	5: "MOD_TIDY",
	6: "GIT_ADD_NEXTTAG",
	7: "GIT_TAG",
}

var Config_Pass_Command_value = map[string]int32{
	"NONE":            0,
	"PROTOC":          1,
	"MOD_INIT":        2,
	"MOD_REQUIRE":     3,
	"MOD_REPLACE":     4,
	"MOD_TIDY":        5,
	"GIT_ADD_NEXTTAG": 6,
	"GIT_TAG":         7,
}

func (x Config_Pass_Command) String() string {
	return proto.EnumName(Config_Pass_Command_name, int32(x))
}

func (Config_Pass_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 1, 0}
}

type Config_Generator_Paths int32

const (
	Config_Generator_Import          Config_Generator_Paths = 0
	Config_Generator_Source_Relative Config_Generator_Paths = 1
)

var Config_Generator_Paths_name = map[int32]string{
	0: "Import",
	1: "Source_Relative",
}

var Config_Generator_Paths_value = map[string]int32{
	"Import":          0,
	"Source_Relative": 1,
}

func (x Config_Generator_Paths) String() string {
	return proto.EnumName(Config_Generator_Paths_name, int32(x))
}

func (Config_Generator_Paths) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 0}
}

// map<string,string> importMap = 5;
//
type Config_Generator_Go_Plugins int32

const (
	Config_Generator_Go_None Config_Generator_Go_Plugins = 0
	Config_Generator_Go_Grpc Config_Generator_Go_Plugins = 1
)

var Config_Generator_Go_Plugins_name = map[int32]string{
	0: "None",
	1: "Grpc",
}

var Config_Generator_Go_Plugins_value = map[string]int32{
	"None": 0,
	"Grpc": 1,
}

func (x Config_Generator_Go_Plugins) String() string {
	return proto.EnumName(Config_Generator_Go_Plugins_name, int32(x))
}

func (Config_Generator_Go_Plugins) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 1, 0}
}

type Config struct {
	HostOwner string `protobuf:"bytes,1,opt,name=host_owner,json=hostOwner,proto3" json:"host_owner,omitempty"`
	//TODO protobuf validators eg [(validator.maditory) = true]
	RepoName string   `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	Includes []string `protobuf:"bytes,3,rep,name=includes,proto3" json:"includes,omitempty"`
	RelImps  []string `protobuf:"bytes,4,rep,name=rel_imps,json=relImps,proto3" json:"rel_imps,omitempty"`
	Require  []string `protobuf:"bytes,5,rep,name=require,proto3" json:"require,omitempty"`
	SrcDir   string   `protobuf:"bytes,6,opt,name=src_dir,json=srcDir,proto3" json:"src_dir,omitempty"`
	// string   output_root              = 7
	// [(wxio.extensions.store.help) = "output root directory eg ."];
	Pass                 []*Config_Pass         `protobuf:"bytes,8,rep,name=pass,proto3" json:"pass,omitempty"`
	PluginPath           []string               `protobuf:"bytes,9,rep,name=plugin_path,json=pluginPath,proto3" json:"plugin_path,omitempty"`
	PluginOutputDir      []*Config_PluginOutDir `protobuf:"bytes,10,rep,name=plugin_output_dir,json=pluginOutputDir,proto3" json:"plugin_output_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetHostOwner() string {
	if m != nil {
		return m.HostOwner
	}
	return ""
}

func (m *Config) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *Config) GetIncludes() []string {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Config) GetRelImps() []string {
	if m != nil {
		return m.RelImps
	}
	return nil
}

func (m *Config) GetRequire() []string {
	if m != nil {
		return m.Require
	}
	return nil
}

func (m *Config) GetSrcDir() string {
	if m != nil {
		return m.SrcDir
	}
	return ""
}

func (m *Config) GetPass() []*Config_Pass {
	if m != nil {
		return m.Pass
	}
	return nil
}

func (m *Config) GetPluginPath() []string {
	if m != nil {
		return m.PluginPath
	}
	return nil
}

func (m *Config) GetPluginOutputDir() []*Config_PluginOutDir {
	if m != nil {
		return m.PluginOutputDir
	}
	return nil
}

//
type Config_PluginOutDir struct {
	Path                 string                      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	OutType              Config_PluginOutDir_OutType `protobuf:"varint,2,opt,name=out_type,json=outType,proto3,enum=godna.Config_PluginOutDir_OutType" json:"out_type,omitempty"`
	Generator            []*Config_Generator         `protobuf:"bytes,3,rep,name=generator,proto3" json:"generator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Config_PluginOutDir) Reset()         { *m = Config_PluginOutDir{} }
func (m *Config_PluginOutDir) String() string { return proto.CompactTextString(m) }
func (*Config_PluginOutDir) ProtoMessage()    {}
func (*Config_PluginOutDir) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 0}
}

func (m *Config_PluginOutDir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_PluginOutDir.Unmarshal(m, b)
}
func (m *Config_PluginOutDir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_PluginOutDir.Marshal(b, m, deterministic)
}
func (m *Config_PluginOutDir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_PluginOutDir.Merge(m, src)
}
func (m *Config_PluginOutDir) XXX_Size() int {
	return xxx_messageInfo_Config_PluginOutDir.Size(m)
}
func (m *Config_PluginOutDir) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_PluginOutDir.DiscardUnknown(m)
}

var xxx_messageInfo_Config_PluginOutDir proto.InternalMessageInfo

func (m *Config_PluginOutDir) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Config_PluginOutDir) GetOutType() Config_PluginOutDir_OutType {
	if m != nil {
		return m.OutType
	}
	return Config_PluginOutDir_NOTHING_SPECIAL
}

func (m *Config_PluginOutDir) GetGenerator() []*Config_Generator {
	if m != nil {
		return m.Generator
	}
	return nil
}

type Config_Pass struct {
	Cmd                  []Config_Pass_Command `protobuf:"varint,1,rep,packed,name=cmd,proto3,enum=godna.Config_Pass_Command" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Config_Pass) Reset()         { *m = Config_Pass{} }
func (m *Config_Pass) String() string { return proto.CompactTextString(m) }
func (*Config_Pass) ProtoMessage()    {}
func (*Config_Pass) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 1}
}

func (m *Config_Pass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Pass.Unmarshal(m, b)
}
func (m *Config_Pass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Pass.Marshal(b, m, deterministic)
}
func (m *Config_Pass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Pass.Merge(m, src)
}
func (m *Config_Pass) XXX_Size() int {
	return xxx_messageInfo_Config_Pass.Size(m)
}
func (m *Config_Pass) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Pass.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Pass proto.InternalMessageInfo

func (m *Config_Pass) GetCmd() []Config_Pass_Command {
	if m != nil {
		return m.Cmd
	}
	return nil
}

type Config_Generator struct {
	Plugin               *Config_Generator_Plugin `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Config_Generator) Reset()         { *m = Config_Generator{} }
func (m *Config_Generator) String() string { return proto.CompactTextString(m) }
func (*Config_Generator) ProtoMessage()    {}
func (*Config_Generator) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2}
}

func (m *Config_Generator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator.Unmarshal(m, b)
}
func (m *Config_Generator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator.Marshal(b, m, deterministic)
}
func (m *Config_Generator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator.Merge(m, src)
}
func (m *Config_Generator) XXX_Size() int {
	return xxx_messageInfo_Config_Generator.Size(m)
}
func (m *Config_Generator) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator proto.InternalMessageInfo

func (m *Config_Generator) GetPlugin() *Config_Generator_Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

//
type Config_Generator_Plugin struct {
	// Types that are valid to be assigned to Cmd:
	//	*Config_Generator_Plugin_Go
	//	*Config_Generator_Plugin_Micro
	//	*Config_Generator_Plugin_GrpcGateway
	//	*Config_Generator_Plugin_Swagger
	Cmd                  isConfig_Generator_Plugin_Cmd `protobuf_oneof:"cmd"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Config_Generator_Plugin) Reset()         { *m = Config_Generator_Plugin{} }
func (m *Config_Generator_Plugin) String() string { return proto.CompactTextString(m) }
func (*Config_Generator_Plugin) ProtoMessage()    {}
func (*Config_Generator_Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 0}
}

func (m *Config_Generator_Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator_Plugin.Unmarshal(m, b)
}
func (m *Config_Generator_Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator_Plugin.Marshal(b, m, deterministic)
}
func (m *Config_Generator_Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator_Plugin.Merge(m, src)
}
func (m *Config_Generator_Plugin) XXX_Size() int {
	return xxx_messageInfo_Config_Generator_Plugin.Size(m)
}
func (m *Config_Generator_Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator_Plugin proto.InternalMessageInfo

type isConfig_Generator_Plugin_Cmd interface {
	isConfig_Generator_Plugin_Cmd()
}

type Config_Generator_Plugin_Go struct {
	Go *Config_Generator_Go `protobuf:"bytes,1,opt,name=go,proto3,oneof"`
}

type Config_Generator_Plugin_Micro struct {
	Micro *Config_Generator_Micro `protobuf:"bytes,2,opt,name=micro,proto3,oneof"`
}

type Config_Generator_Plugin_GrpcGateway struct {
	GrpcGateway *Config_Generator_Grpc_Gateway `protobuf:"bytes,3,opt,name=grpc_gateway,json=grpcGateway,proto3,oneof"`
}

type Config_Generator_Plugin_Swagger struct {
	Swagger *Config_Generator_Swagger `protobuf:"bytes,4,opt,name=swagger,proto3,oneof"`
}

func (*Config_Generator_Plugin_Go) isConfig_Generator_Plugin_Cmd() {}

func (*Config_Generator_Plugin_Micro) isConfig_Generator_Plugin_Cmd() {}

func (*Config_Generator_Plugin_GrpcGateway) isConfig_Generator_Plugin_Cmd() {}

func (*Config_Generator_Plugin_Swagger) isConfig_Generator_Plugin_Cmd() {}

func (m *Config_Generator_Plugin) GetCmd() isConfig_Generator_Plugin_Cmd {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *Config_Generator_Plugin) GetGo() *Config_Generator_Go {
	if x, ok := m.GetCmd().(*Config_Generator_Plugin_Go); ok {
		return x.Go
	}
	return nil
}

func (m *Config_Generator_Plugin) GetMicro() *Config_Generator_Micro {
	if x, ok := m.GetCmd().(*Config_Generator_Plugin_Micro); ok {
		return x.Micro
	}
	return nil
}

func (m *Config_Generator_Plugin) GetGrpcGateway() *Config_Generator_Grpc_Gateway {
	if x, ok := m.GetCmd().(*Config_Generator_Plugin_GrpcGateway); ok {
		return x.GrpcGateway
	}
	return nil
}

func (m *Config_Generator_Plugin) GetSwagger() *Config_Generator_Swagger {
	if x, ok := m.GetCmd().(*Config_Generator_Plugin_Swagger); ok {
		return x.Swagger
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Config_Generator_Plugin) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Config_Generator_Plugin_Go)(nil),
		(*Config_Generator_Plugin_Micro)(nil),
		(*Config_Generator_Plugin_GrpcGateway)(nil),
		(*Config_Generator_Plugin_Swagger)(nil),
	}
}

type Config_Generator_Go struct {
	Plugins []Config_Generator_Go_Plugins `protobuf:"varint,1,rep,packed,name=plugins,proto3,enum=godna.Config_Generator_Go_Plugins" json:"plugins,omitempty"`
	Paths   Config_Generator_Paths        `protobuf:"varint,2,opt,name=paths,proto3,enum=godna.Config_Generator_Paths" json:"paths,omitempty"`
	// string import_prefix = 3;
	AnnotateCode         bool     `protobuf:"varint,4,opt,name=annotate_code,json=annotateCode,proto3" json:"annotate_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config_Generator_Go) Reset()         { *m = Config_Generator_Go{} }
func (m *Config_Generator_Go) String() string { return proto.CompactTextString(m) }
func (*Config_Generator_Go) ProtoMessage()    {}
func (*Config_Generator_Go) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 1}
}

func (m *Config_Generator_Go) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator_Go.Unmarshal(m, b)
}
func (m *Config_Generator_Go) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator_Go.Marshal(b, m, deterministic)
}
func (m *Config_Generator_Go) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator_Go.Merge(m, src)
}
func (m *Config_Generator_Go) XXX_Size() int {
	return xxx_messageInfo_Config_Generator_Go.Size(m)
}
func (m *Config_Generator_Go) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator_Go.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator_Go proto.InternalMessageInfo

func (m *Config_Generator_Go) GetPlugins() []Config_Generator_Go_Plugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *Config_Generator_Go) GetPaths() Config_Generator_Paths {
	if m != nil {
		return m.Paths
	}
	return Config_Generator_Import
}

func (m *Config_Generator_Go) GetAnnotateCode() bool {
	if m != nil {
		return m.AnnotateCode
	}
	return false
}

type Config_Generator_Micro struct {
	Paths                Config_Generator_Paths `protobuf:"varint,2,opt,name=paths,proto3,enum=godna.Config_Generator_Paths" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Config_Generator_Micro) Reset()         { *m = Config_Generator_Micro{} }
func (m *Config_Generator_Micro) String() string { return proto.CompactTextString(m) }
func (*Config_Generator_Micro) ProtoMessage()    {}
func (*Config_Generator_Micro) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 2}
}

func (m *Config_Generator_Micro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator_Micro.Unmarshal(m, b)
}
func (m *Config_Generator_Micro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator_Micro.Marshal(b, m, deterministic)
}
func (m *Config_Generator_Micro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator_Micro.Merge(m, src)
}
func (m *Config_Generator_Micro) XXX_Size() int {
	return xxx_messageInfo_Config_Generator_Micro.Size(m)
}
func (m *Config_Generator_Micro) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator_Micro.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator_Micro proto.InternalMessageInfo

func (m *Config_Generator_Micro) GetPaths() Config_Generator_Paths {
	if m != nil {
		return m.Paths
	}
	return Config_Generator_Import
}

type Config_Generator_Grpc_Gateway struct {
	Paths                Config_Generator_Paths `protobuf:"varint,2,opt,name=paths,proto3,enum=godna.Config_Generator_Paths" json:"paths,omitempty"`
	RegisterFuncSuffix   string                 `protobuf:"bytes,3,opt,name=register_func_suffix,json=registerFuncSuffix,proto3" json:"register_func_suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Config_Generator_Grpc_Gateway) Reset()         { *m = Config_Generator_Grpc_Gateway{} }
func (m *Config_Generator_Grpc_Gateway) String() string { return proto.CompactTextString(m) }
func (*Config_Generator_Grpc_Gateway) ProtoMessage()    {}
func (*Config_Generator_Grpc_Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 3}
}

func (m *Config_Generator_Grpc_Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator_Grpc_Gateway.Unmarshal(m, b)
}
func (m *Config_Generator_Grpc_Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator_Grpc_Gateway.Marshal(b, m, deterministic)
}
func (m *Config_Generator_Grpc_Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator_Grpc_Gateway.Merge(m, src)
}
func (m *Config_Generator_Grpc_Gateway) XXX_Size() int {
	return xxx_messageInfo_Config_Generator_Grpc_Gateway.Size(m)
}
func (m *Config_Generator_Grpc_Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator_Grpc_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator_Grpc_Gateway proto.InternalMessageInfo

func (m *Config_Generator_Grpc_Gateway) GetPaths() Config_Generator_Paths {
	if m != nil {
		return m.Paths
	}
	return Config_Generator_Import
}

func (m *Config_Generator_Grpc_Gateway) GetRegisterFuncSuffix() string {
	if m != nil {
		return m.RegisterFuncSuffix
	}
	return ""
}

type Config_Generator_Swagger struct {
	Logtostderr          bool     `protobuf:"varint,1,opt,name=logtostderr,proto3" json:"logtostderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config_Generator_Swagger) Reset()         { *m = Config_Generator_Swagger{} }
func (m *Config_Generator_Swagger) String() string { return proto.CompactTextString(m) }
func (*Config_Generator_Swagger) ProtoMessage()    {}
func (*Config_Generator_Swagger) Descriptor() ([]byte, []int) {
	return fileDescriptor_738f7b83cb645dfe, []int{0, 2, 4}
}

func (m *Config_Generator_Swagger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Generator_Swagger.Unmarshal(m, b)
}
func (m *Config_Generator_Swagger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Generator_Swagger.Marshal(b, m, deterministic)
}
func (m *Config_Generator_Swagger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Generator_Swagger.Merge(m, src)
}
func (m *Config_Generator_Swagger) XXX_Size() int {
	return xxx_messageInfo_Config_Generator_Swagger.Size(m)
}
func (m *Config_Generator_Swagger) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Generator_Swagger.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Generator_Swagger proto.InternalMessageInfo

func (m *Config_Generator_Swagger) GetLogtostderr() bool {
	if m != nil {
		return m.Logtostderr
	}
	return false
}

func init() {
	proto.RegisterEnum("godna.Config_PluginOutDir_OutType", Config_PluginOutDir_OutType_name, Config_PluginOutDir_OutType_value)
	proto.RegisterEnum("godna.Config_Pass_Command", Config_Pass_Command_name, Config_Pass_Command_value)
	proto.RegisterEnum("godna.Config_Generator_Paths", Config_Generator_Paths_name, Config_Generator_Paths_value)
	proto.RegisterEnum("godna.Config_Generator_Go_Plugins", Config_Generator_Go_Plugins_name, Config_Generator_Go_Plugins_value)
	proto.RegisterType((*Config)(nil), "godna.Config")
	proto.RegisterType((*Config_PluginOutDir)(nil), "godna.Config.PluginOutDir")
	proto.RegisterType((*Config_Pass)(nil), "godna.Config.Pass")
	proto.RegisterType((*Config_Generator)(nil), "godna.Config.Generator")
	proto.RegisterType((*Config_Generator_Plugin)(nil), "godna.Config.Generator.Plugin")
	proto.RegisterType((*Config_Generator_Go)(nil), "godna.Config.Generator.Go")
	proto.RegisterType((*Config_Generator_Micro)(nil), "godna.Config.Generator.Micro")
	proto.RegisterType((*Config_Generator_Grpc_Gateway)(nil), "godna.Config.Generator.Grpc_Gateway")
	proto.RegisterType((*Config_Generator_Swagger)(nil), "godna.Config.Generator.Swagger")
}

func init() { proto.RegisterFile("dna/config/config.v1.proto", fileDescriptor_738f7b83cb645dfe) }

var fileDescriptor_738f7b83cb645dfe = []byte{
	// 1339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xf6, 0x26, 0xf1, 0xdf, 0x38, 0xb4, 0x66, 0x40, 0x62, 0x65, 0xa9, 0x30, 0x72, 0x41, 0x09,
	0xaa, 0xbd, 0x4e, 0x53, 0x15, 0x2e, 0x68, 0x25, 0xf2, 0xe3, 0x3a, 0x2b, 0x35, 0x76, 0xd8, 0x18,
	0x09, 0xa8, 0xaa, 0x65, 0xb2, 0x3b, 0x5e, 0x0f, 0xec, 0xce, 0x6c, 0x67, 0x66, 0x9d, 0x58, 0x42,
	0x82, 0x07, 0xe0, 0x29, 0x78, 0x06, 0x73, 0x55, 0x6e, 0xb9, 0xe5, 0x86, 0xb7, 0xb1, 0x7a, 0x81,
	0x66, 0x76, 0xdd, 0x04, 0x2a, 0x4b, 0xa8, 0x37, 0xc9, 0xfc, 0x9c, 0xf3, 0x9d, 0x73, 0xbe, 0xf3,
	0x9d, 0x59, 0x83, 0x56, 0xc8, 0x70, 0x2f, 0xe0, 0x6c, 0x42, 0xa3, 0xe2, 0x9f, 0x33, 0xbb, 0xef,
	0xa4, 0x82, 0x2b, 0x0e, 0xcb, 0x11, 0x0f, 0x19, 0x6e, 0xf5, 0xb4, 0x89, 0x54, 0x5c, 0x90, 0x1e,
	0xb9, 0x52, 0x84, 0x49, 0xca, 0x99, 0xec, 0xed, 0xed, 0xed, 0xed, 0x9d, 0x8f, 0x47, 0x5e, 0xbf,
	0x37, 0xc3, 0x31, 0x0d, 0xb1, 0xa2, 0x9c, 0xe5, 0x7e, 0xed, 0x3f, 0x20, 0xa8, 0x1c, 0x19, 0x2c,
	0x18, 0x03, 0x30, 0xe5, 0x52, 0xf9, 0xfc, 0x92, 0x11, 0x61, 0x5b, 0xc8, 0xda, 0xad, 0x1f, 0x9e,
	0x2e, 0x17, 0xb6, 0xab, 0x4f, 0x11, 0x66, 0x21, 0x32, 0x37, 0x88, 0x4f, 0x90, 0x9a, 0x12, 0x24,
	0x48, 0xca, 0x57, 0xeb, 0x88, 0x30, 0x22, 0xb0, 0x22, 0x21, 0x0a, 0x78, 0x48, 0x10, 0x89, 0x1c,
	0x14, 0x51, 0x35, 0xcd, 0x2e, 0x9c, 0x80, 0x27, 0xbd, 0x84, 0x06, 0x82, 0x4b, 0x3e, 0x51, 0x5e,
	0x5d, 0x43, 0x8d, 0x34, 0x0a, 0x3c, 0x05, 0x75, 0x8d, 0xe0, 0x33, 0x9c, 0x10, 0x7b, 0xc3, 0x04,
	0xdb, 0x5b, 0x2e, 0xec, 0x8e, 0x81, 0xd5, 0x87, 0x6b, 0xb1, 0x51, 0xc4, 0xbb, 0x33, 0xa9, 0x37,
	0x5e, 0x4d, 0x5b, 0x0f, 0x71, 0x42, 0xe0, 0x08, 0xd4, 0x28, 0x0b, 0xe2, 0x2c, 0x24, 0xd2, 0xde,
	0x44, 0x9b, 0xbb, 0xf5, 0xc3, 0x07, 0xcb, 0x85, 0xdd, 0x33, 0x45, 0x06, 0x28, 0xc5, 0x6a, 0x8a,
	0xba, 0xae, 0x76, 0x7e, 0x86, 0x76, 0x1c, 0xe7, 0x3a, 0xa5, 0xae, 0x26, 0x6c, 0x46, 0x58, 0xc8,
	0xc5, 0x0e, 0x7a, 0xee, 0xbd, 0x06, 0x81, 0x21, 0xa8, 0x09, 0x12, 0xfb, 0x34, 0x49, 0xa5, 0xbd,
	0x65, 0x00, 0xdd, 0xe5, 0xc2, 0xee, 0xa7, 0x82, 0x4c, 0xe8, 0x95, 0xce, 0x2d, 0xe2, 0x4e, 0xc2,
	0x43, 0x14, 0xf3, 0x00, 0xc7, 0x9a, 0x8c, 0x18, 0x07, 0x24, 0x21, 0x4c, 0x49, 0xa7, 0x08, 0x23,
	0x89, 0x98, 0xd1, 0x80, 0xc8, 0xee, 0x0e, 0x7a, 0x8e, 0xba, 0x48, 0xf1, 0x90, 0x23, 0x41, 0x12,
	0x3e, 0x23, 0x5e, 0x55, 0x90, 0xd8, 0x4d, 0x52, 0x09, 0xbf, 0x07, 0x55, 0x41, 0x5e, 0x64, 0x54,
	0x10, 0xbb, 0x6c, 0x82, 0x3c, 0x59, 0x2e, 0xec, 0xc3, 0x98, 0x4a, 0x95, 0x87, 0x40, 0x3a, 0x04,
	0x09, 0xa9, 0x42, 0xdd, 0xc2, 0xee, 0x71, 0x81, 0x7e, 0x83, 0xdf, 0x88, 0xc7, 0x98, 0x45, 0x2f,
	0x7a, 0x2f, 0xbe, 0x9c, 0xdd, 0x77, 0xf6, 0x9c, 0xcf, 0x75, 0x21, 0x2b, 0x58, 0xf8, 0x23, 0xa8,
	0x4a, 0x11, 0xf8, 0x21, 0x15, 0x76, 0xc5, 0xb0, 0xec, 0x2d, 0x17, 0xf6, 0x50, 0xf2, 0x4c, 0x04,
	0x04, 0x85, 0x54, 0x90, 0x40, 0x71, 0x31, 0xd7, 0xa0, 0x6f, 0xd0, 0x92, 0xeb, 0x28, 0x15, 0xfc,
	0x07, 0x12, 0x28, 0xd4, 0x45, 0x8c, 0x2b, 0xa2, 0xff, 0xa0, 0x82, 0xa4, 0x10, 0x61, 0x89, 0x30,
	0xea, 0xba, 0x5e, 0x45, 0x8a, 0xe0, 0x98, 0x0a, 0xf8, 0xca, 0x02, 0x5b, 0x29, 0x96, 0xd2, 0xae,
	0xa1, 0xcd, 0xdd, 0xc6, 0x3e, 0x74, 0x8c, 0x2a, 0x9d, 0x5c, 0x60, 0xce, 0x19, 0x96, 0xf2, 0xf0,
	0x6f, 0x6b, 0xb9, 0xb0, 0xff, 0xb2, 0x02, 0x9e, 0x24, 0x9c, 0x21, 0x49, 0x52, 0x9c, 0x37, 0x77,
	0x55, 0xb3, 0xbe, 0xc0, 0x2c, 0x94, 0x48, 0x71, 0x24, 0x32, 0xe6, 0xa0, 0x33, 0x2e, 0x25, 0xbd,
	0x88, 0xc9, 0xf5, 0x0d, 0x16, 0x04, 0xe5, 0x8d, 0xed, 0x24, 0x3c, 0xa4, 0x8c, 0x2a, 0xfd, 0xbf,
	0x28, 0x38, 0x5f, 0x9a, 0x6e, 0xe8, 0xa5, 0xa2, 0xe1, 0xbc, 0x13, 0x51, 0xa5, 0x70, 0x84, 0x76,
	0x43, 0x32, 0xc1, 0x59, 0xac, 0xd0, 0xb3, 0xf6, 0xff, 0x71, 0x6f, 0x77, 0x50, 0xbb, 0x40, 0xd0,
	0xcb, 0x1c, 0xa4, 0x38, 0x2c, 0xac, 0x4d, 0xf3, 0xdb, 0xcf, 0x3f, 0xf5, 0x4c, 0xd5, 0xf0, 0x37,
	0x0b, 0x34, 0xd2, 0x38, 0x8b, 0x28, 0xf3, 0xb5, 0xe2, 0xec, 0xba, 0x69, 0xe9, 0x2f, 0xba, 0xe2,
	0x9f, 0xb4, 0x4a, 0xcd, 0x10, 0x19, 0x2d, 0xf2, 0x09, 0xc2, 0x28, 0x37, 0x8d, 0x34, 0xf7, 0x79,
	0x3e, 0xdd, 0x88, 0xb0, 0xee, 0xf0, 0xe0, 0xb4, 0xff, 0x58, 0x1b, 0xf5, 0x14, 0xef, 0x25, 0xf3,
	0x0b, 0xca, 0xb0, 0x98, 0x3b, 0xe8, 0x34, 0x0b, 0xa6, 0x08, 0xc7, 0x92, 0x23, 0x99, 0x92, 0x80,
	0x4e, 0xe6, 0xa8, 0xdb, 0x2d, 0x46, 0x83, 0x8b, 0x0e, 0x0a, 0x39, 0x91, 0x79, 0x6f, 0x92, 0x34,
	0x9e, 0x23, 0xaa, 0x10, 0xa2, 0x12, 0xa5, 0x82, 0x48, 0xc2, 0x94, 0x07, 0xf2, 0x50, 0x67, 0x58,
	0x4d, 0xe1, 0x4b, 0x0b, 0xbc, 0x5b, 0x24, 0xc9, 0x33, 0x95, 0x66, 0xca, 0x68, 0x03, 0x98, 0x86,
	0xb5, 0xfe, 0xd3, 0x30, 0x63, 0x36, 0xca, 0xd4, 0x31, 0x15, 0x87, 0x72, 0xb9, 0xb0, 0x79, 0xee,
	0x71, 0x43, 0x37, 0x13, 0x2e, 0xcc, 0x94, 0xe6, 0x88, 0xb9, 0xf6, 0x77, 0xd2, 0x8b, 0x1d, 0x53,
	0xea, 0x0e, 0xc7, 0x29, 0xdd, 0x31, 0x36, 0xf7, 0x1d, 0x34, 0xe0, 0x1d, 0x34, 0x10, 0x69, 0xd0,
	0x41, 0xa7, 0x5a, 0x66, 0xc6, 0xa2, 0x1f, 0x70, 0x39, 0x57, 0x24, 0xe9, 0x0e, 0xb0, 0x22, 0x97,
	0x78, 0x6e, 0x0e, 0xf7, 0x1d, 0x74, 0x7e, 0x89, 0xa3, 0x88, 0x08, 0xef, 0x76, 0xba, 0x4a, 0x21,
	0x35, 0x59, 0xb4, 0xfe, 0xb4, 0xc0, 0xf6, 0xcd, 0xb4, 0x20, 0xd4, 0x8a, 0x53, 0xd3, 0xfc, 0xbd,
	0xf2, 0xcc, 0x1a, 0x3e, 0x06, 0x35, 0x9e, 0x29, 0x5f, 0xcd, 0xd3, 0xfc, 0x69, 0xb9, 0xb5, 0xdf,
	0x5e, 0x5f, 0x98, 0x33, 0xca, 0xd4, 0x78, 0x9e, 0x12, 0xaf, 0xca, 0xf3, 0x05, 0x7c, 0x08, 0xea,
	0xaf, 0x99, 0x35, 0x8f, 0x49, 0x63, 0xff, 0x83, 0x7f, 0xfb, 0x0f, 0x56, 0xd7, 0xde, 0xb5, 0x65,
	0xfb, 0x1e, 0xa8, 0x16, 0x50, 0xf0, 0x3d, 0x70, 0x7b, 0x38, 0x1a, 0x9f, 0xb8, 0xc3, 0x81, 0x7f,
	0x7e, 0xd6, 0x3f, 0x72, 0x0f, 0x9e, 0x36, 0x4b, 0xb0, 0x01, 0xaa, 0x83, 0x91, 0x7f, 0x3a, 0x3a,
	0x3e, 0x6f, 0x5a, 0xad, 0xdf, 0x2d, 0xb0, 0xa5, 0xe7, 0x01, 0x76, 0xc0, 0x66, 0x90, 0x84, 0xb6,
	0x85, 0x36, 0x77, 0x6f, 0xbd, 0xc1, 0x3f, 0x96, 0xd2, 0x39, 0xca, 0x65, 0xef, 0x69, 0xb3, 0xf6,
	0xcf, 0xa0, 0x5a, 0xec, 0x61, 0x0d, 0x6c, 0x0d, 0x47, 0xc3, 0x7e, 0xb3, 0x04, 0x01, 0xa8, 0x9c,
	0x79, 0xa3, 0xf1, 0xe8, 0xa8, 0x69, 0xc1, 0x6d, 0x50, 0x3b, 0x1d, 0x1d, 0xfb, 0xee, 0xd0, 0x1d,
	0x37, 0x37, 0xe0, 0x6d, 0xd0, 0xd0, 0x3b, 0xaf, 0xff, 0xd5, 0xd7, 0xae, 0xd7, 0x6f, 0x6e, 0x5e,
	0x1f, 0x9c, 0x3d, 0x3d, 0x38, 0xea, 0x37, 0xb7, 0x56, 0xf6, 0x63, 0xf7, 0xf8, 0xdb, 0x66, 0x59,
	0xe7, 0x3d, 0x70, 0xc7, 0xfe, 0xc1, 0xf1, 0xb1, 0x3f, 0xec, 0x7f, 0x33, 0x1e, 0x1f, 0x0c, 0x9a,
	0x15, 0x93, 0xb7, 0x3b, 0xf6, 0xf5, 0xa6, 0xda, 0xfa, 0xb5, 0x02, 0xea, 0xaf, 0xab, 0x87, 0x9f,
	0x81, 0x4a, 0xde, 0x20, 0x43, 0x7f, 0x63, 0xff, 0xc3, 0x35, 0x34, 0x15, 0x84, 0x7b, 0x85, 0x75,
	0xeb, 0x95, 0x05, 0x2a, 0xf9, 0x11, 0xec, 0x80, 0x8d, 0x88, 0x17, 0xee, 0xad, 0x75, 0xee, 0x03,
	0x7e, 0x52, 0xf2, 0x36, 0x22, 0x0e, 0x1f, 0x82, 0xb2, 0x79, 0xa3, 0x4c, 0x5b, 0x1b, 0xfb, 0x77,
	0xd6, 0x39, 0x18, 0x85, 0x9d, 0x94, 0xbc, 0xdc, 0x1a, 0xba, 0x60, 0x3b, 0x12, 0x69, 0xe0, 0x47,
	0xb9, 0xca, 0xec, 0x4d, 0xe3, 0xfd, 0xf1, 0xda, 0x70, 0xda, 0xb6, 0x50, 0xe4, 0x49, 0xc9, 0x6b,
	0x68, 0xdf, 0x62, 0x0b, 0xbf, 0x00, 0x55, 0x99, 0x8b, 0xd3, 0xde, 0x32, 0x28, 0x1f, 0xad, 0x43,
	0x29, 0x34, 0x7c, 0x52, 0xf2, 0x56, 0x1e, 0x87, 0x65, 0xd3, 0xec, 0xd6, 0x4b, 0x0b, 0x6c, 0x0c,
	0x38, 0x7c, 0x04, 0xaa, 0xc5, 0xd8, 0x14, 0xed, 0x6f, 0xaf, 0xaf, 0xbf, 0x60, 0x50, 0x7a, 0x2b,
	0x17, 0xf8, 0x00, 0x94, 0xb5, 0xd8, 0x65, 0xa1, 0xf0, 0xb5, 0x54, 0xe8, 0xa1, 0x97, 0x5e, 0x6e,
	0x0b, 0xef, 0x82, 0x77, 0x30, 0x63, 0x5c, 0x61, 0x45, 0x7c, 0xfd, 0x05, 0x35, 0x35, 0xd4, 0xbc,
	0xed, 0xd5, 0xe1, 0x11, 0x0f, 0x49, 0xfb, 0x0e, 0xa8, 0x16, 0xd1, 0x8c, 0xc8, 0x38, 0x23, 0xcd,
	0x92, 0x5e, 0x69, 0x5a, 0x9a, 0x56, 0xeb, 0x11, 0x28, 0x1b, 0x7a, 0xdf, 0x2a, 0x83, 0x56, 0x06,
	0xb6, 0x6f, 0xd2, 0xfb, 0x76, 0x65, 0xec, 0x81, 0xf7, 0x05, 0x89, 0xa8, 0x54, 0x44, 0xf8, 0x93,
	0x8c, 0x05, 0xbe, 0xcc, 0x26, 0x13, 0x7a, 0x65, 0xfa, 0x5a, 0xf7, 0xe0, 0xea, 0xee, 0x49, 0xc6,
	0x82, 0x73, 0x73, 0xd3, 0xba, 0x07, 0xaa, 0x45, 0x3f, 0x20, 0x02, 0x8d, 0x98, 0x47, 0x8a, 0x4b,
	0x15, 0x12, 0x91, 0xff, 0xd0, 0xa9, 0x79, 0x37, 0x8f, 0xda, 0xbb, 0xa0, 0x6c, 0xc2, 0xe9, 0xc9,
	0x72, 0x93, 0x94, 0x0b, 0xd5, 0x2c, 0xe9, 0xd9, 0x38, 0x37, 0x5f, 0x4d, 0xdf, 0x23, 0x31, 0x56,
	0x74, 0x46, 0x9a, 0xd6, 0xe1, 0x27, 0xdf, 0xdd, 0xbd, 0xf1, 0x21, 0xbe, 0xbc, 0xa2, 0xbc, 0x67,
	0xf2, 0xef, 0xa5, 0x17, 0xbd, 0xeb, 0x5f, 0x6b, 0x17, 0x15, 0xf3, 0xbe, 0x3f, 0xf8, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0x09, 0x39, 0x3f, 0xc2, 0x09, 0x00, 0x00,
}
