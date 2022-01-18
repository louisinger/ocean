# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wallet/wallet.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13wallet/wallet.proto\x12\x06wallet\"\x10\n\x0eGenSeedRequest\"f\n\x0cGenSeedReply\x12)\n\x10signing_mnemonic\x18\x01 \x01(\tR\x0fsigningMnemonic\x12+\n\x11\x62linding_mnemonic\x18\x02 \x01(\tR\x10\x62lindingMnemonic\"\x89\x01\n\x13\x43reateWalletRequest\x12)\n\x10signing_mnemonic\x18\x01 \x01(\tR\x0fsigningMnemonic\x12+\n\x11\x62linding_mnemonic\x18\x02 \x01(\tR\x10\x62lindingMnemonic\x12\x1a\n\x08password\x18\x03 \x01(\x0cR\x08password\"\x13\n\x11\x43reateWalletReply\"+\n\rUnlockRequest\x12\x1a\n\x08password\x18\x01 \x01(\x0cR\x08password\"\r\n\x0bUnlockReply\"e\n\x15\x43hangePasswordRequest\x12)\n\x10\x63urrent_password\x18\x01 \x01(\x0cR\x0f\x63urrentPassword\x12!\n\x0cnew_password\x18\x02 \x01(\x0cR\x0bnewPassword\"\x15\n\x13\x43hangePasswordReply\"\x8a\x01\n\x14RestoreWalletRequest\x12)\n\x10signing_mnemonic\x18\x01 \x01(\tR\x0fsigningMnemonic\x12+\n\x11\x62linding_mnemonic\x18\x02 \x01(\tR\x10\x62lindingMnemonic\x12\x1a\n\x08password\x18\x03 \x01(\x0cR\x08password\"\x14\n\x12RestoreWalletReply\"\x0f\n\rStatusRequest\"\x81\x01\n\x0bStatusReply\x12\x32\n\x06status\x18\x01 \x01(\x0e\x32\x1a.wallet.StatusReply.StatusR\x06status\">\n\x06Status\x12\x0f\n\x0bNOT_CREATED\x10\x00\x12\n\n\x06\x43LOSED\x10\x01\x12\r\n\tRESTORING\x10\x02\x12\x08\n\x04OPEN\x10\x03\x32\x95\x03\n\rWalletService\x12\x37\n\x07GenSeed\x12\x16.wallet.GenSeedRequest\x1a\x14.wallet.GenSeedReply\x12\x46\n\x0c\x43reateWallet\x12\x1b.wallet.CreateWalletRequest\x1a\x19.wallet.CreateWalletReply\x12\x34\n\x06Unlock\x12\x15.wallet.UnlockRequest\x1a\x13.wallet.UnlockReply\x12L\n\x0e\x43hangePassword\x12\x1d.wallet.ChangePasswordRequest\x1a\x1b.wallet.ChangePasswordReply\x12I\n\rRestoreWallet\x12\x1c.wallet.RestoreWalletRequest\x1a\x1a.wallet.RestoreWalletReply\x12\x34\n\x06Status\x12\x15.wallet.StatusRequest\x1a\x13.wallet.StatusReplyB\x92\x01\n\ncom.walletB\x0bWalletProtoP\x01Z?github.com/vulpemventures/ocean/api-spec/protobuf/gen/go/wallet\xa2\x02\x03WXX\xaa\x02\x06Wallet\xca\x02\x06Wallet\xe2\x02\x12Wallet\\GPBMetadata\xea\x02\x06Walletb\x06proto3')



_GENSEEDREQUEST = DESCRIPTOR.message_types_by_name['GenSeedRequest']
_GENSEEDREPLY = DESCRIPTOR.message_types_by_name['GenSeedReply']
_CREATEWALLETREQUEST = DESCRIPTOR.message_types_by_name['CreateWalletRequest']
_CREATEWALLETREPLY = DESCRIPTOR.message_types_by_name['CreateWalletReply']
_UNLOCKREQUEST = DESCRIPTOR.message_types_by_name['UnlockRequest']
_UNLOCKREPLY = DESCRIPTOR.message_types_by_name['UnlockReply']
_CHANGEPASSWORDREQUEST = DESCRIPTOR.message_types_by_name['ChangePasswordRequest']
_CHANGEPASSWORDREPLY = DESCRIPTOR.message_types_by_name['ChangePasswordReply']
_RESTOREWALLETREQUEST = DESCRIPTOR.message_types_by_name['RestoreWalletRequest']
_RESTOREWALLETREPLY = DESCRIPTOR.message_types_by_name['RestoreWalletReply']
_STATUSREQUEST = DESCRIPTOR.message_types_by_name['StatusRequest']
_STATUSREPLY = DESCRIPTOR.message_types_by_name['StatusReply']
_STATUSREPLY_STATUS = _STATUSREPLY.enum_types_by_name['Status']
GenSeedRequest = _reflection.GeneratedProtocolMessageType('GenSeedRequest', (_message.Message,), {
  'DESCRIPTOR' : _GENSEEDREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.GenSeedRequest)
  })
_sym_db.RegisterMessage(GenSeedRequest)

GenSeedReply = _reflection.GeneratedProtocolMessageType('GenSeedReply', (_message.Message,), {
  'DESCRIPTOR' : _GENSEEDREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.GenSeedReply)
  })
_sym_db.RegisterMessage(GenSeedReply)

CreateWalletRequest = _reflection.GeneratedProtocolMessageType('CreateWalletRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEWALLETREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.CreateWalletRequest)
  })
_sym_db.RegisterMessage(CreateWalletRequest)

CreateWalletReply = _reflection.GeneratedProtocolMessageType('CreateWalletReply', (_message.Message,), {
  'DESCRIPTOR' : _CREATEWALLETREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.CreateWalletReply)
  })
_sym_db.RegisterMessage(CreateWalletReply)

UnlockRequest = _reflection.GeneratedProtocolMessageType('UnlockRequest', (_message.Message,), {
  'DESCRIPTOR' : _UNLOCKREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.UnlockRequest)
  })
_sym_db.RegisterMessage(UnlockRequest)

UnlockReply = _reflection.GeneratedProtocolMessageType('UnlockReply', (_message.Message,), {
  'DESCRIPTOR' : _UNLOCKREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.UnlockReply)
  })
_sym_db.RegisterMessage(UnlockReply)

ChangePasswordRequest = _reflection.GeneratedProtocolMessageType('ChangePasswordRequest', (_message.Message,), {
  'DESCRIPTOR' : _CHANGEPASSWORDREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.ChangePasswordRequest)
  })
_sym_db.RegisterMessage(ChangePasswordRequest)

ChangePasswordReply = _reflection.GeneratedProtocolMessageType('ChangePasswordReply', (_message.Message,), {
  'DESCRIPTOR' : _CHANGEPASSWORDREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.ChangePasswordReply)
  })
_sym_db.RegisterMessage(ChangePasswordReply)

RestoreWalletRequest = _reflection.GeneratedProtocolMessageType('RestoreWalletRequest', (_message.Message,), {
  'DESCRIPTOR' : _RESTOREWALLETREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.RestoreWalletRequest)
  })
_sym_db.RegisterMessage(RestoreWalletRequest)

RestoreWalletReply = _reflection.GeneratedProtocolMessageType('RestoreWalletReply', (_message.Message,), {
  'DESCRIPTOR' : _RESTOREWALLETREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.RestoreWalletReply)
  })
_sym_db.RegisterMessage(RestoreWalletReply)

StatusRequest = _reflection.GeneratedProtocolMessageType('StatusRequest', (_message.Message,), {
  'DESCRIPTOR' : _STATUSREQUEST,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.StatusRequest)
  })
_sym_db.RegisterMessage(StatusRequest)

StatusReply = _reflection.GeneratedProtocolMessageType('StatusReply', (_message.Message,), {
  'DESCRIPTOR' : _STATUSREPLY,
  '__module__' : 'wallet.wallet_pb2'
  # @@protoc_insertion_point(class_scope:wallet.StatusReply)
  })
_sym_db.RegisterMessage(StatusReply)

_WALLETSERVICE = DESCRIPTOR.services_by_name['WalletService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\ncom.walletB\013WalletProtoP\001Z?github.com/vulpemventures/ocean/api-spec/protobuf/gen/go/wallet\242\002\003WXX\252\002\006Wallet\312\002\006Wallet\342\002\022Wallet\\GPBMetadata\352\002\006Wallet'
  _GENSEEDREQUEST._serialized_start=31
  _GENSEEDREQUEST._serialized_end=47
  _GENSEEDREPLY._serialized_start=49
  _GENSEEDREPLY._serialized_end=151
  _CREATEWALLETREQUEST._serialized_start=154
  _CREATEWALLETREQUEST._serialized_end=291
  _CREATEWALLETREPLY._serialized_start=293
  _CREATEWALLETREPLY._serialized_end=312
  _UNLOCKREQUEST._serialized_start=314
  _UNLOCKREQUEST._serialized_end=357
  _UNLOCKREPLY._serialized_start=359
  _UNLOCKREPLY._serialized_end=372
  _CHANGEPASSWORDREQUEST._serialized_start=374
  _CHANGEPASSWORDREQUEST._serialized_end=475
  _CHANGEPASSWORDREPLY._serialized_start=477
  _CHANGEPASSWORDREPLY._serialized_end=498
  _RESTOREWALLETREQUEST._serialized_start=501
  _RESTOREWALLETREQUEST._serialized_end=639
  _RESTOREWALLETREPLY._serialized_start=641
  _RESTOREWALLETREPLY._serialized_end=661
  _STATUSREQUEST._serialized_start=663
  _STATUSREQUEST._serialized_end=678
  _STATUSREPLY._serialized_start=681
  _STATUSREPLY._serialized_end=810
  _STATUSREPLY_STATUS._serialized_start=748
  _STATUSREPLY_STATUS._serialized_end=810
  _WALLETSERVICE._serialized_start=813
  _WALLETSERVICE._serialized_end=1218
# @@protoc_insertion_point(module_scope)
