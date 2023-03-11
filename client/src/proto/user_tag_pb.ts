// @generated by protoc-gen-es v1.1.0 with parameter "target=ts"
// @generated from file user_tag.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message proto.ListUserTagRequest
 */
export class ListUserTagRequest extends Message<ListUserTagRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  constructor(data?: PartialMessage<ListUserTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.ListUserTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserTagRequest {
    return new ListUserTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserTagRequest {
    return new ListUserTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserTagRequest {
    return new ListUserTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListUserTagRequest | PlainMessage<ListUserTagRequest> | undefined, b: ListUserTagRequest | PlainMessage<ListUserTagRequest> | undefined): boolean {
    return proto3.util.equals(ListUserTagRequest, a, b);
  }
}

/**
 * @generated from message proto.ListUserTagResponse
 */
export class ListUserTagResponse extends Message<ListUserTagResponse> {
  /**
   * @generated from field: repeated proto.UserTagResponse user_tag_list = 1;
   */
  userTagList: UserTagResponse[] = [];

  constructor(data?: PartialMessage<ListUserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.ListUserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_tag_list", kind: "message", T: UserTagResponse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserTagResponse {
    return new ListUserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserTagResponse {
    return new ListUserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserTagResponse {
    return new ListUserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListUserTagResponse | PlainMessage<ListUserTagResponse> | undefined, b: ListUserTagResponse | PlainMessage<ListUserTagResponse> | undefined): boolean {
    return proto3.util.equals(ListUserTagResponse, a, b);
  }
}

/**
 * @generated from message proto.UserTagResponse
 */
export class UserTagResponse extends Message<UserTagResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: string tag_name = 3;
   */
  tagName = "";

  /**
   * @generated from field: bool has_group = 4;
   */
  hasGroup = false;

  /**
   * @generated from field: int64 group_id = 5;
   */
  groupId = protoInt64.zero;

  /**
   * @generated from field: string grant_limit = 6;
   */
  grantLimit = "";

  constructor(data?: PartialMessage<UserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.UserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tag_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "has_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "group_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "grant_limit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserTagResponse {
    return new UserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserTagResponse {
    return new UserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserTagResponse {
    return new UserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UserTagResponse | PlainMessage<UserTagResponse> | undefined, b: UserTagResponse | PlainMessage<UserTagResponse> | undefined): boolean {
    return proto3.util.equals(UserTagResponse, a, b);
  }
}

/**
 * @generated from message proto.GetUserTagRequest
 */
export class GetUserTagRequest extends Message<GetUserTagRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  constructor(data?: PartialMessage<GetUserTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.GetUserTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserTagRequest {
    return new GetUserTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserTagRequest {
    return new GetUserTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserTagRequest {
    return new GetUserTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserTagRequest | PlainMessage<GetUserTagRequest> | undefined, b: GetUserTagRequest | PlainMessage<GetUserTagRequest> | undefined): boolean {
    return proto3.util.equals(GetUserTagRequest, a, b);
  }
}

/**
 * @generated from message proto.GetUserTagResponse
 */
export class GetUserTagResponse extends Message<GetUserTagResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: string tag_name = 3;
   */
  tagName = "";

  /**
   * @generated from field: bool has_group = 4;
   */
  hasGroup = false;

  /**
   * @generated from field: int64 group_id = 5;
   */
  groupId = protoInt64.zero;

  /**
   * @generated from field: string grant_limit = 6;
   */
  grantLimit = "";

  constructor(data?: PartialMessage<GetUserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.GetUserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tag_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "has_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "group_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "grant_limit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserTagResponse {
    return new GetUserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserTagResponse {
    return new GetUserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserTagResponse {
    return new GetUserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserTagResponse | PlainMessage<GetUserTagResponse> | undefined, b: GetUserTagResponse | PlainMessage<GetUserTagResponse> | undefined): boolean {
    return proto3.util.equals(GetUserTagResponse, a, b);
  }
}

/**
 * @generated from message proto.RegisterUserTagRequest
 */
export class RegisterUserTagRequest extends Message<RegisterUserTagRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  /**
   * @generated from field: string tag_name = 2;
   */
  tagName = "";

  /**
   * @generated from field: bool has_group = 3;
   */
  hasGroup = false;

  /**
   * @generated from field: int64 group_id = 4;
   */
  groupId = protoInt64.zero;

  /**
   * @generated from field: string grant_limit = 5;
   */
  grantLimit = "";

  constructor(data?: PartialMessage<RegisterUserTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.RegisterUserTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tag_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "has_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "group_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "grant_limit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterUserTagRequest {
    return new RegisterUserTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterUserTagRequest {
    return new RegisterUserTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterUserTagRequest {
    return new RegisterUserTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterUserTagRequest | PlainMessage<RegisterUserTagRequest> | undefined, b: RegisterUserTagRequest | PlainMessage<RegisterUserTagRequest> | undefined): boolean {
    return proto3.util.equals(RegisterUserTagRequest, a, b);
  }
}

/**
 * @generated from message proto.RegisterUserTagResponse
 */
export class RegisterUserTagResponse extends Message<RegisterUserTagResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: string tag_name = 3;
   */
  tagName = "";

  /**
   * @generated from field: bool has_group = 4;
   */
  hasGroup = false;

  /**
   * @generated from field: int64 group_id = 5;
   */
  groupId = protoInt64.zero;

  /**
   * @generated from field: string grant_limit = 6;
   */
  grantLimit = "";

  constructor(data?: PartialMessage<RegisterUserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.RegisterUserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tag_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "has_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "group_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "grant_limit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterUserTagResponse {
    return new RegisterUserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterUserTagResponse {
    return new RegisterUserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterUserTagResponse {
    return new RegisterUserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterUserTagResponse | PlainMessage<RegisterUserTagResponse> | undefined, b: RegisterUserTagResponse | PlainMessage<RegisterUserTagResponse> | undefined): boolean {
    return proto3.util.equals(RegisterUserTagResponse, a, b);
  }
}

/**
 * @generated from message proto.UpdateUserTagRequest
 */
export class UpdateUserTagRequest extends Message<UpdateUserTagRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  constructor(data?: PartialMessage<UpdateUserTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.UpdateUserTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserTagRequest {
    return new UpdateUserTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserTagRequest {
    return new UpdateUserTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserTagRequest {
    return new UpdateUserTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserTagRequest | PlainMessage<UpdateUserTagRequest> | undefined, b: UpdateUserTagRequest | PlainMessage<UpdateUserTagRequest> | undefined): boolean {
    return proto3.util.equals(UpdateUserTagRequest, a, b);
  }
}

/**
 * @generated from message proto.UpdateUserTagResponse
 */
export class UpdateUserTagResponse extends Message<UpdateUserTagResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: string tag_name = 3;
   */
  tagName = "";

  /**
   * @generated from field: bool has_group = 4;
   */
  hasGroup = false;

  /**
   * @generated from field: int64 group_id = 5;
   */
  groupId = protoInt64.zero;

  /**
   * @generated from field: string grant_limit = 6;
   */
  grantLimit = "";

  constructor(data?: PartialMessage<UpdateUserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.UpdateUserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tag_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "has_group", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "group_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "grant_limit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserTagResponse {
    return new UpdateUserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserTagResponse {
    return new UpdateUserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserTagResponse {
    return new UpdateUserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserTagResponse | PlainMessage<UpdateUserTagResponse> | undefined, b: UpdateUserTagResponse | PlainMessage<UpdateUserTagResponse> | undefined): boolean {
    return proto3.util.equals(UpdateUserTagResponse, a, b);
  }
}

/**
 * @generated from message proto.DeleteUserTagRequest
 */
export class DeleteUserTagRequest extends Message<DeleteUserTagRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  constructor(data?: PartialMessage<DeleteUserTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.DeleteUserTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserTagRequest {
    return new DeleteUserTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserTagRequest {
    return new DeleteUserTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserTagRequest {
    return new DeleteUserTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserTagRequest | PlainMessage<DeleteUserTagRequest> | undefined, b: DeleteUserTagRequest | PlainMessage<DeleteUserTagRequest> | undefined): boolean {
    return proto3.util.equals(DeleteUserTagRequest, a, b);
  }
}

/**
 * @generated from message proto.DeleteUserTagResponse
 */
export class DeleteUserTagResponse extends Message<DeleteUserTagResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  constructor(data?: PartialMessage<DeleteUserTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.DeleteUserTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserTagResponse {
    return new DeleteUserTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserTagResponse {
    return new DeleteUserTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserTagResponse {
    return new DeleteUserTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserTagResponse | PlainMessage<DeleteUserTagResponse> | undefined, b: DeleteUserTagResponse | PlainMessage<DeleteUserTagResponse> | undefined): boolean {
    return proto3.util.equals(DeleteUserTagResponse, a, b);
  }
}

