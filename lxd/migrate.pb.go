// Code generated by protoc-gen-go.
// source: lxd/migrate.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	lxd/migrate.proto

It has these top-level messages:
	IDMapType
	Config
	Device
	Snapshot
	MigrationHeader
	MigrationControl
*/
package main

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MigrationFSType int32

const (
	MigrationFSType_RSYNC MigrationFSType = 0
	MigrationFSType_BTRFS MigrationFSType = 1
	MigrationFSType_ZFS   MigrationFSType = 2
)

var MigrationFSType_name = map[int32]string{
	0: "RSYNC",
	1: "BTRFS",
	2: "ZFS",
}
var MigrationFSType_value = map[string]int32{
	"RSYNC": 0,
	"BTRFS": 1,
	"ZFS":   2,
}

func (x MigrationFSType) Enum() *MigrationFSType {
	p := new(MigrationFSType)
	*p = x
	return p
}
func (x MigrationFSType) String() string {
	return proto.EnumName(MigrationFSType_name, int32(x))
}
func (x *MigrationFSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MigrationFSType_value, data, "MigrationFSType")
	if err != nil {
		return err
	}
	*x = MigrationFSType(value)
	return nil
}

type CRIUType int32

const (
	CRIUType_CRIU_RSYNC CRIUType = 0
	CRIUType_PHAUL      CRIUType = 1
)

var CRIUType_name = map[int32]string{
	0: "CRIU_RSYNC",
	1: "PHAUL",
}
var CRIUType_value = map[string]int32{
	"CRIU_RSYNC": 0,
	"PHAUL":      1,
}

func (x CRIUType) Enum() *CRIUType {
	p := new(CRIUType)
	*p = x
	return p
}
func (x CRIUType) String() string {
	return proto.EnumName(CRIUType_name, int32(x))
}
func (x *CRIUType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CRIUType_value, data, "CRIUType")
	if err != nil {
		return err
	}
	*x = CRIUType(value)
	return nil
}

type IDMapType struct {
	Isuid            *bool  `protobuf:"varint,1,req,name=isuid" json:"isuid,omitempty"`
	Isgid            *bool  `protobuf:"varint,2,req,name=isgid" json:"isgid,omitempty"`
	Hostid           *int32 `protobuf:"varint,3,req,name=hostid" json:"hostid,omitempty"`
	Nsid             *int32 `protobuf:"varint,4,req,name=nsid" json:"nsid,omitempty"`
	Maprange         *int32 `protobuf:"varint,5,req,name=maprange" json:"maprange,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IDMapType) Reset()         { *m = IDMapType{} }
func (m *IDMapType) String() string { return proto.CompactTextString(m) }
func (*IDMapType) ProtoMessage()    {}

func (m *IDMapType) GetIsuid() bool {
	if m != nil && m.Isuid != nil {
		return *m.Isuid
	}
	return false
}

func (m *IDMapType) GetIsgid() bool {
	if m != nil && m.Isgid != nil {
		return *m.Isgid
	}
	return false
}

func (m *IDMapType) GetHostid() int32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *IDMapType) GetNsid() int32 {
	if m != nil && m.Nsid != nil {
		return *m.Nsid
	}
	return 0
}

func (m *IDMapType) GetMaprange() int32 {
	if m != nil && m.Maprange != nil {
		return *m.Maprange
	}
	return 0
}

type Config struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}

func (m *Config) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Device struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config           []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Snapshot struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config           []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	Profiles         []string  `protobuf:"bytes,3,rep,name=profiles" json:"profiles,omitempty"`
	Ephemeral        *bool     `protobuf:"varint,4,req,name=ephemeral" json:"ephemeral,omitempty"`
	Devices          []*Device `protobuf:"bytes,5,rep,name=devices" json:"devices,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}

func (m *Snapshot) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Snapshot) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Snapshot) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Snapshot) GetEphemeral() bool {
	if m != nil && m.Ephemeral != nil {
		return *m.Ephemeral
	}
	return false
}

func (m *Snapshot) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type MigrationHeader struct {
	Fs               *MigrationFSType `protobuf:"varint,1,req,name=fs,enum=main.MigrationFSType" json:"fs,omitempty"`
	Criu             *CRIUType        `protobuf:"varint,2,opt,name=criu,enum=main.CRIUType" json:"criu,omitempty"`
	Idmap            []*IDMapType     `protobuf:"bytes,3,rep,name=idmap" json:"idmap,omitempty"`
	SnapshotNames    []string         `protobuf:"bytes,4,rep,name=snapshotNames" json:"snapshotNames,omitempty"`
	Snapshots        []*Snapshot      `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MigrationHeader) Reset()         { *m = MigrationHeader{} }
func (m *MigrationHeader) String() string { return proto.CompactTextString(m) }
func (*MigrationHeader) ProtoMessage()    {}

func (m *MigrationHeader) GetFs() MigrationFSType {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return MigrationFSType_RSYNC
}

func (m *MigrationHeader) GetCriu() CRIUType {
	if m != nil && m.Criu != nil {
		return *m.Criu
	}
	return CRIUType_CRIU_RSYNC
}

func (m *MigrationHeader) GetIdmap() []*IDMapType {
	if m != nil {
		return m.Idmap
	}
	return nil
}

func (m *MigrationHeader) GetSnapshotNames() []string {
	if m != nil {
		return m.SnapshotNames
	}
	return nil
}

func (m *MigrationHeader) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type MigrationControl struct {
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// optional failure message if sending a failure
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MigrationControl) Reset()         { *m = MigrationControl{} }
func (m *MigrationControl) String() string { return proto.CompactTextString(m) }
func (*MigrationControl) ProtoMessage()    {}

func (m *MigrationControl) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *MigrationControl) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("main.MigrationFSType", MigrationFSType_name, MigrationFSType_value)
	proto.RegisterEnum("main.CRIUType", CRIUType_name, CRIUType_value)
}
