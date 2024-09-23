// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: coinbase/crypto/rosetta/types/account_identifer.proto

// The stable release for rosetta types

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AccountIdentifier The account_identifier uniquely identifies an account within a network. All
// fields in the account_identifier are utilized to determine this uniqueness (including the
// metadata field, if populated).
type AccountIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address may be a cryptographic public key (or some encoding of it) or a provided
	// username.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The sub account identifier
	SubAccount *SubAccountIdentifier `protobuf:"bytes,2,opt,name=sub_account,json=subAccount,proto3" json:"sub_account,omitempty"`
	// Blockchains that utilize a username model (where the address is not a derivative of a
	// cryptographic public key) should specify the public key(s) owned by the address in metadata.
	Metadata map[string]*anypb.Any `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountIdentifier) Reset() {
	*x = AccountIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIdentifier) ProtoMessage() {}

func (x *AccountIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIdentifier.ProtoReflect.Descriptor instead.
func (*AccountIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescGZIP(), []int{0}
}

func (x *AccountIdentifier) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AccountIdentifier) GetSubAccount() *SubAccountIdentifier {
	if x != nil {
		return x.SubAccount
	}
	return nil
}

func (x *AccountIdentifier) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance).
// The sub_account_identifier should specify which state (if applicable) an account instantiation refers to.
type SubAccountIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SubAccount address may be a cryptographic value or some other identifier (ex: bonded) that uniquely
	// specifies a SubAccount.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// If the SubAccount address is not sufficient to uniquely specify a SubAccount, any other identifying
	// information can be stored here. It is important to note that two SubAccounts with identical addresses
	// but differing metadata will not be considered equal by clients.
	Metadata map[string]*anypb.Any `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubAccountIdentifier) Reset() {
	*x = SubAccountIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubAccountIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubAccountIdentifier) ProtoMessage() {}

func (x *SubAccountIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubAccountIdentifier.ProtoReflect.Descriptor instead.
func (*SubAccountIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescGZIP(), []int{1}
}

func (x *SubAccountIdentifier) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SubAccountIdentifier) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_coinbase_crypto_rosetta_types_account_identifer_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x54, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65,
	0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f,
	0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x47, 0x5a, 0x45, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescData = file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coinbase_crypto_rosetta_types_account_identifer_proto_goTypes = []interface{}{
	(*AccountIdentifier)(nil),    // 0: coinbase.crypto.rosetta.types.AccountIdentifier
	(*SubAccountIdentifier)(nil), // 1: coinbase.crypto.rosetta.types.SubAccountIdentifier
	nil,                          // 2: coinbase.crypto.rosetta.types.AccountIdentifier.MetadataEntry
	nil,                          // 3: coinbase.crypto.rosetta.types.SubAccountIdentifier.MetadataEntry
	(*anypb.Any)(nil),            // 4: google.protobuf.Any
}
var file_coinbase_crypto_rosetta_types_account_identifer_proto_depIdxs = []int32{
	1, // 0: coinbase.crypto.rosetta.types.AccountIdentifier.sub_account:type_name -> coinbase.crypto.rosetta.types.SubAccountIdentifier
	2, // 1: coinbase.crypto.rosetta.types.AccountIdentifier.metadata:type_name -> coinbase.crypto.rosetta.types.AccountIdentifier.MetadataEntry
	3, // 2: coinbase.crypto.rosetta.types.SubAccountIdentifier.metadata:type_name -> coinbase.crypto.rosetta.types.SubAccountIdentifier.MetadataEntry
	4, // 3: coinbase.crypto.rosetta.types.AccountIdentifier.MetadataEntry.value:type_name -> google.protobuf.Any
	4, // 4: coinbase.crypto.rosetta.types.SubAccountIdentifier.MetadataEntry.value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_account_identifer_proto_init() }
func file_coinbase_crypto_rosetta_types_account_identifer_proto_init() {
	if File_coinbase_crypto_rosetta_types_account_identifer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountIdentifier); i {
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
		file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubAccountIdentifier); i {
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
			RawDescriptor: file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_account_identifer_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_account_identifer_proto_depIdxs,
		MessageInfos:      file_coinbase_crypto_rosetta_types_account_identifer_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_account_identifer_proto = out.File
	file_coinbase_crypto_rosetta_types_account_identifer_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_account_identifer_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_account_identifer_proto_depIdxs = nil
}
