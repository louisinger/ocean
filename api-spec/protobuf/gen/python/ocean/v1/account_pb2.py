# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocean/v1/account.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ocean.v1 import types_pb2 as ocean_dot_v1_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16ocean/v1/account.proto\x12\x08ocean.v1\x1a\x14ocean/v1/types.proto\"/\n\x19\x43reateAccountBIP44Request\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"x\n\x1a\x43reateAccountBIP44Response\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12#\n\raccount_index\x18\x02 \x01(\rR\x0c\x61\x63\x63ountIndex\x12\x12\n\x04xpub\x18\x03 \x01(\tR\x04xpub\"0\n\x1a\x43reateAccountCustomRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"}\n\x1b\x43reateAccountCustomResponse\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12\'\n\x0f\x64\x65rivation_path\x18\x02 \x01(\tR\x0e\x64\x65rivationPath\x12\x12\n\x04xpub\x18\x03 \x01(\tR\x04xpub\"n\n\x19SetAccountTemplateRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12.\n\x08template\x18\x02 \x01(\x0b\x32\x12.ocean.v1.TemplateR\x08template\"\x1c\n\x1aSetAccountTemplateResponse\"c\n\x14\x44\x65riveAddressRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12(\n\x10num_of_addresses\x18\x02 \x01(\x04R\x0enumOfAddresses\"5\n\x15\x44\x65riveAddressResponse\x12\x1c\n\taddresses\x18\x01 \x03(\tR\taddresses\"i\n\x1a\x44\x65riveChangeAddressRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12(\n\x10num_of_addresses\x18\x02 \x01(\x04R\x0enumOfAddresses\";\n\x1b\x44\x65riveChangeAddressResponse\x12\x1c\n\taddresses\x18\x01 \x03(\tR\taddresses\"9\n\x14ListAddressesRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\"5\n\x15ListAddressesResponse\x12\x1c\n\taddresses\x18\x01 \x03(\tR\taddresses\"Q\n\x0e\x42\x61lanceRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12\x1c\n\taddresses\x18\x03 \x03(\tR\taddresses\"\xa6\x01\n\x0f\x42\x61lanceResponse\x12@\n\x07\x62\x61lance\x18\x01 \x03(\x0b\x32&.ocean.v1.BalanceResponse.BalanceEntryR\x07\x62\x61lance\x1aQ\n\x0c\x42\x61lanceEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12+\n\x05value\x18\x02 \x01(\x0b\x32\x15.ocean.v1.BalanceInfoR\x05value:\x02\x38\x01\"S\n\x10ListUtxosRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12\x1c\n\taddresses\x18\x03 \x03(\tR\taddresses\"\x81\x01\n\x11ListUtxosResponse\x12\x38\n\x0fspendable_utxos\x18\x01 \x01(\x0b\x32\x0f.ocean.v1.UtxosR\x0espendableUtxos\x12\x32\n\x0clocked_utxos\x18\x02 \x01(\x0b\x32\x0f.ocean.v1.UtxosR\x0blockedUtxos\"9\n\x14\x44\x65leteAccountRequest\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\"\x17\n\x15\x44\x65leteAccountResponse2\x96\x06\n\x0e\x41\x63\x63ountService\x12_\n\x12\x43reateAccountBIP44\x12#.ocean.v1.CreateAccountBIP44Request\x1a$.ocean.v1.CreateAccountBIP44Response\x12\x62\n\x13\x43reateAccountCustom\x12$.ocean.v1.CreateAccountCustomRequest\x1a%.ocean.v1.CreateAccountCustomResponse\x12_\n\x12SetAccountTemplate\x12#.ocean.v1.SetAccountTemplateRequest\x1a$.ocean.v1.SetAccountTemplateResponse\x12P\n\rDeriveAddress\x12\x1e.ocean.v1.DeriveAddressRequest\x1a\x1f.ocean.v1.DeriveAddressResponse\x12\x62\n\x13\x44\x65riveChangeAddress\x12$.ocean.v1.DeriveChangeAddressRequest\x1a%.ocean.v1.DeriveChangeAddressResponse\x12P\n\rListAddresses\x12\x1e.ocean.v1.ListAddressesRequest\x1a\x1f.ocean.v1.ListAddressesResponse\x12>\n\x07\x42\x61lance\x12\x18.ocean.v1.BalanceRequest\x1a\x19.ocean.v1.BalanceResponse\x12\x44\n\tListUtxos\x12\x1a.ocean.v1.ListUtxosRequest\x1a\x1b.ocean.v1.ListUtxosResponse\x12P\n\rDeleteAccount\x12\x1e.ocean.v1.DeleteAccountRequest\x1a\x1f.ocean.v1.DeleteAccountResponseB\xa5\x01\n\x0c\x63om.ocean.v1B\x0c\x41\x63\x63ountProtoP\x01ZFgithub.com/vulpemventures/ocean/api-spec/protobuf/gen/ocean/v1;oceanv1\xa2\x02\x03OXX\xaa\x02\x08Ocean.V1\xca\x02\x08Ocean\\V1\xe2\x02\x14Ocean\\V1\\GPBMetadata\xea\x02\tOcean::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ocean.v1.account_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\014com.ocean.v1B\014AccountProtoP\001ZFgithub.com/vulpemventures/ocean/api-spec/protobuf/gen/ocean/v1;oceanv1\242\002\003OXX\252\002\010Ocean.V1\312\002\010Ocean\\V1\342\002\024Ocean\\V1\\GPBMetadata\352\002\tOcean::V1'
  _BALANCERESPONSE_BALANCEENTRY._options = None
  _BALANCERESPONSE_BALANCEENTRY._serialized_options = b'8\001'
  _CREATEACCOUNTBIP44REQUEST._serialized_start=58
  _CREATEACCOUNTBIP44REQUEST._serialized_end=105
  _CREATEACCOUNTBIP44RESPONSE._serialized_start=107
  _CREATEACCOUNTBIP44RESPONSE._serialized_end=227
  _CREATEACCOUNTCUSTOMREQUEST._serialized_start=229
  _CREATEACCOUNTCUSTOMREQUEST._serialized_end=277
  _CREATEACCOUNTCUSTOMRESPONSE._serialized_start=279
  _CREATEACCOUNTCUSTOMRESPONSE._serialized_end=404
  _SETACCOUNTTEMPLATEREQUEST._serialized_start=406
  _SETACCOUNTTEMPLATEREQUEST._serialized_end=516
  _SETACCOUNTTEMPLATERESPONSE._serialized_start=518
  _SETACCOUNTTEMPLATERESPONSE._serialized_end=546
  _DERIVEADDRESSREQUEST._serialized_start=548
  _DERIVEADDRESSREQUEST._serialized_end=647
  _DERIVEADDRESSRESPONSE._serialized_start=649
  _DERIVEADDRESSRESPONSE._serialized_end=702
  _DERIVECHANGEADDRESSREQUEST._serialized_start=704
  _DERIVECHANGEADDRESSREQUEST._serialized_end=809
  _DERIVECHANGEADDRESSRESPONSE._serialized_start=811
  _DERIVECHANGEADDRESSRESPONSE._serialized_end=870
  _LISTADDRESSESREQUEST._serialized_start=872
  _LISTADDRESSESREQUEST._serialized_end=929
  _LISTADDRESSESRESPONSE._serialized_start=931
  _LISTADDRESSESRESPONSE._serialized_end=984
  _BALANCEREQUEST._serialized_start=986
  _BALANCEREQUEST._serialized_end=1067
  _BALANCERESPONSE._serialized_start=1070
  _BALANCERESPONSE._serialized_end=1236
  _BALANCERESPONSE_BALANCEENTRY._serialized_start=1155
  _BALANCERESPONSE_BALANCEENTRY._serialized_end=1236
  _LISTUTXOSREQUEST._serialized_start=1238
  _LISTUTXOSREQUEST._serialized_end=1321
  _LISTUTXOSRESPONSE._serialized_start=1324
  _LISTUTXOSRESPONSE._serialized_end=1453
  _DELETEACCOUNTREQUEST._serialized_start=1455
  _DELETEACCOUNTREQUEST._serialized_end=1512
  _DELETEACCOUNTRESPONSE._serialized_start=1514
  _DELETEACCOUNTRESPONSE._serialized_end=1537
  _ACCOUNTSERVICE._serialized_start=1540
  _ACCOUNTSERVICE._serialized_end=2330
# @@protoc_insertion_point(module_scope)
