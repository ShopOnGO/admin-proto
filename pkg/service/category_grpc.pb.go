// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: category.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CategoryService_CreateCategory_FullMethodName        = "/proto.CategoryService/CreateCategory"
	CategoryService_GetFeaturedCategories_FullMethodName = "/proto.CategoryService/GetFeaturedCategories"
	CategoryService_FindCategoryByName_FullMethodName    = "/proto.CategoryService/FindCategoryByName"
	CategoryService_FindCategoryByID_FullMethodName      = "/proto.CategoryService/FindCategoryByID"
	CategoryService_UpdateCategory_FullMethodName        = "/proto.CategoryService/UpdateCategory"
	CategoryService_DeleteCategory_FullMethodName        = "/proto.CategoryService/DeleteCategory"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Описание gRPC-сервиса
type CategoryServiceClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	GetFeaturedCategories(ctx context.Context, in *GetFeaturedCategoriesRequest, opts ...grpc.CallOption) (*GetFeaturedCategoriesResponse, error)
	FindCategoryByName(ctx context.Context, in *FindCategoryByNameRequest, opts ...grpc.CallOption) (*FindCategoryByNameResponse, error)
	FindCategoryByID(ctx context.Context, in *FindCategoryByIDRequest, opts ...grpc.CallOption) (*FindCategoryByIDResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryByNameRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetFeaturedCategories(ctx context.Context, in *GetFeaturedCategoriesRequest, opts ...grpc.CallOption) (*GetFeaturedCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFeaturedCategoriesResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetFeaturedCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) FindCategoryByName(ctx context.Context, in *FindCategoryByNameRequest, opts ...grpc.CallOption) (*FindCategoryByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCategoryByNameResponse)
	err := c.cc.Invoke(ctx, CategoryService_FindCategoryByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) FindCategoryByID(ctx context.Context, in *FindCategoryByIDRequest, opts ...grpc.CallOption) (*FindCategoryByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCategoryByIDResponse)
	err := c.cc.Invoke(ctx, CategoryService_FindCategoryByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryByNameRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility.
//
// Описание gRPC-сервиса
type CategoryServiceServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	GetFeaturedCategories(context.Context, *GetFeaturedCategoriesRequest) (*GetFeaturedCategoriesResponse, error)
	FindCategoryByName(context.Context, *FindCategoryByNameRequest) (*FindCategoryByNameResponse, error)
	FindCategoryByID(context.Context, *FindCategoryByIDRequest) (*FindCategoryByIDResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryByNameRequest) (*DeleteCategoryResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryServiceServer struct{}

func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetFeaturedCategories(context.Context, *GetFeaturedCategoriesRequest) (*GetFeaturedCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeaturedCategories not implemented")
}
func (UnimplementedCategoryServiceServer) FindCategoryByName(context.Context, *FindCategoryByNameRequest) (*FindCategoryByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategoryByName not implemented")
}
func (UnimplementedCategoryServiceServer) FindCategoryByID(context.Context, *FindCategoryByIDRequest) (*FindCategoryByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategoryByID not implemented")
}
func (UnimplementedCategoryServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategory(context.Context, *DeleteCategoryByNameRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}
func (UnimplementedCategoryServiceServer) testEmbeddedByValue()                         {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedCategoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetFeaturedCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeaturedCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetFeaturedCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetFeaturedCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetFeaturedCategories(ctx, req.(*GetFeaturedCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_FindCategoryByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCategoryByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindCategoryByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_FindCategoryByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindCategoryByName(ctx, req.(*FindCategoryByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_FindCategoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCategoryByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindCategoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_FindCategoryByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindCategoryByID(ctx, req.(*FindCategoryByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "GetFeaturedCategories",
			Handler:    _CategoryService_GetFeaturedCategories_Handler,
		},
		{
			MethodName: "FindCategoryByName",
			Handler:    _CategoryService_FindCategoryByName_Handler,
		},
		{
			MethodName: "FindCategoryByID",
			Handler:    _CategoryService_FindCategoryByID_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CategoryService_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}
