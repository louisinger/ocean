# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocean/v1/types.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14ocean/v1/types.proto\x12\x08ocean.v1\"Q\n\tBuildInfo\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12\x16\n\x06\x63ommit\x18\x02 \x01(\tR\x06\x63ommit\x12\x12\n\x04\x64\x61te\x18\x03 \x01(\tR\x04\x64\x61te\"v\n\x0b\x41\x63\x63ountInfo\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n\x05index\x18\x02 \x01(\rR\x05index\x12\'\n\x0f\x64\x65rivation_path\x18\x03 \x01(\tR\x0e\x64\x65rivationPath\x12\x14\n\x05xpubs\x18\x04 \x03(\tR\x05xpubs\"\xb7\x01\n\x0b\x42\x61lanceInfo\x12+\n\x11\x63onfirmed_balance\x18\x01 \x01(\x04R\x10\x63onfirmedBalance\x12/\n\x13unconfirmed_balance\x18\x02 \x01(\x04R\x12unconfirmedBalance\x12%\n\x0elocked_balance\x18\x03 \x01(\x04R\rlockedBalance\x12#\n\rtotal_balance\x18\x04 \x01(\x04R\x0ctotalBalance\"\x93\x01\n\x05Input\x12\x12\n\x04txid\x18\x01 \x01(\tR\x04txid\x12\x14\n\x05index\x18\x02 \x01(\rR\x05index\x12\x16\n\x06script\x18\x03 \x01(\tR\x06script\x12%\n\x0escriptsig_size\x18\x04 \x01(\x04R\rscriptsigSize\x12!\n\x0cwitness_size\x18\x05 \x01(\x04R\x0bwitnessSize\"\xa0\x01\n\x0eUnblindedInput\x12\x14\n\x05index\x18\x01 \x01(\rR\x05index\x12\x14\n\x05\x61sset\x18\x02 \x01(\tR\x05\x61sset\x12\x16\n\x06\x61mount\x18\x03 \x01(\x04R\x06\x61mount\x12#\n\rasset_blinder\x18\x04 \x01(\tR\x0c\x61ssetBlinder\x12%\n\x0e\x61mount_blinder\x18\x05 \x01(\tR\ramountBlinder\"\x91\x01\n\x06Output\x12\x14\n\x05\x61sset\x18\x01 \x01(\tR\x05\x61sset\x12\x16\n\x06\x61mount\x18\x02 \x01(\x04R\x06\x61mount\x12\x18\n\x07\x61\x64\x64ress\x18\x03 \x01(\tR\x07\x61\x64\x64ress\x12\x16\n\x06script\x18\x04 \x01(\tR\x06script\x12\'\n\x0f\x62linding_pubkey\x18\x05 \x01(\tR\x0e\x62lindingPubkey\"P\n\x05Utxos\x12!\n\x0c\x61\x63\x63ount_name\x18\x01 \x01(\tR\x0b\x61\x63\x63ountName\x12$\n\x05utxos\x18\x02 \x03(\x0b\x32\x0e.ocean.v1.UtxoR\x05utxos\"m\n\nUtxoStatus\x12\x12\n\x04txid\x18\x01 \x01(\tR\x04txid\x12\x35\n\nblock_info\x18\x02 \x01(\x0b\x32\x16.ocean.v1.BlockDetailsR\tblockInfo\x12\x14\n\x05txhex\x18\x03 \x01(\tR\x05txhex\"\x80\x03\n\x04Utxo\x12\x12\n\x04txid\x18\x01 \x01(\tR\x04txid\x12\x14\n\x05index\x18\x02 \x01(\rR\x05index\x12\x14\n\x05\x61sset\x18\x03 \x01(\tR\x05\x61sset\x12\x14\n\x05value\x18\x04 \x01(\x04R\x05value\x12\x16\n\x06script\x18\x05 \x01(\tR\x06script\x12#\n\rasset_blinder\x18\x06 \x01(\tR\x0c\x61ssetBlinder\x12#\n\rvalue_blinder\x18\x07 \x01(\tR\x0cvalueBlinder\x12!\n\x0c\x61\x63\x63ount_name\x18\x08 \x01(\tR\x0b\x61\x63\x63ountName\x12\x37\n\x0cspent_status\x18\t \x01(\x0b\x32\x14.ocean.v1.UtxoStatusR\x0bspentStatus\x12?\n\x10\x63onfirmed_status\x18\n \x01(\x0b\x32\x14.ocean.v1.UtxoStatusR\x0f\x63onfirmedStatus\x12#\n\rredeem_script\x18\x0b \x01(\tR\x0credeemScript\"X\n\x0c\x42lockDetails\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\x12\x16\n\x06height\x18\x02 \x01(\x04R\x06height\x12\x1c\n\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"\xc5\x01\n\x08Template\x12\x31\n\x06\x66ormat\x18\x01 \x01(\x0e\x32\x19.ocean.v1.Template.FormatR\x06\x66ormat\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\"p\n\x06\x46ormat\x12\x16\n\x12\x46ORMAT_UNSPECIFIED\x10\x00\x12\x15\n\x11\x46ORMAT_DESCRIPTOR\x10\x01\x12\x15\n\x11\x46ORMAT_MINISCRIPT\x10\x02\x12\x10\n\x0c\x46ORMAT_IONIO\x10\x03\x12\x0e\n\nFORMAT_RAW\x10\x04*\x87\x01\n\x0bTxEventType\x12\x1d\n\x19TX_EVENT_TYPE_UNSPECIFIED\x10\x00\x12\x1d\n\x19TX_EVENT_TYPE_BROADCASTED\x10\x01\x12\x1d\n\x19TX_EVENT_TYPE_UNCONFIRMED\x10\x02\x12\x1b\n\x17TX_EVENT_TYPE_CONFIRMED\x10\x03*\xbd\x01\n\rUtxoEventType\x12\x1f\n\x1bUTXO_EVENT_TYPE_UNSPECIFIED\x10\x00\x12\x17\n\x13UTXO_EVENT_TYPE_NEW\x10\x01\x12\x1d\n\x19UTXO_EVENT_TYPE_CONFIRMED\x10\x02\x12\x1a\n\x16UTXO_EVENT_TYPE_LOCKED\x10\x03\x12\x1c\n\x18UTXO_EVENT_TYPE_UNLOCKED\x10\x04\x12\x19\n\x15UTXO_EVENT_TYPE_SPENT\x10\x05*w\n\x10WebhookEventType\x12\"\n\x1eWEBHOOK_EVENT_TYPE_UNSPECIFIED\x10\x00\x12\"\n\x1eWEBHOOK_EVENT_TYPE_TRANSACTION\x10\x01\x12\x1b\n\x17WEBHOOK_EVENT_TYPE_UTXO\x10\x02\x42\xa3\x01\n\x0c\x63om.ocean.v1B\nTypesProtoP\x01ZFgithub.com/vulpemventures/ocean/api-spec/protobuf/gen/ocean/v1;oceanv1\xa2\x02\x03OXX\xaa\x02\x08Ocean.V1\xca\x02\x08Ocean\\V1\xe2\x02\x14Ocean\\V1\\GPBMetadata\xea\x02\tOcean::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ocean.v1.types_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\014com.ocean.v1B\nTypesProtoP\001ZFgithub.com/vulpemventures/ocean/api-spec/protobuf/gen/ocean/v1;oceanv1\242\002\003OXX\252\002\010Ocean.V1\312\002\010Ocean\\V1\342\002\024Ocean\\V1\\GPBMetadata\352\002\tOcean::V1'
  _TXEVENTTYPE._serialized_start=1755
  _TXEVENTTYPE._serialized_end=1890
  _UTXOEVENTTYPE._serialized_start=1893
  _UTXOEVENTTYPE._serialized_end=2082
  _WEBHOOKEVENTTYPE._serialized_start=2084
  _WEBHOOKEVENTTYPE._serialized_end=2203
  _BUILDINFO._serialized_start=34
  _BUILDINFO._serialized_end=115
  _ACCOUNTINFO._serialized_start=117
  _ACCOUNTINFO._serialized_end=235
  _BALANCEINFO._serialized_start=238
  _BALANCEINFO._serialized_end=421
  _INPUT._serialized_start=424
  _INPUT._serialized_end=571
  _UNBLINDEDINPUT._serialized_start=574
  _UNBLINDEDINPUT._serialized_end=734
  _OUTPUT._serialized_start=737
  _OUTPUT._serialized_end=882
  _UTXOS._serialized_start=884
  _UTXOS._serialized_end=964
  _UTXOSTATUS._serialized_start=966
  _UTXOSTATUS._serialized_end=1075
  _UTXO._serialized_start=1078
  _UTXO._serialized_end=1462
  _BLOCKDETAILS._serialized_start=1464
  _BLOCKDETAILS._serialized_end=1552
  _TEMPLATE._serialized_start=1555
  _TEMPLATE._serialized_end=1752
  _TEMPLATE_FORMAT._serialized_start=1640
  _TEMPLATE_FORMAT._serialized_end=1752
# @@protoc_insertion_point(module_scope)
