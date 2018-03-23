// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/config.proto

/*
Package fleetspeak_daemonservice is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/config.proto
	fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/messages.proto

It has these top-level messages:
	Config
	StdOutputData
*/
package fleetspeak_daemonservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The configuration information expected by daemonservice.Factory in
// ClientServiceConfig.config.
type Config struct {
	Argv []string `protobuf:"bytes,1,rep,name=argv" json:"argv,omitempty"`
	// If set, process will be killed after this much inactivity. Any message to
	// or from the process, and any stdin/stderr output counts as inactivity.
	InactivityTimeout *google_protobuf.Duration `protobuf:"bytes,2,opt,name=inactivity_timeout,json=inactivityTimeout" json:"inactivity_timeout,omitempty"`
	// If set, start the process only when there is a message for it to work on.
	// Forced to true when inactivity timeout is set.
	LazyStart bool `protobuf:"varint,3,opt,name=lazy_start,json=lazyStart" json:"lazy_start,omitempty"`
	// By default, daemon services report resource usage every 10 minutes. This
	// flag disables this if set.
	DisableResourceMonitoring bool `protobuf:"varint,4,opt,name=disable_resource_monitoring,json=disableResourceMonitoring" json:"disable_resource_monitoring,omitempty"`
	// How many samples to aggregate into a report when monitoring resource usage.
	// If unset, defaults to 20.
	ResourceMonitoringSampleSize int32 `protobuf:"varint,5,opt,name=resource_monitoring_sample_size,json=resourceMonitoringSampleSize" json:"resource_monitoring_sample_size,omitempty"`
	// How long to wait between resource monitoring samples. If unset, defaults to
	// 30.
	ResourceMonitoringSamplePeriodSeconds int32 `protobuf:"varint,6,opt,name=resource_monitoring_sample_period_seconds,json=resourceMonitoringSamplePeriodSeconds" json:"resource_monitoring_sample_period_seconds,omitempty"`
	// If set, Fleetspeak will kill and restart the child if it exceeds this
	// memory limit, in bytes.
	MemoryLimit int64 `protobuf:"varint,7,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	// If set, Fleetspeak will monitor child's heartbeat messages and kill
	// unresponsive processes. The values below should be set to configure the
	// heartbeat monitoring.
	MonitorHeartbeats bool `protobuf:"varint,8,opt,name=monitor_heartbeats,json=monitorHeartbeats" json:"monitor_heartbeats,omitempty"`
	// How long to wait for initial heartbeat.
	HeartbeatUnresponsiveGracePeriodSeconds int32 `protobuf:"varint,9,opt,name=heartbeat_unresponsive_grace_period_seconds,json=heartbeatUnresponsiveGracePeriodSeconds" json:"heartbeat_unresponsive_grace_period_seconds,omitempty"`
	// How long to wait for subsequent heartbeats.
	HeartbeatUnresponsiveKillPeriodSeconds int32             `protobuf:"varint,10,opt,name=heartbeat_unresponsive_kill_period_seconds,json=heartbeatUnresponsiveKillPeriodSeconds" json:"heartbeat_unresponsive_kill_period_seconds,omitempty"`
	StdParams                              *Config_StdParams `protobuf:"bytes,11,opt,name=std_params,json=stdParams" json:"std_params,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

func (m *Config) GetInactivityTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.InactivityTimeout
	}
	return nil
}

func (m *Config) GetLazyStart() bool {
	if m != nil {
		return m.LazyStart
	}
	return false
}

func (m *Config) GetDisableResourceMonitoring() bool {
	if m != nil {
		return m.DisableResourceMonitoring
	}
	return false
}

func (m *Config) GetResourceMonitoringSampleSize() int32 {
	if m != nil {
		return m.ResourceMonitoringSampleSize
	}
	return 0
}

func (m *Config) GetResourceMonitoringSamplePeriodSeconds() int32 {
	if m != nil {
		return m.ResourceMonitoringSamplePeriodSeconds
	}
	return 0
}

func (m *Config) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *Config) GetMonitorHeartbeats() bool {
	if m != nil {
		return m.MonitorHeartbeats
	}
	return false
}

func (m *Config) GetHeartbeatUnresponsiveGracePeriodSeconds() int32 {
	if m != nil {
		return m.HeartbeatUnresponsiveGracePeriodSeconds
	}
	return 0
}

func (m *Config) GetHeartbeatUnresponsiveKillPeriodSeconds() int32 {
	if m != nil {
		return m.HeartbeatUnresponsiveKillPeriodSeconds
	}
	return 0
}

func (m *Config) GetStdParams() *Config_StdParams {
	if m != nil {
		return m.StdParams
	}
	return nil
}

// If set, we forward stderr and stdout data to the server as messages with:
//
// message_type="StdOutput"
// data=<fleetspeak.daemonservice.StdOutputData>
type Config_StdParams struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// A message will be sent when we have flush_bytes queued, or when we
	// have bytes flush_time_seconds old.
	FlushBytes       int32 `protobuf:"varint,2,opt,name=flush_bytes,json=flushBytes" json:"flush_bytes,omitempty"`
	FlushTimeSeconds int32 `protobuf:"varint,3,opt,name=flush_time_seconds,json=flushTimeSeconds" json:"flush_time_seconds,omitempty"`
}

func (m *Config_StdParams) Reset()                    { *m = Config_StdParams{} }
func (m *Config_StdParams) String() string            { return proto.CompactTextString(m) }
func (*Config_StdParams) ProtoMessage()               {}
func (*Config_StdParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Config_StdParams) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Config_StdParams) GetFlushBytes() int32 {
	if m != nil {
		return m.FlushBytes
	}
	return 0
}

func (m *Config_StdParams) GetFlushTimeSeconds() int32 {
	if m != nil {
		return m.FlushTimeSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "fleetspeak.daemonservice.Config")
	proto.RegisterType((*Config_StdParams)(nil), "fleetspeak.daemonservice.Config.StdParams")
}

func init() {
	proto.RegisterFile("fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x86, 0x15, 0xba, 0x95, 0xe5, 0x94, 0x0b, 0xe6, 0x2b, 0x6f, 0x7c, 0x2c, 0x20, 0x01, 0x61,
	0x40, 0x22, 0xc1, 0x3d, 0x17, 0x7c, 0x88, 0x21, 0xbe, 0xa6, 0x14, 0x24, 0x84, 0x90, 0x2c, 0x37,
	0x39, 0xcd, 0xac, 0x39, 0x76, 0x64, 0x3b, 0x95, 0x5a, 0x89, 0x7f, 0xc6, 0x8f, 0x43, 0x71, 0xb2,
	0x94, 0xb2, 0x95, 0x3b, 0xeb, 0x3d, 0xcf, 0xfb, 0xc4, 0x8e, 0x0e, 0x7c, 0x99, 0x4b, 0x44, 0x67,
	0x6b, 0xe4, 0xe7, 0xa9, 0x35, 0x79, 0x9a, 0x4b, 0x81, 0xca, 0xa5, 0x05, 0xc7, 0x4a, 0x2b, 0x8b,
	0x66, 0x21, 0x72, 0x4c, 0x6b, 0xa3, 0x9d, 0x4e, 0xd7, 0x24, 0xdb, 0x1c, 0xe7, 0x5a, 0xcd, 0x45,
	0x99, 0x78, 0x8a, 0xd0, 0x35, 0x96, 0x6c, 0x60, 0x87, 0x77, 0x4b, 0xad, 0x4b, 0xd9, 0xdb, 0x66,
	0xcd, 0x3c, 0x2d, 0x1a, 0xc3, 0x9d, 0xd0, 0xaa, 0x6b, 0xde, 0xff, 0x3d, 0x86, 0xf1, 0x6b, 0xaf,
	0x22, 0x04, 0x76, 0xb8, 0x29, 0x17, 0x34, 0x88, 0x46, 0x71, 0x98, 0xf9, 0x33, 0x39, 0x01, 0x22,
	0x14, 0xcf, 0x9d, 0x58, 0x08, 0xb7, 0x64, 0x4e, 0x54, 0xa8, 0x1b, 0x47, 0xaf, 0x45, 0x41, 0x3c,
	0x79, 0x7e, 0x90, 0x74, 0xee, 0xe4, 0xc2, 0x9d, 0xbc, 0xe9, 0xdd, 0xd9, 0xfe, 0xba, 0xf4, 0xb5,
	0xeb, 0x90, 0x3b, 0x00, 0x92, 0xaf, 0x96, 0xcc, 0x3a, 0x6e, 0x1c, 0x1d, 0x45, 0x41, 0xbc, 0x97,
	0x85, 0x6d, 0x32, 0x6d, 0x03, 0xf2, 0x12, 0x6e, 0x15, 0xc2, 0xf2, 0x99, 0x44, 0x66, 0xd0, 0xea,
	0xc6, 0xe4, 0xc8, 0x2a, 0xad, 0x84, 0xd3, 0x46, 0xa8, 0x92, 0xee, 0x78, 0xfe, 0xa0, 0x47, 0xb2,
	0x9e, 0xf8, 0x34, 0x00, 0xe4, 0x2d, 0x1c, 0x5d, 0xd1, 0x63, 0x96, 0x57, 0xb5, 0x44, 0x66, 0xc5,
	0x0a, 0xe9, 0x6e, 0x14, 0xc4, 0xbb, 0xd9, 0x6d, 0x73, 0xa9, 0x3c, 0xf5, 0xd0, 0x54, 0xac, 0x90,
	0x7c, 0x87, 0xc7, 0xff, 0xd1, 0xd4, 0x68, 0x84, 0x2e, 0x98, 0xc5, 0x5c, 0xab, 0xc2, 0xd2, 0xb1,
	0x17, 0x3e, 0xd8, 0x26, 0x3c, 0xf5, 0xf4, 0xb4, 0x83, 0xc9, 0x3d, 0xb8, 0x51, 0x61, 0xa5, 0xcd,
	0x92, 0x49, 0x51, 0x09, 0x47, 0xaf, 0x47, 0x41, 0x3c, 0xca, 0x26, 0x5d, 0xf6, 0xb1, 0x8d, 0xc8,
	0x33, 0x20, 0xfd, 0x37, 0xd9, 0x19, 0x72, 0xe3, 0x66, 0xc8, 0x9d, 0xa5, 0x7b, 0xfe, 0xe9, 0xfb,
	0xfd, 0xe4, 0x64, 0x18, 0x90, 0x9f, 0xf0, 0x64, 0xc0, 0x58, 0xa3, 0x0c, 0xda, 0x5a, 0x2b, 0x2b,
	0x16, 0xc8, 0x4a, 0xc3, 0xf3, 0x4b, 0xb7, 0x0d, 0xfd, 0x6d, 0x1f, 0x0d, 0x95, 0x6f, 0x7f, 0x35,
	0xde, 0xb5, 0x85, 0xcd, 0xfb, 0xfe, 0x80, 0xe3, 0x2d, 0xf6, 0x73, 0x21, 0xe5, 0xbf, 0x72, 0xf0,
	0xf2, 0x87, 0x57, 0xca, 0x3f, 0x08, 0x29, 0x37, 0xdd, 0xef, 0x01, 0xac, 0x2b, 0x58, 0xcd, 0x0d,
	0xaf, 0x2c, 0x9d, 0xf8, 0x6d, 0x3a, 0x4e, 0xb6, 0xed, 0x70, 0xd2, 0xed, 0x67, 0x32, 0x75, 0xc5,
	0xa9, 0x6f, 0x64, 0xa1, 0xbd, 0x38, 0x1e, 0xfe, 0x82, 0x70, 0xc8, 0xdb, 0x7f, 0xdc, 0x77, 0x98,
	0xe2, 0x15, 0xd2, 0x20, 0x0a, 0xe2, 0x30, 0x9b, 0xf4, 0xd9, 0x67, 0x5e, 0x21, 0x39, 0x82, 0xc9,
	0x5c, 0x36, 0xf6, 0x8c, 0xcd, 0x96, 0x0e, 0xad, 0xdf, 0xe4, 0xdd, 0x0c, 0x7c, 0xf4, 0xaa, 0x4d,
	0xc8, 0x53, 0x20, 0x1d, 0xd0, 0x2e, 0xfb, 0xf0, 0xbe, 0x91, 0xe7, 0x6e, 0xfa, 0x49, 0xbb, 0xd1,
	0xfd, 0x4b, 0x66, 0x63, 0xbf, 0xfb, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xc1, 0x66,
	0x90, 0xd2, 0x03, 0x00, 0x00,
}
