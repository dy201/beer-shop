// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ShopInterfaceClient is the client API for ShopInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopInterfaceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error)
	ListAddress(ctx context.Context, in *ListAddressReq, opts ...grpc.CallOption) (*ListAddressReply, error)
	CreateAddress(ctx context.Context, in *CreateAddressReq, opts ...grpc.CallOption) (*CreateAddressReply, error)
	GetAddress(ctx context.Context, in *GetAddressReq, opts ...grpc.CallOption) (*GetAddressReply, error)
	ListCard(ctx context.Context, in *ListCardReq, opts ...grpc.CallOption) (*ListCardReply, error)
	CreateCard(ctx context.Context, in *CreateCardReq, opts ...grpc.CallOption) (*CreateCardReply, error)
	GetCard(ctx context.Context, in *GetCardReq, opts ...grpc.CallOption) (*GetCardReply, error)
	DeleteCard(ctx context.Context, in *DeleteCardReq, opts ...grpc.CallOption) (*DeleteCardReply, error)
	ListBeer(ctx context.Context, in *ListBeerReq, opts ...grpc.CallOption) (*ListBeerReply, error)
	GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error)
	ListCartItem(ctx context.Context, in *ListCartItemReq, opts ...grpc.CallOption) (*ListCartItemReply, error)
	AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*AddCartItemReply, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error)
	ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderReply, error)
}

type shopInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopInterfaceClient(cc grpc.ClientConnInterface) ShopInterfaceClient {
	return &shopInterfaceClient{cc}
}

func (c *shopInterfaceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListAddress(ctx context.Context, in *ListAddressReq, opts ...grpc.CallOption) (*ListAddressReply, error) {
	out := new(ListAddressReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) CreateAddress(ctx context.Context, in *CreateAddressReq, opts ...grpc.CallOption) (*CreateAddressReply, error) {
	out := new(CreateAddressReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetAddress(ctx context.Context, in *GetAddressReq, opts ...grpc.CallOption) (*GetAddressReply, error) {
	out := new(GetAddressReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListCard(ctx context.Context, in *ListCardReq, opts ...grpc.CallOption) (*ListCardReply, error) {
	out := new(ListCardReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) CreateCard(ctx context.Context, in *CreateCardReq, opts ...grpc.CallOption) (*CreateCardReply, error) {
	out := new(CreateCardReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetCard(ctx context.Context, in *GetCardReq, opts ...grpc.CallOption) (*GetCardReply, error) {
	out := new(GetCardReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) DeleteCard(ctx context.Context, in *DeleteCardReq, opts ...grpc.CallOption) (*DeleteCardReply, error) {
	out := new(DeleteCardReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/DeleteCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListBeer(ctx context.Context, in *ListBeerReq, opts ...grpc.CallOption) (*ListBeerReply, error) {
	out := new(ListBeerReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error) {
	out := new(GetBeerReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/GetBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListCartItem(ctx context.Context, in *ListCartItemReq, opts ...grpc.CallOption) (*ListCartItemReply, error) {
	out := new(ListCartItemReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*AddCartItemReply, error) {
	out := new(AddCartItemReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/AddCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopInterfaceServer is the server API for ShopInterface service.
// All implementations must embed UnimplementedShopInterfaceServer
// for forward compatibility
type ShopInterfaceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	ListAddress(context.Context, *ListAddressReq) (*ListAddressReply, error)
	CreateAddress(context.Context, *CreateAddressReq) (*CreateAddressReply, error)
	GetAddress(context.Context, *GetAddressReq) (*GetAddressReply, error)
	ListCard(context.Context, *ListCardReq) (*ListCardReply, error)
	CreateCard(context.Context, *CreateCardReq) (*CreateCardReply, error)
	GetCard(context.Context, *GetCardReq) (*GetCardReply, error)
	DeleteCard(context.Context, *DeleteCardReq) (*DeleteCardReply, error)
	ListBeer(context.Context, *ListBeerReq) (*ListBeerReply, error)
	GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error)
	ListCartItem(context.Context, *ListCartItemReq) (*ListCartItemReply, error)
	AddCartItem(context.Context, *AddCartItemReq) (*AddCartItemReply, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	mustEmbedUnimplementedShopInterfaceServer()
}

// UnimplementedShopInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedShopInterfaceServer struct {
}

func (UnimplementedShopInterfaceServer) Register(context.Context, *RegisterReq) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedShopInterfaceServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedShopInterfaceServer) Logout(context.Context, *LogoutReq) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedShopInterfaceServer) ListAddress(context.Context, *ListAddressReq) (*ListAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddress not implemented")
}
func (UnimplementedShopInterfaceServer) CreateAddress(context.Context, *CreateAddressReq) (*CreateAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedShopInterfaceServer) GetAddress(context.Context, *GetAddressReq) (*GetAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedShopInterfaceServer) ListCard(context.Context, *ListCardReq) (*ListCardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCard not implemented")
}
func (UnimplementedShopInterfaceServer) CreateCard(context.Context, *CreateCardReq) (*CreateCardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedShopInterfaceServer) GetCard(context.Context, *GetCardReq) (*GetCardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedShopInterfaceServer) DeleteCard(context.Context, *DeleteCardReq) (*DeleteCardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}
func (UnimplementedShopInterfaceServer) ListBeer(context.Context, *ListBeerReq) (*ListBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBeer not implemented")
}
func (UnimplementedShopInterfaceServer) GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeer not implemented")
}
func (UnimplementedShopInterfaceServer) ListCartItem(context.Context, *ListCartItemReq) (*ListCartItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCartItem not implemented")
}
func (UnimplementedShopInterfaceServer) AddCartItem(context.Context, *AddCartItemReq) (*AddCartItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCartItem not implemented")
}
func (UnimplementedShopInterfaceServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedShopInterfaceServer) ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedShopInterfaceServer) mustEmbedUnimplementedShopInterfaceServer() {}

// UnsafeShopInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopInterfaceServer will
// result in compilation errors.
type UnsafeShopInterfaceServer interface {
	mustEmbedUnimplementedShopInterfaceServer()
}

func RegisterShopInterfaceServer(s grpc.ServiceRegistrar, srv ShopInterfaceServer) {
	s.RegisterService(&ShopInterface_ServiceDesc, srv)
}

func _ShopInterface_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListAddress(ctx, req.(*ListAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/CreateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).CreateAddress(ctx, req.(*CreateAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetAddress(ctx, req.(*GetAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListCard(ctx, req.(*ListCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).CreateCard(ctx, req.(*CreateCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetCard(ctx, req.(*GetCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/DeleteCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).DeleteCard(ctx, req.(*DeleteCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListBeer(ctx, req.(*ListBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/GetBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetBeer(ctx, req.(*GetBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListCartItem(ctx, req.(*ListCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_AddCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).AddCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/AddCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).AddCartItem(ctx, req.(*AddCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListOrder(ctx, req.(*ListOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopInterface_ServiceDesc is the grpc.ServiceDesc for ShopInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.interface.v1.ShopInterface",
	HandlerType: (*ShopInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ShopInterface_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ShopInterface_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ShopInterface_Logout_Handler,
		},
		{
			MethodName: "ListAddress",
			Handler:    _ShopInterface_ListAddress_Handler,
		},
		{
			MethodName: "CreateAddress",
			Handler:    _ShopInterface_CreateAddress_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _ShopInterface_GetAddress_Handler,
		},
		{
			MethodName: "ListCard",
			Handler:    _ShopInterface_ListCard_Handler,
		},
		{
			MethodName: "CreateCard",
			Handler:    _ShopInterface_CreateCard_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _ShopInterface_GetCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _ShopInterface_DeleteCard_Handler,
		},
		{
			MethodName: "ListBeer",
			Handler:    _ShopInterface_ListBeer_Handler,
		},
		{
			MethodName: "GetBeer",
			Handler:    _ShopInterface_GetBeer_Handler,
		},
		{
			MethodName: "ListCartItem",
			Handler:    _ShopInterface_ListCartItem_Handler,
		},
		{
			MethodName: "AddCartItem",
			Handler:    _ShopInterface_AddCartItem_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _ShopInterface_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _ShopInterface_ListOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/shop.proto",
}
