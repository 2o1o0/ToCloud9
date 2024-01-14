// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: guilds.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GuildService_GetGuildInfo_FullMethodName         = "/v1.GuildService/GetGuildInfo"
	GuildService_GetRosterInfo_FullMethodName        = "/v1.GuildService/GetRosterInfo"
	GuildService_InviteMember_FullMethodName         = "/v1.GuildService/InviteMember"
	GuildService_InviteAccepted_FullMethodName       = "/v1.GuildService/InviteAccepted"
	GuildService_Leave_FullMethodName                = "/v1.GuildService/Leave"
	GuildService_Kick_FullMethodName                 = "/v1.GuildService/Kick"
	GuildService_SetMessageOfTheDay_FullMethodName   = "/v1.GuildService/SetMessageOfTheDay"
	GuildService_SetGuildInfo_FullMethodName         = "/v1.GuildService/SetGuildInfo"
	GuildService_SetMemberPublicNote_FullMethodName  = "/v1.GuildService/SetMemberPublicNote"
	GuildService_SetMemberOfficerNote_FullMethodName = "/v1.GuildService/SetMemberOfficerNote"
	GuildService_UpdateRank_FullMethodName           = "/v1.GuildService/UpdateRank"
	GuildService_AddRank_FullMethodName              = "/v1.GuildService/AddRank"
	GuildService_DeleteLastRank_FullMethodName       = "/v1.GuildService/DeleteLastRank"
	GuildService_PromoteMember_FullMethodName        = "/v1.GuildService/PromoteMember"
	GuildService_DemoteMember_FullMethodName         = "/v1.GuildService/DemoteMember"
	GuildService_SendGuildMessage_FullMethodName     = "/v1.GuildService/SendGuildMessage"
)

// GuildServiceClient is the client API for GuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuildServiceClient interface {
	GetGuildInfo(ctx context.Context, in *GetInfoParams, opts ...grpc.CallOption) (*GetInfoResponse, error)
	GetRosterInfo(ctx context.Context, in *GetRosterInfoParams, opts ...grpc.CallOption) (*GetRosterInfoResponse, error)
	InviteMember(ctx context.Context, in *InviteMemberParams, opts ...grpc.CallOption) (*InviteMemberResponse, error)
	InviteAccepted(ctx context.Context, in *InviteAcceptedParams, opts ...grpc.CallOption) (*InviteAcceptedResponse, error)
	Leave(ctx context.Context, in *LeaveParams, opts ...grpc.CallOption) (*LeaveResponse, error)
	Kick(ctx context.Context, in *KickParams, opts ...grpc.CallOption) (*KickResponse, error)
	SetMessageOfTheDay(ctx context.Context, in *SetMessageOfTheDayParams, opts ...grpc.CallOption) (*SetMessageOfTheDayResponse, error)
	SetGuildInfo(ctx context.Context, in *SetGuildInfoParams, opts ...grpc.CallOption) (*SetGuildInfoResponse, error)
	SetMemberPublicNote(ctx context.Context, in *SetNoteParams, opts ...grpc.CallOption) (*SetNoteResponse, error)
	SetMemberOfficerNote(ctx context.Context, in *SetNoteParams, opts ...grpc.CallOption) (*SetNoteResponse, error)
	UpdateRank(ctx context.Context, in *RankUpdateParams, opts ...grpc.CallOption) (*RankUpdateResponse, error)
	AddRank(ctx context.Context, in *AddRankParams, opts ...grpc.CallOption) (*AddRankResponse, error)
	DeleteLastRank(ctx context.Context, in *DeleteLastRankParams, opts ...grpc.CallOption) (*DeleteLastRankResponse, error)
	PromoteMember(ctx context.Context, in *PromoteDemoteParams, opts ...grpc.CallOption) (*PromoteDemoteResponse, error)
	DemoteMember(ctx context.Context, in *PromoteDemoteParams, opts ...grpc.CallOption) (*PromoteDemoteResponse, error)
	SendGuildMessage(ctx context.Context, in *SendGuildMessageParams, opts ...grpc.CallOption) (*SendGuildMessageResponse, error)
}

type guildServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuildServiceClient(cc grpc.ClientConnInterface) GuildServiceClient {
	return &guildServiceClient{cc}
}

func (c *guildServiceClient) GetGuildInfo(ctx context.Context, in *GetInfoParams, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, GuildService_GetGuildInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) GetRosterInfo(ctx context.Context, in *GetRosterInfoParams, opts ...grpc.CallOption) (*GetRosterInfoResponse, error) {
	out := new(GetRosterInfoResponse)
	err := c.cc.Invoke(ctx, GuildService_GetRosterInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) InviteMember(ctx context.Context, in *InviteMemberParams, opts ...grpc.CallOption) (*InviteMemberResponse, error) {
	out := new(InviteMemberResponse)
	err := c.cc.Invoke(ctx, GuildService_InviteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) InviteAccepted(ctx context.Context, in *InviteAcceptedParams, opts ...grpc.CallOption) (*InviteAcceptedResponse, error) {
	out := new(InviteAcceptedResponse)
	err := c.cc.Invoke(ctx, GuildService_InviteAccepted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) Leave(ctx context.Context, in *LeaveParams, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, GuildService_Leave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) Kick(ctx context.Context, in *KickParams, opts ...grpc.CallOption) (*KickResponse, error) {
	out := new(KickResponse)
	err := c.cc.Invoke(ctx, GuildService_Kick_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) SetMessageOfTheDay(ctx context.Context, in *SetMessageOfTheDayParams, opts ...grpc.CallOption) (*SetMessageOfTheDayResponse, error) {
	out := new(SetMessageOfTheDayResponse)
	err := c.cc.Invoke(ctx, GuildService_SetMessageOfTheDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) SetGuildInfo(ctx context.Context, in *SetGuildInfoParams, opts ...grpc.CallOption) (*SetGuildInfoResponse, error) {
	out := new(SetGuildInfoResponse)
	err := c.cc.Invoke(ctx, GuildService_SetGuildInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) SetMemberPublicNote(ctx context.Context, in *SetNoteParams, opts ...grpc.CallOption) (*SetNoteResponse, error) {
	out := new(SetNoteResponse)
	err := c.cc.Invoke(ctx, GuildService_SetMemberPublicNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) SetMemberOfficerNote(ctx context.Context, in *SetNoteParams, opts ...grpc.CallOption) (*SetNoteResponse, error) {
	out := new(SetNoteResponse)
	err := c.cc.Invoke(ctx, GuildService_SetMemberOfficerNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) UpdateRank(ctx context.Context, in *RankUpdateParams, opts ...grpc.CallOption) (*RankUpdateResponse, error) {
	out := new(RankUpdateResponse)
	err := c.cc.Invoke(ctx, GuildService_UpdateRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) AddRank(ctx context.Context, in *AddRankParams, opts ...grpc.CallOption) (*AddRankResponse, error) {
	out := new(AddRankResponse)
	err := c.cc.Invoke(ctx, GuildService_AddRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) DeleteLastRank(ctx context.Context, in *DeleteLastRankParams, opts ...grpc.CallOption) (*DeleteLastRankResponse, error) {
	out := new(DeleteLastRankResponse)
	err := c.cc.Invoke(ctx, GuildService_DeleteLastRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) PromoteMember(ctx context.Context, in *PromoteDemoteParams, opts ...grpc.CallOption) (*PromoteDemoteResponse, error) {
	out := new(PromoteDemoteResponse)
	err := c.cc.Invoke(ctx, GuildService_PromoteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) DemoteMember(ctx context.Context, in *PromoteDemoteParams, opts ...grpc.CallOption) (*PromoteDemoteResponse, error) {
	out := new(PromoteDemoteResponse)
	err := c.cc.Invoke(ctx, GuildService_DemoteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildServiceClient) SendGuildMessage(ctx context.Context, in *SendGuildMessageParams, opts ...grpc.CallOption) (*SendGuildMessageResponse, error) {
	out := new(SendGuildMessageResponse)
	err := c.cc.Invoke(ctx, GuildService_SendGuildMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuildServiceServer is the server API for GuildService service.
// All implementations must embed UnimplementedGuildServiceServer
// for forward compatibility
type GuildServiceServer interface {
	GetGuildInfo(context.Context, *GetInfoParams) (*GetInfoResponse, error)
	GetRosterInfo(context.Context, *GetRosterInfoParams) (*GetRosterInfoResponse, error)
	InviteMember(context.Context, *InviteMemberParams) (*InviteMemberResponse, error)
	InviteAccepted(context.Context, *InviteAcceptedParams) (*InviteAcceptedResponse, error)
	Leave(context.Context, *LeaveParams) (*LeaveResponse, error)
	Kick(context.Context, *KickParams) (*KickResponse, error)
	SetMessageOfTheDay(context.Context, *SetMessageOfTheDayParams) (*SetMessageOfTheDayResponse, error)
	SetGuildInfo(context.Context, *SetGuildInfoParams) (*SetGuildInfoResponse, error)
	SetMemberPublicNote(context.Context, *SetNoteParams) (*SetNoteResponse, error)
	SetMemberOfficerNote(context.Context, *SetNoteParams) (*SetNoteResponse, error)
	UpdateRank(context.Context, *RankUpdateParams) (*RankUpdateResponse, error)
	AddRank(context.Context, *AddRankParams) (*AddRankResponse, error)
	DeleteLastRank(context.Context, *DeleteLastRankParams) (*DeleteLastRankResponse, error)
	PromoteMember(context.Context, *PromoteDemoteParams) (*PromoteDemoteResponse, error)
	DemoteMember(context.Context, *PromoteDemoteParams) (*PromoteDemoteResponse, error)
	SendGuildMessage(context.Context, *SendGuildMessageParams) (*SendGuildMessageResponse, error)
	mustEmbedUnimplementedGuildServiceServer()
}

// UnimplementedGuildServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGuildServiceServer struct {
}

func (UnimplementedGuildServiceServer) GetGuildInfo(context.Context, *GetInfoParams) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildInfo not implemented")
}
func (UnimplementedGuildServiceServer) GetRosterInfo(context.Context, *GetRosterInfoParams) (*GetRosterInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosterInfo not implemented")
}
func (UnimplementedGuildServiceServer) InviteMember(context.Context, *InviteMemberParams) (*InviteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteMember not implemented")
}
func (UnimplementedGuildServiceServer) InviteAccepted(context.Context, *InviteAcceptedParams) (*InviteAcceptedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteAccepted not implemented")
}
func (UnimplementedGuildServiceServer) Leave(context.Context, *LeaveParams) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedGuildServiceServer) Kick(context.Context, *KickParams) (*KickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kick not implemented")
}
func (UnimplementedGuildServiceServer) SetMessageOfTheDay(context.Context, *SetMessageOfTheDayParams) (*SetMessageOfTheDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMessageOfTheDay not implemented")
}
func (UnimplementedGuildServiceServer) SetGuildInfo(context.Context, *SetGuildInfoParams) (*SetGuildInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGuildInfo not implemented")
}
func (UnimplementedGuildServiceServer) SetMemberPublicNote(context.Context, *SetNoteParams) (*SetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMemberPublicNote not implemented")
}
func (UnimplementedGuildServiceServer) SetMemberOfficerNote(context.Context, *SetNoteParams) (*SetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMemberOfficerNote not implemented")
}
func (UnimplementedGuildServiceServer) UpdateRank(context.Context, *RankUpdateParams) (*RankUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRank not implemented")
}
func (UnimplementedGuildServiceServer) AddRank(context.Context, *AddRankParams) (*AddRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRank not implemented")
}
func (UnimplementedGuildServiceServer) DeleteLastRank(context.Context, *DeleteLastRankParams) (*DeleteLastRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLastRank not implemented")
}
func (UnimplementedGuildServiceServer) PromoteMember(context.Context, *PromoteDemoteParams) (*PromoteDemoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteMember not implemented")
}
func (UnimplementedGuildServiceServer) DemoteMember(context.Context, *PromoteDemoteParams) (*PromoteDemoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteMember not implemented")
}
func (UnimplementedGuildServiceServer) SendGuildMessage(context.Context, *SendGuildMessageParams) (*SendGuildMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGuildMessage not implemented")
}
func (UnimplementedGuildServiceServer) mustEmbedUnimplementedGuildServiceServer() {}

// UnsafeGuildServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuildServiceServer will
// result in compilation errors.
type UnsafeGuildServiceServer interface {
	mustEmbedUnimplementedGuildServiceServer()
}

func RegisterGuildServiceServer(s grpc.ServiceRegistrar, srv GuildServiceServer) {
	s.RegisterService(&GuildService_ServiceDesc, srv)
}

func _GuildService_GetGuildInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).GetGuildInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_GetGuildInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).GetGuildInfo(ctx, req.(*GetInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_GetRosterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRosterInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).GetRosterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_GetRosterInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).GetRosterInfo(ctx, req.(*GetRosterInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_InviteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteMemberParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).InviteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_InviteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).InviteMember(ctx, req.(*InviteMemberParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_InviteAccepted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteAcceptedParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).InviteAccepted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_InviteAccepted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).InviteAccepted(ctx, req.(*InviteAcceptedParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_Leave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).Leave(ctx, req.(*LeaveParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_Kick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).Kick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_Kick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).Kick(ctx, req.(*KickParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_SetMessageOfTheDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMessageOfTheDayParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).SetMessageOfTheDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_SetMessageOfTheDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).SetMessageOfTheDay(ctx, req.(*SetMessageOfTheDayParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_SetGuildInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGuildInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).SetGuildInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_SetGuildInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).SetGuildInfo(ctx, req.(*SetGuildInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_SetMemberPublicNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).SetMemberPublicNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_SetMemberPublicNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).SetMemberPublicNote(ctx, req.(*SetNoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_SetMemberOfficerNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).SetMemberOfficerNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_SetMemberOfficerNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).SetMemberOfficerNote(ctx, req.(*SetNoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_UpdateRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).UpdateRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_UpdateRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).UpdateRank(ctx, req.(*RankUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_AddRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).AddRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_AddRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).AddRank(ctx, req.(*AddRankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_DeleteLastRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLastRankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).DeleteLastRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_DeleteLastRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).DeleteLastRank(ctx, req.(*DeleteLastRankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_PromoteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteDemoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).PromoteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_PromoteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).PromoteMember(ctx, req.(*PromoteDemoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_DemoteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteDemoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).DemoteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_DemoteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).DemoteMember(ctx, req.(*PromoteDemoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildService_SendGuildMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGuildMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildServiceServer).SendGuildMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuildService_SendGuildMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildServiceServer).SendGuildMessage(ctx, req.(*SendGuildMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

// GuildService_ServiceDesc is the grpc.ServiceDesc for GuildService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuildService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GuildService",
	HandlerType: (*GuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGuildInfo",
			Handler:    _GuildService_GetGuildInfo_Handler,
		},
		{
			MethodName: "GetRosterInfo",
			Handler:    _GuildService_GetRosterInfo_Handler,
		},
		{
			MethodName: "InviteMember",
			Handler:    _GuildService_InviteMember_Handler,
		},
		{
			MethodName: "InviteAccepted",
			Handler:    _GuildService_InviteAccepted_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _GuildService_Leave_Handler,
		},
		{
			MethodName: "Kick",
			Handler:    _GuildService_Kick_Handler,
		},
		{
			MethodName: "SetMessageOfTheDay",
			Handler:    _GuildService_SetMessageOfTheDay_Handler,
		},
		{
			MethodName: "SetGuildInfo",
			Handler:    _GuildService_SetGuildInfo_Handler,
		},
		{
			MethodName: "SetMemberPublicNote",
			Handler:    _GuildService_SetMemberPublicNote_Handler,
		},
		{
			MethodName: "SetMemberOfficerNote",
			Handler:    _GuildService_SetMemberOfficerNote_Handler,
		},
		{
			MethodName: "UpdateRank",
			Handler:    _GuildService_UpdateRank_Handler,
		},
		{
			MethodName: "AddRank",
			Handler:    _GuildService_AddRank_Handler,
		},
		{
			MethodName: "DeleteLastRank",
			Handler:    _GuildService_DeleteLastRank_Handler,
		},
		{
			MethodName: "PromoteMember",
			Handler:    _GuildService_PromoteMember_Handler,
		},
		{
			MethodName: "DemoteMember",
			Handler:    _GuildService_DemoteMember_Handler,
		},
		{
			MethodName: "SendGuildMessage",
			Handler:    _GuildService_SendGuildMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guilds.proto",
}
