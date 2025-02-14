// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: wallet/wallet.proto

package wallet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusReply_Status int32

const (
	// there is no data about wallet seeds
	StatusReply_NOT_CREATED StatusReply_Status = 0
	// wallet is closed and it needs to be opened with correct password in order to be used
	StatusReply_CLOSED StatusReply_Status = 1
	// wallet is trying to restore information's about unspents and used addresses
	StatusReply_RESTORING StatusReply_Status = 2
	// wallet is opened and ready to use
	StatusReply_OPEN StatusReply_Status = 3
)

// Enum value maps for StatusReply_Status.
var (
	StatusReply_Status_name = map[int32]string{
		0: "NOT_CREATED",
		1: "CLOSED",
		2: "RESTORING",
		3: "OPEN",
	}
	StatusReply_Status_value = map[string]int32{
		"NOT_CREATED": 0,
		"CLOSED":      1,
		"RESTORING":   2,
		"OPEN":        3,
	}
)

func (x StatusReply_Status) Enum() *StatusReply_Status {
	p := new(StatusReply_Status)
	*p = x
	return p
}

func (x StatusReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_wallet_wallet_proto_enumTypes[0].Descriptor()
}

func (StatusReply_Status) Type() protoreflect.EnumType {
	return &file_wallet_wallet_proto_enumTypes[0]
}

func (x StatusReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusReply_Status.Descriptor instead.
func (StatusReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{11, 0}
}

type GenSeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenSeedRequest) Reset() {
	*x = GenSeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSeedRequest) ProtoMessage() {}

func (x *GenSeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSeedRequest.ProtoReflect.Descriptor instead.
func (*GenSeedRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{0}
}

type GenSeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signing_mnemonic is 24 length phrase used to derive addresses and scripts
	SigningMnemonic string `protobuf:"bytes,1,opt,name=signing_mnemonic,json=signingMnemonic,proto3" json:"signing_mnemonic,omitempty"`
	// blinding_mnemonic is 24 length phrase used to derive blinding keys
	BlindingMnemonic string `protobuf:"bytes,2,opt,name=blinding_mnemonic,json=blindingMnemonic,proto3" json:"blinding_mnemonic,omitempty"`
}

func (x *GenSeedReply) Reset() {
	*x = GenSeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSeedReply) ProtoMessage() {}

func (x *GenSeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSeedReply.ProtoReflect.Descriptor instead.
func (*GenSeedReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *GenSeedReply) GetSigningMnemonic() string {
	if x != nil {
		return x.SigningMnemonic
	}
	return ""
}

func (x *GenSeedReply) GetBlindingMnemonic() string {
	if x != nil {
		return x.BlindingMnemonic
	}
	return ""
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signing_mnemonic is 24 length phrase used to derive addresses and scripts
	SigningMnemonic string `protobuf:"bytes,1,opt,name=signing_mnemonic,json=signingMnemonic,proto3" json:"signing_mnemonic,omitempty"`
	// blinding_mnemonic is 24 length phrase used to derive blinding keys
	BlindingMnemonic string `protobuf:"bytes,2,opt,name=blinding_mnemonic,json=blindingMnemonic,proto3" json:"blinding_mnemonic,omitempty"`
	// password is used to encrypt HD Wallet wallet. After creation, password is required to unlock the wallet daemon.
	Password []byte `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWalletRequest) GetSigningMnemonic() string {
	if x != nil {
		return x.SigningMnemonic
	}
	return ""
}

func (x *CreateWalletRequest) GetBlindingMnemonic() string {
	if x != nil {
		return x.BlindingMnemonic
	}
	return ""
}

func (x *CreateWalletRequest) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

type CreateWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWalletReply) Reset() {
	*x = CreateWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletReply) ProtoMessage() {}

func (x *CreateWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletReply.ProtoReflect.Descriptor instead.
func (*CreateWalletReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{3}
}

type UnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// password is used to decrypt HD Wallet wallet and start wallet daemon.
	Password []byte `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UnlockRequest) Reset() {
	*x = UnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockRequest) ProtoMessage() {}

func (x *UnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockRequest.ProtoReflect.Descriptor instead.
func (*UnlockRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockRequest) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

type UnlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlockReply) Reset() {
	*x = UnlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockReply) ProtoMessage() {}

func (x *UnlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockReply.ProtoReflect.Descriptor instead.
func (*UnlockReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{5}
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// current_password should be the current valid passphrase used to unlock wallet daemon.
	CurrentPassword []byte `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	// new_password should be the new passphrase that will be needed to unlock the wallet daemon.
	NewPassword []byte `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *ChangePasswordRequest) GetCurrentPassword() []byte {
	if x != nil {
		return x.CurrentPassword
	}
	return nil
}

func (x *ChangePasswordRequest) GetNewPassword() []byte {
	if x != nil {
		return x.NewPassword
	}
	return nil
}

type ChangePasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangePasswordReply) Reset() {
	*x = ChangePasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReply) ProtoMessage() {}

func (x *ChangePasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReply.ProtoReflect.Descriptor instead.
func (*ChangePasswordReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{7}
}

type RestoreWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signing_mnemonic is 24 length phrase used to derive addresses and scripts
	SigningMnemonic string `protobuf:"bytes,1,opt,name=signing_mnemonic,json=signingMnemonic,proto3" json:"signing_mnemonic,omitempty"`
	// blinding_mnemonic is 24 length phrase used to derive blinding keys
	BlindingMnemonic string `protobuf:"bytes,2,opt,name=blinding_mnemonic,json=blindingMnemonic,proto3" json:"blinding_mnemonic,omitempty"`
	// password is used to decrypt HD Wallet wallet and start wallet daemon.
	Password []byte `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RestoreWalletRequest) Reset() {
	*x = RestoreWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreWalletRequest) ProtoMessage() {}

func (x *RestoreWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreWalletRequest.ProtoReflect.Descriptor instead.
func (*RestoreWalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *RestoreWalletRequest) GetSigningMnemonic() string {
	if x != nil {
		return x.SigningMnemonic
	}
	return ""
}

func (x *RestoreWalletRequest) GetBlindingMnemonic() string {
	if x != nil {
		return x.BlindingMnemonic
	}
	return ""
}

func (x *RestoreWalletRequest) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

type RestoreWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestoreWalletReply) Reset() {
	*x = RestoreWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreWalletReply) ProtoMessage() {}

func (x *RestoreWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreWalletReply.ProtoReflect.Descriptor instead.
func (*RestoreWalletReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{9}
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{10}
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusReply_Status `protobuf:"varint,1,opt,name=status,proto3,enum=wallet.StatusReply_Status" json:"status,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{11}
}

func (x *StatusReply) GetStatus() StatusReply_Status {
	if x != nil {
		return x.Status
	}
	return StatusReply_NOT_CREATED
}

var File_wallet_wallet_proto protoreflect.FileDescriptor

var file_wallet_wallet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x66, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x6c,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x6c,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x03, 0x32, 0x95, 0x03, 0x0a, 0x0d, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x07,
	0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a,
	0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x12, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x8f, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c,
	0x70, 0x65, 0x6d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0xa2, 0x02,
	0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0xca, 0x02, 0x06,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0xe2, 0x02, 0x12, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_wallet_proto_rawDescOnce sync.Once
	file_wallet_wallet_proto_rawDescData = file_wallet_wallet_proto_rawDesc
)

func file_wallet_wallet_proto_rawDescGZIP() []byte {
	file_wallet_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_wallet_proto_rawDescData)
	})
	return file_wallet_wallet_proto_rawDescData
}

var file_wallet_wallet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wallet_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_wallet_wallet_proto_goTypes = []interface{}{
	(StatusReply_Status)(0),       // 0: wallet.StatusReply.Status
	(*GenSeedRequest)(nil),        // 1: wallet.GenSeedRequest
	(*GenSeedReply)(nil),          // 2: wallet.GenSeedReply
	(*CreateWalletRequest)(nil),   // 3: wallet.CreateWalletRequest
	(*CreateWalletReply)(nil),     // 4: wallet.CreateWalletReply
	(*UnlockRequest)(nil),         // 5: wallet.UnlockRequest
	(*UnlockReply)(nil),           // 6: wallet.UnlockReply
	(*ChangePasswordRequest)(nil), // 7: wallet.ChangePasswordRequest
	(*ChangePasswordReply)(nil),   // 8: wallet.ChangePasswordReply
	(*RestoreWalletRequest)(nil),  // 9: wallet.RestoreWalletRequest
	(*RestoreWalletReply)(nil),    // 10: wallet.RestoreWalletReply
	(*StatusRequest)(nil),         // 11: wallet.StatusRequest
	(*StatusReply)(nil),           // 12: wallet.StatusReply
}
var file_wallet_wallet_proto_depIdxs = []int32{
	0,  // 0: wallet.StatusReply.status:type_name -> wallet.StatusReply.Status
	1,  // 1: wallet.WalletService.GenSeed:input_type -> wallet.GenSeedRequest
	3,  // 2: wallet.WalletService.CreateWallet:input_type -> wallet.CreateWalletRequest
	5,  // 3: wallet.WalletService.Unlock:input_type -> wallet.UnlockRequest
	7,  // 4: wallet.WalletService.ChangePassword:input_type -> wallet.ChangePasswordRequest
	9,  // 5: wallet.WalletService.RestoreWallet:input_type -> wallet.RestoreWalletRequest
	11, // 6: wallet.WalletService.Status:input_type -> wallet.StatusRequest
	2,  // 7: wallet.WalletService.GenSeed:output_type -> wallet.GenSeedReply
	4,  // 8: wallet.WalletService.CreateWallet:output_type -> wallet.CreateWalletReply
	6,  // 9: wallet.WalletService.Unlock:output_type -> wallet.UnlockReply
	8,  // 10: wallet.WalletService.ChangePassword:output_type -> wallet.ChangePasswordReply
	10, // 11: wallet.WalletService.RestoreWallet:output_type -> wallet.RestoreWalletReply
	12, // 12: wallet.WalletService.Status:output_type -> wallet.StatusReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_wallet_wallet_proto_init() }
func file_wallet_wallet_proto_init() {
	if File_wallet_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSeedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSeedReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreWalletReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_wallet_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_wallet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_wallet_proto_depIdxs,
		EnumInfos:         file_wallet_wallet_proto_enumTypes,
		MessageInfos:      file_wallet_wallet_proto_msgTypes,
	}.Build()
	File_wallet_wallet_proto = out.File
	file_wallet_wallet_proto_rawDesc = nil
	file_wallet_wallet_proto_goTypes = nil
	file_wallet_wallet_proto_depIdxs = nil
}
