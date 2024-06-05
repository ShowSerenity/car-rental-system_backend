// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: cars.proto

package cars_service_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CarService_GetCars_FullMethodName   = "/cars.CarService/GetCars"
	CarService_GetCar_FullMethodName    = "/cars.CarService/GetCar"
	CarService_AddCar_FullMethodName    = "/cars.CarService/AddCar"
	CarService_UpdateCar_FullMethodName = "/cars.CarService/UpdateCar"
	CarService_DeleteCar_FullMethodName = "/cars.CarService/DeleteCar"
)

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarServiceClient interface {
	GetCars(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*GetCarsResponse, error)
	GetCar(ctx context.Context, in *GetCarRequest, opts ...grpc.CallOption) (*Car, error)
	AddCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*Car, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) GetCars(ctx context.Context, in *GetCarsRequest, opts ...grpc.CallOption) (*GetCarsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCarsResponse)
	err := c.cc.Invoke(ctx, CarService_GetCars_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) GetCar(ctx context.Context, in *GetCarRequest, opts ...grpc.CallOption) (*Car, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Car)
	err := c.cc.Invoke(ctx, CarService_GetCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) AddCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*Car, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Car)
	err := c.cc.Invoke(ctx, CarService_AddCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*Car, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Car)
	err := c.cc.Invoke(ctx, CarService_UpdateCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CarService_DeleteCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServiceServer is the server API for CarService service.
// All implementations must embed UnimplementedCarServiceServer
// for forward compatibility
type CarServiceServer interface {
	GetCars(context.Context, *GetCarsRequest) (*GetCarsResponse, error)
	GetCar(context.Context, *GetCarRequest) (*Car, error)
	AddCar(context.Context, *Car) (*Car, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*Car, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCarServiceServer()
}

// UnimplementedCarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (UnimplementedCarServiceServer) GetCars(context.Context, *GetCarsRequest) (*GetCarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCars not implemented")
}
func (UnimplementedCarServiceServer) GetCar(context.Context, *GetCarRequest) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}
func (UnimplementedCarServiceServer) AddCar(context.Context, *Car) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCar not implemented")
}
func (UnimplementedCarServiceServer) UpdateCar(context.Context, *UpdateCarRequest) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (UnimplementedCarServiceServer) DeleteCar(context.Context, *DeleteCarRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (UnimplementedCarServiceServer) mustEmbedUnimplementedCarServiceServer() {}

// UnsafeCarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServiceServer will
// result in compilation errors.
type UnsafeCarServiceServer interface {
	mustEmbedUnimplementedCarServiceServer()
}

func RegisterCarServiceServer(s grpc.ServiceRegistrar, srv CarServiceServer) {
	s.RegisterService(&CarService_ServiceDesc, srv)
}

func _CarService_GetCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarService_GetCars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetCars(ctx, req.(*GetCarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_GetCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarService_GetCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetCar(ctx, req.(*GetCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_AddCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Car)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).AddCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarService_AddCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).AddCar(ctx, req.(*Car))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarService_UpdateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CarService_DeleteCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CarService_ServiceDesc is the grpc.ServiceDesc for CarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cars.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCars",
			Handler:    _CarService_GetCars_Handler,
		},
		{
			MethodName: "GetCar",
			Handler:    _CarService_GetCar_Handler,
		},
		{
			MethodName: "AddCar",
			Handler:    _CarService_AddCar_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _CarService_UpdateCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _CarService_DeleteCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cars.proto",
}
