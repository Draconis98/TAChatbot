# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: qa.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08qa.proto\x1a\x1bgoogle/protobuf/empty.proto\"&\n\x02QA\x12\x10\n\x08question\x18\x01 \x01(\t\x12\x0e\n\x06\x61nswer\x18\x02 \x01(\t23\n\x08QAServer\x12\'\n\x06PushQA\x12\x03.QA\x1a\x16.google.protobuf.Empty\"\x00\x42\x06Z\x04/apib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'qa_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\004/api'
  _globals['_QA']._serialized_start=41
  _globals['_QA']._serialized_end=79
  _globals['_QASERVER']._serialized_start=81
  _globals['_QASERVER']._serialized_end=132
# @@protoc_insertion_point(module_scope)
