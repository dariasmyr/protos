// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: sso/sso.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountStatus int32

const (
	AccountStatus_ACTIVE   AccountStatus = 0
	AccountStatus_INACTIVE AccountStatus = 1
	AccountStatus_DELETED  AccountStatus = 2
)

// Enum value maps for AccountStatus.
var (
	AccountStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
		2: "DELETED",
	}
	AccountStatus_value = map[string]int32{
		"ACTIVE":   0,
		"INACTIVE": 1,
		"DELETED":  2,
	}
)

func (x AccountStatus) Enum() *AccountStatus {
	p := new(AccountStatus)
	*p = x
	return p
}

func (x AccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_sso_proto_enumTypes[0].Descriptor()
}

func (AccountStatus) Type() protoreflect.EnumType {
	return &file_sso_sso_proto_enumTypes[0]
}

func (x AccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountStatus.Descriptor instead.
func (AccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

type AccountRole int32

const (
	AccountRole_USER  AccountRole = 0
	AccountRole_ADMIN AccountRole = 1
)

// Enum value maps for AccountRole.
var (
	AccountRole_name = map[int32]string{
		0: "USER",
		1: "ADMIN",
	}
	AccountRole_value = map[string]int32{
		"USER":  0,
		"ADMIN": 1,
	}
)

func (x AccountRole) Enum() *AccountRole {
	p := new(AccountRole)
	*p = x
	return p
}

func (x AccountRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountRole) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_sso_proto_enumTypes[1].Descriptor()
}

func (AccountRole) Type() protoreflect.EnumType {
	return &file_sso_sso_proto_enumTypes[1]
}

func (x AccountRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountRole.Descriptor instead.
func (AccountRole) EnumDescriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

type RegisterClientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppName       string                 `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Secret        string                 `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	RedirectUrl   string                 `protobuf:"bytes,3,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	mi := &file_sso_sso_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClientRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RegisterClientRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *RegisterClientRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type RegisterClientResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         int64                  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	mi := &file_sso_sso_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClientResponse) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         int64                  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role          AccountRole            `protobuf:"varint,4,opt,name=role,proto3,enum=auth.AccountRole" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_sso_sso_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetRole() AccountRole {
	if x != nil {
		return x.Role
	}
	return AccountRole_USER
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_sso_sso_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_sso_sso_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_sso_sso_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_sso_sso_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OldPassword   string                 `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword   string                 `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	mi := &file_sso_sso_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[7]
	if x != nil {
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
	return file_sso_sso_proto_rawDescGZIP(), []int{7}
}

func (x *ChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordResponse) Reset() {
	*x = ChangePasswordResponse{}
	mi := &file_sso_sso_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordResponse) ProtoMessage() {}

func (x *ChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{8}
}

func (x *ChangePasswordResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ChangeStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Status        AccountStatus          `protobuf:"varint,2,opt,name=status,proto3,enum=auth.AccountStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	mi := &file_sso_sso_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{9}
}

func (x *ChangeStatusRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ChangeStatusRequest) GetStatus() AccountStatus {
	if x != nil {
		return x.Status
	}
	return AccountStatus_ACTIVE
}

type ChangeStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Status        AccountStatus          `protobuf:"varint,2,opt,name=status,proto3,enum=auth.AccountStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeStatusResponse) Reset() {
	*x = ChangeStatusResponse{}
	mi := &file_sso_sso_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusResponse) ProtoMessage() {}

func (x *ChangeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusResponse.ProtoReflect.Descriptor instead.
func (*ChangeStatusResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{10}
}

func (x *ChangeStatusResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ChangeStatusResponse) GetStatus() AccountStatus {
	if x != nil {
		return x.Status
	}
	return AccountStatus_ACTIVE
}

type GetActiveAccountSessionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sessions      []*Session             `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActiveAccountSessionsResponse) Reset() {
	*x = GetActiveAccountSessionsResponse{}
	mi := &file_sso_sso_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActiveAccountSessionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveAccountSessionsResponse) ProtoMessage() {}

func (x *GetActiveAccountSessionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveAccountSessionsResponse.ProtoReflect.Descriptor instead.
func (*GetActiveAccountSessionsResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{11}
}

func (x *GetActiveAccountSessionsResponse) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type RefreshAccountSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt     int64                  `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshAccountSessionResponse) Reset() {
	*x = RefreshAccountSessionResponse{}
	mi := &file_sso_sso_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshAccountSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccountSessionResponse) ProtoMessage() {}

func (x *RefreshAccountSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccountSessionResponse.ProtoReflect.Descriptor instead.
func (*RefreshAccountSessionResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{12}
}

func (x *RefreshAccountSessionResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RefreshAccountSessionResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshAccountSessionResponse) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type ValidateAccountSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Valid         bool                   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateAccountSessionResponse) Reset() {
	*x = ValidateAccountSessionResponse{}
	mi := &file_sso_sso_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateAccountSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAccountSessionResponse) ProtoMessage() {}

func (x *ValidateAccountSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAccountSessionResponse.ProtoReflect.Descriptor instead.
func (*ValidateAccountSessionResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{13}
}

func (x *ValidateAccountSessionResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type RevokeAccountSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeAccountSessionRequest) Reset() {
	*x = RevokeAccountSessionRequest{}
	mi := &file_sso_sso_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeAccountSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAccountSessionRequest) ProtoMessage() {}

func (x *RevokeAccountSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAccountSessionRequest.ProtoReflect.Descriptor instead.
func (*RevokeAccountSessionRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{14}
}

func (x *RevokeAccountSessionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RevokeAccountSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         bool                   `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	Revoked       bool                   `protobuf:"varint,2,opt,name=revoked,proto3" json:"revoked,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeAccountSessionResponse) Reset() {
	*x = RevokeAccountSessionResponse{}
	mi := &file_sso_sso_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeAccountSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeAccountSessionResponse) ProtoMessage() {}

func (x *RevokeAccountSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeAccountSessionResponse.ProtoReflect.Descriptor instead.
func (*RevokeAccountSessionResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{15}
}

func (x *RevokeAccountSessionResponse) GetToken() bool {
	if x != nil {
		return x.Token
	}
	return false
}

func (x *RevokeAccountSessionResponse) GetRevoked() bool {
	if x != nil {
		return x.Revoked
	}
	return false
}

type Session struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AccountId        int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Token            string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken     string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	UserAgent        string                 `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	IpAddress        string                 `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	ExpiresAt        int64                  `protobuf:"varint,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	RefreshExpiresAt int64                  `protobuf:"varint,7,opt,name=refresh_expires_at,json=refreshExpiresAt,proto3" json:"refresh_expires_at,omitempty"`
	CreatedAt        int64                  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Revoked          bool                   `protobuf:"varint,10,opt,name=revoked,proto3" json:"revoked,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Session) Reset() {
	*x = Session{}
	mi := &file_sso_sso_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{16}
}

func (x *Session) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Session) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Session) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Session) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Session) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *Session) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Session) GetRefreshExpiresAt() int64 {
	if x != nil {
		return x.RefreshExpiresAt
	}
	return 0
}

func (x *Session) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Session) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Session) GetRevoked() bool {
	if x != nil {
		return x.Revoked
	}
	return false
}

var File_sso_sso_proto protoreflect.FileDescriptor

const file_sso_sso_proto_rawDesc = "" +
	"\n" +
	"\rsso/sso.proto\x12\x04auth\x1a\x1bgoogle/protobuf/empty.proto\"m\n" +
	"\x15RegisterClientRequest\x12\x19\n" +
	"\bapp_name\x18\x01 \x01(\tR\aappName\x12\x16\n" +
	"\x06secret\x18\x02 \x01(\tR\x06secret\x12!\n" +
	"\fredirect_url\x18\x03 \x01(\tR\vredirectUrl\"/\n" +
	"\x16RegisterClientResponse\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\x03R\x05appId\"\x81\x01\n" +
	"\x0fRegisterRequest\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\x03R\x05appId\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12%\n" +
	"\x04role\x18\x04 \x01(\x0e2\x11.auth.AccountRoleR\x04role\"1\n" +
	"\x10RegisterResponse\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\x03R\taccountId\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"i\n" +
	"\rLoginResponse\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\x03R\taccountId\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12#\n" +
	"\rrefresh_token\x18\x03 \x01(\tR\frefreshToken\"*\n" +
	"\x0eLogoutResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"]\n" +
	"\x15ChangePasswordRequest\x12!\n" +
	"\fold_password\x18\x01 \x01(\tR\voldPassword\x12!\n" +
	"\fnew_password\x18\x02 \x01(\tR\vnewPassword\"2\n" +
	"\x16ChangePasswordResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"a\n" +
	"\x13ChangeStatusRequest\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\x03R\taccountId\x12+\n" +
	"\x06status\x18\x02 \x01(\x0e2\x13.auth.AccountStatusR\x06status\"b\n" +
	"\x14ChangeStatusResponse\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\x03R\taccountId\x12+\n" +
	"\x06status\x18\x02 \x01(\x0e2\x13.auth.AccountStatusR\x06status\"M\n" +
	" GetActiveAccountSessionsResponse\x12)\n" +
	"\bsessions\x18\x01 \x03(\v2\r.auth.SessionR\bsessions\"y\n" +
	"\x1dRefreshAccountSessionResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x12\x1d\n" +
	"\n" +
	"expires_at\x18\x03 \x01(\x03R\texpiresAt\"6\n" +
	"\x1eValidateAccountSessionResponse\x12\x14\n" +
	"\x05valid\x18\x01 \x01(\bR\x05valid\"3\n" +
	"\x1bRevokeAccountSessionRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"N\n" +
	"\x1cRevokeAccountSessionResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\bR\x05token\x12\x18\n" +
	"\arevoked\x18\x02 \x01(\bR\arevoked\"\xc6\x02\n" +
	"\aSession\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\x03R\taccountId\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12#\n" +
	"\rrefresh_token\x18\x03 \x01(\tR\frefreshToken\x12\x1d\n" +
	"\n" +
	"user_agent\x18\x04 \x01(\tR\tuserAgent\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x05 \x01(\tR\tipAddress\x12\x1d\n" +
	"\n" +
	"expires_at\x18\x06 \x01(\x03R\texpiresAt\x12,\n" +
	"\x12refresh_expires_at\x18\a \x01(\x03R\x10refreshExpiresAt\x12\x1d\n" +
	"\n" +
	"created_at\x18\b \x01(\x03R\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\t \x01(\x03R\tupdatedAt\x12\x18\n" +
	"\arevoked\x18\n" +
	" \x01(\bR\arevoked*6\n" +
	"\rAccountStatus\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x00\x12\f\n" +
	"\bINACTIVE\x10\x01\x12\v\n" +
	"\aDELETED\x10\x02*\"\n" +
	"\vAccountRole\x12\b\n" +
	"\x04USER\x10\x00\x12\t\n" +
	"\x05ADMIN\x10\x012\x8c\x03\n" +
	"\x04Auth\x12K\n" +
	"\x0eRegisterClient\x12\x1b.auth.RegisterClientRequest\x1a\x1c.auth.RegisterClientResponse\x129\n" +
	"\bRegister\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x120\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x126\n" +
	"\x06Logout\x12\x16.google.protobuf.Empty\x1a\x14.auth.LogoutResponse\x12K\n" +
	"\x0eChangePassword\x12\x1b.auth.ChangePasswordRequest\x1a\x1c.auth.ChangePasswordResponse\x12E\n" +
	"\fChangeStatus\x12\x19.auth.ChangeStatusRequest\x1a\x1a.auth.ChangeStatusResponse2\xd7\x02\n" +
	"\bSessions\x12S\n" +
	"\x11GetActiveSessions\x12\x16.google.protobuf.Empty\x1a&.auth.GetActiveAccountSessionsResponse\x12M\n" +
	"\x0eRefreshSession\x12\x16.google.protobuf.Empty\x1a#.auth.RefreshAccountSessionResponse\x12O\n" +
	"\x0fValidateSession\x12\x16.google.protobuf.Empty\x1a$.auth.ValidateAccountSessionResponse\x12V\n" +
	"\rRevokeSession\x12!.auth.RevokeAccountSessionRequest\x1a\".auth.RevokeAccountSessionResponseB\tZ\asso/genb\x06proto3"

var (
	file_sso_sso_proto_rawDescOnce sync.Once
	file_sso_sso_proto_rawDescData []byte
)

func file_sso_sso_proto_rawDescGZIP() []byte {
	file_sso_sso_proto_rawDescOnce.Do(func() {
		file_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)))
	})
	return file_sso_sso_proto_rawDescData
}

var file_sso_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_sso_sso_proto_goTypes = []any{
	(AccountStatus)(0),                       // 0: auth.AccountStatus
	(AccountRole)(0),                         // 1: auth.AccountRole
	(*RegisterClientRequest)(nil),            // 2: auth.RegisterClientRequest
	(*RegisterClientResponse)(nil),           // 3: auth.RegisterClientResponse
	(*RegisterRequest)(nil),                  // 4: auth.RegisterRequest
	(*RegisterResponse)(nil),                 // 5: auth.RegisterResponse
	(*LoginRequest)(nil),                     // 6: auth.LoginRequest
	(*LoginResponse)(nil),                    // 7: auth.LoginResponse
	(*LogoutResponse)(nil),                   // 8: auth.LogoutResponse
	(*ChangePasswordRequest)(nil),            // 9: auth.ChangePasswordRequest
	(*ChangePasswordResponse)(nil),           // 10: auth.ChangePasswordResponse
	(*ChangeStatusRequest)(nil),              // 11: auth.ChangeStatusRequest
	(*ChangeStatusResponse)(nil),             // 12: auth.ChangeStatusResponse
	(*GetActiveAccountSessionsResponse)(nil), // 13: auth.GetActiveAccountSessionsResponse
	(*RefreshAccountSessionResponse)(nil),    // 14: auth.RefreshAccountSessionResponse
	(*ValidateAccountSessionResponse)(nil),   // 15: auth.ValidateAccountSessionResponse
	(*RevokeAccountSessionRequest)(nil),      // 16: auth.RevokeAccountSessionRequest
	(*RevokeAccountSessionResponse)(nil),     // 17: auth.RevokeAccountSessionResponse
	(*Session)(nil),                          // 18: auth.Session
	(*emptypb.Empty)(nil),                    // 19: google.protobuf.Empty
}
var file_sso_sso_proto_depIdxs = []int32{
	1,  // 0: auth.RegisterRequest.role:type_name -> auth.AccountRole
	0,  // 1: auth.ChangeStatusRequest.status:type_name -> auth.AccountStatus
	0,  // 2: auth.ChangeStatusResponse.status:type_name -> auth.AccountStatus
	18, // 3: auth.GetActiveAccountSessionsResponse.sessions:type_name -> auth.Session
	2,  // 4: auth.Auth.RegisterClient:input_type -> auth.RegisterClientRequest
	4,  // 5: auth.Auth.Register:input_type -> auth.RegisterRequest
	6,  // 6: auth.Auth.Login:input_type -> auth.LoginRequest
	19, // 7: auth.Auth.Logout:input_type -> google.protobuf.Empty
	9,  // 8: auth.Auth.ChangePassword:input_type -> auth.ChangePasswordRequest
	11, // 9: auth.Auth.ChangeStatus:input_type -> auth.ChangeStatusRequest
	19, // 10: auth.Sessions.GetActiveSessions:input_type -> google.protobuf.Empty
	19, // 11: auth.Sessions.RefreshSession:input_type -> google.protobuf.Empty
	19, // 12: auth.Sessions.ValidateSession:input_type -> google.protobuf.Empty
	16, // 13: auth.Sessions.RevokeSession:input_type -> auth.RevokeAccountSessionRequest
	3,  // 14: auth.Auth.RegisterClient:output_type -> auth.RegisterClientResponse
	5,  // 15: auth.Auth.Register:output_type -> auth.RegisterResponse
	7,  // 16: auth.Auth.Login:output_type -> auth.LoginResponse
	8,  // 17: auth.Auth.Logout:output_type -> auth.LogoutResponse
	10, // 18: auth.Auth.ChangePassword:output_type -> auth.ChangePasswordResponse
	12, // 19: auth.Auth.ChangeStatus:output_type -> auth.ChangeStatusResponse
	13, // 20: auth.Sessions.GetActiveSessions:output_type -> auth.GetActiveAccountSessionsResponse
	14, // 21: auth.Sessions.RefreshSession:output_type -> auth.RefreshAccountSessionResponse
	15, // 22: auth.Sessions.ValidateSession:output_type -> auth.ValidateAccountSessionResponse
	17, // 23: auth.Sessions.RevokeSession:output_type -> auth.RevokeAccountSessionResponse
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_sso_sso_proto_init() }
func file_sso_sso_proto_init() {
	if File_sso_sso_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_sso_sso_proto_goTypes,
		DependencyIndexes: file_sso_sso_proto_depIdxs,
		EnumInfos:         file_sso_sso_proto_enumTypes,
		MessageInfos:      file_sso_sso_proto_msgTypes,
	}.Build()
	File_sso_sso_proto = out.File
	file_sso_sso_proto_goTypes = nil
	file_sso_sso_proto_depIdxs = nil
}
